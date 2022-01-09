package homeloan_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHomeloan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Homeloan Suite")
}
