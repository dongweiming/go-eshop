package api

import (
	"log"
	"strings"

	"github.com/dongweiming/go-eshop/eshop"
)

func GetAllGames(country interface{}) []AlgoliaItem {
	var (
		ok bool
		c int
	)
	switch t := country.(type) {
	case string:
		c, ok = eshop.CountryMap[strings.ToUpper(t)]
		if !ok {
			log.Fatal("Wrong country type!")
		}
	case int:
		c = t
	}
	// US/BR/MX/CA
	if _, ok = eshop.AlgoliaIndexMap[c]; ok {
		var items []AlgoliaItem
		for order, _ := range eshop.OrderByMap {
			items_ := GetAlgoliaItems(c, "", order)
			items = append(items, items_...)
		}
		return items
	}
	return nil
}
