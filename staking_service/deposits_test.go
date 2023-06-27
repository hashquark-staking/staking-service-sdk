package staking_service

import (
	"fmt"
	"log"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/require"
)

func TestDepositProcess(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	listResult, code, msg, err := stakingService.ListDeposits(1, 10)
	fmt.Println(listResult, code, msg, err)

	depositResult, code, msg, err := stakingService.GetDepositDetails(listResult.List[0].TxHash)
	fmt.Println(depositResult, code, msg, err)
	pp.Println(depositResult)

	createResult, code, msg, err := stakingService.CreateUser("deposit_test", "test1@example.com", "0x12345678")
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, createResult)
	require.Equal(t, "ok", msg)

	assignmentResult, code, msg, err := stakingService.AssignDepositedValidators(depositResult.TxHash, ValidatorAssignmentParams{
		Assignments: []Assignment{
			{
				UID:            "deposit_test",
				ValidatorCount: len(depositResult.Validators),
			},
		},
	})
	fmt.Println(assignmentResult, code, msg, err)
	pp.Println(assignmentResult)

	detailsResult, code, msg, err := stakingService.GetUserDetails("deposit_test")
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, detailsResult)
	require.Len(t, detailsResult.Validators, len(depositResult.Validators))
	require.Equal(t, "ok", msg)

	deleteResult, code, msg, err := stakingService.DeleteUser("deposit_test")
	require.Nil(t, err)
	require.Equal(t, uint(0), code)
	require.NotNil(t, deleteResult)
	require.Equal(t, "ok", msg)
}
