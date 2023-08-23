package staking_service

import (
	"fmt"
	"log"
	"testing"
)

func TestValidatorProcess(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	pubkey := "0x93133d5f5c4062ef8f0c8de4637f9e64eaf88f6de9f602625301d3cda466955a8983450342cd0cfbb317c9b93b0004c8"
	params := ValidatorExitParam{
		Epoch:     0,
		Broadcast: 1,
	}
	_, code, msg, err := stakingService.ValidatorExit(pubkey, params)
	fmt.Println(err, code, msg)

	//validatorListParams := ValidatorListRequestParam{
	//	Uid:        "mkx",
	//	PageParams: PageParams{PageNum: 1, PageSize: 5},
	//}
	//list, code, msg, err := stakingService.ValidatorList(validatorListParams)
	//fmt.Println(err, code, msg)
	//for _, l := range list.List {
	//	fmt.Println("validator:", l)
	//}
}
