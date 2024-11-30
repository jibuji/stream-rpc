package generator

import (
	"bytes"
	"io"
	"text/template"
)

var clientStubTemplate = template.Must(template.New("client").Parse(`
// Code generated by stream-rpc. DO NOT EDIT.
package {{.PackageName}}

import (
	rpc "github.com/jibuji/go-stream-rpc"
)

type {{.ServiceName}}Client struct {
	peer *rpc.RpcPeer
}

func New{{.ServiceName}}Client(peer *rpc.RpcPeer) *{{.ServiceName}}Client {
	return &{{.ServiceName}}Client{peer: peer}
}

{{range .Methods}}
func (c *{{$.ServiceName}}Client) {{.Name}}(req *{{.InputType}}) *{{.OutputType}} {
	resp := &{{.OutputType}}{}
	err := c.peer.Call("{{$.ServiceName}}.{{.Name}}", req, resp)
	if err != nil {
		return nil
	}
	return resp
}
{{end}}
`))

var serverStubTemplate = template.Must(template.New("server").Parse(`
// Code generated by stream-rpc. DO NOT EDIT.
package {{.PackageName}}

import (
	rpc "github.com/jibuji/go-stream-rpc"
	"context"
)

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations
type Unimplemented{{.ServiceName}}Server struct {}

type {{.ServiceName}}Server interface {
	{{range .Methods}}
	{{.Name}}(context.Context, *{{.InputType}}) (*{{.OutputType}})
	{{end}}
}

type {{.ServiceName}}ServerImpl struct {
	impl {{.ServiceName}}Server
}

func Register{{.ServiceName}}Server(peer *rpc.RpcPeer, impl {{.ServiceName}}Server) {
	server := &{{.ServiceName}}ServerImpl{impl: impl}
	peer.RegisterService("{{.ServiceName}}", server)
}

{{range .Methods}}
func (s *Unimplemented{{$.ServiceName}}Server) {{.Name}}(ctx context.Context, req *{{.InputType}}) (*{{.OutputType}}) {
	return nil
}
{{end}}

{{range .Methods}}
func (s *{{$.ServiceName}}ServerImpl) {{.Name}}(ctx context.Context, req *{{.InputType}}) (*{{.OutputType}}) {
	return s.impl.{{.Name}}(ctx, req)
}
{{end}}
`))

var skeletonTemplate = template.Must(template.New("skeleton").Parse(`
// Code generated by protoc-gen-stream-rpc. DO NOT EDIT.
package service

import (
	"context"
	proto "{{.ProtoPackage}}"
)

// {{.ServiceName}}Service implements the {{.ServiceName}} service
type {{.ServiceName}}Service struct {
	proto.Unimplemented{{.ServiceName}}Server
}

{{range .Methods}}
// {{.Name}} implements {{$.ServiceName}}Server
func (s *{{$.ServiceName}}Service) {{.Name}}(ctx context.Context, req *proto.{{.InputType}}) *proto.{{.OutputType}} {
	// TODO: Implement your logic here
	return &proto.{{.OutputType}}{}
}
{{end}}
`))

var methodTemplate = template.Must(template.New("method").Parse(`func (s *{{$.ServiceName}}Service) {{.Name}}(ctx context.Context, req *proto.{{.InputType}}) *proto.{{.OutputType}} {
	// TODO: Implement your logic here
	return &proto.{{.OutputType}}{}
}
`))

func GenerateSkeleton(w io.Writer, data TemplateData, existingMethods map[string]*ExistingMethod) error {
	var buf bytes.Buffer

	// Write header
	headerTemplate := `// Code generated by protoc-gen-stream-rpc. DO NOT EDIT.
package service

import (
	"context"
	proto "{{.ProtoPackage}}"
)

// {{.ServiceName}}Service implements the {{.ServiceName}} service
type {{.ServiceName}}Service struct {
	proto.Unimplemented{{.ServiceName}}Server
}
`
	if err := template.Must(template.New("header").Parse(headerTemplate)).Execute(&buf, data); err != nil {
		return err
	}

	// Write methods
	for _, method := range data.Methods {
		buf.WriteString("\n")
		if existing, ok := existingMethods[method.Name]; ok && existing.SignatureMatches(method) {
			// Keep existing implementation
			buf.WriteString(existing.Body)
		} else {
			// Generate new method
			methodData := struct {
				ServiceName string
				Method
			}{
				ServiceName: data.ServiceName,
				Method:      method,
			}
			if err := methodTemplate.Execute(&buf, methodData); err != nil {
				return err
			}
		}
	}

	_, err := w.Write(buf.Bytes())
	return err
}
