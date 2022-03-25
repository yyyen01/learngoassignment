package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type crClient struct {
	handRing ClientInterface
}

func (c crClient) post() {

	cr := c.handRing.ReadPostInformation()
	data, _ := json.Marshal(cr)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8081/post", "application/json", r)
	if err != nil {
		log.Println("WARNING: register fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("Posting Return：", string(data))
	}
}

func (c crClient) list() {
	resp, err := http.Get("http://localhost:8081/list")
	if err != nil {
		log.Println("WARNING: Listing fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("Listing Return：", string(data))
	}
}

func (c *crClient) delete(personId uint32) {
	url := fmt.Sprintf("http://localhost:8081/delete/%d", personId)
	log.Println("url", url)
	rep, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println("WARNING: Delete request creation fails:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(rep)
	if err != nil {
		log.Println("WARNING: Delete execution fails:", err)
		return
	}

	defer resp.Body.Close()

	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("Delete Return：", string(data))
	}
}

func main() {
	crCli := &crClient{handRing: &fakeCircleInterface{
		personId:     124,
		personName:   "Yen2",
		sex:          "F",
		content:      "This is my post",
		atTimeHeight: 1.5,
		atTimeWeight: 56.0,
		atTimeAge:    20,
	},
	}
	//Post status in circle
	crCli.post()
	//List top posts in circle
	crCli.list()
	//list post under this person id
	crCli.delete(crCli.handRing.GetPersonId())
	crCli.list()
}
