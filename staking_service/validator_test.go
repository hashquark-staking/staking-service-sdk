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
	pubkey := "0xada97cb94e155a65757451f91278b151ecfbdba0d20b45e01bbb69915f5c3774aeff97bc6ebdfce9466204c3fca67a56"
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
