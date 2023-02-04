// Code generated by protoc-gen-mk. DO NOT EDIT.
package v1

import (
	"log"
	"net"
	"net/rpc"
)

type LoginRequest struct {
	Account  string
	Password string
}

type LoginReply struct {
	User *User
}

type LogoutRequest struct {
}

type LogoutReply struct {
}

type User struct {
	Id     int64
	Name   string
	Locked bool
}

type DemoServiceInterface interface {
	Login(in *LoginRequest, out *LoginReply) error
	Logout(in *LogoutRequest, out *LogoutReply) error
}

func (s *server) RegisterDemoService(svc DemoServiceInterface) {
	s.services["Demo"] = svc
}

type DemoClient struct {
	client *rpc.Client
}

func NewDemoClient(network, addr string) (*DemoClient, error) {
	conn, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &DemoClient{
		client: conn,
	}, nil
}

var _ DemoServiceInterface = (*DemoClient)(nil)

func (c *DemoClient) Login(in *LoginRequest, out *LoginReply) error {
	return c.client.Call("Demo.Login", in, out)
}

func (c *DemoClient) Logout(in *LogoutRequest, out *LogoutReply) error {
	return c.client.Call("Demo.Logout", in, out)
}

type server struct {
	services map[string]any
}

func NewServer() *server {
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