package models

type Product struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    ImageURL    string  `json:"image_url"`
    Provider    string  `json:"provider"`
    Category    string  `json:"category"`
    Brand       string  `json:"brand"`
    URL         string  `json:"product_url"`
}
