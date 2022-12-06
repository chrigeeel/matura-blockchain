package core

import (
	"math/big"
	"time"
)

func (c Chain) CalculateTarget() []byte {
	target := new(big.Int)
	target.SetBytes(InitialTarget)

	if len(c)-1 < BlocksBetweenTargetCalculation {
		return InitialTarget
	}

	time1 := time.Unix(c[len(c)-1-BlocksBetweenTargetCalculation].Header.Time, 0)
	time2 := time.Unix(c[len(c)-1].Header.Time, 0)

	timeDifference := time2.Sub(time1)

	multiplier := float64(timeDifference) / float64(TimeBetweenBlocks*BlocksBetweenTargetCalculation)

	target.Mul(target, big.NewInt(int64(multiplier*10000)))
	target.Div(target, big.NewInt(10000))

	return target.Bytes()
}
