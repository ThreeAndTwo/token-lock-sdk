package sdk

import (
	"math/big"
	"time"
)

func fmtTS2Date(ts int64) string {
	return time.Unix(ts, 0).Format(time.RFC3339)
}

func float2BigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}
