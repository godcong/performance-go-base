package reflect

import (
	"reflect"
	"sync"
	"testing"
)

type T struct {
	A int
	B int
}

func BenchmarkCallDeep(b *testing.B) {
	var t T
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(reflect.Zero(v.Type()).Interface(), v.Interface())
	}
}

func BenchmarkCallZero(b *testing.B) {
	var t T
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isZero(v)
	}
}

func BenchmarkCallDeepArr(b *testing.B) {
	var t []T
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(reflect.Zero(v.Type()).Interface(), v.Interface())
	}
}

func BenchmarkCallDeepBig(b *testing.B) {
	var t []T = make([]T, 10000)
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(reflect.Zero(v.Type()).Interface(), v.Interface())
	}
}

func BenchmarkCallZeroArr(b *testing.B) {
	var t []T
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isZero(v)
	}
}

func BenchmarkCallZeroBig(b *testing.B) {
	var t []T = make([]T, 10000)
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isZero(v)
	}
}

var once sync.Once
var zeros [reflect.UnsafePointer + 1]func(v reflect.Value) bool

func BenchmarkCallPreZero(b *testing.B) {
	once.Do(func() {
		zeros = [reflect.UnsafePointer + 1]func(v reflect.Value) bool{
			reflect.Bool:    func(v reflect.Value) bool { return !v.Bool() },
			reflect.Int:     func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int8:    func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int16:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int32:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int64:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Float32: func(v reflect.Value) bool { return v.Float() == 0 },
			reflect.Float64: func(v reflect.Value) bool { return v.Float() == 0 },
			reflect.Uint:    func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint8:   func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint16:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint32:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint64:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.String:  func(v reflect.Value) bool { return v.String() == "" },
			reflect.Slice: func(v reflect.Value) bool {
				switch v.Type().Elem().Kind() {
				case reflect.Struct:
					return true
				default:
					return v.Len() == 0
				}
			},
			reflect.Struct: func(v reflect.Value) bool { return true },
			reflect.Map:    func(v reflect.Value) bool { return v.Len() == 0 },
		}
	})
	var t T
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeros[v.Kind()](v)
	}
}

func BenchmarkCallPreZeroArr(b *testing.B) {
	once.Do(func() {
		zeros = [reflect.UnsafePointer + 1]func(v reflect.Value) bool{
			reflect.Bool:    func(v reflect.Value) bool { return !v.Bool() },
			reflect.Int:     func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int8:    func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int16:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int32:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int64:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Float32: func(v reflect.Value) bool { return v.Float() == 0 },
			reflect.Float64: func(v reflect.Value) bool { return v.Float() == 0 },
			reflect.Uint:    func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint8:   func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint16:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint32:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint64:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.String:  func(v reflect.Value) bool { return v.String() == "" },
			reflect.Slice: func(v reflect.Value) bool {
				switch v.Type().Elem().Kind() {
				case reflect.Struct:
					return true
				default:
					return v.Len() == 0
				}
			},
			reflect.Struct: func(v reflect.Value) bool { return true },
			reflect.Map:    func(v reflect.Value) bool { return v.Len() == 0 },
		}
	})
	var t []T
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeros[v.Kind()](v)
	}
}

func BenchmarkCallPreZeroBig(b *testing.B) {
	once.Do(func() {
		zeros = [reflect.UnsafePointer + 1]func(v reflect.Value) bool{
			reflect.Bool:    func(v reflect.Value) bool { return !v.Bool() },
			reflect.Int:     func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int8:    func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int16:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int32:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Int64:   func(v reflect.Value) bool { return v.Int() == 0 },
			reflect.Float32: func(v reflect.Value) bool { return v.Float() == 0 },
			reflect.Float64: func(v reflect.Value) bool { return v.Float() == 0 },
			reflect.Uint:    func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint8:   func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint16:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint32:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.Uint64:  func(v reflect.Value) bool { return v.Uint() == 0 },
			reflect.String:  func(v reflect.Value) bool { return v.String() == "" },
			reflect.Slice: func(v reflect.Value) bool {
				switch v.Type().Elem().Kind() {
				case reflect.Struct:
					return true
				default:
					return v.Len() == 0
				}
			},
			reflect.Struct: func(v reflect.Value) bool { return true },
			reflect.Map:    func(v reflect.Value) bool { return v.Len() == 0 },
		}
	})
	var t = make([]T, 10000)
	v := reflect.ValueOf(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zeros[v.Kind()](v)
	}
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return !(v.Float() != 0)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Slice:
		switch v.Type().Elem().Kind() {
		case reflect.Struct:
			// always assume the structs in the slice is empty and can be filled
			// the actually struct filling logic should take care of the rest
			return true
		default:
			return v.Len() == 0
		}
	case reflect.String:
		return v.String() == ""
	}
	return true
}
