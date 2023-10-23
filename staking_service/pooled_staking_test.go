package staking_service

import (
	"fmt"
	"log"
	"testing"
)

func TestGetPoolRewards(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	result, code, msg, err := stakingService.GetPoolRewards(pointer("20230918"), pointer(uint(10)))
	fmt.Println(result, code, msg, err)
}

func TestGetUserStakingInfo(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	result, code, msg, err := stakingService.GetUserPooledStakingInfo("0x6210438324d8F8bdA3d3413a6DF533efDB505036")
	fmt.Println(result, code, msg, err)
}

func TestGetPooledStakingInfo(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	result, code, msg, err := stakingService.GetPooledStakingInfo()
	fmt.Println(result, code, msg, err)
}

func TestGetWithdrawRequestInfo(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	result, code, msg, err := stakingService.GetWithdrawRequestInfo(10)
	fmt.Println(result, code, msg, err)
}

func TestGetUserWithdrawPossible(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	result, code, msg, err := stakingService.GetUserWithdrawPossible("0x91bC9Eefca5BdF2dB609879AeDc579c70a0Ae901")
	fmt.Println(result, code, msg, err)
}

func TestGetTransactionHistory(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	result, code, msg, err := stakingService.GetTransactionHistory("0x91bC9Eefca5BdF2dB609879AeDc579c70a0Ae901", 0, 1, 3)
	fmt.Println(result, code, msg, err)
	for _, transaction := range result.List {
		fmt.Println("transaction:", transaction)
	}
}

func pointer[T any](s T) *T {
	return &s
}
