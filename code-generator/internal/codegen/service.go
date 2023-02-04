// explain the services part of .proto file
package codegen

import "google.golang.org/protobuf/compiler/protogen"

func buildServices(services []*protogen.Service, pbMeta *File) error {
	for _, v := range services {
		service, err := buildService(v)
		if err != nil {
			return err
		}
		pbMeta.ServiceList = append(pbMeta.ServiceList, service)
	}
	return nil
}

func buildService(svc *protogen.Service) (*Service, error) {
	pbSvc := Service{}
	pbSvc.ServiceName = svc.GoName
	methods, err := buildMethods(svc.Methods)
	if err != nil {
		return nil, err
	}
	pbSvc.MethodList = methods
	return &pbSvc, nil
}

func buildMethods(methods []*protogen.Method) ([]*Method, error) {
	pbMethods := make([]*Method, 0, len(methods))
	for _, method := range methods {
		pbMethod, err := buildMethod(method)
		if err != nil {
			return nil, err
		}
		pbMethods = append(pbMethods, pbMethod)
	}
	return pbMethods, nil
}

func buildMethod(method *protogen.Method) (*Method, error) {
	pbMethod := &Method{}
	pbMethod.MethodName = method.GoName
	pbMethod.RequestName = method.Input.GoIdent.GoName
	pbMethod.ReplyName = method.Output.GoIdent.GoName
	return pbMethod, nil
}
