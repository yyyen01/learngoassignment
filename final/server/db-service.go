package main

import (
	"assignment/learngoassignment/final/pkg/apis"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func NewDBService(conn *gorm.DB) *dbService {
	if conn == nil {
		log.Fatal("Connection is nil")
	}
	return &dbService{conn: conn}
}

type dbService struct {
	conn *gorm.DB
}

func (s dbService) register(info *apis.UserInfo) error {
	resp := s.conn.Create(info)

	if err := resp.Error; err != nil {
		fmt.Printf("Failure when creating user record for %+v：%v\n", info, err)
		return err
	}
	return nil
}

func (s dbService) login(info *apis.UserInfo) error {
	dbinfo := &apis.UserInfo{}
	resp := s.conn.Where("account_no = ? and password = ?", info.AccountNo, info.Password).Find(dbinfo)

	if err := resp.Error; err != nil {
		log.Println("Unable to find the account number：", err)
		return err
	}

	if resp.RowsAffected == 0 {
		log.Println("No row found. Login Fail")
		return errors.New("Credential not correct")
	}

	log.Println("Nickname :", dbinfo.NickName)
	info.NickName = dbinfo.NickName
	info.Id = dbinfo.Id

	return nil
}
