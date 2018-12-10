package go_expvar

import (
	"expvar"
	"github.com/MovieStoreGuy/polygraph"
	"gopkg.in/yaml.v2"
	"reflect"
)

type exporter struct {

}

func NewExporter() polygraph.Exporter {
	return &exporter{}
}

func (e *exporter) Configure(data []byte) error {
	return yaml.Unmarshal(data, e)
}

func (e *exporter) Export(obj interface, metricName ...string) {

	if reflect.ValueOf(obj).Kind() != reflect.Struct {

	}
	interalObj := reflect.ValueOf(obj)
	if interalObj.Kind() == reflect.Ptr {
		
	}
}

func (e *exporter) Start() error {

	return nil
}