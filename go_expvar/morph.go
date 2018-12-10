package go_expvar

import "fmt"

type morph struct {
	variable interface{}
}

func (m *morph) Set(obj interface{}, metricName string) {
	_ = metricName
	m.variable = obj
}


func (m *morph) String() string {
	return fmt.Sprint(m.variable)
}