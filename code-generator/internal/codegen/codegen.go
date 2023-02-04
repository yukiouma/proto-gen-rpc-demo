package codegen

import (
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	extention = ".mk.go"
)

func NewGenerator(tmplString string) (*Generator, error) {
	tmpl, err := template.New("codegen").Parse(tmplString)
	if err != nil {
		return nil, err
	}
	return &Generator{tmpl: tmpl}, nil
}

type Generator struct {
	tmpl *template.Template
}

func (g *Generator) Generate(plugin *protogen.Plugin) error {
	for _, f := range plugin.Files {
		generatedFile := plugin.NewGeneratedFile(
			f.GeneratedFilenamePrefix+extention,
			f.GoImportPath,
		)
		metadata := File{
			PackageName: string(f.GoPackageName),
			MessageList: make([]*Message, 0),
			ServiceList: make([]*Service, 0),
		}
		if err := buildServices(f.Services, &metadata); err != nil {
			return err
		}
		if err := buildMessages(f.Messages, &metadata); err != nil {
			return err
		}
		g.tmpl.Execute(generatedFile, metadata)
	}
	return nil
}
