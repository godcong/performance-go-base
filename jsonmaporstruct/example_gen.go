package jsonmaporstruct

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Object) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Xvlbzgbaic":
			z.Xvlbzgbaic, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Xvlbzgbaic")
				return
			}
		case "Krbemfdzdc":
			z.Krbemfdzdc, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Krbemfdzdc")
				return
			}
		case "Rzlntxyeuc":
			z.Rzlntxyeuc, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Rzlntxyeuc")
				return
			}
		case "Ctzkjkziva":
			z.Ctzkjkziva, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Ctzkjkziva")
				return
			}
		case "Orsufumaps":
			z.Orsufumaps, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Orsufumaps")
				return
			}
		case "Hyevwbtcml":
			z.Hyevwbtcml, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Hyevwbtcml")
				return
			}
		case "Baatlyhdao":
			z.Baatlyhdao, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Baatlyhdao")
				return
			}
		case "Fkfohsvvxs":
			z.Fkfohsvvxs, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Fkfohsvvxs")
				return
			}
		case "Pqwarpxptp":
			z.Pqwarpxptp, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Pqwarpxptp")
				return
			}
		case "Orvaukawww":
			z.Orvaukawww, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Orvaukawww")
				return
			}
		case "Object2":
			err = z.Object2.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Object2")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Object) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "Xvlbzgbaic"
	err = en.Append(0x8b, 0xaa, 0x58, 0x76, 0x6c, 0x62, 0x7a, 0x67, 0x62, 0x61, 0x69, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Xvlbzgbaic)
	if err != nil {
		err = msgp.WrapError(err, "Xvlbzgbaic")
		return
	}
	// write "Krbemfdzdc"
	err = en.Append(0xaa, 0x4b, 0x72, 0x62, 0x65, 0x6d, 0x66, 0x64, 0x7a, 0x64, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Krbemfdzdc)
	if err != nil {
		err = msgp.WrapError(err, "Krbemfdzdc")
		return
	}
	// write "Rzlntxyeuc"
	err = en.Append(0xaa, 0x52, 0x7a, 0x6c, 0x6e, 0x74, 0x78, 0x79, 0x65, 0x75, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Rzlntxyeuc)
	if err != nil {
		err = msgp.WrapError(err, "Rzlntxyeuc")
		return
	}
	// write "Ctzkjkziva"
	err = en.Append(0xaa, 0x43, 0x74, 0x7a, 0x6b, 0x6a, 0x6b, 0x7a, 0x69, 0x76, 0x61)
	if err != nil {
		return
	}
	err = en.WriteString(z.Ctzkjkziva)
	if err != nil {
		err = msgp.WrapError(err, "Ctzkjkziva")
		return
	}
	// write "Orsufumaps"
	err = en.Append(0xaa, 0x4f, 0x72, 0x73, 0x75, 0x66, 0x75, 0x6d, 0x61, 0x70, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Orsufumaps)
	if err != nil {
		err = msgp.WrapError(err, "Orsufumaps")
		return
	}
	// write "Hyevwbtcml"
	err = en.Append(0xaa, 0x48, 0x79, 0x65, 0x76, 0x77, 0x62, 0x74, 0x63, 0x6d, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Hyevwbtcml)
	if err != nil {
		err = msgp.WrapError(err, "Hyevwbtcml")
		return
	}
	// write "Baatlyhdao"
	err = en.Append(0xaa, 0x42, 0x61, 0x61, 0x74, 0x6c, 0x79, 0x68, 0x64, 0x61, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteString(z.Baatlyhdao)
	if err != nil {
		err = msgp.WrapError(err, "Baatlyhdao")
		return
	}
	// write "Fkfohsvvxs"
	err = en.Append(0xaa, 0x46, 0x6b, 0x66, 0x6f, 0x68, 0x73, 0x76, 0x76, 0x78, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Fkfohsvvxs)
	if err != nil {
		err = msgp.WrapError(err, "Fkfohsvvxs")
		return
	}
	// write "Pqwarpxptp"
	err = en.Append(0xaa, 0x50, 0x71, 0x77, 0x61, 0x72, 0x70, 0x78, 0x70, 0x74, 0x70)
	if err != nil {
		return
	}
	err = en.WriteString(z.Pqwarpxptp)
	if err != nil {
		err = msgp.WrapError(err, "Pqwarpxptp")
		return
	}
	// write "Orvaukawww"
	err = en.Append(0xaa, 0x4f, 0x72, 0x76, 0x61, 0x75, 0x6b, 0x61, 0x77, 0x77, 0x77)
	if err != nil {
		return
	}
	err = en.WriteString(z.Orvaukawww)
	if err != nil {
		err = msgp.WrapError(err, "Orvaukawww")
		return
	}
	// write "Object2"
	err = en.Append(0xa7, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x32)
	if err != nil {
		return
	}
	err = z.Object2.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Object2")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Object) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "Xvlbzgbaic"
	o = append(o, 0x8b, 0xaa, 0x58, 0x76, 0x6c, 0x62, 0x7a, 0x67, 0x62, 0x61, 0x69, 0x63)
	o = msgp.AppendString(o, z.Xvlbzgbaic)
	// string "Krbemfdzdc"
	o = append(o, 0xaa, 0x4b, 0x72, 0x62, 0x65, 0x6d, 0x66, 0x64, 0x7a, 0x64, 0x63)
	o = msgp.AppendString(o, z.Krbemfdzdc)
	// string "Rzlntxyeuc"
	o = append(o, 0xaa, 0x52, 0x7a, 0x6c, 0x6e, 0x74, 0x78, 0x79, 0x65, 0x75, 0x63)
	o = msgp.AppendString(o, z.Rzlntxyeuc)
	// string "Ctzkjkziva"
	o = append(o, 0xaa, 0x43, 0x74, 0x7a, 0x6b, 0x6a, 0x6b, 0x7a, 0x69, 0x76, 0x61)
	o = msgp.AppendString(o, z.Ctzkjkziva)
	// string "Orsufumaps"
	o = append(o, 0xaa, 0x4f, 0x72, 0x73, 0x75, 0x66, 0x75, 0x6d, 0x61, 0x70, 0x73)
	o = msgp.AppendString(o, z.Orsufumaps)
	// string "Hyevwbtcml"
	o = append(o, 0xaa, 0x48, 0x79, 0x65, 0x76, 0x77, 0x62, 0x74, 0x63, 0x6d, 0x6c)
	o = msgp.AppendString(o, z.Hyevwbtcml)
	// string "Baatlyhdao"
	o = append(o, 0xaa, 0x42, 0x61, 0x61, 0x74, 0x6c, 0x79, 0x68, 0x64, 0x61, 0x6f)
	o = msgp.AppendString(o, z.Baatlyhdao)
	// string "Fkfohsvvxs"
	o = append(o, 0xaa, 0x46, 0x6b, 0x66, 0x6f, 0x68, 0x73, 0x76, 0x76, 0x78, 0x73)
	o = msgp.AppendString(o, z.Fkfohsvvxs)
	// string "Pqwarpxptp"
	o = append(o, 0xaa, 0x50, 0x71, 0x77, 0x61, 0x72, 0x70, 0x78, 0x70, 0x74, 0x70)
	o = msgp.AppendString(o, z.Pqwarpxptp)
	// string "Orvaukawww"
	o = append(o, 0xaa, 0x4f, 0x72, 0x76, 0x61, 0x75, 0x6b, 0x61, 0x77, 0x77, 0x77)
	o = msgp.AppendString(o, z.Orvaukawww)
	// string "Object2"
	o = append(o, 0xa7, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x32)
	o, err = z.Object2.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Object2")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Object) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Xvlbzgbaic":
			z.Xvlbzgbaic, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Xvlbzgbaic")
				return
			}
		case "Krbemfdzdc":
			z.Krbemfdzdc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Krbemfdzdc")
				return
			}
		case "Rzlntxyeuc":
			z.Rzlntxyeuc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Rzlntxyeuc")
				return
			}
		case "Ctzkjkziva":
			z.Ctzkjkziva, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ctzkjkziva")
				return
			}
		case "Orsufumaps":
			z.Orsufumaps, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Orsufumaps")
				return
			}
		case "Hyevwbtcml":
			z.Hyevwbtcml, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hyevwbtcml")
				return
			}
		case "Baatlyhdao":
			z.Baatlyhdao, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Baatlyhdao")
				return
			}
		case "Fkfohsvvxs":
			z.Fkfohsvvxs, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Fkfohsvvxs")
				return
			}
		case "Pqwarpxptp":
			z.Pqwarpxptp, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Pqwarpxptp")
				return
			}
		case "Orvaukawww":
			z.Orvaukawww, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Orvaukawww")
				return
			}
		case "Object2":
			bts, err = z.Object2.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Object2")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Object) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.Xvlbzgbaic) + 11 + msgp.StringPrefixSize + len(z.Krbemfdzdc) + 11 + msgp.StringPrefixSize + len(z.Rzlntxyeuc) + 11 + msgp.StringPrefixSize + len(z.Ctzkjkziva) + 11 + msgp.StringPrefixSize + len(z.Orsufumaps) + 11 + msgp.StringPrefixSize + len(z.Hyevwbtcml) + 11 + msgp.StringPrefixSize + len(z.Baatlyhdao) + 11 + msgp.StringPrefixSize + len(z.Fkfohsvvxs) + 11 + msgp.StringPrefixSize + len(z.Pqwarpxptp) + 11 + msgp.StringPrefixSize + len(z.Orvaukawww) + 8 + z.Object2.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Object2) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Xvlbzgbaic":
			z.Xvlbzgbaic, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Xvlbzgbaic")
				return
			}
		case "Krbemfdzdc":
			z.Krbemfdzdc, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Krbemfdzdc")
				return
			}
		case "Rzlntxyeuc":
			z.Rzlntxyeuc, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Rzlntxyeuc")
				return
			}
		case "Ctzkjkziva":
			z.Ctzkjkziva, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Ctzkjkziva")
				return
			}
		case "Orsufumaps":
			z.Orsufumaps, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Orsufumaps")
				return
			}
		case "Hyevwbtcml":
			z.Hyevwbtcml, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Hyevwbtcml")
				return
			}
		case "Baatlyhdao":
			z.Baatlyhdao, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Baatlyhdao")
				return
			}
		case "Fkfohsvvxs":
			z.Fkfohsvvxs, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Fkfohsvvxs")
				return
			}
		case "Pqwarpxptp":
			z.Pqwarpxptp, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Pqwarpxptp")
				return
			}
		case "Orvaukawww":
			z.Orvaukawww, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Orvaukawww")
				return
			}
		case "Object3":
			err = z.Object3.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Object3")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Object2) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "Xvlbzgbaic"
	err = en.Append(0x8b, 0xaa, 0x58, 0x76, 0x6c, 0x62, 0x7a, 0x67, 0x62, 0x61, 0x69, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Xvlbzgbaic)
	if err != nil {
		err = msgp.WrapError(err, "Xvlbzgbaic")
		return
	}
	// write "Krbemfdzdc"
	err = en.Append(0xaa, 0x4b, 0x72, 0x62, 0x65, 0x6d, 0x66, 0x64, 0x7a, 0x64, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Krbemfdzdc)
	if err != nil {
		err = msgp.WrapError(err, "Krbemfdzdc")
		return
	}
	// write "Rzlntxyeuc"
	err = en.Append(0xaa, 0x52, 0x7a, 0x6c, 0x6e, 0x74, 0x78, 0x79, 0x65, 0x75, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Rzlntxyeuc)
	if err != nil {
		err = msgp.WrapError(err, "Rzlntxyeuc")
		return
	}
	// write "Ctzkjkziva"
	err = en.Append(0xaa, 0x43, 0x74, 0x7a, 0x6b, 0x6a, 0x6b, 0x7a, 0x69, 0x76, 0x61)
	if err != nil {
		return
	}
	err = en.WriteString(z.Ctzkjkziva)
	if err != nil {
		err = msgp.WrapError(err, "Ctzkjkziva")
		return
	}
	// write "Orsufumaps"
	err = en.Append(0xaa, 0x4f, 0x72, 0x73, 0x75, 0x66, 0x75, 0x6d, 0x61, 0x70, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Orsufumaps)
	if err != nil {
		err = msgp.WrapError(err, "Orsufumaps")
		return
	}
	// write "Hyevwbtcml"
	err = en.Append(0xaa, 0x48, 0x79, 0x65, 0x76, 0x77, 0x62, 0x74, 0x63, 0x6d, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Hyevwbtcml)
	if err != nil {
		err = msgp.WrapError(err, "Hyevwbtcml")
		return
	}
	// write "Baatlyhdao"
	err = en.Append(0xaa, 0x42, 0x61, 0x61, 0x74, 0x6c, 0x79, 0x68, 0x64, 0x61, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteString(z.Baatlyhdao)
	if err != nil {
		err = msgp.WrapError(err, "Baatlyhdao")
		return
	}
	// write "Fkfohsvvxs"
	err = en.Append(0xaa, 0x46, 0x6b, 0x66, 0x6f, 0x68, 0x73, 0x76, 0x76, 0x78, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Fkfohsvvxs)
	if err != nil {
		err = msgp.WrapError(err, "Fkfohsvvxs")
		return
	}
	// write "Pqwarpxptp"
	err = en.Append(0xaa, 0x50, 0x71, 0x77, 0x61, 0x72, 0x70, 0x78, 0x70, 0x74, 0x70)
	if err != nil {
		return
	}
	err = en.WriteString(z.Pqwarpxptp)
	if err != nil {
		err = msgp.WrapError(err, "Pqwarpxptp")
		return
	}
	// write "Orvaukawww"
	err = en.Append(0xaa, 0x4f, 0x72, 0x76, 0x61, 0x75, 0x6b, 0x61, 0x77, 0x77, 0x77)
	if err != nil {
		return
	}
	err = en.WriteString(z.Orvaukawww)
	if err != nil {
		err = msgp.WrapError(err, "Orvaukawww")
		return
	}
	// write "Object3"
	err = en.Append(0xa7, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x33)
	if err != nil {
		return
	}
	err = z.Object3.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Object3")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Object2) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "Xvlbzgbaic"
	o = append(o, 0x8b, 0xaa, 0x58, 0x76, 0x6c, 0x62, 0x7a, 0x67, 0x62, 0x61, 0x69, 0x63)
	o = msgp.AppendString(o, z.Xvlbzgbaic)
	// string "Krbemfdzdc"
	o = append(o, 0xaa, 0x4b, 0x72, 0x62, 0x65, 0x6d, 0x66, 0x64, 0x7a, 0x64, 0x63)
	o = msgp.AppendString(o, z.Krbemfdzdc)
	// string "Rzlntxyeuc"
	o = append(o, 0xaa, 0x52, 0x7a, 0x6c, 0x6e, 0x74, 0x78, 0x79, 0x65, 0x75, 0x63)
	o = msgp.AppendString(o, z.Rzlntxyeuc)
	// string "Ctzkjkziva"
	o = append(o, 0xaa, 0x43, 0x74, 0x7a, 0x6b, 0x6a, 0x6b, 0x7a, 0x69, 0x76, 0x61)
	o = msgp.AppendString(o, z.Ctzkjkziva)
	// string "Orsufumaps"
	o = append(o, 0xaa, 0x4f, 0x72, 0x73, 0x75, 0x66, 0x75, 0x6d, 0x61, 0x70, 0x73)
	o = msgp.AppendString(o, z.Orsufumaps)
	// string "Hyevwbtcml"
	o = append(o, 0xaa, 0x48, 0x79, 0x65, 0x76, 0x77, 0x62, 0x74, 0x63, 0x6d, 0x6c)
	o = msgp.AppendString(o, z.Hyevwbtcml)
	// string "Baatlyhdao"
	o = append(o, 0xaa, 0x42, 0x61, 0x61, 0x74, 0x6c, 0x79, 0x68, 0x64, 0x61, 0x6f)
	o = msgp.AppendString(o, z.Baatlyhdao)
	// string "Fkfohsvvxs"
	o = append(o, 0xaa, 0x46, 0x6b, 0x66, 0x6f, 0x68, 0x73, 0x76, 0x76, 0x78, 0x73)
	o = msgp.AppendString(o, z.Fkfohsvvxs)
	// string "Pqwarpxptp"
	o = append(o, 0xaa, 0x50, 0x71, 0x77, 0x61, 0x72, 0x70, 0x78, 0x70, 0x74, 0x70)
	o = msgp.AppendString(o, z.Pqwarpxptp)
	// string "Orvaukawww"
	o = append(o, 0xaa, 0x4f, 0x72, 0x76, 0x61, 0x75, 0x6b, 0x61, 0x77, 0x77, 0x77)
	o = msgp.AppendString(o, z.Orvaukawww)
	// string "Object3"
	o = append(o, 0xa7, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x33)
	o, err = z.Object3.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Object3")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Object2) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Xvlbzgbaic":
			z.Xvlbzgbaic, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Xvlbzgbaic")
				return
			}
		case "Krbemfdzdc":
			z.Krbemfdzdc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Krbemfdzdc")
				return
			}
		case "Rzlntxyeuc":
			z.Rzlntxyeuc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Rzlntxyeuc")
				return
			}
		case "Ctzkjkziva":
			z.Ctzkjkziva, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ctzkjkziva")
				return
			}
		case "Orsufumaps":
			z.Orsufumaps, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Orsufumaps")
				return
			}
		case "Hyevwbtcml":
			z.Hyevwbtcml, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hyevwbtcml")
				return
			}
		case "Baatlyhdao":
			z.Baatlyhdao, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Baatlyhdao")
				return
			}
		case "Fkfohsvvxs":
			z.Fkfohsvvxs, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Fkfohsvvxs")
				return
			}
		case "Pqwarpxptp":
			z.Pqwarpxptp, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Pqwarpxptp")
				return
			}
		case "Orvaukawww":
			z.Orvaukawww, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Orvaukawww")
				return
			}
		case "Object3":
			bts, err = z.Object3.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Object3")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Object2) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.Xvlbzgbaic) + 11 + msgp.StringPrefixSize + len(z.Krbemfdzdc) + 11 + msgp.StringPrefixSize + len(z.Rzlntxyeuc) + 11 + msgp.StringPrefixSize + len(z.Ctzkjkziva) + 11 + msgp.StringPrefixSize + len(z.Orsufumaps) + 11 + msgp.StringPrefixSize + len(z.Hyevwbtcml) + 11 + msgp.StringPrefixSize + len(z.Baatlyhdao) + 11 + msgp.StringPrefixSize + len(z.Fkfohsvvxs) + 11 + msgp.StringPrefixSize + len(z.Pqwarpxptp) + 11 + msgp.StringPrefixSize + len(z.Orvaukawww) + 8 + z.Object3.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Object3) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Xvlbzgbaic":
			z.Xvlbzgbaic, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Xvlbzgbaic")
				return
			}
		case "Krbemfdzdc":
			z.Krbemfdzdc, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Krbemfdzdc")
				return
			}
		case "Rzlntxyeuc":
			z.Rzlntxyeuc, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Rzlntxyeuc")
				return
			}
		case "Ctzkjkziva":
			z.Ctzkjkziva, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Ctzkjkziva")
				return
			}
		case "Orsufumaps":
			z.Orsufumaps, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Orsufumaps")
				return
			}
		case "Hyevwbtcml":
			z.Hyevwbtcml, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Hyevwbtcml")
				return
			}
		case "Baatlyhdao":
			z.Baatlyhdao, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Baatlyhdao")
				return
			}
		case "Fkfohsvvxs":
			z.Fkfohsvvxs, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Fkfohsvvxs")
				return
			}
		case "Pqwarpxptp":
			z.Pqwarpxptp, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Pqwarpxptp")
				return
			}
		case "Orvaukawww":
			z.Orvaukawww, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Orvaukawww")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Object3) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "Xvlbzgbaic"
	err = en.Append(0x8a, 0xaa, 0x58, 0x76, 0x6c, 0x62, 0x7a, 0x67, 0x62, 0x61, 0x69, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Xvlbzgbaic)
	if err != nil {
		err = msgp.WrapError(err, "Xvlbzgbaic")
		return
	}
	// write "Krbemfdzdc"
	err = en.Append(0xaa, 0x4b, 0x72, 0x62, 0x65, 0x6d, 0x66, 0x64, 0x7a, 0x64, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Krbemfdzdc)
	if err != nil {
		err = msgp.WrapError(err, "Krbemfdzdc")
		return
	}
	// write "Rzlntxyeuc"
	err = en.Append(0xaa, 0x52, 0x7a, 0x6c, 0x6e, 0x74, 0x78, 0x79, 0x65, 0x75, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Rzlntxyeuc)
	if err != nil {
		err = msgp.WrapError(err, "Rzlntxyeuc")
		return
	}
	// write "Ctzkjkziva"
	err = en.Append(0xaa, 0x43, 0x74, 0x7a, 0x6b, 0x6a, 0x6b, 0x7a, 0x69, 0x76, 0x61)
	if err != nil {
		return
	}
	err = en.WriteString(z.Ctzkjkziva)
	if err != nil {
		err = msgp.WrapError(err, "Ctzkjkziva")
		return
	}
	// write "Orsufumaps"
	err = en.Append(0xaa, 0x4f, 0x72, 0x73, 0x75, 0x66, 0x75, 0x6d, 0x61, 0x70, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Orsufumaps)
	if err != nil {
		err = msgp.WrapError(err, "Orsufumaps")
		return
	}
	// write "Hyevwbtcml"
	err = en.Append(0xaa, 0x48, 0x79, 0x65, 0x76, 0x77, 0x62, 0x74, 0x63, 0x6d, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Hyevwbtcml)
	if err != nil {
		err = msgp.WrapError(err, "Hyevwbtcml")
		return
	}
	// write "Baatlyhdao"
	err = en.Append(0xaa, 0x42, 0x61, 0x61, 0x74, 0x6c, 0x79, 0x68, 0x64, 0x61, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteString(z.Baatlyhdao)
	if err != nil {
		err = msgp.WrapError(err, "Baatlyhdao")
		return
	}
	// write "Fkfohsvvxs"
	err = en.Append(0xaa, 0x46, 0x6b, 0x66, 0x6f, 0x68, 0x73, 0x76, 0x76, 0x78, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Fkfohsvvxs)
	if err != nil {
		err = msgp.WrapError(err, "Fkfohsvvxs")
		return
	}
	// write "Pqwarpxptp"
	err = en.Append(0xaa, 0x50, 0x71, 0x77, 0x61, 0x72, 0x70, 0x78, 0x70, 0x74, 0x70)
	if err != nil {
		return
	}
	err = en.WriteString(z.Pqwarpxptp)
	if err != nil {
		err = msgp.WrapError(err, "Pqwarpxptp")
		return
	}
	// write "Orvaukawww"
	err = en.Append(0xaa, 0x4f, 0x72, 0x76, 0x61, 0x75, 0x6b, 0x61, 0x77, 0x77, 0x77)
	if err != nil {
		return
	}
	err = en.WriteString(z.Orvaukawww)
	if err != nil {
		err = msgp.WrapError(err, "Orvaukawww")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Object3) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "Xvlbzgbaic"
	o = append(o, 0x8a, 0xaa, 0x58, 0x76, 0x6c, 0x62, 0x7a, 0x67, 0x62, 0x61, 0x69, 0x63)
	o = msgp.AppendString(o, z.Xvlbzgbaic)
	// string "Krbemfdzdc"
	o = append(o, 0xaa, 0x4b, 0x72, 0x62, 0x65, 0x6d, 0x66, 0x64, 0x7a, 0x64, 0x63)
	o = msgp.AppendString(o, z.Krbemfdzdc)
	// string "Rzlntxyeuc"
	o = append(o, 0xaa, 0x52, 0x7a, 0x6c, 0x6e, 0x74, 0x78, 0x79, 0x65, 0x75, 0x63)
	o = msgp.AppendString(o, z.Rzlntxyeuc)
	// string "Ctzkjkziva"
	o = append(o, 0xaa, 0x43, 0x74, 0x7a, 0x6b, 0x6a, 0x6b, 0x7a, 0x69, 0x76, 0x61)
	o = msgp.AppendString(o, z.Ctzkjkziva)
	// string "Orsufumaps"
	o = append(o, 0xaa, 0x4f, 0x72, 0x73, 0x75, 0x66, 0x75, 0x6d, 0x61, 0x70, 0x73)
	o = msgp.AppendString(o, z.Orsufumaps)
	// string "Hyevwbtcml"
	o = append(o, 0xaa, 0x48, 0x79, 0x65, 0x76, 0x77, 0x62, 0x74, 0x63, 0x6d, 0x6c)
	o = msgp.AppendString(o, z.Hyevwbtcml)
	// string "Baatlyhdao"
	o = append(o, 0xaa, 0x42, 0x61, 0x61, 0x74, 0x6c, 0x79, 0x68, 0x64, 0x61, 0x6f)
	o = msgp.AppendString(o, z.Baatlyhdao)
	// string "Fkfohsvvxs"
	o = append(o, 0xaa, 0x46, 0x6b, 0x66, 0x6f, 0x68, 0x73, 0x76, 0x76, 0x78, 0x73)
	o = msgp.AppendString(o, z.Fkfohsvvxs)
	// string "Pqwarpxptp"
	o = append(o, 0xaa, 0x50, 0x71, 0x77, 0x61, 0x72, 0x70, 0x78, 0x70, 0x74, 0x70)
	o = msgp.AppendString(o, z.Pqwarpxptp)
	// string "Orvaukawww"
	o = append(o, 0xaa, 0x4f, 0x72, 0x76, 0x61, 0x75, 0x6b, 0x61, 0x77, 0x77, 0x77)
	o = msgp.AppendString(o, z.Orvaukawww)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Object3) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Xvlbzgbaic":
			z.Xvlbzgbaic, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Xvlbzgbaic")
				return
			}
		case "Krbemfdzdc":
			z.Krbemfdzdc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Krbemfdzdc")
				return
			}
		case "Rzlntxyeuc":
			z.Rzlntxyeuc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Rzlntxyeuc")
				return
			}
		case "Ctzkjkziva":
			z.Ctzkjkziva, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ctzkjkziva")
				return
			}
		case "Orsufumaps":
			z.Orsufumaps, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Orsufumaps")
				return
			}
		case "Hyevwbtcml":
			z.Hyevwbtcml, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hyevwbtcml")
				return
			}
		case "Baatlyhdao":
			z.Baatlyhdao, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Baatlyhdao")
				return
			}
		case "Fkfohsvvxs":
			z.Fkfohsvvxs, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Fkfohsvvxs")
				return
			}
		case "Pqwarpxptp":
			z.Pqwarpxptp, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Pqwarpxptp")
				return
			}
		case "Orvaukawww":
			z.Orvaukawww, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Orvaukawww")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Object3) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.Xvlbzgbaic) + 11 + msgp.StringPrefixSize + len(z.Krbemfdzdc) + 11 + msgp.StringPrefixSize + len(z.Rzlntxyeuc) + 11 + msgp.StringPrefixSize + len(z.Ctzkjkziva) + 11 + msgp.StringPrefixSize + len(z.Orsufumaps) + 11 + msgp.StringPrefixSize + len(z.Hyevwbtcml) + 11 + msgp.StringPrefixSize + len(z.Baatlyhdao) + 11 + msgp.StringPrefixSize + len(z.Fkfohsvvxs) + 11 + msgp.StringPrefixSize + len(z.Pqwarpxptp) + 11 + msgp.StringPrefixSize + len(z.Orvaukawww)
	return
}
