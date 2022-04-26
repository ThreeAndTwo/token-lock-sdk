package sdk

import "math/big"

type ISDK interface {
	GetLockedCountForToken() (*big.Int, error)
	GetTokenLocks(lockId *big.Int) (*TokenLockerInfo, error)
}
