package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/mozillazg/request"
)

// URLItemList url
const URLItemList = "https://www.dota2.com.cn/itemscategory/json"

// Item struct
type Item struct {
	ID     int    `json:"id"`
	NameSC string `json:"name_sc"`
	NameEN string `json:"name_en"`
	Cost   int    `json:"cost"`
}

// GetID func
func (i Item) GetID() int {
	return i.ID
}

// ItemList struct
type ItemList struct {
	List []Item `json:"item_list"`
}

// GetItemList func
func GetItemList() (ItemList, error) {
	var list ItemList
	resp, err := request.Get(URLItemList, nil)
	if err != nil {
		return list, err
	}
	j, err := resp.Json()
	if err != nil {
		return list, err
	}
	if j.Get("status").MustString() != "success" {
		return list, fmt.Errorf("Item List Error")
	}
	categories := []string{"basic", "upgrade", "neutral"}

	for _, category := range categories {
		for i := range j.Get("result").Get(category).MustArray() {
			for k := range j.Get("result").Get(category).GetIndex(i).Get("items").MustArray() {
				var item Item
				item.ID, _ = strconv.Atoi(j.Get("result").Get(category).GetIndex(i).Get("items").GetIndex(k).Get("item_id").MustString())
				item.NameEN = j.Get("result").Get(category).GetIndex(i).Get("items").GetIndex(k).Get("name").MustString()
				item.NameSC = j.Get("result").Get(category).GetIndex(i).Get("items").GetIndex(k).Get("name_loc").MustString()
				item.Cost, _ = strconv.Atoi(j.Get("result").Get(category).GetIndex(i).Get("items").GetIndex(k).Get("cost").MustString())
				list.List = append(list.List, item)
			}
		}
	}

	SortList(list.List)

	return list, nil
}

// WriteItemList func
func WriteItemList(list ItemList) error {
	file, err := os.OpenFile("data/item.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")
	return enc.Encode(list)
}
