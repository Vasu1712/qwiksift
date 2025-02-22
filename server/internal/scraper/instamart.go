package scraper

import (
    "fmt"
    "strings"
    "strconv"

    "github.com/gocolly/colly/v2"
    "github.com/Vasu1712/qwiksift/server/internal/models"
)

type InstamartScraper struct {
    collector *colly.Collector
}

func NewInstamartScraper() *InstamartScraper {
    c := colly.NewCollector(
        colly.UserAgent("Mozilla/5.0 (compatible; Qwiksift/1.0 +https://qwiksift.com)"),
    )
    
    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        Parallelism: 1,
        Delay:       5 * time.Second,
    })

    return &InstamartScraper{
        collector: c,
    }
}

func (s *InstamartScraper) Scrape() ([]models.Product, error) {
    var products []models.Product
    baseURL := "https://www.swiggy.com"

    s.collector.OnHTML("div.product-item", func(e *colly.HTMLElement) {
        priceStr := strings.TrimSpace(e.ChildText("div.price"))
        price, _ := strconv.ParseFloat(strings.ReplaceAll(priceStr, "₹", ""), 64)

        product := models.Product{
            Name:     strings.TrimSpace(e.ChildText("h3.name")),
            Price:    price,
            ImageURL: e.ChildAttr("img", "src"),
            Provider: "Instamart",
            URL:      baseURL + e.ChildAttr("a", "href"),
        }
        
        products = append(products, product)
    })

    s.collector.OnError(func(r *colly.Response, err error) {
        fmt.Printf("Instamart scraper error: %v\n", err)
    })

    err := s.collector.Visit(baseURL + "/instamart")
    if err != nil {
        return nil, err
    }

    return products, nil
}

func (s *InstamartScraper) Platform() string {
    return "instamart"
}
