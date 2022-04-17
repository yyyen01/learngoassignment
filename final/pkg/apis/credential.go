package apis

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
)

func NewCredentialRecord(filePath string) *credentialRecord {
	return &credentialRecord{filePath: filePath}
}

type credentialRecord struct {
	filePath string
}

func (c *credentialRecord) ReadCredential() (info *UserInfo, err error) {
	data, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		log.Println("读取文件失败：", err)
		return
	}
	info = &UserInfo{}
	err = proto.Unmarshal(data, info)
	if err != nil {
		log.Println("Error in unmarshalling :", err)
		return nil, err
	}
	log.Println("读取出来的内容是：", info.String())
	return info, nil

}

func (c *credentialRecord) RemoveCredentialFile() error {
	log.Println("Deleting ", c.filePath)
	err := os.Remove(c.filePath)
	if err != nil {
		log.Println("Error in removing file :", err)
		return err
	}
	return nil
}

func (c *credentialRecord) SaveCredential(info *UserInfo) error {
	data, err := proto.Marshal(info)
	if err != nil {
		log.Println("marshal 出错：", err)
		return err
	}
	if err := c.writeFileWithAppendProtobuf(data); err != nil {
		log.Println("写入PROTOBUF时出错：", err)
		return err
	}
	return nil
}

func (c *credentialRecord) writeFileWithAppendProtobuf(data []byte) error {
	file, err := os.OpenFile(c.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission settings
	if err != nil {
		log.Println("无法打开文件", c.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	//_, err = file.Write([]byte(base64.StdEncoding.EncodeToString(data)))
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}
