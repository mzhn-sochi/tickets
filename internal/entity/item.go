package entity

type Item struct {
	Product     string  `json:"product" db:"product"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Measure     struct {
		Amount float64 `json:"amount" db:"amount"`
		Unit   string  `json:"unit" db:"unit"`
	} `json:"measure"`
	Overprice uint `json:"overprice" db:"overprice"`
}
