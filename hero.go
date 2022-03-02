package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/mozillazg/request"
)

// URLHeroList url
const URLHeroList = "https://www.dota2.com.cn/datafeed/heroList"

// Hero struct
type Hero struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NameSC string `json:"name_sc"`
	NameEN string `json:"name_en"`
	Img    string `json:"img"`
}

// HeroList struct
type HeroList struct {
	List []Hero `json:"hero_list"`
}

// GetHeroList func
func GetHeroList() (HeroList, error) {
	var list HeroList

	resp, err := request.Get(URLHeroList, nil)
	if err != nil {
		return list, err
	}
	j, err := resp.Json()
	if err != nil {
		return list, err
	}
	if j.Get("status").MustString() != "success" {
		return list, fmt.Errorf("Hero List Error")
	}

	for i := range j.Get("result").Get("heroes").MustArray() {
		var hero Hero
		hero.ID = j.Get("result").Get("heroes").GetIndex(i).Get("id").MustInt()
		hero.Name = j.Get("result").Get("heroes").GetIndex(i).Get("name").MustString()
		hero.NameSC = j.Get("result").Get("heroes").GetIndex(i).Get("name_loc").MustString()
		hero.NameEN = j.Get("result").Get("heroes").GetIndex(i).Get("name_english_loc").MustString()
		hero.Img = j.Get("result").Get("heroes").GetIndex(i).Get("crops_img").MustString()
		list.List = append(list.List, hero)
	}

	sort.Slice(list.List, func(i, j int) bool {
		return list.List[i].ID < list.List[j].ID
	})

	return list, nil
}

// WriteHeroList func
func WriteHeroList(list HeroList) error {
	file, err := os.OpenFile("data/hero.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")
	return enc.Encode(list)
}
