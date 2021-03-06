package go_expvar

import (
	"expvar"
	"fmt"
	"github.com/MovieStoreGuy/polygraph"
	"gopkg.in/yaml.v2"
	"net/http"
	"reflect"
)

type exporter struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewExporter() polygraph.Exporter {
	return &exporter{}
}

func (e *exporter) Configure(data []byte) error {
	return yaml.Unmarshal(data, e)
}

func (e *exporter) PublishStruct(ref interface{}) {
	obj := reflect.ValueOf(ref)
	if obj.Kind() != reflect.Ptr || obj.Elem().Kind() != reflect.Struct {
		return
	}
	obj = obj.Elem()
	for i := 0; i < obj.NumField(); i++ {
		nested := obj.Field(i)
		if !nested.IsValid() || !nested.CanAddr() || !nested.Addr().CanInterface() {
			continue
		}
	Check:
		switch nested.Kind() {
		case reflect.Map, reflect.Slice, reflect.Invalid:
			// Do nothing
		case reflect.Ptr:
			if !nested.IsValid() || !nested.Addr().CanInterface() {
				continue
			}
			nested = nested.Elem()
			goto Check
		case reflect.Struct:
			e.PublishStruct(nested.Addr().Interface())
		default:
			if tag, exist := obj.Type().Field(i).Tag.Lookup(polygraph.Tag); exist {
				if !nested.IsValid() || !nested.Addr().CanInterface() {
					continue
				}
				e.PublishVariable(nested.Addr().Interface(), tag)
			}
		}
	}
}

func (e *exporter) Start() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", e.Host, e.Port), nil)
}

func (e *exporter) PublishVariable(ref interface{}, label string) {
	if ref == nil {
		return
	}
	metric := &morph{}
	metric.Set(ref, label)
	expvar.Publish(label, metric)
}
