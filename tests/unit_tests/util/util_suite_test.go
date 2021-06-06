package util_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"

)

func TestUtilityMethods(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test utility")
}
