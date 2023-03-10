// Code generated by protoc-gen-mk. DO NOT EDIT.
package v1

import (
	"log"
	"net"
	"net/rpc"
)

type LoginRequest struct {
	account  string
	password string
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func (m *LoginRequest) Account() (data string) {
	if m == nil {
		return
	}
	data = m.account
	return
}

func (m *LoginRequest) SetAccount(val string) {
	if m == nil {
		return
	}
	m.account = val
}

func (m *LoginRequest) Password() (data string) {
	if m == nil {
		return
	}
	data = m.password
	return
}

func (m *LoginRequest) SetPassword(val string) {
	if m == nil {
		return
	}
	m.password = val
}

type LoginReply struct {
	user *User
}

func NewLoginReply() *LoginReply {
	return &LoginReply{}
}

func (m *LoginReply) User() (data *User) {
	if m == nil {
		return
	}
	data = m.user
	return
}

func (m *LoginReply) SetUser(val *User) {
	if m == nil {
		return
	}
	m.user = val
}

type LogoutRequest struct {
}

func NewLogoutRequest() *LogoutRequest {
	return &LogoutRequest{}
}

type LogoutReply struct {
}

func NewLogoutReply() *LogoutReply {
	return &LogoutReply{}
}

type User struct {
	id     int64
	name   string
	locked bool
}

func NewUser() *User {
	return &User{}
}

func (m *User) Id() (data int64) {
	if m == nil {
		return
	}
	data = m.id
	return
}

func (m *User) SetId(val int64) {
	if m == nil {
		return
	}
	m.id = val
}

func (m *User) Name() (data string) {
	if m == nil {
		return
	}
	data = m.name
	return
}

func (m *User) SetName(val string) {
	if m == nil {
		return
	}
	m.name = val
}

func (m *User) Locked() (data bool) {
	if m == nil {
		return
	}
	data = m.locked
	return
}

func (m *User) SetLocked(val bool) {
	if m == nil {
		return
	}
	m.locked = val
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
