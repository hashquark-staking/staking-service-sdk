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
	sign, err := stakingService.sign("1234qwer1234qwer", 123218932179378219)
	fmt.Println(sign, err)
}
