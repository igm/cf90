package main

import (
	"errors"
	"fmt"
)

type Listable interface {
	Title() string
	Len() int
	Render(int) string
}

type Selectable interface {
	Listable
	Selection() string
}

func choose(s Selectable) (selectedIndex int, err error) {
	list(s)

	fmt.Print(s.Selection())
	_, err = fmt.Scanf("%d\n", &selectedIndex)
	selectedIndex--
	if err != nil {
		return
	}
	if selectedIndex < 0 || selectedIndex >= s.Len() {
		err = errors.New("Incorrect selection.")
	}
	return
}

func list(s Listable) {
	fmt.Printf("    %s\n", s.Title())
	itemsCount := s.Len()
	for i := 0; i < itemsCount; i++ {
		fmt.Printf("(%d) %s\n", i+1, s.Render(i))
	}
}

func enterText(label string) (input string, err error) {
	fmt.Print(label)
	_, err = fmt.Scanf("%s\n", &input)
	return
}
