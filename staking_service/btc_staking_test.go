package staking_service

import (
	"fmt"
	"log"
	"testing"
)

func TestGenerateStakeRequest(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	params := GenerateStakingRequestParams{
		StakerPublicKeyHex: "e1aa76e3c7620afbb6740432eccd404c299a4c029271036a68b72be93c4aef7d",
		StakingAmount:      50008,
		StakingTimeBlocks:  64000,
		TxInclusionHeight:  216000,
	}
	resp, code, msg, err := stakingService.GenerateStakingRequest(params)
	fmt.Println(resp, code, msg, err)

}

func TestBroadcastStakeRequest(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	params := GenerateBroadcastStakingRequestParams{
		StakingTransactionHex: "02000000000101097fe5ddcc4ba4b2714889e3c005bc2b09f7107c84d832b7f07f580bda8d559d0000000000fdffffff0384f50c0000000000160014b139a1656453fffe20d1d3c576a35ded5753133058c30000000000002251207d0c0a07612ee5e3e653b6a8cbb881262f13f87921a146c3d1d303a55be104040000000000000000496a476262743400e1aa76e3c7620afbb6740432eccd404c299a4c029271036a68b72be93c4aef7df712efa8037fa13cf57388ab32731ff706f3028142b3800e132f77b02bab1015fa00024730440220120132157c200e266ae2e2daa3d02137e25e6ee84616fc6b2a4d885046780c7e022062b1c471ded1db9a7aae060f20f21a151d07cac8d2283a66b5531cd06a42f822012103de6a290267169edefe53be88925e7bb25c636fe5d1342c160dd8fabd3e7bf6dd00000000",
		RequestID:             "8663d7366e5c4dbd84b6d3cd6c6005f6",
	}
	resp, code, msg, err := stakingService.GenerateBroadcastStakingRequest(params)
	fmt.Println(resp, code, msg, err)

}
