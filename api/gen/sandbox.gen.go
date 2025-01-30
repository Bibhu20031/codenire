// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package api

// ImageConfig defines model for ImageConfig.
type ImageConfig struct {
	CompileCmd    string                   `json:"CompileCmd"`
	Description   string                   `json:"Description"`
	Name          string                   `json:"Name"`
	Options       ImageConfigOption        `json:"Options"`
	RunCmd        string                   `json:"RunCmd"`
	ScriptOptions ImageConfigScriptOptions `json:"ScriptOptions"`
	Version       *string                  `json:"Version,omitempty"`
}

// ImageConfigList defines model for ImageConfigList.
type ImageConfigList = []ImageConfig

// ImageConfigOption defines model for ImageConfigOption.
type ImageConfigOption struct {
	CompileTTL *int `json:"CompileTTL,omitempty"`
	RunTTL     *int `json:"RunTTL,omitempty"`
}

// ImageConfigScriptOptions defines model for ImageConfigScriptOptions.
type ImageConfigScriptOptions struct {
	SourceFile string `json:"SourceFile"`
}

// SandboxRequest defines model for SandboxRequest.
type SandboxRequest struct {
	Args string `json:"args"`

	// Binary files in tar archive encoded with base64
	Binary string `json:"binary"`
	SandId string `json:"sandId"`
}

// SandboxResponse defines model for SandboxResponse.
type SandboxResponse struct {
	Error    *string `json:"error,omitempty"`
	ExitCode int     `json:"exitCode"`
	Stderr   []byte  `json:"stderr"`
	Stdout   []byte  `json:"stdout"`
}

// RunSandboxJSONRequestBody defines body for RunSandbox for application/json ContentType.
type RunSandboxJSONRequestBody = SandboxRequest
