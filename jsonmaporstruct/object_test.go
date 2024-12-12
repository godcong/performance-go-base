package jsonmaporstruct

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
	_ "github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/tinylib/msgp/msgp"
	"github.com/valyala/fastjson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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
	for i := 0; i < b.N; i++ {
		var obj Object
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalToObjectMarshal 解码到结构体并编码
func BenchmarkUnmarshalToObjectMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := json.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalToAny 解码到 any
func BenchmarkUnmarshalToAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj any
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalToAnyMarshal 解码到 any 并编码
func BenchmarkUnmarshalToAnyMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj any
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := json.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithJsoniter 解码到结构体使用 jsoniter
func BenchmarkUnmarshalWithJsoniter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		if err := jsoniter.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithJsoniterMarshal 解码到结构体使用 jsoniter 并编码
func BenchmarkUnmarshalWithJsoniterMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		if err := jsoniter.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := jsoniter.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithSonic 解码到结构体使用 sonic
func BenchmarkUnmarshalWithSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		if err := sonic.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithSonicMarshal 解码到结构体使用 sonic 并编码
func BenchmarkUnmarshalWithSonicMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		if err := sonic.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := sonic.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithFastjson 解码到结构体使用 fastjson
func BenchmarkUnmarshalWithFastjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		if _, err := p.ParseBytes(jsonData); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithFastjsonMarshal 解码到结构体使用 fastjson 并编码
func BenchmarkUnmarshalWithFastjsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		v, err := p.ParseBytes(jsonData)
		if err != nil {
			b.Fatal(err)
		}
		_ = v.MarshalTo(nil)

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

// BenchmarkUnmarshalWithProtojsonMarshal 解码到结构体使用 protojson 并编码
func BenchmarkUnmarshalWithProtojsonMarshal(b *testing.B) {
	var obj ProtoObject
	for i := 0; i < b.N; i++ {
		if err := protojson.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := protojson.Marshal(&obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var protoData []byte
var bufMsgpb bytes.Buffer
var bufGob bytes.Buffer
var lexer jlexer.Lexer

func init() {
	var obj Object
	var pobj ProtoObject
	_ = json.Unmarshal(jsonData, &obj)
	_ = protojson.Unmarshal(jsonData, &pobj)
	protoData, _ = proto.Marshal(&pobj)
	_ = msgp.Encode(&bufMsgpb, &obj)
	encoder := gob.NewEncoder(&bufGob)
	_ = encoder.Encode(obj)
	lexer = jlexer.Lexer{Data: jsonData}
}

// BenchmarkUnmarshalWithEasyjson 解码到结构体使用 easyjson
func BenchmarkUnmarshalWithEasyjson(b *testing.B) {
	var tmp *jlexer.Lexer
	for i := 0; i < b.N; i++ {
		var obj EasyObject
		tmp = &jlexer.Lexer{Data: jsonData}
		obj.UnmarshalEasyJSON(tmp)
		if tmp.Error() != nil {
			b.Fatal(tmp.Error())
		}
	}
}

// BenchmarkUnmarshalWithEasyjsonMarshal 解码到结构体使用 easyjson 并编码
func BenchmarkUnmarshalWithEasyjsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj EasyObject
		obj.UnmarshalEasyJSON(&lexer)
		_, err := json.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalWithEasyjsonSTD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj EasyObject
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithEasyjsonSTDMarshal 解码到结构体使用 easyjson 并编码
func BenchmarkUnmarshalWithEasyjsonSTDMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj EasyObject
		if err := json.Unmarshal(jsonData, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := json.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithEasyjsonOneInstance 解码到结构体使用 easyjson
func BenchmarkUnmarshalWithEasyjsonOneInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj EasyObject
		obj.UnmarshalEasyJSON(&lexer)
	}
}

// BenchmarkUnmarshalWithEasyjsonOneInstanceMarshal 解码到结构体使用 easyjson 并编码
func BenchmarkUnmarshalWithEasyjsonOneInstanceMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj EasyObject
		obj.UnmarshalEasyJSON(&lexer)
		tmpW := jwriter.Writer{}
		obj.MarshalEasyJSON(&tmpW)
		if tmpW.Error != nil {
			b.Fatal(tmpW.Error)
		}
	}
}

// BenchmarkUnmarshalWithProto 解码到结构体使用 proto
func BenchmarkUnmarshalWithProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj ProtoObject
		if err := proto.Unmarshal(protoData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithProtoMarshal 解码到结构体使用 proto 并编码
func BenchmarkUnmarshalWithProtoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj ProtoObject
		if err := proto.Unmarshal(protoData, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := proto.Marshal(&obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithMsgpb 解码到结构体使用 proto
func BenchmarkUnmarshalWithMsgpb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		buf := bytes.NewBuffer(bufMsgpb.Bytes())
		if err := msgp.Decode(buf, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithMsgpbMarshal 解码到结构体使用 proto 并编码
func BenchmarkUnmarshalWithMsgpbMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		buf := bytes.NewBuffer(bufMsgpb.Bytes())
		if err := msgp.Decode(buf, &obj); err != nil {
			b.Fatal(err)
		}
		_, err := json.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithGob 解码到结构体使用 gob
func BenchmarkUnmarshalWithGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		buf := bytes.NewBuffer(bufGob.Bytes())
		decoder := gob.NewDecoder(buf)
		if err := decoder.Decode(&obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshalWithGobMarshal 解码到结构体使用 gob 并编码
func BenchmarkUnmarshalWithGobMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var obj Object
		buf := bytes.NewBuffer(bufGob.Bytes())
		decoder := gob.NewDecoder(buf)
		if err := decoder.Decode(&obj); err != nil {
			b.Fatal(err)
		}
		var bufW bytes.Buffer
		encoder := gob.NewEncoder(&bufW)
		err := encoder.Encode(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}
