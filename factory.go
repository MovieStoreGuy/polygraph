package polygraph

// Exporter will enable users to export their metrics to their desired
// target platform for monitoring
type Exporter interface {

	// Configure reads the yaml data and will update the underlying exporter
	Configure(data []byte) error

	// Export will accept a reference to variable to covert into a morph
	// type and prepare it to be added to the underlying metrics provider
	// Example:
	//
	//   type A struct {
	//      count int `monitor:"a.count"`
	//   }
	//   a := A{}
	//   exporter.PublishStruct(&a)
	PublishStruct(obj interface{})

	// PublishVariable will ensure that the target variable will be exported
	// to the correct format for the underlying metrics provider
	// Example:
	//
	//   hitCount := 0
	//   exporter.PublishVariable(&hitCount, "server.ingress.hitCount")
	//   // ...
	PublishVariable(obj interface{}, label string)


	// Start will ensure that the underlying metrics exporter
	// is running based off the configuration set on the factory
	Start() error
}