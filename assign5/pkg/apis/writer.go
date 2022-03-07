package apis

import (
	"encoding/json"
	"log"
	"os"
)

type Writer interface {
	Write(p Person) error
	CreateOutputFile(outputfile string) *os.File
}

type JsonWriter struct {
	file *os.File
}

func (writer *JsonWriter) CreateOutputFile(outputfile string) *os.File {
	file, err := os.Create(outputfile)
	if err != nil {
		log.Fatalln("Unable to create output file : %+v", err)
	}
	writer.file = file
	return file
}

func (writer *JsonWriter) Write(p Person) error {
	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("Unable to marshal into Json : %+v", err)
		return err
	}
	data = append(data, []byte("\n")...)
	if writer.file == nil {
		log.Printf("file is nil")
	}
	_, err = writer.file.Write(data)

	if err != nil {
		log.Printf("Unable to Write to file : %+v", err)
		return err
	}
	return nil
}

func NewJsonWriter(outputFile string) JsonWriter {
	return JsonWriter{}
}
