// Package protovalidate implements the functions, types, and interfaces for the module.
package protovalidate

//go:generate protoc -I. --go_out=paths=source_relative:. *.proto
//go:generate protoc -I. --validate_out=paths=source_relative,lang=go:.  *.proto
