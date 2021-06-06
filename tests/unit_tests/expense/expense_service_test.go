package expense_tests

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/wallet-tracky/Golang-backend/dto"
	"github.com/wallet-tracky/Golang-backend/mocks"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/service"
	"github.com/wallet-tracky/Golang-backend/util"
)

var _ = Describe("Service:- Get expense Service", func() {

	mockController := gomock.NewController(GinkgoT())

	defer mockController.Finish()

	Context("Positive scenario", func() {
		It("Should not return error if valid inputs are given", func() {

			fromdate,_:=time.Parse(util.DateLayout,"2021-05-19")
			toDate,_:=time.Parse(util.DateLayout,"2021-05-26")
			categories:=make([]string,0)
			pageable:=dto.Pageable{OffSet: 0,Limit: 5}

			mockRepository := mocks.NewMockExpenseRepository(mockController)

			expenseModel := make([]model.Expense, 2)
			mockRepository.EXPECT().FindByUserIdAndSpendDateBetween(1,
				fromdate,
				toDate.Add(time.Hour*24), //To get inclusive expense of lastday, service adds 1 day with todate
				pageable).Return(&expenseModel)

			expenseService := service.NewExpenseService(mockRepository)



			newExpense := expenseService.FindAllExpenseOfUser(1,fromdate,toDate,categories,pageable)

			gomega.Expect(len(newExpense)).To(gomega.Equal(2)) //length of model and response list are same
		})
	})

})

