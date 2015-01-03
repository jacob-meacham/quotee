package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "github.com/sclevine/agouti/core"

	"testing"
)

func TestQuotee(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Quotee Suite")
}

//var agoutiDriver WebDriver

var _ = BeforeSuite(func() {
	//var err error

	//agoutiDriver, err = PhantomJS()
	// agoutiDriver, err = Selenium()
	// agoutiDriver, err = Chrome()

	//Expect(err).NotTo(HaveOccurred())
	//Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	//agoutiDriver.Stop()
})
