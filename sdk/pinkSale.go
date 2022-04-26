package sdk

import (
	"fmt"
	pink "github.com/ThreeAndTwo/token-lock-sdk/contracts/pinkSale"
	"github.com/deng00/ethutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type pinkSale struct {
	client *pink.Pink
	token  string
	lockID *big.Int
}

func (p *pinkSale) GetLockedCountForToken() (*big.Int, error) {
	opts := ethutils.MakeCallOpts(common.Address{})
	count, err := p.client.TotalLockCountForToken(opts, common.HexToAddress(p.token))
	if err != nil {
		return nil, err
	}

	p.lockID = count
	return count, err
}

func (p *pinkSale) GetTokenLocks(lockID *big.Int) (*TokenLockerInfo, error) {
	opts := ethutils.MakeCallOpts(common.Address{})
	lockParams, err := p.client.GetLocksForToken(opts, common.HexToAddress(p.token), big.NewInt(0), p.lockID)
	if err != nil {
		return nil, err
	}

	info, err := p.client.CumulativeLockInfo(opts, common.HexToAddress(p.token))
	if err != nil {
		return nil, err
	}

	lockInfo := &TokenLockerInfo{
		Platform:          PlatformPinkSale,
		TokenAddress:      p.token,
		TokenLockedAmount: info.Amount,
		TokenLockedUsd:    0, // should be calc
		TokenLockedRate:   0, // should be calc
	}

	lockTime := int64(0)
	unlockTime := int64(0)
	for k, param := range lockParams {
		if k == 0 {
			lockInfo.Owner = param.Owner.Hex()
			lockTime = param.LockDate.Int64()
		}

		if param.UnlockDate.Int64() > unlockTime {
			unlockTime = param.UnlockDate.Int64()
		}
	}

	lockInfo.TokenLockedTime = fmtTS2Date(lockTime)
	lockInfo.TokenUnlockTime = fmtTS2Date(unlockTime)
	return lockInfo, nil
}

func newPinkSale(_chain ChainName, _client *ethclient.Client, _token string) (*pinkSale, error) {
	contract := common.Address{}
	switch _chain {
	case BSC:
		contract = common.HexToAddress(string(PinkSaleBSC))
	default:
		return nil, fmt.Errorf("unSupport Chain for %s by pinkSale", _chain)
	}

	_pink, err := pink.NewPink(contract, _client)
	if err != nil {
		return nil, fmt.Errorf("new pinkSale client error: %s", err)
	}
	return &pinkSale{client: _pink, token: _token}, nil
}
