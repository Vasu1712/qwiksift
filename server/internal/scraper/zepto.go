package scraper

import (
    "fmt"
    "strings"
    "strconv"

    "github.com/gocolly/colly/v2"
    "github.com/Vasu1712/qwiksift/server/internal/models"
)

type ZeptoScraper struct {
    collector *colly.Collector
}

func NewZeptoScraper() *ZeptoScraper {
    c := colly.NewCollector(
        colly.UserAgent("Mozilla/5.0 (compatible; Qwiksift/1.0 +https://qwiksift.com)"),
    )
    
    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        Parallelism: 1,
        Delay:       5 * time.Second,
    })

    return &ZeptoScraper{
        collector: c,
    }
}

func (s *ZeptoScraper) Scrape() ([]models.Product, error) {
    var products []models.Product
    baseURL := "https://www.zeptonow.com"

    s.collector.OnHTML("div.product-item", func(e *colly.HTMLElement) {
        priceStr := strings.TrimSpace(e.ChildText("div.price-box"))
        price, _ := strconv.ParseFloat(strings.ReplaceAll(priceStr, "₹", ""), 64)

        product := models.Product{
            Name:     strings.TrimSpace(e.ChildText("h2.product-name")),
            Price:    price,
            ImageURL: e.ChildAttr("img", "data-src"),
            Provider: "Zepto",
            URL:      baseURL + e.ChildAttr("a", "href"),
        }
        
        products = append(products, product)
    })

    s.collector.OnError(func(r *colly.Response, err error) {
        fmt.Printf("Zepto scraper error: %v\n", err)
    })

    err := s.collector.Visit(baseURL + "/products")
    if err != nil {
        return nil, err
    }

    return products, nil
}

func (s *ZeptoScraper) Platform() string {
    return "zepto"
}
