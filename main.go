package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	simplepb "goProjects/bufferedFiles"
	"io/ioutil"
	"log"
)

func main() {
	doAll()
}

func doAll() {
	sm := &simplepb.SimpleMessage{
		Name:     "Pushpan",
		Id:       777,
		IsSimple: true,
	}
	fmt.Println(sm)
	err := writeIntoFile("simple.bin",sm)
	if err != nil {
		return
	}
	newSm := &simplepb.SimpleMessage{}
	err = readFromFile("simple.bin",newSm)
	if err != nil {
		return
	}
	fmt.Println("--------")
	fmt.Println(newSm)
}

func writeIntoFile(fileName string ,sm *simplepb.SimpleMessage ) error{

	payload, err := proto.Marshal(sm)
	if err != nil {
		log.Fatalln("Error in converting data into bytes",err)
		return err
	}
	err = ioutil.WriteFile(fileName,payload,0667)
	if err != nil {
		log.Fatalln("Error in writing data into File",err)
		return err
	}
	fmt.Println("Data is added to file")
	return nil

}

func readFromFile(fileName string, sm *simplepb.SimpleMessage) error{

	payload,err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error in reading file",err)
		return err
	}
	err = proto.Unmarshal(payload,sm)
	if err != nil {
		log.Fatalln("Error in Converting data",err)
		return err
	}
	fmt.Println("Data is Added from File")
	return nil

}