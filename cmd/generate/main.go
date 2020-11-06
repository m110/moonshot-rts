package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/m110/moonshot-rts/internal/objects"
	"github.com/m110/moonshot-rts/internal/tiles"
	"github.com/m110/moonshot-rts/internal/units"
)

const rootPkg = "github.com/m110/moonshot-rts/"
const prefix = "*components."

const interfaceTemplate = `
func ({{.Short}} {{.Struct}}) Get{{.TypeName}}() {{.Prefix}}{{.TypeName}} {
	return {{.Short}}.{{.TypeName}}
}
`

type interfaceData struct {
	Short    string
	Struct   string
	TypeName string
	Prefix   string
}

var structs = []interface{}{
	objects.Building{},
	objects.Object{},
	objects.Overlay{},
	objects.Panel{},
	objects.PanelButton{},
	objects.ProgressBar{},
	tiles.Tile{},
	units.Unit{},
}

func main() {
	content := map[string]string{}

	for _, s := range structs {
		t := reflect.TypeOf(s)

		tpl, err := template.New("").Parse(interfaceTemplate)
		if err != nil {
			panic(err)
		}

		pkg := strings.Replace(t.PkgPath(), rootPkg, "", 1)
		pkg += "/getters_gen.go"

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			if strings.HasPrefix(f.Type.String(), prefix) {
				data := interfaceData{
					Short:    string(strings.ToLower(t.Name())[0]),
					Struct:   t.Name(),
					TypeName: strings.Replace(f.Type.String(), prefix, "", 1),
					Prefix:   prefix,
				}

				b := bytes.Buffer{}
				err = tpl.Execute(&b, data)
				if err != nil {
					panic(err)
				}

				content[pkg] += b.String()
			}
		}
	}

	for path, c := range content {
		o, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		elements := strings.Split(path, "/")
		pkg := elements[len(elements)-2]

		header := fmt.Sprintf("package %v\n\nimport \"%vinternal/components\"\n", pkg, rootPkg)

		_, err = o.WriteString(header + c)
		if err != nil {
			panic(err)
		}
	}
}
