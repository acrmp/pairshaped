package pair_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPair(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pair Suite")
}
