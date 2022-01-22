package api

import (
	"fmt"
	"log"

	"github.com/dongweiming/go-eshop/eshop"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type AlgoliaItem struct {
	Nid string `json:"nsuid"` // Nsuid
	Title string `json:"title"`
	Url string `json:"url"`
	Desc string `json:"description"`
	Slug string `json:"slug"`
	Developers []string `json:"developers"`
	Genres []string `json:"genres"`
	Publishers []string `json:"publishers"`
	Image string `json:"horizontalHeaderImage"`
	GeneralFilters []string `json:"generalFilters"` // 占位
	ReleaseDate string `json:"releaseDateDisplay"`
	// lowestPrice/msrp 价格
}

func Search(country int, filter string, page, per_page, order int) ([]AlgoliaItem, bool) {
	index, ok := eshop.AlgoliaIndexMap[country]
	if !ok {
		print("Please use the known country constant such as `eshop.US` for country param!")
		return nil, false
	}
	if order != eshop.ORDER_FEATURE {
		index += "_" + eshop.OrderByMap[order]
	}
	client := search.NewClient(eshop.AlgoliaID, eshop.AlgoliaKey)

	params := []interface{}{
		opt.Page(page),
		opt.Facets("generalFilters,platform,availability,genres"),
		opt.HitsPerPage(per_page),
	}

	if filter != "" {
		params = append(params, opt.FacetFilterAnd(
			opt.FacetFilter("platform:Nintendo Switch"),
			opt.FacetFilter(filter),
		))
	} else {
		params = append(params, opt.FacetFilter("platform:Nintendo Switch"))
	}

	res, err := client.InitIndex(index).Search("", params...)
	if err != nil {
		log.Fatal(err)
	}

	var games []AlgoliaItem

	err = res.UnmarshalHits(&games)
	if err != nil {
		log.Fatal(err)
	}

	return games, res.NbPages <= page
}

func GetAlgoliaItems(country int, filter string, order int) []AlgoliaItem {
	var page, stop = 0, false
	var games_, games []AlgoliaItem

	for !stop {
		games_, stop = Search(country, filter, page, eshop.PerPage, order)
		games = append(games, games_...)
		page++
	}
	return games
}

func GetAllDealItems(country int, order int) []AlgoliaItem {
	return GetAlgoliaItems(country, "generalFilters:Deals", order)
}

func GetAllItems(country int, order int) []AlgoliaItem {
	return GetAlgoliaItems(country, "", order)
}
