package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/taskcluster/slugid-go/slugid"
	tcclient "github.com/taskcluster/taskcluster-client-go"
	"github.com/taskcluster/taskcluster-client-go/secrets"
)

func main() {
	sshSecret := make(map[string]string)
	if len(os.Args) != 2 {
		log.Fatal("Usage: " + os.Args[0] + " WORKER_TYPE_DIRECTORY")
	}
	workerTypeDir := os.Args[1]
	absFile, err := filepath.Abs(workerTypeDir)
	if err != nil {
		log.Fatalf("File/directory '%v' could not be read due to '%s'", workerTypeDir, err)
	}
	files, err := ioutil.ReadDir(workerTypeDir)
	if err != nil {
		log.Fatalf("File/directory '%v' (%v) could not be read due to '%s'", workerTypeDir, absFile, err)
	}

	workerType := filepath.Base(absFile)
	secretName := "project/taskcluster/aws-provisioner-v1/worker-types/ssh-keys/" + workerType

	tcCreds := &tcclient.Credentials{
		ClientID:    os.Getenv("TASKCLUSTER_CLIENT_ID"),
		AccessToken: os.Getenv("TASKCLUSTER_ACCESS_TOKEN"),
		Certificate: os.Getenv("TASKCLUSTER_CERTIFICATE"),
	}

	cd := &tcclient.Client{
		Credentials:  tcCreds,
		BaseURL:      "https://aws-provisioner.taskcluster.net/v1",
		Authenticate: true,
	}
	var wt map[string]interface{}
	_, _, err = cd.APICall(nil, "GET", "/worker-type/"+url.QueryEscape(workerType), &wt, nil)
	if err != nil {
		log.Fatal(err)
	}
	regions := wt["regions"].([]interface{})
	oldAMICount := 0
	newAMICount := 0

	delete(wt, "lastModified")
	delete(wt, "workerType")
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".id_rsa") {
			region := f.Name()[:len(f.Name())-7]
			bytes, err := ioutil.ReadFile(filepath.Join(workerTypeDir, f.Name()))
			if err != nil {
				log.Fatalf("Problem reading file %v", filepath.Join(workerTypeDir, f.Name()))
			}
			sshSecret[region] = string(bytes)
		}
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".latest-ami") {
			newAMICount++
			tokens := strings.Split(f.Name(), ".")
			region := tokens[0]
			newAmi := tokens[1]
			oldAmi := ""
			for i := range regions {
				regionObj := regions[i].(map[string]interface{})
				if regionObj["region"] == region {
					launchSpec := regionObj["launchSpec"].(map[string]interface{})
					oldAmi = launchSpec["ImageId"].(string)
					launchSpec["ImageId"] = newAmi
					oldAMICount++
				}
			}
			if newAMICount < oldAMICount {
				log.Fatal(fmt.Errorf("Did not find ami specification in worker type %v for region %v", workerType, region))
			}
			if newAMICount > oldAMICount {
				log.Fatal(fmt.Errorf("Found multiple AMIs in worker type %v for region %v", workerType, region))
			}
			log.Print("Old AMI for worker type " + workerType + " region " + region + ": " + oldAmi)
			log.Print("New AMI for worker type " + workerType + " region " + region + ": " + newAmi)
			if oldAmi == newAmi {
				log.Print("WARNING: No change to AMI used in workert type " + workerType + " for region " + region + " (" + oldAmi + ")")
			} else {
				log.Print("Worker type " + workerType + " region " + region + " updated to use " + newAmi)
			}
		}
	}

	secretsMap := wt["secrets"].(map[string]interface{})
	genericWorker := secretsMap["generic-worker"].(map[string]interface{})
	config := genericWorker["config"].(map[string]interface{})
	oldDeploymentID := config["deploymentId"].(string)
	newDeploymentID := slugid.Nice()
	config["deploymentId"] = newDeploymentID
	log.Print("Old deployment ID: " + oldDeploymentID)
	log.Print("New deployment ID: " + newDeploymentID)

	if newAMICount != len(regions) {
		log.Printf("WARNING: not updating all AMIs for worker type %v - only %v of %v", workerType, newAMICount, len(regions))
	}

	mySecrets := secrets.New(tcCreds)

	secBytes, err := json.Marshal(sshSecret)
	if err != nil {
		log.Fatalf("Could not convert secret %#v to json: %v", sshSecret, err)
	}

	err = mySecrets.Set(
		secretName,
		&secrets.Secret{
			Expires: tcclient.Time(time.Now().AddDate(1, 0, 0)),
			Secret:  json.RawMessage(secBytes),
		},
	)
	if err != nil {
		log.Printf("Problem publishing new secrets: %v", err)
	}
	s, err := mySecrets.Get(secretName)
	if err != nil {
		log.Fatalf("Error retrieving secret: %v", err)
	}
	log.Print("Secret name:  " + secretName)
	log.Print("Secret value: " + string(s.Secret))
	log.Print("Expiry:       " + s.Expires.String())
	_, _, err = cd.APICall(wt, "POST", "/worker-type/"+url.QueryEscape(workerType)+"/update", new(interface{}), nil)
	if err != nil {
		log.Fatal(err)
	}
}
