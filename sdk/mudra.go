package sdk

import (
	"fmt"
	mudRa "github.com/ThreeAndTwo/token-lock-sdk/contracts/mudra"
	"github.com/deng00/ethutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// yahoo!
// not yet open source for contract
// https://mudra.website/

type mudra struct {
	client *mudRa.Mudra
	token  string
	lockID *big.Int
}

func newMudra(_chain ChainName, _client *ethclient.Client, _token string) (*mudra, error) {
	contract := common.Address{}
	switch _chain {
	case BSC:
		contract = common.HexToAddress(string(MudraBSC))
	default:
		return nil, fmt.Errorf("unSupport Chain for %s by mudra", _chain)
	}

	_mudra, err := mudRa.NewMudra(contract, _client)
	if err != nil {
		return nil, fmt.Errorf("new mudra client error: %s", err)
	}
	return &mudra{client: _mudra, token: _token}, nil
}

func (m *mudra) GetLockedCountForToken() (*big.Int, error) {
	opts := ethutils.MakeCallOpts(common.HexToAddress(""))
	nonce, err := m.client.LockNonce(opts)

	if err != nil {
		return nil, err
	}
	return nonce, fmt.Errorf("unSupport")
}

func (m *mudra) GetTokenLocks(lockID *big.Int) (*TokenLockerInfo, error) {
	opts := ethutils.MakeCallOpts(common.Address{})
	locks, err := m.client.TokenLocks(opts, lockID)

	if err != nil {
		return nil, err
	}

	lockInfo := &TokenLockerInfo{
		Platform:          PlatformMudra,
		Owner:             locks.Owner.Hex(),
		TokenAddress:      m.token,
		TokenLockedAmount: locks.TokenAmount,
		TokenLockedUsd:    0, // should be calc
		TokenLockedRate:   0, // should be calc
		TokenUnlockTime:   locks.UnlockTime.String(),
	}
	return lockInfo, nil
}
