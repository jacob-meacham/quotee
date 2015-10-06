package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

var _ = Describe("Home", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should show a quote on the page", func() {
		By("appearing", func() {
			//Expect(true).To(BeTrue())
			Expect(page.Navigate("http://localhost:3000")).To(Succeed())
			//Eventually(page.Find(".quote")).ShouldNot(BeNil())
			//Expect(page.Find(".author")).ToNot(BeNil())
		})
	})
})
