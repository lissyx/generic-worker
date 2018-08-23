// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package main

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	Artifact struct {

		// Explicitly set the value of the HTTP `Content-Type` response header when the artifact(s)
		// is/are served over HTTP(S). If not provided (this property is optional) the worker will
		// guess the content type of artifacts based on the filename extension of the file storing
		// the artifact content. It does this by looking at the system filename-to-mimetype mappings
		// defined in the Windows registry. Note, setting `contentType` on a directory artifact will
		// apply the same contentType to all files contained in the directory.
		//
		// See [mime.TypeByExtension](https://godoc.org/mime#TypeByExtension).
		//
		// Since: generic-worker 10.4.0
		ContentType string `json:"contentType,omitempty"`

		// Date when artifact should expire must be in the future, no earlier than task deadline, but
		// no later than task expiry. If not set, defaults to task expiry.
		//
		// Since: generic-worker 1.0.0
		Expires tcclient.Time `json:"expires,omitempty"`

		// Name of the artifact, as it will be published. If not set, `path` will be used.
		// Conventionally (although not enforced) path elements are forward slash separated. Example:
		// `public/build/a/house`. Note, no scopes are required to read artifacts beginning `public/`.
		// Artifact names not beginning `public/` are scope-protected (caller requires scopes to
		// download the artifact). See the Queue documentation for more information.
		//
		// Since: generic-worker 8.1.0
		Name string `json:"name,omitempty"`

		// Relative path of the file/directory from the task directory. Note this is not an absolute
		// path as is typically used in docker-worker, since the absolute task directory name is not
		// known when the task is submitted. Example: `dist\regedit.exe`. It doesn't matter if
		// forward slashes or backslashes are used.
		//
		// Since: generic-worker 1.0.0
		Path string `json:"path"`

		// Artifacts can be either an individual `file` or a `directory` containing
		// potentially multiple files with recursively included subdirectories.
		//
		// Since: generic-worker 1.0.0
		//
		// Possible values:
		//   * "file"
		//   * "directory"
		Type string `json:"type"`
	}

	// Requires scope `queue:get-artifact:<artifact-name>`.
	//
	// Since: generic-worker 5.4.0
	ArtifactContent struct {

		// Max length: 1024
		Artifact string `json:"artifact"`

		// The required SHA 256 of the content body.
		//
		// Since: generic-worker 10.8.0
		//
		// Syntax:     ^[a-f0-9]{64}$
		Sha256 string `json:"sha256,omitempty"`

		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}

	// By default tasks will be resolved with `state/reasonResolved`: `completed/completed`
	// if all task commands have a zero exit code, or `failed/failed` if any command has a
	// non-zero exit code. This payload property allows customsation of the task resolution
	// based on exit code of task commands.
	ExitCodeHandling struct {

		// Exit codes for any command in the task payload to cause this task to
		// be resolved as `exception/intermittent-task`. Typically the Queue
		// will then schedule a new run of the existing `taskId` (rerun) if not
		// all task runs have been exhausted.
		//
		// See [itermittent tasks](https://docs.taskcluster.net/docs/reference/platform/taskcluster-queue/docs/worker-interaction#intermittent-tasks) for more detail.
		//
		// Since: generic-worker 10.10.0
		Retry []int64 `json:"retry,omitempty"`
	}

	// Feature flags enable additional functionality.
	//
	// Since: generic-worker 5.3.0
	FeatureFlags struct {

		// An artifact named `public/chainOfTrust.json.asc` should be generated
		// which will include information for downstream tasks to build a level
		// of trust for the artifacts produced by the task and the environment
		// it ran in.
		//
		// Since: generic-worker 5.3.0
		ChainOfTrust bool `json:"chainOfTrust,omitempty"`

		// Runs commands with UAC elevation. Only set to true when UAC is
		// enabled on the worker and Administrative privileges are required by
		// task commands. When UAC is disabled on the worker, task commands will
		// already run with full user privileges, and therefore a value of true
		// will result in a malformed-payload task exception.
		//
		// A value of true does not add the task user to the `Administrators`
		// group - see the `osGroups` property for that. Typically
		// `task.payload.osGroups` should include an Administrative group, such
		// as `Administrators`, when setting to true.
		//
		// For security, `runAsAdministrator` feature cannot be used in
		// conjunction with `chainOfTrust` feature.
		//
		// Requires scope
		// `generic-worker:run-as-administrator:<provisionerId>/<workerType>`.
		//
		// Since: generic-worker 10.11.0
		RunAsAdministrator bool `json:"runAsAdministrator,omitempty"`

		// The taskcluster proxy provides an easy and safe way to make authenticated
		// taskcluster requests within the scope(s) of a particular task. See
		// [the github project](https://github.com/taskcluster/taskcluster-proxy) for more information.
		//
		// Since: generic-worker 10.6.0
		TaskclusterProxy bool `json:"taskclusterProxy,omitempty"`
	}

	FileMount struct {

		// One of:
		//   * ArtifactContent
		//   * URLContent
		Content json.RawMessage `json:"content"`

		// The filesystem location to mount the file.
		//
		// Since: generic-worker 5.4.0
		File string `json:"file"`
	}

	// This schema defines the structure of the `payload` property referred to in a
	// Taskcluster Task definition.
	GenericWorkerPayload struct {

		// Artifacts to be published.
		//
		// Since: generic-worker 1.0.0
		Artifacts []Artifact `json:"artifacts,omitempty"`

		// One entry per command (consider each entry to be interpreted as a full line of
		// a Windows™ .bat file). For example:
		// ```
		// [
		//   "set",
		//   "echo hello world > hello_world.txt",
		//   "set GOPATH=C:\\Go"
		// ]
		// ```
		//
		// Since: generic-worker 0.0.1
		Command []string `json:"command"`

		// Env vars must be string to __string__ mappings (not number or boolean). For example:
		// ```
		// {
		//   "PATH": "C:\\Windows\\system32;C:\\Windows",
		//   "GOOS": "windows",
		//   "FOO_ENABLE": "true",
		//   "BAR_TOTAL": "3"
		// }
		// ```
		//
		// Since: generic-worker 0.0.1
		Env map[string]string `json:"env,omitempty"`

		// Feature flags enable additional functionality.
		//
		// Since: generic-worker 5.3.0
		Features FeatureFlags `json:"features,omitempty"`

		// Maximum time the task container can run in seconds.
		//
		// Since: generic-worker 0.0.1
		//
		// Mininum:    1
		// Maximum:    86400
		MaxRunTime int64 `json:"maxRunTime"`

		// Directories and/or files to be mounted.
		//
		// Since: generic-worker 5.4.0
		Mounts []json.RawMessage `json:"mounts,omitempty"`

		// By default tasks will be resolved with `state/reasonResolved`: `completed/completed`
		// if all task commands have a zero exit code, or `failed/failed` if any command has a
		// non-zero exit code. This payload property allows customsation of the task resolution
		// based on exit code of task commands.
		OnExitStatus ExitCodeHandling `json:"onExitStatus,omitempty"`

		// A list of OS Groups that the task user should be a member of. Requires scope
		// `generic-worker:os-group:<provisionerId>/<workerType>/<os-group>` for each
		// group listed.
		//
		// Since: generic-worker 6.0.0
		OSGroups []string `json:"osGroups,omitempty"`

		// Specifies an artifact name for publishing RDP connection information.
		//
		// Since this is potentially sensitive data, care should be taken to publish
		// to a suitably locked down path, such as
		// `login-identity/<login-identity>/rdpinfo.json` which is only readable for
		// the given login identity (for example
		// `login-identity/mozilla-ldap/pmoore@mozilla.com/rdpinfo.json`). See the
		// [artifact namespace guide](https://docs.taskcluster.net/manual/design/namespaces#artifacts) for more information.
		//
		// Use of this feature requires scope
		// `generic-worker:allow-rdp:<provisionerId>/<workerType>` which must be
		// declared as a task scope.
		//
		// The RDP connection data is published during task startup so that a user
		// may interact with the running task.
		//
		// The task environment will be retained for 12 hours after the task
		// completes, to enable an interactive user to perform investigative tasks.
		// After these 12 hours, the worker will delete the task's Windows user
		// account, and then continue with other tasks.
		//
		// No guarantees are given about the resolution status of the interactive
		// task, since the task is inherently non-reproducible and no automation
		// should rely on this value.
		//
		// Since: generic-worker 10.5.0
		RdpInfo string `json:"rdpInfo,omitempty"`

		// URL of a service that can indicate tasks superseding this one; the current `taskId`
		// will be appended as a query argument `taskId`. The service should return an object with
		// a `supersedes` key containing a list of `taskId`s, including the supplied `taskId`. The
		// tasks should be ordered such that each task supersedes all tasks appearing later in the
		// list.
		//
		// See [superseding](https://docs.taskcluster.net/reference/platform/taskcluster-queue/docs/superseding) for more detail.
		//
		// Since: generic-worker 10.2.2
		SupersederURL string `json:"supersederUrl,omitempty"`
	}

	ReadOnlyDirectory struct {

		// One of:
		//   * ArtifactContent
		//   * URLContent
		Content json.RawMessage `json:"content"`

		// The filesystem location to mount the directory volume.
		//
		// Since: generic-worker 5.4.0
		Directory string `json:"directory"`

		// Archive format of content for read only directory.
		//
		// Since: generic-worker 5.4.0
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "zip"
		Format string `json:"format"`
	}

	// URL to download content from.
	//
	// Since: generic-worker 5.4.0
	URLContent struct {

		// The required SHA 256 of the content body.
		//
		// Since: generic-worker 10.8.0
		//
		// Syntax:     ^[a-f0-9]{64}$
		Sha256 string `json:"sha256,omitempty"`

		// URL to download content from.
		//
		// Since: generic-worker 5.4.0
		URL string `json:"url"`
	}

	WritableDirectoryCache struct {

		// Implies a read/write cache directory volume. A unique name for the
		// cache volume. Requires scope `generic-worker:cache:<cache-name>`.
		// Note if this cache is loaded from an artifact, you will also require
		// scope `queue:get-artifact:<artifact-name>` to use this cache.
		//
		// Since: generic-worker 5.4.0
		CacheName string `json:"cacheName"`

		// One of:
		//   * ArtifactContent
		//   * URLContent
		Content json.RawMessage `json:"content,omitempty"`

		// The filesystem location to mount the directory volume.
		//
		// Since: generic-worker 5.4.0
		Directory string `json:"directory"`

		// Archive format of the preloaded content (if `content` provided).
		//
		// Since: generic-worker 5.4.0
		//
		// Possible values:
		//   * "rar"
		//   * "tar.bz2"
		//   * "tar.gz"
		//   * "zip"
		Format string `json:"format,omitempty"`
	}
)

