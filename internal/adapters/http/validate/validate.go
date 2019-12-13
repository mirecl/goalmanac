package validate

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

// Validater - интерфейс для валидации endpoint по POST
type Validater interface {
	Validate(body []byte) (*gojsonschema.Result, error)
}

// CreateValidater - валидаторы
var (
	Create *Schema
	Change *Schema
)

// Schema ...
type Schema struct {
	schema gojsonschema.JSONLoader
}

func init() {
	// Чтение схемы для валидации создания события
	crFile, err := loadFile("createEvent.json")
	if err != nil {
		log.WithFields(log.Fields{"type": "cmd"}).Errorln(err.Error())
		os.Exit(0)
	}
	Create = &Schema{schema: crFile}

	// Чтение схемы для валидации изменения события
	chFile, err := loadFile("changeEvent.json")
	if err != nil {
		log.WithFields(log.Fields{"type": "cmd"}).Errorln(err.Error())
		os.Exit(0)
	}
	Change = &Schema{schema: chFile}
}

// Validate ...
func (c *Schema) Validate(body []byte) (*gojsonschema.Result, error) {
	loader := gojsonschema.NewBytesLoader(body)
	result, err := gojsonschema.Validate(c.schema, loader)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func loadFile(file string) (gojsonschema.JSONLoader, error) {
	_, paths, _, _ := runtime.Caller(0)
	fileS := path.Join(paths, "../", file)
	s, err := ioutil.ReadFile(fileS)
	if err != nil {
		return nil, err
	}
	schema := gojsonschema.NewBytesLoader(s)
	return schema, nil
}