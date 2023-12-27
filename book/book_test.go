package book_test

import (
	"github.com/MobarakHsn/Ginkgo/book"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Book.isCheap()", func() {
	Context("When is the book is cheap", func() {
		It("returns true", func() {
			book := book.Book{Price: 499}
			response := book.IsCheap()
			Expect(response).To(BeTrue())
		})
	})

	Context("When is the book is not cheap", func() {
		AfterEach(func() {
			log.Println("Okay, not cheap")
		})
		BeforeEach(func() {
			log.Println("Not cheap")
		})

		It("returns false", func() {
			book := book.Book{Price: 503}
			response := book.IsCheap()
			Expect(response).To(BeFalse())
		})

	})

	DescribeTable("IsCheap table test",
		func(price int, expectedResponse bool) {
			b := book.Book{Price: price}
			Expect(b.IsCheap()).To(Equal(expectedResponse))
		},
		Entry("When a book is cheap", 500, true),
		Entry("When a book is not cheap", 555, false))
})
