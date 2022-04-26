package sdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SDK struct {
	chainName ChainName
	client    *ethclient.Client
	platform  Platform
	token     string
}

func NewSDK(_chainName ChainName, _rpc string, _platform Platform, _token string) *SDK {
	client, _ := ethclient.Dial(_rpc)
	return &SDK{chainName: _chainName, client: client, platform: _platform, token: _token}
}

func (s *SDK) check() bool {
	return s.chainName != ETH && s.chainName != BSC || s.token == "" || s.client == nil
}

func (s *SDK) SdkGetter() (ISDK, error) {
	if s.check() {
		return nil, fmt.Errorf("sdk config error")
	}
	switch s.platform {
	case PlatformMudra:
		return newMudra(s.chainName, s.client, s.token)
	case PlatformPinkSale:
		return newPinkSale(s.chainName, s.client, s.token)
	case PlatformTeamFinance:
		return newTeamFinance(s.chainName, s.token)
	default:
		return nil, fmt.Errorf("unSupport protocol")
	}
}
