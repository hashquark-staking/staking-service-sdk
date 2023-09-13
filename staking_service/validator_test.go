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
	pubkey := "0xb68e15ec601a6504b25c8a48cdb40861c40414fbf27d7451a09523f13e1efe267eb218d72122eb302b73e9bb93d55d44"
	params := ValidatorExitParam{
		Epoch:     0,
		Broadcast: 1,
	}
	res, code, msg, err := stakingService.ValidatorExit(pubkey, params)
	fmt.Println(res, err, code, msg)

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
