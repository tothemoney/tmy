package clique

import (
	"math/big"
	"time"
)

var (
	SecondsPerYear uint  = 31_536_000
	NetworkStarted int64 = 1648562363
	Years                = 10
	//Budget               = *big.NewInt(2_000_000e18)

	YearAmount = map[uint]big.Int{
		1:  *big.NewInt(0).Mul(big.NewInt(4e18), big.NewInt(10000)),
		2:  *big.NewInt(0).Mul(big.NewInt(8e18), big.NewInt(10000)),
		3:  *big.NewInt(0).Mul(big.NewInt(1.2e18), big.NewInt(100000)),
		4:  *big.NewInt(0).Mul(big.NewInt(1.6e18), big.NewInt(100000)),
		5:  *big.NewInt(0).Mul(big.NewInt(2e18), big.NewInt(100000)),
		6:  *big.NewInt(0).Mul(big.NewInt(2.4e18), big.NewInt(100000)),
		7:  *big.NewInt(0).Mul(big.NewInt(2.6e18), big.NewInt(100000)),
		8:  *big.NewInt(0).Mul(big.NewInt(3e18), big.NewInt(100000)),
		9:  *big.NewInt(0).Mul(big.NewInt(3e18), big.NewInt(100000)),
		10: *big.NewInt(0).Mul(big.NewInt(3e18), big.NewInt(100000)),
	}
)

type Emission struct {
	YearAmount big.Int
	Period     uint
}

func NewEmission(period uint) (Emission, error) {
	now := time.Now() // current local time
	sec := now.Unix()
	NumberOfYear := (sec - NetworkStarted + int64(SecondsPerYear)) / int64(SecondsPerYear)
	return Emission{
		YearAmount: YearAmount[uint(NumberOfYear)],
		Period:     period,
	}, nil
}

func (e *Emission) BlockReward() *big.Int {
	blocks := e.blocksPerYear()

	blockReward := big.NewInt(0)

	blockReward.Div(&(e.YearAmount), big.NewInt(int64(blocks)))

	return blockReward
}

func (e *Emission) blocksPerYear() uint {
	return SecondsPerYear / e.Period
}

func getBlockReward(period uint) (*big.Int, error) {

	now := time.Now() // current local time
	sec := now.Unix()
	NumberOfYear := (sec - NetworkStarted + int64(SecondsPerYear)) / int64(SecondsPerYear)
	if NumberOfYear > 10 {
		NumberOfYear = 10
	}
	ya := YearAmount[uint(NumberOfYear)]
	blocks := SecondsPerYear / period

	blockReward := big.NewInt(0)

	blockReward.Div(&(ya), big.NewInt(int64(blocks)))

	return blockReward, nil
}
