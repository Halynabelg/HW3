package main

import (
	"fmt"
)

type NewWorld struct {
	Name          string
	Wake_UP_Where string
	In_backpack   string
	Decription    string
	Find          string
}

func (s NewWorld) Print() {
	fmt.Printf(s.Name, s.Wake_UP_Where, s.In_backpack, s.Decription, s.Find)
}
func main() {
	s := NewWorld{
		Name:          "Steven",
		Wake_UP_Where: "At the entrance to the cave",
		In_backpack:   "Matches, flashlight and knife",
		Decription:    "It is dark in the cave, so Steven follows the path that leads from the cave into the forest",
		Find:          "The body of a strange animal",
	}
	s.Print()

	fmt.Println("Enter Safe`s code between 0 and 20:")

	var code int

	fmt.Scan(&code)

	if code > 20 {
		fmt.Println("Error!")
	} else {
		fmt.Printf("Seems your choice was wrong on %d units!\n", 20-code)
	}
}
