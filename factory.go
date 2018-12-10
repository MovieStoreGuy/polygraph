package polygraph

// Exporter will enable users to export their metrics to their desired
// target platform for monitoring
type Exporter interface {

	// Configure reads the yaml data and will update the underlying exporter
	Configure(data []byte) error

	// Export will accept a reference to variable to covert into a morph
	// type and prepare it to be added to the underlying metrics provider
	// The use of metricName is optional since you can pass a struct with structTags applied
	//
	// Example of a basic variable:
	//
	//   version := 3
	//   exporter.CreateMorph(&version, "version.number")
	//
	// Example of a complex struct:
	//
	//   type A struct {
	//      count int `monitor:"a.count"`
	//   }
	//   a := A{}
	//   exporter.CreateMorph(&a) // metricName not need as it will be part of the struct tag
	Export(obj interface{}, metricName ...string)


	// Start will ensure that the underlying metrics exporter
	// is running based off the configuration set on the factory
	Start() error
}