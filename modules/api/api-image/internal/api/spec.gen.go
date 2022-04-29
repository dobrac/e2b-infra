// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
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

	"H4sIAAAAAAAC/8xUTW/UMBD9K9bAMWqWVlxyQ4BgJSoqKnGpKmSSSdaQ2O7MpKha5b8j29km+0UXCQGX",
	"jddjz7yZ9/zWULrOO4tWGIo1cLnCTsflWyJHYeHJeSQxGLdLV2H4VsglGS/GWSjSYRVjGdSOOi1QgLFy",
	"cQ4ZyIPH9BcbJBgy6JBZN0cTbcKPV1nI2AaGIQPCu94QVlDcwFhwc/x2yOAamWOqXeCcAl9MtV91WaEV",
	"Uxsk5WolK1Tj6ScRzLLeDiFobO32C7zB+6/OfVevrpYhpZEWp13I4B4poYYXZ4uzRRiR82i1N1DARdzK",
	"wGtZxVby8NOg7Nd5j7qVlSpXWIa0oX8dQssKCniHAgE9e2c5DeV8sdhP8gnvemRRPzQr7ssSmeu+hSH2",
	"l48d81EMHwyL0m2rHk8eAHI9xQ4BKp0VtDG59r41Zbycf+PEbJJpWBnBLl7UbfuxhuJmDc8JayjgWT5J",
	"Ox91nW/UMQSpjMxqIv0AsbntRq4fe28fFKH0ZLHa7mzI4OVvIv4VvPTmDkBJz6I1LMY2s/KhC93wTIkc",
	"X4F3fICZ14RaUOlNAuXsKHa6R9qj6crxH+PpJFKeoKCM8KsJ/18ef6w/m/+x8c9fSb6eHGJIjLQoeMgh",
	"wv7J3KTjG3bG77KKLkG6Q0Hi+BpMSB6cAzKwuotWNnnW3MyEesxmk9o1vtv/QAJpev9OArH+KRKYtneJ",
	"vtRWNyHF58u5Q25zw8GhfgYAAP//WwXm150HAAA=",
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