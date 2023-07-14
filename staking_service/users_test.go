package staking_service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserProcess(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	createResult, code, msg, err := stakingService.CreateUser("1")
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, createResult)
	require.Equal(t, "ok", msg)

	listResult, code, msg, err := stakingService.ListUsers(1, 10)
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, listResult)
	require.NotEmpty(t, listResult.List)
	require.Equal(t, "ok", msg)

	detailsResult, code, msg, err := stakingService.GetUserDetails("1")
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, detailsResult)
	require.Equal(t, "ok", msg)

	// deleteResult, code, msg, err := stakingService.DeleteUser("1")
	// require.Nil(t, err)
	// require.Equal(t, uint(0), code)
	// require.NotNil(t, deleteResult)
	// require.Equal(t, "ok", msg)
}
