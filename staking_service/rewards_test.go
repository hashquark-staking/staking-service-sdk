package staking_service

import (
	"log"
	"testing"
	"time"

	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/require"
)

func TestRewards(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	startTime, _ := time.Parse(time.RFC3339, "2023-06-21T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-06-28T23:59:59Z")
	dailyResult, code, msg, err := stakingService.DailyRewards("deposit_test", startTime.UnixMilli(), endTime.UnixMilli())
	pp.Println(dailyResult)
	pp.Println(msg)
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, dailyResult)
	require.Len(t, dailyResult, 7)
	require.Equal(t, "ok", msg)

	startTime, _ = time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ = time.Parse(time.RFC3339, "2023-06-28T23:59:59Z")
	monthlyResult, code, msg, err := stakingService.MonthlyRewards("deposit_test", startTime.UnixMilli(), endTime.UnixMilli())
	pp.Println(monthlyResult)
	pp.Println(msg)
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, monthlyResult)
	require.Equal(t, "ok", msg)
}
