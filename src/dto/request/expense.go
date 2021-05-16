package request

type Expense struct {
	Description     string  `json:"description" binding:"required,min=2,max=100"`
	Amount          float64 `json:"amount" binding:"required,gt=0"`
	SpendFrom       string  `json:"spendFrom" binding:"required"`
	Date            string  `json:"date" binding:"required"`
	Category        string  `json:"category" binding:"required"`
	AdditionalNotes string  `json:"additionalNote" binding:"max=150"`
	Image           string  `json:"image"`
	Tag             string  `json:"tag" binding:"max=20"`
	IsCounted       bool    `json:"isCounted"`
	AddedAs         string  `json:"addedAs" binding:"required"`
	IsReviewed      bool    `json:"isReviewed"`
}
