package staking_service

import (
	"fmt"
	"log"
	"testing"
)

func TestGetUnusedKeys(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	unused_keys, code, msg, err := stakingService.GetUnusedKeys()
	fmt.Println(unused_keys, code, msg, err)

}
