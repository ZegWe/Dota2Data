package main

import (
	"fmt"
	dota2 "github.com/zegwe/dota2data"
)

func main() {
	heroList, err := dota2.GetHeroList()
	if err != nil {
		panic(err)
	}
	if err := dota2.WriteHeroList(heroList); err != nil {
		panic(fmt.Errorf("Write Hero List Error: %s", err))
	}

	itemList, err := dota2.GetItemList()
	if err != nil {
		panic(err)
	}
	if err := dota2.WriteItemList(itemList); err != nil {
		panic(fmt.Errorf("Write Hero List Error: %s", err))
	}
	fmt.Println("Sync Dota2 Data Succeed!")
}
