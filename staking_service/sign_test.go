package staking_service

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestSign(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot initial staking service")
	}
	nonce := "0eaff6d4-3e67-48a4-96e6-bbe96c349a33"
	now := int64(1719822308)
	// nonce := uuid.NewString()
	// now := time.Now().Unix()
	sign, err := stakingService.sign(nonce, now)
	fmt.Println("PK: ", PRIVATE_KEY)
	fmt.Println("Address: ", ADDRESS)
	fmt.Println("Nonce: ", nonce)
	fmt.Println("Timestamp: ", now)
	fmt.Println("Sign: ", sign)
}
