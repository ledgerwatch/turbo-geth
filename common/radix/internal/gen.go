package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"text/template"
)

type TemplateVariables struct {
	Type      string
	Humanized string
	TreeType  string
	NilValue  string
}

var Types = []TemplateVariables{
	{
		Type:     "map[common.Hash]uint64",
		TreeType: "RadixMapHash2Uint64",
		NilValue: "nil",
	}, {
		Type:     "map[uint64]int",
		TreeType: "RadixMapUint642Int",
		NilValue: "nil",
	},
	{
		Type:     "uint64",
		TreeType: "RadixUint64",
		NilValue: "0",
	},
}

func main() {
	buf := bytes.NewBuffer(nil)

	fmt.Fprintf(buf, `// Code generated by go generate; DO NOT EDIT.
package radix

import (
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ugorji/go/codec"
)
`)

	for _, el := range Types {
		if err := typedTreeTemplate.Execute(buf, el); err != nil {
			panic(err)
		}
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("typed_gen.go", b, 0644); err != nil {
		panic(err)
	}
}

var typedTreeTemplate = template.Must(template.New("").Parse(`
// string -> {{.Type}}
type {{.TreeType}} struct {
	*Tree
	version string
}

func New{{.TreeType}}() *{{.TreeType}} {
	return &{{.TreeType}}{
		Tree: New(),
		version: "1",
	}
}

func (t *{{.TreeType}}) Get(k string) ({{.Type}}, bool) {
	v, ok := t.Tree.Get(k)
	if ok {
		return v.({{.Type}}), ok
	}
	return {{.NilValue}}, ok
}

func (t *{{.TreeType}}) Insert(k string, v {{.Type}}) {
	t.Tree.Insert(k, v)
}

func (t *{{.TreeType}}) Delete(k string) ({{.Type}}, bool) {
	v, ok := t.Tree.Delete(k)
	if ok {
		return v.({{.Type}}), ok
	}
	return {{.NilValue}}, ok
}

func (t *{{.TreeType}}) DeletePrefix(k string) int {
	return t.Tree.DeletePrefix(k)
}

func (t *{{.TreeType}}) Walk(f func(string, {{.Type}}) bool) {
	t.Tree.Walk(func(k string, v interface{}) bool {
		return f(k, v.({{.Type}}))
	})
}

func (t *{{.TreeType}}) CodecEncodeSelf(e *codec.Encoder) {
	e.MustEncode(t.version)
	e.MustEncode(t.Tree.Len())
	t.Tree.Walk(func(k string, v interface{}) bool {
		e.MustEncode(&k)
		e.MustEncode(&v)
		return false
	})
}

func (t *{{.TreeType}}) CodecDecodeSelf(d *codec.Decoder) {
	var version string
	d.MustDecode(&version)
	if version != t.version {
		panic("unexpected version")
	}
	var amount int
	d.MustDecode(&amount)
	for i := 0; i < amount; i++ {
		var k string
		var v {{.Type}}
		d.MustDecode(&k)
		d.MustDecode(&v)
		t.Insert(k, v)
	}
}
`))
