package sdk

import (
	"math/big"
	"os"
	"testing"
)

func TestNewSDK(t *testing.T) {
	tests := []struct {
		name      string
		chainName ChainName
		rpc       string
		platform  Platform
		token     string
		want      bool
	}{
		{
			name:      "test teamFinance -- ETH",
			chainName: ETH,
			rpc:       os.Getenv("ETH_RPC"),
			platform:  PlatformTeamFinance,
			token:     "0xBC7250C8c3eCA1DfC1728620aF835FCa489bFdf3",
			want:      true,
		},
		{
			name:      "test teamFinance -- BSC",
			chainName: BSC,
			rpc:       os.Getenv("BSC_RPC"),
			platform:  PlatformTeamFinance,
			token:     "0x83090EB52069f5B435FC05DCE85999C985E687b2",
			want:      true,
		},
		{
			name:      "test teamFinance error",
			chainName: ETH,
			rpc:       os.Getenv("ETH_RPC"),
			platform:  PlatformTeamFinance,
			token:     "",
			want:      true,
		},
		//{
		//	name:      "test unicrypt -- ETH",
		//	chainName: ETH,
		//	rpc:       os.Getenv("ETH_RPC"),
		//	platform:  PlatformUnicryot,
		//	token:     "0x8A9dF4f6fF40b98c87d31F70B62F1aF16ab0bE2D",
		//	want:      true,
		//},
		//{
		//	name:      "test unicrypt -- BSC",
		//	chainName: BSC,
		//	rpc:       os.Getenv("BSC_RPC"),
		//	platform:  PlatformUnicryot,
		//	token:     "0x7BE7bc6716D82D87a1A8763E121D1B53c45feD43",
		//	want:      true,
		//},
		//{
		//	name:      "test unicrypt, token is null",
		//	chainName: BSC,
		//	rpc:       os.Getenv("BSC_RPC"),
		//	platform:  PlatformUnicryot,
		//	token:     "",
		//	want:      false,
		//},
		{
			name:      "test pinkSale -- ETH",
			chainName: ETH,
			rpc:       os.Getenv("ETH_RPC"),
			platform:  PlatformPinkSale,
			token:     "",
			want:      true,
		},
		{
			name:      "test pinkSale -- BSC",
			chainName: BSC,
			rpc:       os.Getenv("BSC_RPC"),
			platform:  PlatformPinkSale,
			token:     "0xf97155Cd2fF5EEF54C4Eed786C7BE8102942bD24",
			want:      true,
		},
		{
			name:      "test pinkSale -- BSC, token is null",
			chainName: BSC,
			rpc:       os.Getenv("BSC_RPC"),
			platform:  PlatformPinkSale,
			token:     "",
			want:      true,
		},
		{
			name:      "test mudra -- ETH",
			chainName: ETH,
			rpc:       os.Getenv("ETH_RPC"),
			platform:  PlatformMudra,
			token:     "",
			want:      true,
		},
		{
			name:      "test mudra -- BSC",
			chainName: BSC,
			rpc:       os.Getenv("BSC_RPC"),
			platform:  PlatformMudra,
			token:     "0xef4f58Cb8527EF70759B25CD7DE6b79C722bdc5c",
			want:      true,
		},
		{
			name:      "test mudra -- not data on mudra",
			chainName: BSC,
			rpc:       os.Getenv("BSC_RPC"),
			platform:  PlatformMudra,
			token:     "",
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := NewSDK(tt.chainName, tt.rpc, tt.platform, tt.token)
			getter, err := sdk.SdkGetter()
			if err != nil == tt.want {
				t.Fatalf("sdk getter error:%s", err)
			}

			lockID, err := getter.GetLockedCountForToken()
			if err != nil == tt.want && err.Error() != "unSupport" {
				t.Fatalf("GetLockedCountForToken error:%s", err)
			}
			locks, err := getter.GetTokenLocks(big.NewInt(lockID.Int64() - 1))
			if err != nil == tt.want {
				t.Fatalf("GetTokenLocks error:%s", err)
			}
			t.Logf("locks: %v", locks)
		})
	}
}
