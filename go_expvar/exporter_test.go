package go_expvar

import (
	"expvar"
	"fmt"
	"testing"
)

func TestInsertingVariables(t *testing.T) {
	var (
		counter = 0
		label   = "counter"
	)
	export := NewExporter()
	export.PublishVariable(&counter, label)
	for ; counter < 10; counter++ {
		v := expvar.Get(label)
		if v == nil {
			t.Fatal(`"counter" was not correctly registered in expvar`)
		}
		if v.String() != fmt.Sprint(counter) {
			t.Fatalf("Expected %s to match %s", v.String(), fmt.Sprint(counter))
		}
	}
}

func TestInsertingBasicStruct(t *testing.T) {
	type metricsBlob struct {
		Version   int    `polygraph:"application.version"`
		Accessors int    `polygraph:"application.accessors"`
		hidden    string `polygraph:"hidden"`
	}
	var (
		m = metricsBlob{
			Version: 1,
		}
	)
	export := NewExporter()
	export.PublishStruct(&m)
	if v := expvar.Get("application.version"); v == nil || v.String() != fmt.Sprint(1) {
		t.Fatal(`Was not able to correctly export variable "application.version"`)
	}
	if v := expvar.Get("hidden"); v != nil {
		t.Fatal("expvar has a hidden variable published")
	}
	for ; m.Accessors < 10; m.Accessors++ {
		v := expvar.Get("application.accessors")
		if v == nil {
			t.Fatal(`"application.accessors" was not registered correctly in expvar`)
		}
		if v.String() != fmt.Sprint(m.Accessors) {
			t.Fatalf("Expected %s to match %s", v.String(), fmt.Sprint(m.Accessors))
		}
	}
}

func TestPassingNil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Had to recover from a panic", r)
		}
	}()
	NewExporter().PublishStruct(nil)
}

func TestNestedComplexStructures(t *testing.T) {
	type metricsBlob struct {
		IndirectCount *int     `polygraph:"indirect.count"`
		NotSet        *float32 `polygraph:"not.set"`
		Nested        struct {
			Accessed int `polygraph:"nested.accessor"`
		}
	}
	m := metricsBlob{}
	m.IndirectCount = new(int)
	export := NewExporter()
	export.PublishStruct(&m)
	for i := 0; i < 10; i++ {
		(*m.IndirectCount) = i * 2 + 1
		m.Nested.Accessed = i
		if v := expvar.Get("indirect.count"); v == nil || v.String() != fmt.Sprint((i*2+1)) {
			t.Fatalf(`"indirect.count" was not set correctly, expected %v and got %d`, v, (i*2+1))
		}
		if v := expvar.Get("nested.accessor"); v == nil || v.String() != fmt.Sprint(i) {
			t.Fatal(`"nested.accessor" is not correctly set`)
		}
	}
	if v := expvar.Get("not.set"); v != nil {
		t.Fatal(`"not.set" has been set in expvar`)
	}
}

func TestConfiguration(t *testing.T) {
	const (
		conf = `---
host: ""
port: 8080
`
	)
	if err := NewExporter().Configure([]byte(conf)); err != nil {
		t.Fatal(err)
	}
}