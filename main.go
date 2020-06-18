package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	simplepb "goProjects/bufferedFiles"
	"io/ioutil"
	"log"
)

func main() {
	sm := simplepb.SimpleMessage{
		Name:     "Pushpan",
		Id:       777,
		IsSimple: true,
	}
	doAll(&sm)
}

func doAll(sm *simplepb.SimpleMessage) {

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
	fmt.Println("Data is added")
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
		log.Fatalln("Error in ")
	}
}