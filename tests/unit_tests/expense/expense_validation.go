package expense_tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/middlewares/validator"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)



var _ = Describe("Add expense validation", func() {

	expenseValidator := &validator.ExpenseValidator{}

	Context("Positive Scenario", func() {
		It("Should not throw error for valid inputs", func() {
			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request=&http.Request{}

			expenseRequest := request.Expense{
				Description: "Amazon purchase for mobile",
				Amount: 1,
				SpendFrom: "axis_2635",
				Date: "30-05-1997",
				Category: "LIFE STYLE",
				AdditionalNotes: "Bought from amazon for rachels",
				Image: "may be later",
				Tag: "ONLINE",
				IsCounted: true,
				AddedAs: "manual",
				IsReviewed: false,
			}



			requestBody,_:= json.Marshal(expenseRequest)

			context.Request.Body=ioutil.NopCloser(bytes.NewBuffer(requestBody))



			expenseValidator.ValidateAddExpenseRequest(context)
		})
	})

	Context("Negative Scenario", func() {
		It("Should not return 200 if amount is less that 0", func() {
			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request=&http.Request{}

			expenseRequest := request.Expense{
				Description: "Amazon purchase for mobile",
				Amount: -1,
				SpendFrom: "axis_2635",
				Date: "30-05-1997",
				Category: "LIFE STYLE",
				AdditionalNotes: "Bought from amazon for rachel's birday",
				Image: "may be later",
				Tag: "ONLINE",
				IsCounted: true,
				AddedAs: "manual",
				IsReviewed: false,
			}

			toBytes,_:= json.Marshal(expenseRequest)

			context.Request.Body=ioutil.NopCloser(bytes.NewBuffer(toBytes))

			expenseValidator.ValidateAddExpenseRequest(context)

			gomega.Expect(recorder.Code).ToNot(gomega.Equal(200))
		})
	})
})

var _=Describe("Get expense Validator", func() {
	gin.SetMode(gin.TestMode)
	Context("Positive scenario", func() {
		It("Should not return error on valid inputs", func() {


			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)



			context.Request=&http.Request{}

			context.Request.Header=make(map[string][]string)

			context.Request.Header.Add("userId","1")

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(200))

		})
	})

	Context("Negative Scenarios", func() {
		It("Should return 400(Bad request) when userId is not present", func() {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)


			context.Request=&http.Request{}

			context.Request.Header=make(map[string][]string)

			//context.Request.Header.Add("userId","a") should not header userId to header

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(400))

		})

		It("Should return 400(Bad request) when userId value type is not compatible", func() {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)


			context.Request=&http.Request{}

			context.Request.Header=make(map[string][]string)

			context.Request.Header.Add("userId","a")

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(400))

		})
	})
})
