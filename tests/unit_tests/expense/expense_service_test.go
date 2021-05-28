package expense_tests

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	mockRepository "github.com/wallet-tracky/Golang-backend/mock"
	"github.com/wallet-tracky/Golang-backend/model"
	"github.com/wallet-tracky/Golang-backend/service"
)

var _ = Describe("Service:- Get expense Service", func() {

	mockController := gomock.NewController(GinkgoT())

	defer mockController.Finish()

	Context("Positive scenario", func() {
		It("Should not return error if valid inputs are given", func() {

			mockRepository := mockRepository.NewMockExpenseRepository(mockController)

			expenseModel := make([]model.Expense, 2)
			mockRepository.EXPECT().FindByUserId(1).Return(&expenseModel)

			expenseService := service.NewExpenseService(mockRepository)

			newExpense := expenseService.FindAllExpenseOfUser(1)

			gomega.Expect(len(newExpense)).To(gomega.Equal(2)) //length of model and response list are same
		})
	})

})

