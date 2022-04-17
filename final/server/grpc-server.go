package main

import (
	"assignment/learngoassignment/final/pkg/apis"
	_ "assignment/learngoassignment/final/pkg/apis"
	"fmt"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	startGRPCServer(ctx)
}

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:12345678!@tcp(127.0.0.1:3306)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func register() {

}

func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	conn := connectDb()
	chatServer := NewChatServer(conn)
	apis.RegisterChatServiceServer(s, chatServer)

	go func() {

		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