// Returns json schema for the payload part of the task definition. Please
// note we use a go string and do not load an external file, since we want this
// to be *part of the compiled executable*. If this sat in another file that
// was loaded at runtime, it would not be burned into the build, which would be
// bad for the following two reasons:
//  1) we could no longer distribute a single binary file that didn't require
//     installation/extraction
//  2) the payload schema is specific to the version of the code, therefore
//     should be versioned directly with the code and *frozen on build*.
//
// Run `generic-worker show-payload-schema` to output this schema to standard
// out.
func taskPayloadSchema() string {
	return `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "additionalProperties": false,
  "definitions": {
    "content": {
      "oneOf": [
        {
          "additionalProperties": false,
          "description": "Requires scope ` + "`" + `queue:get-artifact:\u003cartifact-name\u003e` + "`" + `.\n\nSince: generic-worker 5.4.0",
          "properties": {
            "artifact": {
              "maxLength": 1024,
              "type": "string"
            },
            "sha256": {
              "description": "The required SHA 256 of the content body.\n\nSince: generic-worker 10.8.0",
              "pattern": "^[a-f0-9]{64}$",
              "title": "SHA 256",
              "type": "string"
            },
            "taskId": {
              "pattern": "^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$",
              "type": "string"
            }
          },
          "required": [
            "taskId",
            "artifact"
          ],
          "title": "Artifact Content",
          "type": "object"
        },
        {
          "additionalProperties": false,
          "description": "URL to download content from.\n\nSince: generic-worker 5.4.0",
          "properties": {
            "sha256": {
              "description": "The required SHA 256 of the content body.\n\nSince: generic-worker 10.8.0",
              "pattern": "^[a-f0-9]{64}$",
              "title": "SHA 256",
              "type": "string"
            },
            "url": {
              "description": "URL to download content from.\n\nSince: generic-worker 5.4.0",
              "format": "uri",
              "title": "URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ],
          "title": "URL Content",
          "type": "object"
        }
      ]
    },
    "fileMount": {
      "additionalProperties": false,
      "properties": {
        "content": {
          "$ref": "#/definitions/content",
          "description": "Content of the file to be mounted.\n\nSince: generic-worker 5.4.0"
        },
        "file": {
          "description": "The filesystem location to mount the file.\n\nSince: generic-worker 5.4.0",
          "title": "File",
          "type": "string"
        }
      },
      "required": [
        "file",
        "content"
      ],
      "title": "File Mount",
      "type": "object"
    },
    "mount": {
      "oneOf": [
        {
          "$ref": "#/definitions/fileMount"
        },
        {
          "$ref": "#/definitions/writableDirectoryCache"
        },
        {
          "$ref": "#/definitions/readOnlyDirectory"
        }
      ],
      "title": "Mount"
    },
    "readOnlyDirectory": {
      "additionalProperties": false,
      "properties": {
        "content": {
          "$ref": "#/definitions/content",
          "description": "Contents of read only directory.\n\nSince: generic-worker 5.4.0",
          "title": "Content"
        },
        "directory": {
          "description": "The filesystem location to mount the directory volume.\n\nSince: generic-worker 5.4.0",
          "title": "Directory",
          "type": "string"
        },
        "format": {
          "description": "Archive format of content for read only directory.\n\nSince: generic-worker 5.4.0",
          "enum": [
            "rar",
            "tar.bz2",
            "tar.gz",
            "zip"
          ],
          "title": "Format",
          "type": "string"
        }
      },
      "required": [
        "directory",
        "content",
        "format"
      ],
      "title": "Read Only Directory",
      "type": "object"
    },
    "writableDirectoryCache": {
      "additionalProperties": false,
      "dependencies": {
        "content": [
          "format"
        ],
        "format": [
          "content"
        ]
      },
      "properties": {
        "cacheName": {
          "description": "Implies a read/write cache directory volume. A unique name for the\ncache volume. Requires scope ` + "`" + `generic-worker:cache:\u003ccache-name\u003e` + "`" + `.\nNote if this cache is loaded from an artifact, you will also require\nscope ` + "`" + `queue:get-artifact:\u003cartifact-name\u003e` + "`" + ` to use this cache.\n\nSince: generic-worker 5.4.0",
          "title": "Cache Name",
          "type": "string"
        },
        "content": {
          "$ref": "#/definitions/content",
          "description": "Optional content to be preloaded when initially creating the cache\n(if set, ` + "`" + `format` + "`" + ` must also be provided).\n\nSince: generic-worker 5.4.0",
          "title": "Content"
        },
        "directory": {
          "description": "The filesystem location to mount the directory volume.\n\nSince: generic-worker 5.4.0",
          "title": "Directory Volume",
          "type": "string"
        },
        "format": {
          "description": "Archive format of the preloaded content (if ` + "`" + `content` + "`" + ` provided).\n\nSince: generic-worker 5.4.0",
          "enum": [
            "rar",
            "tar.bz2",
            "tar.gz",
            "zip"
          ],
          "title": "Format",
          "type": "string"
        }
      },
      "required": [
        "directory",
        "cacheName"
      ],
      "title": "Writable Directory Cache",
      "type": "object"
    }
  },
  "description": "This schema defines the structure of the ` + "`" + `payload` + "`" + ` property referred to in a\nTaskcluster Task definition.",
  "properties": {
    "artifacts": {
      "description": "Artifacts to be published.\n\nSince: generic-worker 1.0.0",
      "items": {
        "additionalProperties": false,
        "properties": {
          "contentType": {
            "description": "Explicitly set the value of the HTTP ` + "`" + `Content-Type` + "`" + ` response header when the artifact(s)\nis/are served over HTTP(S). If not provided (this property is optional) the worker will\nguess the content type of artifacts based on the filename extension of the file storing\nthe artifact content. It does this by looking at the system filename-to-mimetype mappings\ndefined in the Windows registry. Note, setting ` + "`" + `contentType` + "`" + ` on a directory artifact will\napply the same contentType to all files contained in the directory.\n\nSee [mime.TypeByExtension](https://godoc.org/mime#TypeByExtension).\n\nSince: generic-worker 10.4.0",
            "title": "Content-Type header when serving artifact over HTTP",
            "type": "string"
          },
          "expires": {
            "description": "Date when artifact should expire must be in the future, no earlier than task deadline, but\nno later than task expiry. If not set, defaults to task expiry.\n\nSince: generic-worker 1.0.0",
            "format": "date-time",
            "title": "Expiry date and time",
            "type": "string"
          },
          "name": {
            "description": "Name of the artifact, as it will be published. If not set, ` + "`" + `path` + "`" + ` will be used.\nConventionally (although not enforced) path elements are forward slash separated. Example:\n` + "`" + `public/build/a/house` + "`" + `. Note, no scopes are required to read artifacts beginning ` + "`" + `public/` + "`" + `.\nArtifact names not beginning ` + "`" + `public/` + "`" + ` are scope-protected (caller requires scopes to\ndownload the artifact). See the Queue documentation for more information.\n\nSince: generic-worker 8.1.0",
            "title": "Name of the artifact",
            "type": "string"
          },
          "path": {
            "description": "Relative path of the file/directory from the task directory. Note this is not an absolute\npath as is typically used in docker-worker, since the absolute task directory name is not\nknown when the task is submitted. Example: ` + "`" + `dist\\regedit.exe` + "`" + `. It doesn't matter if\nforward slashes or backslashes are used.\n\nSince: generic-worker 1.0.0",
            "title": "Artifact location",
            "type": "string"
          },
          "type": {
            "description": "Artifacts can be either an individual ` + "`" + `file` + "`" + ` or a ` + "`" + `directory` + "`" + ` containing\npotentially multiple files with recursively included subdirectories.\n\nSince: generic-worker 1.0.0",
            "enum": [
              "file",
              "directory"
            ],
            "title": "Artifact upload type.",
            "type": "string"
          }
        },
        "required": [
          "type",
          "path"
        ],
        "title": "Artifact",
        "type": "object"
      },
      "title": "Artifacts to be published",
      "type": "array"
    },
    "command": {
      "description": "One entry per command (consider each entry to be interpreted as a full line of\na Windows™ .bat file). For example:\n` + "`" + `` + "`" + `` + "`" + `\n[\n  \"set\",\n  \"echo hello world \u003e hello_world.txt\",\n  \"set GOPATH=C:\\\\Go\"\n]\n` + "`" + `` + "`" + `` + "`" + `\n\nSince: generic-worker 0.0.1",
      "items": {
        "type": "string"
      },
      "minItems": 1,
      "title": "Commands to run",
      "type": "array"
    },
    "env": {
      "additionalProperties": {
        "type": "string"
      },
      "description": "Env vars must be string to __string__ mappings (not number or boolean). For example:\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"PATH\": \"C:\\\\Windows\\\\system32;C:\\\\Windows\",\n  \"GOOS\": \"windows\",\n  \"FOO_ENABLE\": \"true\",\n  \"BAR_TOTAL\": \"3\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n\nSince: generic-worker 0.0.1",
      "title": "Env vars",
      "type": "object"
    },
    "features": {
      "additionalProperties": false,
      "description": "Feature flags enable additional functionality.\n\nSince: generic-worker 5.3.0",
      "properties": {
        "chainOfTrust": {
          "description": "An artifact named ` + "`" + `public/chainOfTrust.json.asc` + "`" + ` should be generated\nwhich will include information for downstream tasks to build a level\nof trust for the artifacts produced by the task and the environment\nit ran in.\n\nSince: generic-worker 5.3.0",
          "title": "Enable generation of a openpgp signed Chain of Trust artifact",
          "type": "boolean"
        },
        "runAsAdministrator": {
          "description": "Runs commands with UAC elevation. Only set to true when UAC is\nenabled on the worker and Administrative privileges are required by\ntask commands. When UAC is disabled on the worker, task commands will\nalready run with full user privileges, and therefore a value of true\nwill result in a malformed-payload task exception.\n\nA value of true does not add the task user to the ` + "`" + `Administrators` + "`" + `\ngroup - see the ` + "`" + `osGroups` + "`" + ` property for that. Typically\n` + "`" + `task.payload.osGroups` + "`" + ` should include an Administrative group, such\nas ` + "`" + `Administrators` + "`" + `, when setting to true.\n\nFor security, ` + "`" + `runAsAdministrator` + "`" + ` feature cannot be used in\nconjunction with ` + "`" + `chainOfTrust` + "`" + ` feature.\n\nRequires scope\n` + "`" + `generic-worker:run-as-administrator:\u003cprovisionerId\u003e/\u003cworkerType\u003e` + "`" + `.\n\nSince: generic-worker 10.11.0",
          "title": "Run commands with UAC process elevation",
          "type": "boolean"
        },
        "taskclusterProxy": {
          "description": "The taskcluster proxy provides an easy and safe way to make authenticated\ntaskcluster requests within the scope(s) of a particular task. See\n[the github project](https://github.com/taskcluster/taskcluster-proxy) for more information.\n\nSince: generic-worker 10.6.0",
          "title": "Run [taskcluster-proxy](https://github.com/taskcluster/taskcluster-proxy) to allow tasks to dynamically proxy requests to taskcluster services",
          "type": "boolean"
        }
      },
      "title": "Feature flags",
      "type": "object"
    },
    "maxRunTime": {
      "description": "Maximum time the task container can run in seconds.\n\nSince: generic-worker 0.0.1",
      "maximum": 86400,
      "minimum": 1,
      "multipleOf": 1,
      "title": "Maximum run time in seconds",
      "type": "integer"
    },
    "mounts": {
      "description": "Directories and/or files to be mounted.\n\nSince: generic-worker 5.4.0",
      "items": {
        "$ref": "#/definitions/mount",
        "title": "Mount"
      },
      "type": "array"
    },
    "onExitStatus": {
      "additionalProperties": false,
      "description": "By default tasks will be resolved with ` + "`" + `state/reasonResolved` + "`" + `: ` + "`" + `completed/completed` + "`" + `\nif all task commands have a zero exit code, or ` + "`" + `failed/failed` + "`" + ` if any command has a\nnon-zero exit code. This payload property allows customsation of the task resolution\nbased on exit code of task commands.",
      "properties": {
        "retry": {
          "description": "Exit codes for any command in the task payload to cause this task to\nbe resolved as ` + "`" + `exception/intermittent-task` + "`" + `. Typically the Queue\nwill then schedule a new run of the existing ` + "`" + `taskId` + "`" + ` (rerun) if not\nall task runs have been exhausted.\n\nSee [itermittent tasks](https://docs.taskcluster.net/docs/reference/platform/taskcluster-queue/docs/worker-interaction#intermittent-tasks) for more detail.\n\nSince: generic-worker 10.10.0",
          "items": {
            "minimum": 1,
            "title": "Exit codes",
            "type": "integer"
          },
          "title": "Intermittent task exit codes",
          "type": "array"
        }
      },
      "title": "Exit code handling",
      "type": "object"
    },
    "osGroups": {
      "description": "A list of OS Groups that the task user should be a member of. Requires scope\n` + "`" + `generic-worker:os-group:\u003cprovisionerId\u003e/\u003cworkerType\u003e/\u003cos-group\u003e` + "`" + ` for each\ngroup listed.\n\nSince: generic-worker 6.0.0",
      "items": {
        "type": "string"
      },
      "title": "OS Groups",
      "type": "array"
    },
    "rdpInfo": {
      "description": "Specifies an artifact name for publishing RDP connection information.\n\nSince this is potentially sensitive data, care should be taken to publish\nto a suitably locked down path, such as\n` + "`" + `login-identity/\u003clogin-identity\u003e/rdpinfo.json` + "`" + ` which is only readable for\nthe given login identity (for example\n` + "`" + `login-identity/mozilla-ldap/pmoore@mozilla.com/rdpinfo.json` + "`" + `). See the\n[artifact namespace guide](https://docs.taskcluster.net/manual/design/namespaces#artifacts) for more information.\n\nUse of this feature requires scope\n` + "`" + `generic-worker:allow-rdp:\u003cprovisionerId\u003e/\u003cworkerType\u003e` + "`" + ` which must be\ndeclared as a task scope.\n\nThe RDP connection data is published during task startup so that a user\nmay interact with the running task.\n\nThe task environment will be retained for 12 hours after the task\ncompletes, to enable an interactive user to perform investigative tasks.\nAfter these 12 hours, the worker will delete the task's Windows user\naccount, and then continue with other tasks.\n\nNo guarantees are given about the resolution status of the interactive\ntask, since the task is inherently non-reproducible and no automation\nshould rely on this value.\n\nSince: generic-worker 10.5.0",
      "title": "RDP Info",
      "type": "string"
    },
    "supersederUrl": {
      "description": "URL of a service that can indicate tasks superseding this one; the current ` + "`" + `taskId` + "`" + `\nwill be appended as a query argument ` + "`" + `taskId` + "`" + `. The service should return an object with\na ` + "`" + `supersedes` + "`" + ` key containing a list of ` + "`" + `taskId` + "`" + `s, including the supplied ` + "`" + `taskId` + "`" + `. The\ntasks should be ordered such that each task supersedes all tasks appearing later in the\nlist.\n\nSee [superseding](https://docs.taskcluster.net/reference/platform/taskcluster-queue/docs/superseding) for more detail.\n\nSince: generic-worker 10.2.2",
      "format": "uri",
      "title": "Superseder URL",
      "type": "string"
    }
  },
  "required": [
    "command",
    "maxRunTime"
  ],
  "title": "Generic worker payload",
  "type": "object"
}`
}
