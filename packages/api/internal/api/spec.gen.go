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
