package main

import "assignment/learngoassignment/assign6/pkg/apis"

type ClientInterface interface {
	ReadPostInformation() apis.Circle
	GetPersonId() uint32
}

var _ ClientInterface = &fakeCircleInterface{}

type fakeCircleInterface struct {
	personId     uint32
	personName   string
	sex          string
	content      string
	atTimeHeight float32
	atTimeWeight float32
	atTimeAge    uint32
}

func (f *fakeCircleInterface) ReadPostInformation() apis.Circle {
	cr := apis.Circle{
		PersonId:     f.personId,
		PersonName:   f.personName,
		Sex:          f.sex,
		Content:      f.content,
		AtTimeHeight: f.atTimeHeight,
		AtTimeWeight: f.atTimeWeight,
		AtTimeAge:    f.atTimeAge,
	}

	return cr
}
func (f fakeCircleInterface) GetPersonId() uint32 {
	return f.personId
}
