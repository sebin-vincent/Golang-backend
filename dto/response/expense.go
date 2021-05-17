package response

type Expense struct {
	Id              int     `json:"id"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
	SpendFrom       string  `json:"spendFrom"`
	Category        string  `json:"category"`
	Date            string  `json:"date"`
	IsCounted       bool    `json:"isCounted"`
	AdditionalNotes string  `json:"additionalNotes"`
	Tag             string  `json:"tag"`
	Image           string  `json:"image"`
	AddedAs         string  `json:"addedAs"`
	IsReviewed      bool    `json:"isReviewed"`
}
