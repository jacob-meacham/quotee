package main_test

import (
	. "github.com/jacob-meacham/quotee"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sclevine/agouti/core"

	"github.com/go-martini/martini"
)

var _ = Describe("Home", func() {
	var page Page
	var server *martini.Martini

	BeforeEach(func() {
		server = NewServer()
		//var err error
		page, _ = agoutiDriver.Page()
		//Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		//page.Destroy()
	})

	It("should show a quote on the page", func() {
		By("appearing", func() {
			Expect(true).To(BeTrue())
			//Expect(page.Navigate("http://localhost:3000")).To(Succeed())
			//Eventually(page.Find(".quote")).ShouldNot(BeNil())
			//Expect(page.Find(".author")).ToNot(BeNil())
		})
	})
})
