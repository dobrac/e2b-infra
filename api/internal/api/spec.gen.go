// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
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
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
