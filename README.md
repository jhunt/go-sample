go-sample - A Sampling Statistics Library for Go
================================================

This library implements some statistical sampling under a
reservoir model, whereby a streamed population of measurements can
be analyzed and summarized in constant space.

It leverages two space-saving measures: _reservoir sampling_ and
_Welford's Algorithm_.  The latter calculates the standard
deviation of a stream of inputs, without requiring perpetual
access to the entire series.  The former allows random
sub-sampling of the same stream of inputs, for statistical
summaries that require a full set (i.e. median).

    package main

    import (
      "fmt"
      "math/rand"

      "github.com/jhunt/go-sample"
    )

    func main() {
      r := sample.NewReservoir(100)
      for i := 0; i < 2000; i++ {
        r.Sample(rand.Float64())
      }
      fmt.Printf("minimum is %v\n", r.Minimum())
      fmt.Printf("maximum is %v\n", r.Maximum())
      fmt.Printf(" median is %v\n", r.Median())
      fmt.Printf("st. dev is %v\n", r.Stdev())
    }

(Note: `rand.Float64()` returns a value in the range `[0.0,1.0)`)

The above program outputs something like:

    minimum is 0.0005138155161213613
    maximum is 0.9999968116565763
     median is 0.4810159213036518
    st. dev is 0.28704497179947713

Happy Sampling!
