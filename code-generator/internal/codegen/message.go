// explain the messages part of .proto file
package codegen

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func buildMessages(messages []*protogen.Message, pbMeta *File) error {
	for _, v := range messages {
		message, err := buildMessage(v)
		if err != nil {
			return err
		}

		pbMeta.MessageList = append(pbMeta.MessageList, message)
	}
	return nil
}

func buildMessage(message *protogen.Message) (*Message, error) {
	pbMessage := &Message{}
	pbMessage.MessageName = message.GoIdent.GoName
	fields, err := buildFields(message.Fields)
	if err != nil {
		return nil, err
	}
	pbMessage.Field = fields
	return pbMessage, nil
}

func buildFields(fields []*protogen.Field) ([]*Field, error) {
	pbFields := make([]*Field, 0, len(fields))
	for _, field := range fields {
		pbField, err := buildField(field)
		if err != nil {
			return nil, err
		}
		pbFields = append(pbFields, pbField)
	}
	return pbFields, nil
}

func buildField(field *protogen.Field) (*Field, error) {
	pbField := &Field{}
	pbField.PublicFieldName = field.GoName
	pbField.PrivateFieldName = publicFieldName2PrivateFieldName(pbField.PublicFieldName)
	desc := field.Desc
	if desc.IsList() {
		// TODO
	} else if desc.IsMap() {
		// TODO
	} else if desc.Message() != nil {
		pbField.FieldType = "*" + field.Message.GoIdent.GoName
	} else {
		pbField.FieldType = kindMapper(desc.Kind().GoString())
	}
	return pbField, nil
}

func kindMapper(in string) (out string) {
	switch in {
	case protoreflect.StringKind.GoString():
		out = "string"
	case protoreflect.Int64Kind.GoString():
		out = "int64"
	case protoreflect.Int32Kind.GoString():
		out = "int32"
	case protoreflect.Uint64Kind.GoString():
		out = "uint64"
	case protoreflect.Uint32Kind.GoString():
		out = "uint32"
	case protoreflect.FloatKind.GoString():
		out = "float32"
	case protoreflect.DoubleKind.GoString():
		out = "float64"
	case protoreflect.BoolKind.GoString():
		out = "bool"
	default:
		out = "any"
	}
	return
}

// transfer public field name to priviate field name
//
// for example: publicFieldName2PrivateFieldName("User") = user
func publicFieldName2PrivateFieldName(origin string) string {
	size := len(origin)
	if size == 0 {
		return ""
	}
	firstChar := origin[0]
	if firstChar > 64 && firstChar < 91 {
		return string(firstChar+32) + origin[1:]
	}
	return origin
}
