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
	pubkey := "0x80bbab1ced9cfe9d30ba76701b2c43b41c15fa1fd31115db00de7dd58566fdf28ca8a10b0517f1d2eeb68b8719f83e4b"
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
