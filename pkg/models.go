package handlers

type Item struct {
  Description   string  `json:"shortDescription"`
  Price         string  `json:"price"`
}

type Receipt struct {
  ID            string  `json:"id"`
  Retailer      string  `json:"retailer"`
  PurchaseDate  string  `json:"purchaseDate"`
  PurchaseTime  string  `json:"purchaseTime"`
  Items         []Item  `json:"items"`
}
