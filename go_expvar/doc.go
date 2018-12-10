// Package go_expvar allows for metrics exporting using the expvar package
// This to take note of are:
//  - The labels applied to struct variables must be unique otherwise you
//    risk over writing a previously set variable
//    - Stick to a naming convention like `polygraph:"package.class.(innerStructName)*.variable"`
//  - Any value that is nil once passed through the Publish function will be ignored
//  - Once something is added to Publish method, it can not be removed.
package go_expvar
