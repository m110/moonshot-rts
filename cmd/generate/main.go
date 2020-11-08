package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/m110/moonshot-rts/internal/components"
	"github.com/m110/moonshot-rts/internal/objects"
	"github.com/m110/moonshot-rts/internal/tiles"
	"github.com/m110/moonshot-rts/internal/units"
)

const rootPkg = "github.com/m110/moonshot-rts/"
const prefix = "*components."

const archetypeInterfaceTemplate = `
func ({{.Short}} {{.Struct}}) Get{{.TypeName}}() {{.Prefix}}{{.TypeName}} {
	return {{.Short}}.{{.TypeName}}
}
`

const componentOwnerTemplate = `
type {{.TypeName}}Owner interface {
	Get{{.TypeName}}() *{{.TypeName}}
}
`

type archetypeData struct {
	Short    string
	Struct   string
	TypeName string
	Prefix   string
}

type componentData struct {
	TypeName string
}

var archetypes = []interface{}{
	objects.Building{},
	objects.Object{},
	objects.Overlay{},
	objects.Panel{},
	objects.PanelButton{},
	objects.ProgressBar{},
	tiles.Tile{},
	units.Unit{},
}

var allComponents = []interface{}{
	components.Builder{},
	components.Button{},
	components.Clickable{},
	components.Collider{},
	components.Drawable{},
	components.Movable{},
	components.ProgressBar{},
	components.Selectable{},
	components.Size{},
	components.TimeActions{},
	components.UnitSpawner{},
	components.WorldSpace{},
}

func main() {
	archetypesContent := map[string]string{}

	archetypeTpl, err := template.New("").Parse(archetypeInterfaceTemplate)
	if err != nil {
		panic(err)
	}

	for _, s := range archetypes {
		t := reflect.TypeOf(s)

		pkg := strings.Replace(t.PkgPath(), rootPkg, "", 1)
		pkg += "/getters_gen.go"

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			if strings.HasPrefix(f.Type.String(), prefix) {
				data := archetypeData{
					Short:    string(strings.ToLower(t.Name())[0]),
					Struct:   t.Name(),
					TypeName: strings.Replace(f.Type.String(), prefix, "", 1),
					Prefix:   prefix,
				}

				b := bytes.Buffer{}
				err = archetypeTpl.Execute(&b, data)
				if err != nil {
					panic(err)
				}

				archetypesContent[pkg] += b.String()
			}
		}
	}

	for path, c := range archetypesContent {
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

	componentsContent := map[string]string{}

	componentTpl, err := template.New("").Parse(componentOwnerTemplate)
	if err != nil {
		panic(err)
	}

	for _, c := range allComponents {
		t := reflect.TypeOf(c)

		pkg := strings.Replace(t.PkgPath(), rootPkg, "", 1)
		pkg += "/owners_gen.go"

		names := strings.Split(t.Name(), ".")
		typeName := names[len(names)-1]
		data := componentData{
			TypeName: typeName,
		}

		b := bytes.Buffer{}
		err = componentTpl.Execute(&b, data)
		if err != nil {
			panic(err)
		}

		componentsContent[pkg] += b.String()
	}

	for path, c := range componentsContent {
		o, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		elements := strings.Split(path, "/")
		pkg := elements[len(elements)-2]

		header := fmt.Sprintf("package %v\n", pkg)

		_, err = o.WriteString(header + c)
		if err != nil {
			panic(err)
		}
	}
}
