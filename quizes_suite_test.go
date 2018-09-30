package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQuizes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Quizes Suite")
}
