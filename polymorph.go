package polygraph

// Morph allows you to take your implemented non complex type
// and be able to export to your target metric platform
type Morph interface {

	// Set expects interface to be a pointer to the non complex variable
	// so that it can export it to the target provider.
	Set(obj interface{}, identity string)

	// String formats the object in a format for the target exported platform.
	String() string
}
