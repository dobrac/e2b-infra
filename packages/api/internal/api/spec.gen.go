// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
