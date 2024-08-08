package staking_service

import (
	"fmt"
	"log"
	"testing"
)

func TestAllDelegation(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	chainName := "aptos"
	validatorID := "default"
	delegatorID := "default"
	pageParams := PageParams{
		PageNum:  1,
		PageSize: 10,
	}

	// result1, code, msg, err := stakingService.ListDelegatedValidators(chainName)
	// fmt.Println(result1, code, msg, err)
	// result2, code, msg, err := stakingService.GetDelegatedValidatorInfo(chainName, validatorID)
	// fmt.Println(result2, code, msg, err)

	// result3, code, msg, err := stakingService.ListDelegatesForValidator(chainName, validatorID, pageParams)
	// fmt.Println(result3, code, msg, err)

	// result4, code, msg, err := stakingService.ListRewardsForValidator(chainName, validatorID, pageParams)
	// fmt.Println(result4, code, msg, err)

	result5, code, msg, err := stakingService.ListDelegators(chainName, validatorID)
	fmt.Println(result5, code, msg, err)

	result6, code, msg, err := stakingService.GetDelegatorInfo(chainName, validatorID, delegatorID)
	fmt.Println(result6, code, msg, err)

	result7, code, msg, err := stakingService.GetDelegatorOverview(chainName, validatorID, delegatorID)
	fmt.Println(result7, code, msg, err)

	result8, code, msg, err := stakingService.ListDelegatesForDelegator(chainName, validatorID, delegatorID, pageParams)
	fmt.Println(result8, code, msg, err)

	result9, code, msg, err := stakingService.ListRewardsForDelegator(chainName, validatorID, delegatorID, pageParams)
	fmt.Println(result9, code, msg, err)
}
