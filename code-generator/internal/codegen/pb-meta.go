package codegen

type File struct {
	PackageName string
	ServiceList []*Service
	MessageList []*Message
}

type Method struct {
	MethodName  string
	RequestName string
	ReplyName   string
}

type Service struct {
	ServiceName string
	MethodList  []*Method
}

type Message struct {
	MessageName string
	Field       []*Field
}

type Field struct {
	FieldName string
	FieldType string
}
