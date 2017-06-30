// Package anyplugin provides a standard set of build tags
// for anynet (https://github.com/unixpickle/anynet)
// plugins.
//
// To use anyplugin in your package, simply put this line
// somewhere in your code:
//
//     import _ "github.com/unixpickle/anyplugin"
//
// When you want to use a plugin, for example "cuda", you
// can download the plugin dependencies like so:
//
//     go get -u -tags cuda github.com/unixpickle/anyplugin
//
// Now, you can use the plugin in your project using the
// -tags flag, for example:
//
//     go run -tags cuda *.go
//
// The following tags are currently supported:
//
//     cuda: use cudavec (https://github.com/unixpickle/cudavec)
//
// In the future, tags may be added for cuDNN, OpenCL, and
// others.
package anyplugin
