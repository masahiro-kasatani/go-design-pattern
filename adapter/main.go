package main

import "fmt"

type Reader interface {
	ReadJson()
}

type XmlReader struct{}

func (x *XmlReader) readXml() {
	fmt.Println("Read XML.")
}

type XmlAdapter struct {
	xmlReader *XmlReader
}

func (x *XmlAdapter) ReadJson() {
	x.xmlReader.readXml()
	fmt.Println("Convert XML to JSON.")
}

type JsonReader struct{}

func (j *JsonReader) ReadJson() {
	fmt.Println("Read JSON.")
}

type Client struct{}

func (c *Client) ImportJson(i Reader) {
	i.ReadJson()
	fmt.Println("Client imports JSON.")
}

func main() {
	client := &Client{}
	jsonReader := &JsonReader{}

	client.ImportJson(jsonReader)

	xmlReader := &XmlReader{}
	xmlAdapter := &XmlAdapter{
		xmlReader: xmlReader,
	}

	client.ImportJson(xmlAdapter)
}
