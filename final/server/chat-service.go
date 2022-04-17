package main

import (
	"assignment/learngoassignment/final/pkg/apis"
	"context"
	context2 "golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var _ apis.ChatServiceServer = &chatServer{}

func NewChatServer(conn *gorm.DB) *chatServer {
	db := NewDBService(conn)
	onlineUsers := make([]*apis.UserInfo, 0, 1000)
	return &chatServer{db: db, onlineUsers: onlineUsers}
}

type chatServer struct {
	lock        sync.Mutex
	db          *dbService
	onlineUsers []*apis.UserInfo
}

func (c *chatServer) EndChat(ctx context.Context, number *apis.AccountNumber) (*apis.Status, error) {
	var found bool
	var index int
	c.lock.Lock()
	for i, user := range c.onlineUsers {
		if user.AccountNo == number.AccountNo {
			found = true
			index = i
			break
		}
	}
	if found {
		c.onlineUsers = append(c.onlineUsers[:index], c.onlineUsers[index+1:]...)
	}

	c.lock.Unlock()
	return &apis.Status{Status: true}, nil
}

func (c *chatServer) Login(ctx context2.Context, info *apis.UserInfo) (*apis.UserInfo, error) {
	error := c.db.login(info)

	if error != nil {
		log.Println("Unable to log in : ", error)
		return nil, error
	}
	c.lock.Lock()
	c.onlineUsers = append(c.onlineUsers, info)
	c.lock.Unlock()

	log.Printf("User %s login successfully.", info.NickName)
	return info, nil
}

func (c *chatServer) GetOnlineUsers(ctx context2.Context, number *apis.PageNumber) (*apis.OnlineUsers, error) {
	listOfUsers := &apis.OnlineUsers{}
	start := 2*number.Page + 1
	end := 2*number.Page + 20
	log.Println("Getting 20 users on page ", number.Page)
	c.lock.Lock()
	defer c.lock.Unlock()
	for i, user := range c.onlineUsers {
		if i >= int(start) && i <= int(end) {
			listOfUsers.UserInfo = append(listOfUsers.UserInfo, user)
		}
	}
	return listOfUsers, nil
}

func (c *chatServer) Chat(ctx context2.Context, number *apis.AccountNumber) (*apis.ChatRecords, error) {
	listOfChatRecords, _ := c.ChatHistory(ctx, number)
	log.Println("Start charting")
	return listOfChatRecords, nil
}

func (c *chatServer) ChatHistory(ctx context2.Context, number *apis.AccountNumber) (*apis.ChatRecords, error) {
	listOfChatRecords := &apis.ChatRecords{}
	log.Println("Retrieving last 20 records with account ", number.AccountNo)
	return listOfChatRecords, nil

}

func (c *chatServer) Subscribe(ctx context2.Context, number *apis.AccountNumber) (*apis.Status, error) {
	log.Printf("Adding %s to subscription list", number.AccountNo)
	return &apis.Status{Status: true}, nil
}

func (c *chatServer) Register(ctx context2.Context, info *apis.UserInfo) (*apis.UserInfo, error) {
	t := time.Now()
	info.AccountNo = t.Unix()

	error := c.db.register(info)
	if error != nil {
		log.Println("Error in creating user record :", error)
		return nil, error
	}
	log.Printf("Register successful: %s\n", info.String())

	return info, error
}
