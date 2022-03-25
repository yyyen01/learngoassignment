package crinterface

import "assignment/learngoassignment/assign6/pkg/apis"

type ServerInterface interface {
	PostStatus(c *apis.Circle) error
	DeletePost(id uint32) error
	ListPost() ([]*apis.TopPost, error)
}

type CircleInitInterface interface {
	Init() error
}
