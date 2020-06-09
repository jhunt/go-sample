package sample_test

import (
	"fmt"
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"

	"github.com/jhunt/go-sample"
)

func TestSample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sample Test Suite")
}

var _ = Describe("sample", func() {
	e := 0.001
	NaN := math.NaN()
	scenario("with no samples", 0, ss(), e, NaN, NaN, NaN, NaN)
	scenario("with one sample", 1, ss(42.0), e, 42.0, 42.0, 42.0, NaN)
	scenario("with 2 samples", 2, ss(10.0, 20.0), e, 10.0, 20.0, 15.0, 7.0711)
	scenario("with 3 samples", 3, ss(10.0, 10.0, 15.0), e, 10.0, 15.0, 10.0, 2.8867)
	scenario("with negative samples", 4, ss(5.0, 6.0, -1.0, -28.0), e, -28.0, 6.0, 2.0, 15.9687)

	scenario("with an overflowing reservoir", 5,
		ss(1.0, 2.0, 1.0, 3.0, 2.0, 1.98, 2.1, 3.8, 0.2, 1.7),
		e, 0.2, 3.8, 1.0, 1.023)
})

func ss(l ...float64) []float64 {
	return l
}

func scenario(ctx string, size int, samples []float64, e, min, max, med, std float64) {
	Context(ctx, func() {
		var r sample.Reservoir

		BeforeEach(func() {
			r = sample.NewReservoir(size)
			for _, f := range samples {
				r.Sample(f)
			}
		})

		It(fmt.Sprintf("should have %d samples", size), func() {
			Ω(r.Samples()).Should(Equal(size))
		})
		It(fmt.Sprintf("should have seen %d samples", len(samples)), func() {
			Ω(r.Seen()).Should(Equal(len(samples)))
		})
		It("should calculate a minimum", func() {
			if math.IsNaN(min) {
				Ω(math.IsNaN(r.Minimum())).Should(BeTrue())
			} else {
				Ω(r.Minimum()).Should(BeNumerically("~", min, e))
			}
		})
		It("should calculate a maximum", func() {
			if math.IsNaN(max) {
				Ω(math.IsNaN(r.Maximum())).Should(BeTrue())
			} else {
				Ω(r.Maximum()).Should(BeNumerically("~", max, e))
			}
		})
		It("should calculate a median", func() {
			if math.IsNaN(med) {
				Ω(math.IsNaN(r.Median())).Should(BeTrue())
			} else {
				Ω(r.Median()).Should(BeNumerically("~", med, e))
			}
		})
		It("should calculate a standard deviation", func() {
			if math.IsNaN(std) {
				Ω(math.IsNaN(r.Stdev())).Should(BeTrue())
			} else {
				Ω(r.Stdev()).Should(BeNumerically("~", std, e))
			}
		})

		It("should reset", func() {
			r.Reset()
			Ω(r.Samples()).Should(Equal(0))
			Ω(r.Seen()).Should(Equal(0))
			Ω(math.IsNaN(r.Minimum())).Should(BeTrue())
			Ω(math.IsNaN(r.Maximum())).Should(BeTrue())
			Ω(math.IsNaN(r.Median())).Should(BeTrue())
			Ω(math.IsNaN(r.Stdev())).Should(BeTrue())
		})
	})
}
