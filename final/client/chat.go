package main

import (
	"assignment/learngoassignment/final/pkg/apis"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {

	fmt.Println("=========================================")
	fmt.Println("Welcome to Chat Service")
	fmt.Println("=========================================")
	fmt.Println("Connecting to Chat Server.")

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ctx := context.TODO()
	c := apis.NewChatServiceClient(conn)

	switch os.Args[1] {
	case "register":
		registerOp := flag.NewFlagSet("register", flag.ExitOnError)
		nickname := registerOp.String("nickname", "", "Enter Your Nickname")
		password := registerOp.String("password", "", "Enter Your Password")
		registerOp.Parse(os.Args[2:])
		if *nickname == "" || *password == "" {
			log.Fatal("Please enter nickname and password")
		}
		log.Println("nickname :", *nickname, "  password:", *password)
		register(c, ctx, *nickname, *password)

	case "login":
		loginOp := flag.NewFlagSet("login", flag.ExitOnError)
		account := loginOp.Int64("account", -1, "Enter Your Account")
		password := loginOp.String("password", "", "Enter Your Password")
		loginOp.Parse(os.Args[2:])
		if *account == -1 || *password == "" {
			log.Fatal("Please enter account number and password")
		}
		//log.Println("account :", *account, "  password:", *password)
		logininfo, err := login(c, ctx, *account, *password, true)
		if err != nil {
			fmt.Println("Error in login :", err)
			os.Exit(1)
		}
		fmt.Println("Login Successfully! Welcome back ", logininfo.NickName)
	case "list":
		listOp := flag.NewFlagSet("list", flag.ExitOnError)
		account := listOp.Int64("account", -1, "Enter Your Account")
		page := listOp.Int("page", 1, "Enter the page number")
		listOp.Parse(os.Args[2:])
		if *account == -1 || *page <= 0 {
			log.Fatal("Please enter account number and make sure page number is greater than zero")
		}

		log.Printf("account no: %d , page: %d", *account, *page)

		list(c, ctx, *account, int64(*page))
	case "with":
	case "history":
	case "subscribe":
	case "end":
		endOp := flag.NewFlagSet("end", flag.ExitOnError)
		account := endOp.Int64("account", -1, "Enter Your Account")
		endOp.Parse(os.Args[2:])
		if *account == -1 {
			log.Fatal("Please enter account number ")
		}
		endChat(c, ctx, *account)
	default:
		fmt.Println("Action not supported! Please try again!")
	}

}

func list(c apis.ChatServiceClient, ctx context.Context, accountNo int64, pageNum int64) error {
	error := loginUsingCredentialFile(c, ctx, accountNo)
	if error != nil {
		fmt.Println("Unable to login using Credential File. Please try manual login first ", error)
		return error
	}
	onlineUsers, error := c.GetOnlineUsers(ctx, &apis.PageNumber{Page: pageNum})
	if error != nil {
		fmt.Println("Unable to retrieve online user list :", error)
		return error
	}
	for _, info := range onlineUsers.UserInfo {
		fmt.Printf("AccountNumber : %d , NickName : %s", info.AccountNo, info.NickName)
	}
	return nil

}

func loginUsingCredentialFile(c apis.ChatServiceClient, ctx context.Context, accountNo int64) error {
	info := &apis.UserInfo{
		AccountNo: accountNo,
	}

	error := readCredentialFromFile(info)
	if error != nil {
		log.Println("Not able to read credential File. Please try manual login first.")
		return error
	}

	info, error = login(c, ctx, info.AccountNo, info.Password, false)
	if error != nil {
		log.Println("Not able to login using credential File. Please try manual login first.")
		return error
	}
	return nil

}

func endChat(c apis.ChatServiceClient, ctx context.Context, accountNo int64) error {
	filepath := getCredentialFilePath(accountNo)
	cr := apis.NewCredentialRecord(filepath)
	cr.RemoveCredentialFile()
	account := &apis.AccountNumber{AccountNo: accountNo}
	c.EndChat(ctx, account)
	fmt.Println("Chat Ended successfully")
	return nil
}

func readCredentialFromFile(info *apis.UserInfo) error {
	filepath := getCredentialFilePath(info.AccountNo)
	cr := apis.NewCredentialRecord(filepath)
	info, err := cr.ReadCredential()
	if err != nil {
		return err
	}
	return nil
}

func getCredentialFilePath(accountNumber int64) string {
	return fmt.Sprintf("../cache/client/credential/%d", accountNumber)
}

func register(c apis.ChatServiceClient, ctx context.Context, nickname string, password string) (*apis.UserInfo, error) {
	info := &apis.UserInfo{
		Password: password,
		NickName: nickname,
	}

	ret, err := c.Register(ctx, info)
	if err != nil {
		log.Fatal("Fail to register ：", err)
		return nil, err
	}
	info.AccountNo = ret.AccountNo
	log.Println("Register successfully. Your account number is ", ret.AccountNo)

	return ret, nil
}

func saveCredentialToFile(info *apis.UserInfo) error {
	filepath := getCredentialFilePath(info.AccountNo)
	cr := apis.NewCredentialRecord(filepath)
	err := cr.SaveCredential(info)
	if err != nil {
		return err
	}
	return nil
}

func login(c apis.ChatServiceClient, ctx context.Context, account int64, password string, saveFile bool) (*apis.UserInfo, error) {
	info := &apis.UserInfo{
		Password:  password,
		AccountNo: account,
	}
	ret, err := c.Login(ctx, info)

	if err != nil {
		log.Fatal("Fail to login：", err)
	}
	if saveFile {
		err = saveCredentialToFile(ret)
		if err != nil {
			return ret, err
		}
	}
	log.Println("Login successfully")

	return ret, nil
}
