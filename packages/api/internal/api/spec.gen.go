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
