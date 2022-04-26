package sdk

import (
	"math/big"
	"time"
)

type Contact string

const (
	MudraBSC    Contact = "0xAe7e6CAbad8d80f0b4E1C4DDE2a5dB7201eF1252"
	PinkSaleBSC Contact = "0x7ee058420e5937496F5a2096f04caA7721cF70cc"
)

type ChainName string

const (
	ETH ChainName = "eth"
	BSC ChainName = "bsc"
)

type Platform string

const (
	PlatformTeamFinance Platform = "teamFinance"
	PlatformMudra       Platform = "mudra"
	PlatformPinkSale    Platform = "pinkSale"
)

type TokenLockerInfo struct {
	Platform          Platform `json:"platform"`
	Owner             string   `json:"owner"`
	TokenAddress      string   `json:"token_address"`
	TokenLockedAmount *big.Int `json:"token_locked_amount"`
	TokenLockedUsd    float64  `json:"token_locked_usd"`
	TokenLockedRate   float64  `json:"token_locked_rate"`
	TokenLockedTime   string   `json:"token_locked_time"`
	TokenUnlockTime   string   `json:"token_unlock_time"`
}

type TeamFinanceData struct {
	Data struct {
		PagedData []struct {
			Token struct {
				Network                  string    `json:"network"`
				ChainId                  string    `json:"chainId"`
				IsLiquidityToken         bool      `json:"isLiquidityToken"`
				TokenDecimals            int       `json:"tokenDecimals"`
				TokenName                string    `json:"tokenName"`
				TokenSymbol              string    `json:"tokenSymbol"`
				TokenTotalSupply         float64   `json:"tokenTotalSupply"`
				TokenAddress             string    `json:"tokenAddress"`
				UpdatedAt                time.Time `json:"updatedAt"`
				CreatedAt                time.Time `json:"createdAt"`
				LiquidityLockedInPercent float64   `json:"liquidityLockedInPercent"`
				LiquidityLockedInUsd     float64   `json:"liquidityLockedInUsd"`
				TokenCirculatingSupply   float64   `json:"tokenCirculatingSupply"`
				TokenLocked              float64   `json:"tokenLocked"`
				TokenImage               string    `json:"tokenImage"`
				TokenLockedInUsd         float64   `json:"tokenLockedInUsd"`
				TokenPrice               float64   `json:"tokenPrice"`
				TokenRank                uint64    `json:"tokenRank"`
			} `json:"token"`
			Event struct {
				Network                  string    `json:"network"`
				ChainId                  string    `json:"chainId"`
				IsWithdrawn              bool      `json:"isWithdrawn"`
				LockContractAddress      string    `json:"lockContractAddress"`
				LockAmount               float64   `json:"lockAmount"`
				LockDepositId            int       `json:"lockDepositId"`
				UnlockTime               int64     `json:"unlockTime"`
				WithdrawalAddress        string    `json:"withdrawalAddress"`
				TokenAddress             string    `json:"tokenAddress"`
				TokenTotalSupply         float64   `json:"tokenTotalSupply"`
				LiquidityContractAddress string    `json:"liquidityContractAddress"`
				LiquidityPairAddress     string    `json:"liquidityPairAddress"`
				LiquidityTotalSupply     float64   `json:"liquidityTotalSupply"`
				LiquidityPairReserve     float64   `json:"liquidityPairReserve"`
				LiquidityTokenReserve    float64   `json:"liquidityTokenReserve"`
				UpdatedAt                time.Time `json:"updatedAt"`
				CreatedAt                time.Time `json:"createdAt"`
				SenderAddress            string    `json:"senderAddress"`
				TimeStamp                int64     `json:"timeStamp"`
				TransactionAmount        float64   `json:"transactionAmount"`
				TransactionHash          string    `json:"transactionHash"`
				TransactionIndex         int       `json:"transactionIndex"`
			} `json:"event"`
		} `json:"pagedData"`
		TotalCount int `json:"totalCount"`
	} `json:"data"`
}
