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
