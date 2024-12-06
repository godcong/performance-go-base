// Package protovalidate implements the functions, types, and interfaces for the module.
package protovalidate

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	//"github.com/envoyproxy/protoc-gen-validate/validate"
)

func BenchmarkProtocGenValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := &User2{
			Name: "John Doe",
			Age:  30,
		}
		if err := user.Validate(); err != nil {
			b.Errorf("validation failed: %v", err)
		}
	}
}

func BenchmarkProtoValidateGo(b *testing.B) {
	v, err := protovalidate.New()
	if err != nil {
		b.Fatalf("validation failed: %v", err)
	}
	for i := 0; i < b.N; i++ {
		user := &User1{
			Name: "John Doe",
			Age:  30,
		}
		if err := v.Validate(user); err != nil {
			b.Errorf("validation failed: %v", err)
		}
	}
}
