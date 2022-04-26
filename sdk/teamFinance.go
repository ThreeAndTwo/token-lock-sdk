package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/deng00/req"
	"math/big"
)

// https://www.team.finance/
// not yet open source For contract
// no valid ABI file found not yet

const TeamFinanceURL = `https://team-finance-backend-origdfl2wq-uc.a.run.app/api/app/explorer/search?`

type teamFinance struct {
	chain ChainName
	token string
}

func newTeamFinance(_chain ChainName, _token string) (*teamFinance, error) {
	return &teamFinance{chain: _chain, token: _token}, nil
}

func (tf *teamFinance) GetLockedCountForToken() (*big.Int, error) {
	return nil, fmt.Errorf("unSupport")
}

func (tf *teamFinance) GetTokenLocks(lockId *big.Int) (*TokenLockerInfo, error) {
	url := fmt.Sprintf("%sinput=%s&skip=0&limit=15&order=4", TeamFinanceURL, tf.token)
	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	data := &TeamFinanceData{}
	err = json.Unmarshal(resp.Bytes(), data)
	if err != nil {
		return nil, err
	}

	if data.Data.TotalCount == 0 {
		return nil, err
	}

	info := &TokenLockerInfo{
		Platform:          PlatformTeamFinance,
		Owner:             data.Data.PagedData[0].Event.SenderAddress,
		TokenAddress:      tf.token,
		TokenLockedAmount: float2BigInt(data.Data.PagedData[0].Event.LockAmount),
		TokenLockedRate:   data.Data.PagedData[0].Token.LiquidityLockedInPercent,
		TokenLockedTime:   fmtTS2Date(data.Data.PagedData[0].Event.TimeStamp),
		TokenUnlockTime:   fmtTS2Date(data.Data.PagedData[0].Event.UnlockTime),
	}

	if data.Data.PagedData[0].Token.IsLiquidityToken {
		// LP token
		info.TokenAddress = data.Data.PagedData[0].Event.LiquidityContractAddress
		info.TokenLockedUsd = data.Data.PagedData[0].Token.LiquidityLockedInUsd
	} else {
		// classic token
		info.TokenAddress = data.Data.PagedData[0].Token.TokenAddress
		info.TokenLockedUsd = data.Data.PagedData[0].Token.TokenLockedInUsd
	}
	return info, nil
}
