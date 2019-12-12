package validate

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

// Validater ...
type Validater interface {
	Validate(body []byte) (*gojsonschema.Result, error)
}

// CreateValidater ...
var (
	Create *Schema
	Change *Schema
)

// Schema ...
type Schema struct {
	schema gojsonschema.JSONLoader
}

func init() {
	crFile, err := loadFile("internal/adapters/http/validate/createEvent.json")
	if err != nil {
		fmt.Println(err)
	}
	Create = &Schema{schema: crFile}

	chFile, err := loadFile("internal/adapters/http/validate/changeEvent.json")
	if err != nil {
		fmt.Println(err)
	}
	Change = &Schema{schema: chFile}
}

// Validate ...
func (c *Schema) Validate(body []byte) (*gojsonschema.Result, error) {
	loader := gojsonschema.NewBytesLoader(body)
	result, err := gojsonschema.Validate(c.schema, loader)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return result, nil
}

func loadFile(file string) (gojsonschema.JSONLoader, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	s, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	schema := gojsonschema.NewBytesLoader(s)
	return schema, nil
}
