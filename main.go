package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	complex "goProjects/src/complex"
	enum "goProjects/src/enum"
	simple "goProjects/src/simple"
	"io/ioutil"
	"log"
)

func main() {

	sm := doAll()
	jsonDemo(sm)
	doEnum()
	doComplex()

}

func doAll() *simple.SimpleMessage {

	sm := &simple.SimpleMessage{
		Name:     "Pushpan",
		Id:       777,
		IsSimple: true,
	}
	fmt.Println(sm)
	err := writeIntoFile("simple.bin", sm)
	if err != nil {
		return nil
	}
	newSm := &simple.SimpleMessage{}
	err = readFromFile("simple.bin", newSm)
	if err != nil {
		return nil
	}
	fmt.Println("--------", newSm)
	return newSm

}

func writeIntoFile(fileName string, sm *simple.SimpleMessage) error {

	payload, err := proto.Marshal(sm)
	if err != nil {
		log.Fatalln("Error in converting data into bytes", err)
		return err
	}
	err = ioutil.WriteFile(fileName, payload, 0667)
	if err != nil {
		log.Fatalln("Error in writing data into File", err)
		return err
	}
	fmt.Println("Data is added to file")
	return nil

}

func readFromFile(fileName string, sm *simple.SimpleMessage) error {

	payload, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error in reading file", err)
		return err
	}
	err = proto.Unmarshal(payload, sm)
	if err != nil {
		log.Fatalln("Error in Converting data", err)
		return err
	}
	fmt.Println("Data is Added from File")
	return nil

}

func jsonDemo(sm proto.Message) {

	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simple.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)

}

func toJSON(pb proto.Message) string {

	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out

}

func fromJSON(in string, pb proto.Message) {

	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}

}

func doComplex() {

	cm := complex.ComplexMessage{
		OneDummy: &complex.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complex.DummyMessage{
			&complex.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complex.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}
	fmt.Println(cm)

}

func doEnum() {

	em := enum.EnumMessage{
		Id:           42,
		DayOfTheWeek: enum.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enum.DayOfTheWeek_MONDAY
	fmt.Println(em)

}
