package parser

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Node struct {
	BrowseNodeId   string    `xml:"browseNodeId"`
	BrowseNodeName []string  `xml:"browseNodeName"`
	HasChildren    bool      `xml:"hasChildren"`
	ChildNodes     ChildNode `xml:"childNodes"`
	Nodes          []Node    `xml:",any"`
}

type ChildNode struct {
	Count int      `xml:"count,attr"`
	Id    []string `xml:"id"`
}

func ReadFile(pathToBTG string) []byte {
	file, err := os.Open(pathToBTG)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	filesize := fileInfo.Size()
	buffer := make([]byte, filesize)

	bytesRead, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("bytes read: ", bytesRead)
	return buffer
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type node Node

	return d.DecodeElement((*node)(n), &start)
}
