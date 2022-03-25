package main

import (
	"assignment/learngoassignment/assign6/crinterface"
	"assignment/learngoassignment/assign6/pkg/apis"
	"fmt"
	"sort"
	"sync"
)

var _ crinterface.ServerInterface = &CircleCache{}

const SexMale string = "M"
const SexFemale string = "F"

var sexArray = [2]string{SexMale, SexFemale}

type CircleItem struct {
	ID            uint32
	Timestamp     int64
	PersonID      uint32
	PersonName    string
	Content       string
	AtTimeHeight  float32
	AtTimeWeight  float32
	AtTimeFatRate float32
	Visible       bool
}

type CircleCache struct {
	items     []CircleItem
	itemsLock sync.Mutex
}

func (c2 *CircleCache) PostStatus(c *apis.Circle) error {
	fatrate := calFatRate(c.Sex, c.AtTimeHeight, c.AtTimeWeight, c.AtTimeAge)
	c2.inputRecord(c, fatrate)
	return nil
}

func (c2 *CircleCache) inputRecord(c *apis.Circle, fatRate float32) {
	c2.itemsLock.Lock()
	defer c2.itemsLock.Unlock()

	c2.items = append(c2.items, CircleItem{
		ID:            c.Id,
		Timestamp:     c.Timestamp,
		PersonID:      c.PersonId,
		PersonName:    c.PersonName,
		Content:       c.Content,
		AtTimeHeight:  c.AtTimeHeight,
		AtTimeWeight:  c.AtTimeWeight,
		AtTimeFatRate: fatRate,
		Visible:       c.Visible,
	})

}
func (c2 *CircleCache) DeletePost(persoanlid uint32) error {
	for i, item := range c2.items {
		if item.PersonID == persoanlid {
			c2.items[i].Visible = false
		}
	}
	return nil
}

func (c2 *CircleCache) ListPost() ([]*apis.TopPost, error) {
	c2.itemsLock.Lock()
	defer c2.itemsLock.Unlock()
	sort.Slice(c2.items, func(i, j int) bool {
		return c2.items[i].Timestamp > c2.items[j].Timestamp
	})

	fmt.Printf(" number of record post : %d", len(c2.items))

	count := 0
	out := make([]*apis.TopPost, 0, 10)
	for _, item := range c2.items {
		if item.Visible {
			out = append(out, &apis.TopPost{
				ID:            item.ID,
				Timestamp:     item.Timestamp,
				PersonID:      item.PersonID,
				PersonName:    item.PersonName,
				Content:       item.Content,
				AtTimeHeight:  item.AtTimeHeight,
				AtTimeWeight:  item.AtTimeWeight,
				AtTimeFatRate: item.AtTimeFatRate,
			})
			count++
		}
		if count > 10 {
			break
		}
	}
	fmt.Printf(" Getting post : %+v", out)
	return out, nil
}

func NewCircleCache() *CircleCache {
	return &CircleCache{
		items: make([]CircleItem, 0, 1000),
	}
}

func calFatRate(sex string, tall float32, weight float32, age uint32) (currentFatRate float32) {
	var sexFactor float32

	bmi := weight / (tall * tall)

	if sex == SexFemale {
		sexFactor = 0.0
	} else {
		sexFactor = 1.0
	}
	currentFatRate = (1.2*bmi + 0.23*float32(age) - 5.4 - 10.8*float32(sexFactor)) / 100
	return
}
