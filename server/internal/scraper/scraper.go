package scraper

import (
    "sync"
    "time"
    
    "github.com/Vasu1712/qwiksift/server/internal/models"
)

type Scraper interface {
    Scrape() ([]models.Product, error)
    Platform() string
}

var (
    scrapers = []Scraper{
        NewBlinkitScraper(),
        NewZeptoScraper(),
        NewInstamartScraper(),
    }
    requestDelay = 2 * time.Second // Respectful scraping interval
)

func ScrapeAll() ([]models.Product, error) {
    var wg sync.WaitGroup
    results := make(chan []models.Product)
    errChan := make(chan error, len(scrapers))

    for _, s := range scrapers {
        wg.Add(1)
        go func(s Scraper) {
            defer wg.Done()
            time.Sleep(requestDelay) // Rate limiting
            
            products, err := s.Scrape()
            if err != nil {
                errChan <- err
                return
            }
            results <- products
        }(s)
    }

    go func() {
        wg.Wait()
        close(results)
        close(errChan)
    }()

    var allProducts []models.Product
    for products := range results {
        allProducts = append(allProducts, products...)
    }

    if len(errChan) > 0 {
        return nil, <-errChan
    }

    return allProducts, nil
}
