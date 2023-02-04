package codegen

const GoCodeTemplate = `// Code generated by protoc-gen-mk. DO NOT EDIT.
package {{.PackageName}}

import (
	"log"
	"net"
	"net/rpc"
)

{{range $index, $message := .MessageList}}
type {{$message.MessageName}} struct {
	{{- range $index, $field := $message.Field}}
	{{$field.FieldName}} {{$field.FieldType}}
	{{- end}}
}
{{end}}


{{range $index, $service := .ServiceList}}
type {{.ServiceName}}ServiceInterface interface {
	{{- range $index, $methodName := .MethodList}}
		{{.MethodName}}(in *{{.RequestName}}, out *{{.ReplyName}}) error
	{{- end}}
}


func (s *server) Register{{.ServiceName}}Service(svc {{.ServiceName}}ServiceInterface) {
	s.services["{{.ServiceName}}"] = svc
}

type {{.ServiceName}}Client struct {
	client *rpc.Client
}

func New{{.ServiceName}}Client(network, addr string) (*{{.ServiceName}}Client, error) {
	conn, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &{{.ServiceName}}Client{
		client: conn,
	}, nil
}

var _ {{.ServiceName}}ServiceInterface = (*{{.ServiceName}}Client)(nil)
{{$svrName := .ServiceName}}
{{- range $index, $methodName := .MethodList}}
func (c *{{$svrName}}Client) {{.MethodName}}(in *{{.RequestName}}, out *{{.ReplyName}}) error {
	return c.client.Call("{{$svrName}}.{{.MethodName}}", in, out)
}
{{end}}
{{end}}

type server struct {
	services map[string]any
}

func NewServer() *server{
	return &server{
			services: make(map[string]any, 1),
	}
}

func (s *server) Run(network, addr string) error {
	for name, svc := range s.services {
			if err := rpc.RegisterName(name, svc); err != nil {
					return nil
			}
	}
	lis, err := net.Listen(network, addr)
	if err != nil {
			return err
	}
	for {
			conn, err := lis.Accept()
			if err != nil {
					log.Fatalf("[ERROR]: accept connection failed because of %v", err.Error())
			}
			go func() {
					rpc.ServeConn(conn)
					defer conn.Close()
			}()
	}
}
`