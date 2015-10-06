package main_test

import (
  . "github.com/jacob-meacham/quotee"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/sclevine/agouti"
  "github.com/go-martini/martini"
  
  "testing"
)

func TestQuotee(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Quotee Suite")
}

var agoutiDriver *agouti.WebDriver
var server *martini.Martini

var _ = BeforeSuite(func() {
  agoutiDriver = agouti.PhantomJS()
  server = NewServer()

  Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
  agoutiDriver.Stop()
})
