package api

import (
	"log"

	"github.com/dongweiming/go-eshop/eshop"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

const (
	PerPage = 200
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

func Search(country int, filter string, page, per_page int) ([]AlgoliaItem, bool) {
	index, ok := eshop.AlgoliaIndexMap[country]
	if !ok {
		print("Please use the known country constant such as `eshop.US` for country param!")
		return nil, false
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

func GetAlgoliaItems(country int, filter string) []AlgoliaItem {
	var page, stop = 0, false
	var games_, games []AlgoliaItem
	games_, _ = Search(country, filter, page, eshop.PerPage)
	return games_

	for !stop {
		games_, stop = Search(country, filter, page, eshop.PerPage)
		games = append(games, games_...)
	}
	return games
}

func GetAllDealItems(country int) []AlgoliaItem {
	return GetAlgoliaItems(country, "generalFilters:Deals")
}

func GetAllItems(country int) []AlgoliaItem {
	return GetAlgoliaItems(country, "")
}
