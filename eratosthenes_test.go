package primes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEratosthenes(t *testing.T) {
	Convey("Eratosthene algorithm should work", t, func() {
		max := uint64(10000)
		s := NewEratosthenes(max)
		So(s.Len(), ShouldEqual, 1251)
		numPrimes := 0
		for i := uint64(0); i <= max; i++ {
			if s.IsPrime(i) {
				numPrimes++
				So(isPrime(i), ShouldBeTrue)
			} else {
				So(isPrime(i), ShouldBeFalse)
			}
		}
		So(numPrimes, ShouldEqual, 1229)
		ps := s.ListPrimes()
		So(len(ps), ShouldEqual, 1230)
		for _, p := range ps {
			So(isPrime(p), ShouldBeTrue)
		}
	})
}
