package demo

import (
	v1 "demo/api/demo/v1"
	"fmt"
)

var (
	userData = map[string]*user{
		"a": {1, "A", "123", true},
		"b": {2, "B", "456", false},
	}
)

type user struct {
	Id       int
	Name     string
	Password string
	Locked   bool
}

type demo struct {
	data map[string]*user
}

func NewDemoService() v1.DemoServiceInterface {
	return &demo{
		data: userData,
	}
}

var _ v1.DemoServiceInterface = (*demo)(nil)

func (d *demo) Login(in *v1.LoginRequest, out *v1.LoginReply) error {
	if in == nil {
		return fmt.Errorf("[ERROR]: invalid request")
	}
	acc, pw := in.Account(), in.Password()
	user, ok := d.data[acc]
	if !ok {
		return fmt.Errorf("[ERROR]: invalid user")
	}
	if user.Password != pw {
		return fmt.Errorf("[ERROR]: invalid password")
	}
	replyUser := v1.NewUser()
	replyUser.SetId(int64(user.Id))
	replyUser.SetName(user.Name)
	replyUser.SetLocked(user.Locked)
	out.SetUser(replyUser)
	return nil
}

func (d *demo) Logout(in *v1.LogoutRequest, out *v1.LogoutReply) error {
	return nil
}
