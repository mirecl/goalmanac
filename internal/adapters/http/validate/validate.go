package validate

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

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
	// Создаем валидаторы
	if err := createValidate(); err != nil {
		log.WithFields(log.Fields{"type": "http"}).Errorln(err.Error())
		os.Exit(0)
	}
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
	paths, err := filepath.Abs("./config")
	if err != nil {
		return nil, err
	}
	fileS := path.Join(paths, file)
	s, err := ioutil.ReadFile(fileS)
	if err != nil {
		return nil, err
	}
	schema := gojsonschema.NewBytesLoader(s)
	return schema, nil
}

func createValidate() error {
	// Чтение схемы для валидации создания события
	crFile, err := loadFile("ValidateCreateEvent.json")
	if err != nil {
		return err
	}
	Create = &Schema{schema: crFile}

	// Чтение схемы для валидации изменения события
	chFile, err := loadFile("ValidateChangeEvent.json")
	if err != nil {
		return err
	}
	Change = &Schema{schema: chFile}
	return nil
}
