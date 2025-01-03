// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package jsonmaporstruct

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct(in *jlexer.Lexer, out *EasyObject3) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "xvlbzgbaic":
			out.Xvlbzgbaic = string(in.String())
		case "krbemfdzdc":
			out.Krbemfdzdc = string(in.String())
		case "rzlntxyeuc":
			out.Rzlntxyeuc = string(in.String())
		case "ctzkjkziva":
			out.Ctzkjkziva = string(in.String())
		case "orsufumaps":
			out.Orsufumaps = string(in.String())
		case "hyevwbtcml":
			out.Hyevwbtcml = string(in.String())
		case "baatlyhdao":
			out.Baatlyhdao = string(in.String())
		case "fkfohsvvxs":
			out.Fkfohsvvxs = string(in.String())
		case "pqwarpxptp":
			out.Pqwarpxptp = string(in.String())
		case "orvaukawww":
			out.Orvaukawww = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct(out *jwriter.Writer, in EasyObject3) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"xvlbzgbaic\":"
		out.RawString(prefix[1:])
		out.String(string(in.Xvlbzgbaic))
	}
	{
		const prefix string = ",\"krbemfdzdc\":"
		out.RawString(prefix)
		out.String(string(in.Krbemfdzdc))
	}
	{
		const prefix string = ",\"rzlntxyeuc\":"
		out.RawString(prefix)
		out.String(string(in.Rzlntxyeuc))
	}
	{
		const prefix string = ",\"ctzkjkziva\":"
		out.RawString(prefix)
		out.String(string(in.Ctzkjkziva))
	}
	{
		const prefix string = ",\"orsufumaps\":"
		out.RawString(prefix)
		out.String(string(in.Orsufumaps))
	}
	{
		const prefix string = ",\"hyevwbtcml\":"
		out.RawString(prefix)
		out.String(string(in.Hyevwbtcml))
	}
	{
		const prefix string = ",\"baatlyhdao\":"
		out.RawString(prefix)
		out.String(string(in.Baatlyhdao))
	}
	{
		const prefix string = ",\"fkfohsvvxs\":"
		out.RawString(prefix)
		out.String(string(in.Fkfohsvvxs))
	}
	{
		const prefix string = ",\"pqwarpxptp\":"
		out.RawString(prefix)
		out.String(string(in.Pqwarpxptp))
	}
	{
		const prefix string = ",\"orvaukawww\":"
		out.RawString(prefix)
		out.String(string(in.Orvaukawww))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EasyObject3) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EasyObject3) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EasyObject3) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EasyObject3) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct(l, v)
}
func easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct1(in *jlexer.Lexer, out *EasyObject2) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "xvlbzgbaic":
			out.Xvlbzgbaic = string(in.String())
		case "krbemfdzdc":
			out.Krbemfdzdc = string(in.String())
		case "rzlntxyeuc":
			out.Rzlntxyeuc = string(in.String())
		case "ctzkjkziva":
			out.Ctzkjkziva = string(in.String())
		case "orsufumaps":
			out.Orsufumaps = string(in.String())
		case "hyevwbtcml":
			out.Hyevwbtcml = string(in.String())
		case "baatlyhdao":
			out.Baatlyhdao = string(in.String())
		case "fkfohsvvxs":
			out.Fkfohsvvxs = string(in.String())
		case "pqwarpxptp":
			out.Pqwarpxptp = string(in.String())
		case "orvaukawww":
			out.Orvaukawww = string(in.String())
		case "EasyObject3":
			(out.EasyObject3).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct1(out *jwriter.Writer, in EasyObject2) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"xvlbzgbaic\":"
		out.RawString(prefix[1:])
		out.String(string(in.Xvlbzgbaic))
	}
	{
		const prefix string = ",\"krbemfdzdc\":"
		out.RawString(prefix)
		out.String(string(in.Krbemfdzdc))
	}
	{
		const prefix string = ",\"rzlntxyeuc\":"
		out.RawString(prefix)
		out.String(string(in.Rzlntxyeuc))
	}
	{
		const prefix string = ",\"ctzkjkziva\":"
		out.RawString(prefix)
		out.String(string(in.Ctzkjkziva))
	}
	{
		const prefix string = ",\"orsufumaps\":"
		out.RawString(prefix)
		out.String(string(in.Orsufumaps))
	}
	{
		const prefix string = ",\"hyevwbtcml\":"
		out.RawString(prefix)
		out.String(string(in.Hyevwbtcml))
	}
	{
		const prefix string = ",\"baatlyhdao\":"
		out.RawString(prefix)
		out.String(string(in.Baatlyhdao))
	}
	{
		const prefix string = ",\"fkfohsvvxs\":"
		out.RawString(prefix)
		out.String(string(in.Fkfohsvvxs))
	}
	{
		const prefix string = ",\"pqwarpxptp\":"
		out.RawString(prefix)
		out.String(string(in.Pqwarpxptp))
	}
	{
		const prefix string = ",\"orvaukawww\":"
		out.RawString(prefix)
		out.String(string(in.Orvaukawww))
	}
	{
		const prefix string = ",\"EasyObject3\":"
		out.RawString(prefix)
		(in.EasyObject3).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EasyObject2) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EasyObject2) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EasyObject2) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EasyObject2) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct1(l, v)
}
func easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct2(in *jlexer.Lexer, out *EasyObject) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "xvlbzgbaic":
			out.Xvlbzgbaic = string(in.String())
		case "krbemfdzdc":
			out.Krbemfdzdc = string(in.String())
		case "rzlntxyeuc":
			out.Rzlntxyeuc = string(in.String())
		case "ctzkjkziva":
			out.Ctzkjkziva = string(in.String())
		case "orsufumaps":
			out.Orsufumaps = string(in.String())
		case "hyevwbtcml":
			out.Hyevwbtcml = string(in.String())
		case "baatlyhdao":
			out.Baatlyhdao = string(in.String())
		case "fkfohsvvxs":
			out.Fkfohsvvxs = string(in.String())
		case "pqwarpxptp":
			out.Pqwarpxptp = string(in.String())
		case "orvaukawww":
			out.Orvaukawww = string(in.String())
		case "EasyObject2":
			(out.EasyObject2).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct2(out *jwriter.Writer, in EasyObject) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"xvlbzgbaic\":"
		out.RawString(prefix[1:])
		out.String(string(in.Xvlbzgbaic))
	}
	{
		const prefix string = ",\"krbemfdzdc\":"
		out.RawString(prefix)
		out.String(string(in.Krbemfdzdc))
	}
	{
		const prefix string = ",\"rzlntxyeuc\":"
		out.RawString(prefix)
		out.String(string(in.Rzlntxyeuc))
	}
	{
		const prefix string = ",\"ctzkjkziva\":"
		out.RawString(prefix)
		out.String(string(in.Ctzkjkziva))
	}
	{
		const prefix string = ",\"orsufumaps\":"
		out.RawString(prefix)
		out.String(string(in.Orsufumaps))
	}
	{
		const prefix string = ",\"hyevwbtcml\":"
		out.RawString(prefix)
		out.String(string(in.Hyevwbtcml))
	}
	{
		const prefix string = ",\"baatlyhdao\":"
		out.RawString(prefix)
		out.String(string(in.Baatlyhdao))
	}
	{
		const prefix string = ",\"fkfohsvvxs\":"
		out.RawString(prefix)
		out.String(string(in.Fkfohsvvxs))
	}
	{
		const prefix string = ",\"pqwarpxptp\":"
		out.RawString(prefix)
		out.String(string(in.Pqwarpxptp))
	}
	{
		const prefix string = ",\"orvaukawww\":"
		out.RawString(prefix)
		out.String(string(in.Orvaukawww))
	}
	{
		const prefix string = ",\"EasyObject2\":"
		out.RawString(prefix)
		(in.EasyObject2).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EasyObject) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EasyObject) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7c533daaEncodePerformanceGoBaseJsonmaporstruct2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EasyObject) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EasyObject) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7c533daaDecodePerformanceGoBaseJsonmaporstruct2(l, v)
}
