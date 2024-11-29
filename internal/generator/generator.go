package generator

import (
	"bytes"
	"io"
)

type Method struct {
	Name       string
	InputType  string
	OutputType string
}

type TemplateData struct {
	PackageName  string
	ProtoPackage string
	ServiceName  string
	Methods      []Method
}

func GenerateClient(w io.Writer, data TemplateData) error {
	var buf bytes.Buffer
	if err := clientStubTemplate.Execute(&buf, data); err != nil {
		return err
	}
	_, err := w.Write(buf.Bytes())
	return err
}

func GenerateServer(w io.Writer, data TemplateData) error {
	var buf bytes.Buffer
	if err := serverStubTemplate.Execute(&buf, data); err != nil {
		return err
	}
	_, err := w.Write(buf.Bytes())
	return err
}