package jsonmaporstruct

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/valyala/fastjson"
	"google.golang.org/protobuf/encoding/protojson"
)

// 准备 JSON 数据
var jsonData = []byte(`{
	"xvlbzgbaic": "value1",
	"krbemfdzdc": "value2",
	"rzlntxyeuc": "value3",
	"ctzkjkziva": "value4",
	"orsufumaps": "value5",
	"hyevwbtcml": "value6",
	"baatlyhdao": "value7",
	"fkfohsvvxs": "value8",
	"pqwarpxptp": "value9",
	"orvaukawww": "value10",
	"object2": {
		"xvlbzgbaic": "value1",
		"krbemfdzdc": "value2",
		"rzlntxyeuc": "value3",
		"ctzkjkziva": "value4",
		"orsufumaps": "value5",
		"hyevwbtcml": "value6",
		"baatlyhdao": "value7",
		"fkfohsvvxs": "value8",
		"pqwarpxptp": "value9",
		"orvaukawww": "value10",
		"object3": {
			"xvlbzgbaic": "value1",
			"krbemfdzdc": "value2",
			"rzlntxyeuc": "value3",
			"ctzkjkziva": "value4",
			"orsufumaps": "value5",
			"hyevwbtcml": "value6",
			"baatlyhdao": "value7",
			"fkfohsvvxs": "value8",
			"pqwarpxptp": "value9",
			"orvaukawww": "value10"
		}
	}
}`)

// BenchmarkUnmarshalToObject 解码到结构体
func BenchmarkUnmarshalToObject(b *testing.B) {
	var obj Object
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalToAny 解码到 any
func BenchmarkUnmarshalToAny(b *testing.B) {
	var obj any
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithJsoniter 解码到结构体使用 jsoniter
func BenchmarkUnmarshalWithJsoniter(b *testing.B) {
	var obj Object
	for i := 0; i < b.N; i++ {
		if err := jsoniter.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithSonic 解码到结构体使用 sonic
func BenchmarkUnmarshalWithSonic(b *testing.B) {
	var obj Object
	for i := 0; i < b.N; i++ {
		if err := sonic.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithEasyjson 解码到结构体使用 easyjson
func BenchmarkUnmarshalWithEasyjson(b *testing.B) {
	var obj Object
	for i := 0; i < b.N; i++ {
		if err := easyjson.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithFastjson 解码到结构体使用 fastjson
func BenchmarkUnmarshalWithFastjson(b *testing.B) {
	var p fastjson.Parser
	for i := 0; i < b.N; i++ {
		if _, err := p.ParseBytes(jsonData); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithProtojson 解码到结构体使用 protojson
func BenchmarkUnmarshalWithProtojson(b *testing.B) {
	var obj ProtoObject
	for i := 0; i < b.N; i++ {
		if err := protojson.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}
