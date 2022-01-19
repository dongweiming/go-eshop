package api

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

const (
	parallel = 2
	url = "https://www.metacritic.com/browse/games/release-date/available/switch/metascore?page=%s"
	ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36"
)

type MetacriticItem struct {
	ID int `json:"id"`
	Url string `json:"url"`
	Score float64 `json:"score"`
	Title string `json:"title"`
}

func GetMetacriticItems() []*MetacriticItem {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent(ua),
	)
	page := "0"
	visited := map[string]bool{
		page: true,
	}

	var items []*MetacriticItem

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: parallel})

	c.OnHTML("a.page_num", func(e *colly.HTMLElement) {
		page = e.Text
		if _, ok := visited[page]; !ok {
			c.Visit(fmt.Sprintf(url, page))
			visited[page] = true
		}
	})

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		i := &MetacriticItem{}
		id, error := strconv.Atoi(e.ChildAttr("input.clamp-summary-expand", "id"))
		if error != nil {
			return
		}
		i.ID = id
		score, _ := strconv.ParseFloat(e.DOM.Find("a.metascore_anchor > div").Eq(0).Text(), 64)
		i.Score = score
		i.Title = e.ChildText("a.title > h3")
		i.Url = e.ChildAttr("a.title", "href")
		items = append(items, i)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error: Request ", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(fmt.Sprintf(url, page))

	c.Wait()

	return items
}
