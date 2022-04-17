package main

import (
	"assignment/learngoassignment/assign6/crinterface"
	"assignment/learngoassignment/assign6/pkg/apis"
	"fmt"
	"gorm.io/gorm"
	"log"
)

var _ crinterface.ServerInterface = &dbCircle{}
var _ crinterface.CircleInitInterface = &dbCircle{}

func NewDbCircle(conn *gorm.DB,
	embedCircle crinterface.ServerInterface) crinterface.ServerInterface {
	if conn == nil {
		log.Fatal("Connection is nil")
	}
	if embedCircle == nil {
		log.Fatal("Embedded Circle is nil")
	}
	return &dbCircle{
		conn:        conn,
		embedCircle: embedCircle,
	}
}

type dbCircle struct {
	conn        *gorm.DB
	embedCircle crinterface.ServerInterface
}

func (d dbCircle) Init() error {
	output := make([]*apis.Circle, 0)
	resp := d.conn.Find(&output)
	if resp.Error != nil {
		fmt.Println("Unable to load Data in DB：", resp.Error)
		return resp.Error
	}
	for _, item := range output {
		if err := d.embedCircle.PostStatus(item); err != nil {
			log.Fatalf("Failure when Initialising %d：%v", item.Id, err)
		}
	}
	return nil
}

func (d dbCircle) PostStatus(c *apis.Circle) error {
	resp := d.conn.Create(c)
	if err := resp.Error; err != nil {
		fmt.Printf("Failure when creating record for %+v：%v\n", c, err)
		return err
	}

	fmt.Println("Post status successfully : %+v", c)
	_ = d.embedCircle.PostStatus()
	return nil
}

func (d dbCircle) DeletePost(person_id uint32) error {

	resp := d.conn.Model(&apis.Circle{}).Where("person_id = ?", person_id).Update("visible", 0)

	if err := resp.Error; err != nil {
		fmt.Println("Delete Failed：", err)
		return err
	}
	fmt.Println("Delete successfully")
	_ = d.embedCircle.DeletePost(person_id)
	return nil
}

func (d dbCircle) ListPost() ([]*apis.TopPost, error) {
	return d.embedCircle.ListPost()
}
