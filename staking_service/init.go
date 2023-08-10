package staking_service

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/crypto"
)

type StakingService struct {
	baseURL    string
	privateKey *ecdsa.PrivateKey
	address    string
	c          *req.Client
}

func NewStakingService(baseURL, privateKey, address string) (*StakingService, error) {
	var err error
	stakingService := new(StakingService)
	stakingService.baseURL = baseURL
	stakingService.privateKey, err = crypto.HexToECDSA(privateKey)
	stakingService.address = address
	stakingService.c = req.C()
	stakingService.c.SetBaseURL(baseURL)

	return stakingService, err
}

func (stakingService *StakingService) getBaseRequest() *req.Request {
	u, _ := uuid.NewV4()
	nonce := u.String()
	timestamp := time.Now().Unix()
	signature, _ := stakingService.sign(nonce, timestamp)
	authorization := fmt.Sprintf("Quark-Keccak256-ECDSA Address=%s,Nonce=%s,Timestamp=%d,Signature=%s",
		stakingService.address, nonce, timestamp, signature)
	//Quark-EIP-1911 Address=0x1234567890abcdef1234567890abcdef12345678,Nonce=1cb89e48-5d71-43e6-ba61-4a810b879e5e,Timestamp=1686048892,Signature=0x123456789
	r := stakingService.c.R()
	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("Accept", "application/json")
	r.SetHeader("Authorization", authorization)
	return r
}

func processResponse[T any](res *req.Response) (result *T, code uint, msg string, err error) {
	log.Info().Str("content", res.String()).Msg("Response Content")
	response := new(Response[T])
	err = res.UnmarshalJson(response)
	if err != nil {
		return
	}
	code = response.Code
	msg = response.Msg
	if response.Code == 0 {
		result = &response.Data
	}
	return
}
