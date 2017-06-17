// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package parser

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

func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser(in *jlexer.Lexer, out *CSVTableCollection) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tables":
			if in.IsNull() {
				in.Skip()
				out.Tables = nil
			} else {
				in.Delim('[')
				if out.Tables == nil {
					if !in.IsDelim(']') {
						out.Tables = make([]CSVTable, 0, 2)
					} else {
						out.Tables = []CSVTable{}
					}
				} else {
					out.Tables = (out.Tables)[:0]
				}
				for !in.IsDelim(']') {
					var v1 CSVTable
					(v1).UnmarshalEasyJSON(in)
					out.Tables = append(out.Tables, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser(out *jwriter.Writer, in CSVTableCollection) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tables\":")
	if in.Tables == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Tables {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSVTableCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSVTableCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSVTableCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSVTableCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser(l, v)
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser1(in *jlexer.Lexer, out *CSVTable) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "table":
			out.Content = string(in.String())
		case "url":
			out.URL = string(in.String())
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser1(out *jwriter.Writer, in CSVTable) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"table\":")
	out.String(string(in.Content))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSVTable) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSVTable) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSVTable) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSVTable) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser1(l, v)
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser2(in *jlexer.Lexer, out *Collections) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "collections":
			if in.IsNull() {
				in.Skip()
				out.Collections = nil
			} else {
				in.Delim('[')
				if out.Collections == nil {
					if !in.IsDelim(']') {
						out.Collections = make([]*collection, 0, 8)
					} else {
						out.Collections = []*collection{}
					}
				} else {
					out.Collections = (out.Collections)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *collection
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(collection)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Collections = append(out.Collections, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser2(out *jwriter.Writer, in Collections) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"collections\":")
	if in.Collections == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Collections {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Collections) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Collections) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Collections) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Collections) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser2(l, v)
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser3(in *jlexer.Lexer, out *collection) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]interface{}, 0, 4)
					} else {
						out.Items = []interface{}{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v7 interface{}
					if m, ok := v7.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v7.(json.Unmarshaler); ok {
						m.UnmarshalJSON(in.Raw())
					} else {
						v7 = in.Interface()
					}
					out.Items = append(out.Items, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "count":
			out.Count = int(in.Int())
		case "time":
			out.CreatedAt = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser3(out *jwriter.Writer, in collection) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"items\":")
	if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in.Items {
			if v8 > 0 {
				out.RawByte(',')
			}
			if m, ok := v9.(easyjson.Marshaler); ok {
				m.MarshalEasyJSON(out)
			} else if m, ok := v9.(json.Marshaler); ok {
				out.Raw(m.MarshalJSON())
			} else {
				out.Raw(json.Marshal(v9))
			}
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"count\":")
	out.Int(int(in.Count))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"time\":")
	out.Int64(int64(in.CreatedAt))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v collection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v collection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *collection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *collection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser3(l, v)
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser4(in *jlexer.Lexer, out *Parser) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "format":
			out.Format = string(in.String())
		case "paginatedResults":
			out.PaginatedResults = bool(in.Bool())
		case "collections":
			if in.IsNull() {
				in.Skip()
				out.Payloads = nil
			} else {
				in.Delim('[')
				if out.Payloads == nil {
					if !in.IsDelim(']') {
						out.Payloads = make([]Payload, 0, 1)
					} else {
						out.Payloads = []Payload{}
					}
				} else {
					out.Payloads = (out.Payloads)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Payload
					(v10).UnmarshalEasyJSON(in)
					out.Payloads = append(out.Payloads, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "payloadMD5":
			if in.IsNull() {
				in.Skip()
				out.PayloadMD5 = nil
			} else {
				out.PayloadMD5 = in.Bytes()
			}
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser4(out *jwriter.Writer, in Parser) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"format\":")
	out.String(string(in.Format))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"paginatedResults\":")
	out.Bool(bool(in.PaginatedResults))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"collections\":")
	if in.Payloads == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v12, v13 := range in.Payloads {
			if v12 > 0 {
				out.RawByte(',')
			}
			(v13).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"payloadMD5\":")
	out.Base64Bytes(in.PayloadMD5)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Parser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Parser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Parser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Parser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser4(l, v)
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser5(in *jlexer.Lexer, out *Payload) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fields":
			if in.IsNull() {
				in.Skip()
				out.Fields = nil
			} else {
				in.Delim('[')
				if out.Fields == nil {
					if !in.IsDelim(']') {
						out.Fields = make([]field, 0, 1)
					} else {
						out.Fields = []field{}
					}
				} else {
					out.Fields = (out.Fields)[:0]
				}
				for !in.IsDelim(']') {
					var v16 field
					easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser6(in, &v16)
					out.Fields = append(out.Fields, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "paginator":
			easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser7(in, &out.Paginator)
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser5(out *jwriter.Writer, in Payload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"fields\":")
	if in.Fields == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v17, v18 := range in.Fields {
			if v17 > 0 {
				out.RawByte(',')
			}
			easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser6(out, v18)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"paginator\":")
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser7(out, in.Paginator)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Payload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Payload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Payload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Payload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser5(l, v)
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser7(in *jlexer.Lexer, out *paginator) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "selector":
			out.Selector = string(in.String())
		case "attr":
			out.Attribute = string(in.String())
		case "maxPages":
			out.MaxPages = int(in.Int())
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser7(out *jwriter.Writer, in paginator) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"selector\":")
	out.String(string(in.Selector))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"attr\":")
	out.String(string(in.Attribute))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxPages\":")
	out.Int(int(in.MaxPages))
	out.RawByte('}')
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser6(in *jlexer.Lexer, out *field) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "selector":
			out.Selector = string(in.String())
		case "count":
			out.Count = int(in.Int())
		case "extractor":
			easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser8(in, &out.Extractor)
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser6(out *jwriter.Writer, in field) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"selector\":")
	out.String(string(in.Selector))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"count\":")
	out.Int(int(in.Count))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"extractor\":")
	easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser8(out, in.Extractor)
	out.RawByte('}')
}
func easyjsonBc4ecbcDecodeGithubComSlotixDataflowkitParser8(in *jlexer.Lexer, out *Extractor) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "params":
			if m, ok := out.Params.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Params.(json.Unmarshaler); ok {
				m.UnmarshalJSON(in.Raw())
			} else {
				out.Params = in.Interface()
			}
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
func easyjsonBc4ecbcEncodeGithubComSlotixDataflowkitParser8(out *jwriter.Writer, in Extractor) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	out.String(string(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"params\":")
	if m, ok := in.Params.(easyjson.Marshaler); ok {
		m.MarshalEasyJSON(out)
	} else if m, ok := in.Params.(json.Marshaler); ok {
		out.Raw(m.MarshalJSON())
	} else {
		out.Raw(json.Marshal(in.Params))
	}
	out.RawByte('}')
}
