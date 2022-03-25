package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestMarshalJson(t *testing.T) {
	circle := Circle{
		PersonId:     12,
		PersonName:   "Yen",
		Sex:          "F",
		Content:      "First article",
		AtTimeHeight: 1.7,
		AtTimeWeight: 56,
		AtTimeAge:    40,
		Visible:      false,
	}
	fmt.Printf("#{circle}\n")

	data, err := json.Marshal(circle)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal 的结果是(原生)：", data)
	fmt.Println("marshal 的结果是（string）：", string(data))
}

func TestUnmarshalJson(t *testing.T) {
	data := ` {"person_id":12,"person_name":"Yen","sex":"F","content":"First article","at_time_height":1.7,"at_time_weight":56,"at_time_age":40}`
	circle := Circle{}
	err := json.Unmarshal([]byte(data), &circle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", circle)
}

func TestMarshalProtobuf(t *testing.T) {
	circle := &Circle{
		PersonId:     12,
		PersonName:   "Yen",
		Sex:          "F",
		Content:      "First article",
		AtTimeHeight: 1.7,
		AtTimeWeight: 56,
		AtTimeAge:    40,
		Visible:      false,
	}
	data, err := proto.Marshal(circle)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
	// 通常在非程序交互过程中，要保留原生protobuf，可以直接写入文件。如果想要单行保存，必须转码。
	// 选择的通用转码是：base64
	output64Data := base64.StdEncoding.EncodeToString(data)
	fmt.Println(">>>>>", output64Data)
}
