package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
)

type Distribution struct {
	ring.DistributionParameters
	Std      float64
	AbsBound float64
}

func NewDistribution(params ring.DistributionParameters, logN int) (d Distribution) {
	_ = "STUB: not implemented"
	return *new(Distribution)
}

// Sanity check
