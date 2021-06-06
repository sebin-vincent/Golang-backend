package expense_tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	logger "github.com/sirupsen/logrus"
	"github.com/wallet-tracky/Golang-backend/dto/request"
	"github.com/wallet-tracky/Golang-backend/middlewares/validator"
)

var _ = Describe("Add expense validation", func() {

	expenseValidator := &validator.ExpenseValidator{}

	Context("Positive Scenario", func() {
		It("Should not throw error for valid inputs", func() {
			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request = &http.Request{}

			expenseRequest := request.Expense{
				Description:     "Amazon purchase for mobile",
				Amount:          1,
				SpendFrom:       "axis_2635",
				Date:            "1997-05-30",
				Category:        "LIFE STYLE",
				AdditionalNotes: "Bought from amazon for rachels",
				Image:           "may be later",
				Tag:             "ONLINE",
				IsCounted:       true,
				AddedAs:         "manual",
				IsReviewed:      false,
			}

			requestBody, _ := json.Marshal(expenseRequest)

			context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

			expenseValidator.ValidateAddExpenseRequest(context)
		})
	})

	Context("Negative Scenario", func() {
		It("Should not return 200 if amount is less that 0", func() {
			gin.SetMode(gin.TestMode)

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request = &http.Request{}

			expenseRequest := request.Expense{
				Description:     "Amazon purchase for mobile",
				Amount:          -1,
				SpendFrom:       "axis_2635",
				Date:            "1997-05-30",
				Category:        "LIFE STYLE",
				AdditionalNotes: "Bought from amazon for rachel's birday",
				Image:           "may be later",
				Tag:             "ONLINE",
				IsCounted:       true,
				AddedAs:         "manual",
				IsReviewed:      false,
			}

			toBytes, _ := json.Marshal(expenseRequest)

			context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(toBytes))

			expenseValidator.ValidateAddExpenseRequest(context)

			gomega.Expect(recorder.Code).ToNot(gomega.Equal(200))
		})
	})
})

var _ = Describe("Get expense Validator", func() {
	gin.SetMode(gin.TestMode)
	Context("Positive scenario", func() {
		It("Should return 200 on valid inputs", func() {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			context.Request = &http.Request{}

			queries := make(url.Values)
			queries.Add("offset", "0")
			queries.Add("limit", "5")
			queries.Add("from", "1997-05-30")
			queries.Add("to", "1997-06-21")

			context.Request.Header = make(map[string][]string)

			context.Request.Header.Add("userId", "1")

			context.Request.URL = &url.URL{RawQuery: queries.Encode()}

			logger.Info("Queries :", context.Request.URL)

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(200))

		})
	})

	Context("Negative Scenarios", func() {
		It("Should return status 400(Bad request) when userId is not present", func() {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			context.Request = &http.Request{}

			context.Request.Header = make(map[string][]string)

			context.Request.URL = &url.URL{}

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(400))

		})

		It("Should return 400(Bad request) when userId value type is not compatible", func() {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			context.Request = &http.Request{}

			context.Request.Header = make(map[string][]string)

			context.Request.URL = &url.URL{}

			context.Request.Header.Add("userId", "a")

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(400))

		})

		It("Should return status 400 if 'from' date is not parsable", func() {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request = &http.Request{}
			context.Request.Header = make(http.Header)

			queryParams := make(url.Values)

			queryParams.Add("offset", "0")
			queryParams.Add("limit", "5")
			queryParams.Add("from", "30-05-1997") //Incorrect formate for date
			queryParams.Add("to", "1997-05-30")   //correct

			context.Request.URL = &url.URL{RawQuery: queryParams.Encode()}

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(400))

		})


		It("Should return status 400 if limit is not posetive integer", func() {

			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Request = &http.Request{}
			context.Request.Header = make(http.Header)

			queryParams := make(url.Values)

			queryParams.Add("offset", "0")
			queryParams.Add("limit", "-5")
			queryParams.Add("from", "1997-05-23") //Incorrect formate for date
			queryParams.Add("to", "1997-05-30")   //correct

			context.Request.URL = &url.URL{RawQuery: queryParams.Encode()}

			validator.ExpenseValidator{}.ValidateGetExpensesOfUser(context)

			gomega.Expect(recorder.Code).To(gomega.Equal(400))

		})
	})
})
