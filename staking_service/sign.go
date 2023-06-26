package staking_service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (stakingService *StakingService) sign(nonce string, timestamp int64) (string, error) {
	bigTime := big.NewInt(timestamp)

	hash := crypto.Keccak256Hash(
		common.HexToAddress(nonce).Bytes(),
		common.BigToHash(bigTime).Bytes(),
	)
	sign, err := crypto.Sign(hash.Bytes(), stakingService.privateKey)
	if err != nil {
		return "", err
	}
	return common.Bytes2Hex(sign), nil
}
