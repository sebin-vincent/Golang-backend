package expense_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"

)

func TestExpenseService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Expense Service")
}
