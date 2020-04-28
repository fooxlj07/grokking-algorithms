package chapter3

import "fmt"

type Box struct {
	name    string
	content interface{}
}

func FindKey(box Box) {
	if k, ok := box.content.(string); ok && k == "key" {
		fmt.Println(box.name)
	}
	if b, ok := box.content.(Box); ok {
		FindKey(b)
	}
	if boxes, ok := box.content.([]Box); ok {
		for _, b := range boxes {
			FindKey(b)
		}
	}
}

func BuildBox() Box {
	boxA := Box{name: "A"}
	boxB := Box{name: "B", content: Box{name: "E"}}
	boxC := Box{name: "C", content: Box{name: "D", content: "key"}}
	boxA.content = []Box{boxB, boxC}
	return boxA
}
