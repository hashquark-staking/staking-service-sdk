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
		StakerPublicKeyHex: "",
		StakingAmount:      100000,
		StakingTimeBlocks:  1024,
		TxInclusionHeight:  0,
	}
	resp, code, msg, err := stakingService.GenerateStakingRequest(params)
	fmt.Println(resp, code, msg, err)

}
