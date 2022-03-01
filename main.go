package main

import "fmt"

func main() {
	heroList, err := GetHeroList()
	if err != nil {
		panic(err)
	}
	if err := WriteHeroList(heroList); err != nil {
		panic(fmt.Errorf("Write Hero List Error: %s", err))
	}

	itemList, err := GetItemList()
	if err != nil {
		panic(err)
	}
	if err := WriteItemList(itemList); err != nil {
		panic(fmt.Errorf("Write Hero List Error: %s", err))
	}
}
