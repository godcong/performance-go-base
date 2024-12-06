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

var vali, err = protovalidate.New(protovalidate.WithFailFast(true))

func init() {
	if err != nil {
		panic(err)
	}
}

func BenchmarkProtoValidateGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := &User1{
			Name: "John Doe",
			Age:  30,
		}
		if err := vali.Validate(user); err != nil {
			b.Errorf("validation failed: %v", err)
		}
	}
}

func BenchmarkProtoValidateAll(b *testing.B) {
	b.Run("ProtoValidateGo", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				user := &User1{
					Name: "John Doe",
					Age:  30,
				}
				if err := vali.Validate(user); err != nil {
					b.Errorf("validation failed: %v", err)
				}
			}
		})
	})
	b.Run("ProtocGenValidate", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				user := &User2{
					Name: "John Doe",
					Age:  30,
				}
				if err := user.Validate(); err != nil {
					b.Errorf("validation failed: %v", err)
				}
			}
		})
	})
}
