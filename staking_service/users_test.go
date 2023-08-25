package staking_service

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestUserProcess(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	//createResult, code, msg, err := stakingService.CreateUser("mkx", []string{"0x2B3779A253dB55B98eCED3EF427992740C17db17"})
	//require.Nil(t, err)
	//require.Equal(t, uint(0), code)
	//require.NotNil(t, createResult)
	//require.Equal(t, "ok", msg)

	depositAResult, code, msg, err := stakingService.AddDepositAddress("mkx", []string{"0xD29f1f0810f3E679c2D589e5A72EFB4639B606E3"})
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, depositAResult)
	require.Equal(t, "ok", msg)

	//listResult, code, msg, err := stakingService.ListUsers(1, 10)
	//require.Nil(t, err)
	//require.Equal(t, uint(0), code)
	//require.NotNil(t, listResult)
	//require.NotEmpty(t, listResult.List)
	//require.Equal(t, "ok", msg)
	//
	//detailsResult, code, msg, err := stakingService.GetUserDetails("1")
	//require.Nil(t, err)
	//require.Equal(t, uint(0), code)
	//require.NotNil(t, detailsResult)
	//require.Equal(t, "ok", msg)

	// deleteResult, code, msg, err := stakingService.DeleteUser("1")
	// require.Nil(t, err)
	// require.Equal(t, uint(0), code)
	// require.NotNil(t, deleteResult)
	// require.Equal(t, "ok", msg)
}
