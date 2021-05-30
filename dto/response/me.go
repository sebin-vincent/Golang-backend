package response

type Me struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	AnnualIncome int    `json:"annualIncome"`
	IsEnabled    bool   `json:"IsEnabled"`
}
