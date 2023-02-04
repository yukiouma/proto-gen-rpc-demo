package main

import (
	v1 "demo/api/demo/v1"
	"log"
)

const (
	network = "tcp"
	addr    = "localhost:8000"
)

func main() {
	var (
		loginReq  *v1.LoginRequest
		loginRes  *v1.LoginReply
		logoutReq *v1.LogoutRequest
		logoutRes *v1.LogoutReply
	)
	client, err := v1.NewDemoClient(network, addr)
	if err != nil {
		log.Fatalf("[ERROR]: failed to connect server because: %v", err)
	}
	loginReq = v1.NewLoginRequest()
	loginReq.SetAccount("a")
	loginReq.SetPassword("123")
	loginRes = v1.NewLoginReply()
	if err := client.Login(loginReq, loginRes); err != nil {
		log.Fatalf("[ERROR]: login failed because: %v", err)
	}
	log.Printf("[INFO]: welcome, %v", loginRes.User.Name)

	logoutReq = v1.NewLogoutRequest()
	logoutRes = v1.NewLogoutReply()
	if err := client.Logout(logoutReq, logoutRes); err != nil {
		log.Fatalf("[ERROR]: logout failed because: %v", err)
	}

	log.Printf("[INFO]: bye, %v", loginRes.User.Name)

	loginReq = v1.NewLoginRequest()
	loginReq.SetAccount("b")
	loginReq.SetPassword("456")
	loginRes = v1.NewLoginReply()
	if err := client.Login(loginReq, loginRes); err != nil {
		log.Fatalf("[ERROR]: login failed because: %v", err)
	}
	log.Printf("[INFO]: welcome, %v", loginRes.User.Name)

	logoutReq = v1.NewLogoutRequest()
	logoutRes = v1.NewLogoutReply()
	if err := client.Logout(logoutReq, logoutRes); err != nil {
		log.Fatalf("[ERROR]: logout failed because: %v", err)
	}

	log.Printf("[INFO]: bye, %v", loginRes.User.Name)

	loginReq = v1.NewLoginRequest()
	loginReq.SetAccount("b")
	loginReq.SetPassword("1456")
	loginRes = v1.NewLoginReply()
	if err := client.Login(loginReq, loginRes); err != nil {
		log.Fatalf("[ERROR]: login failed because: %v", err)
	}
}
