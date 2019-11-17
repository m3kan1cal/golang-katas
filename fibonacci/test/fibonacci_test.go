package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/golang-katas/fibonacci/internal"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fibonacci Suite")
}

var _ = Describe("Fibonnaci", func() {
	var fibonacci internal.Fibonacci

	DescribeTable("Calculating Fibonacci numbers",
		func(steps int, expected int) {
			Expect(fibonacci.Calculate(steps)).To(Equal(expected))
		},
		Entry("fib(0)", 0, 0),
		Entry("fib(1)", 1, 1),
		Entry("fib(2)", 2, 1),
		Entry("fib(3)", 3, 2),
		Entry("fib(4)", 4, 3),
		Entry("fib(5)", 5, 5),
	)

})
