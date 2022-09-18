package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// RootIds for each of Amazon's primary categories.
var RootIds = []string{
	"2619526011", "2617942011", "15690151", "165797011",
	"11055981", "1000", "4991426011", "493964",
	"7141124011", "7147444011", "7147443011", "7147442011",
	"7147441011", "7147440011", "2864120011", "16310211",
	"11260433011", "3760931", "1063498", "16310161", "133141011",
	"3238155011", "9479199011", "599872", "2350150011", "2625374011",
	"624868011", "301668", "11965861", "1084128",
	"541966", "2619534011", "409488", "3375301",
	"468240", "165795011", "2858778011", "10677470011",
	"11846801", "2983386011", "2335753011",
}

func SearchNextLevelById(id string, filePath string) {
	var buffer = bytes.NewBuffer(ReadFile(filePath))
	dec := xml.NewDecoder(buffer)

	var nodes Node
	err := dec.Decode(&nodes)
	if err != nil {
		panic(err)
	}

	for _, n := range nodes.Nodes {
		if id == n.BrowseNodeId {
			if n.HasChildren {
				for _, id := range n.ChildNodes.Id {
					node := GetNodeById(nodes.Nodes, id)
					if HasChildren(nodes.Nodes, id) {
						fmt.Println(node.BrowseNodeName, id, true, node.ChildNodes.Count)
					} else {
						fmt.Println(node.BrowseNodeName, id, false)
					}
				}
			}
		}
	}
}

func PrintAllNodes(filePath string) {

	var buffer = bytes.NewBuffer(ReadFile(filePath))
	dec := xml.NewDecoder(buffer)

	var nodes Node
	err := dec.Decode(&nodes)
	if err != nil {
		panic(err)
	}

	walk([]Node{nodes}, func(n Node) bool {
		for _, element := range RootIds {
			if element == n.BrowseNodeId {
				if n.HasChildren {
					CheckChildren(nodes.Nodes, n.ChildNodes.Id)
				}
			}
		}
		return true
	})
}

// Recursively walk the tree and call the callback function for each node.
func walk(nodes []Node, f func(Node) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}

// CheckChildren Recursively checks if the children have children and walks through the tree
func CheckChildren(nodes []Node, ids []string) Node {
	for _, id := range ids {
		for _, nodes := range nodes {
			if id == nodes.BrowseNodeId {
				if nodes.HasChildren {
					fmt.Println(nodes.ChildNodes.Count, nodes.BrowseNodeName, nodes.BrowseNodeId)
					fmt.Println(nodes.ChildNodes.Id)
					return CheckChildren(nodes.Nodes, nodes.ChildNodes.Id)
				} else {
					fmt.Println(nodes.BrowseNodeName, nodes.BrowseNodeId, "has no children")
					return nodes
				}
			}
		}
	}
	return Node{}
}

func HasChildren(nodes []Node, id string) bool {
	for _, n := range nodes {
		if id == n.BrowseNodeId {
			if n.HasChildren {
				return true
			}
		}
	}
	return false
}

func GetNodeById(nodes []Node, id string) Node {
	for _, n := range nodes {
		if id == n.BrowseNodeId {
			return n
		}
	}
	return Node{}
}
