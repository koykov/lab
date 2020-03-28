package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/loader"
)

var (
	processed map[string]bool
	pkgName   string
)

func main() {
	pkgName = "github.com/koykov/inspector/testobj"

	processed = make(map[string]bool)

	var conf loader.Config
	conf.Import(pkgName)
	prog, err := conf.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	pkg := prog.Package(pkgName)
	for e, t := range pkg.Info.Types {
		if t.IsType() {
			u := t.Type.Underlying()
			if typ, ok := u.(*types.Struct); ok {
				if idnt, ok := e.(*ast.Ident); ok {
					if _, ok := processed[idnt.Name]; ok {
						continue
					}
					processed[idnt.Name] = true
					fmt.Println("struct", idnt.Name)
					for i := 0; i < typ.NumFields(); i++ {
						f := typ.Field(i)
						parseField(f, 1)
					}
				}
			}
		}
	}
}

func parseType(t types.Type, depth int) {
	u := t.Underlying()
	if p, ok := u.(*types.Pointer); ok {
		u = p.Elem().Underlying()
		parseType(u, depth)
		return
	}
	if s, ok := u.(*types.Struct); ok {
		for i := 0; i < s.NumFields(); i++ {
			f := s.Field(i)
			parseField(f, depth+1)
		}
		return
	}
	if m, ok := u.(*types.Map); ok {
		fmt.Println(strings.Repeat(">", depth+1), "key", m.Key(), "val")
		v := m.Elem().Underlying()
		parseType(v, depth+1)
		return
	}
	if s, ok := u.(*types.Slice); ok {
		v := s.Elem()
		parseType(v, depth+1)
		return
	}

	// all other cases
	fmt.Println(strings.Repeat(">", depth+1), u.String())
}

func parseField(f *types.Var, depth int) {
	ts := f.Type().String()
	fmt.Println(strings.Repeat(">", depth), f.Name(), ts)
	t := f.Type()
	parseType(t, depth)
}
