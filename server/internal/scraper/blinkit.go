package scraper

import (
    "fmt"
    "strings"
    "strconv"

    "github.com/gocolly/colly/v2"
    "github.com/Vasu1712/qwiksift/server/internal/models"
)

type BlinkitScraper struct {
    collector *colly.Collector
}

func NewBlinkitScraper() *BlinkitScraper {
    c := colly.NewCollector(
        colly.UserAgent("Mozilla/5.0 (compatible; Qwiksift/1.0 +https://qwiksift.com)"),
        colly.AllowURLRevisit(),
    )
    
    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        Parallelism: 1,
        Delay:       5 * time.Second,
    })

    return &BlinkitScraper{
        collector: c,
    }
}

func (s *BlinkitScraper) Scrape() ([]models.Product, error) {
    var products []models.Product
    baseURL := "https://blinkit.com"

    s.collector.OnHTML("div.product-card", func(e *colly.HTMLElement) {
        priceStr := strings.TrimSpace(e.ChildText("div.price"))
        price, _ := strconv.ParseFloat(strings.ReplaceAll(priceStr, "₹", ""), 64)

        product := models.Product{
            Name:     strings.TrimSpace(e.ChildText("h3.title")),
            Price:    price,
            ImageURL: e.ChildAttr("img", "src"),
            Provider: "Blinkit",
            URL:      baseURL + e.ChildAttr("a", "href"),
        }
        
        products = append(products, product)
    })

    s.collector.OnError(func(r *colly.Response, err error) {
        fmt.Printf("Blinkit scraper error: %v\n", err)
    })

    err := s.collector.Visit(baseURL + "/cn/fruits-vegetables/cid/15")
    if err != nil {
        return nil, err
    }

    return products, nil
}

func (s *BlinkitScraper) Platform() string {
    return "blinkit"
}
