package staking_service

import (
	"fmt"
	"log"
	"math/big"
	"testing"
)

func TestDepositProcess(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	listResult, code, msg, err := stakingService.DepositData(DepositDataRequestParams{
		Uid:               "mkx",
		Quantity:          2,
		WithdrawalAddress: "0x2B3779A253dB55B98eCED3EF427992740C17db17",
	})
	fmt.Println(listResult, code, msg, err)
	if err == nil {
		for _, list := range listResult.Ethereum.DepositData {
			fmt.Println("---pubkey:", list.Pubkey)
			fmt.Println("---withdraw:", list.WithdrawalCredential)
			fmt.Println("---signature:", list.Signature)
			fmt.Println("---depositDataRoot:", list.DepositDataRoot)

		}
	}

	//listResult, code, msg, err := stakingService.ListDeposits(1, 10)
	//fmt.Println(listResult, code, msg, err)
	//
	//depositResult, code, msg, err := stakingService.GetDepositDetails(listResult.List[0].TxHash)
	//fmt.Println(depositResult, code, msg, err)
	//pp.Println(depositResult)
	//
	//createResult, code, msg, err := stakingService.CreateUser("deposit_test")
	//require.Nil(t, err)
	//require.Equal(t, uint(0), code)
	//require.NotNil(t, createResult)
	//require.Equal(t, "ok", msg)
	//
	//assignmentResult, code, msg, err := stakingService.AssignDepositedValidators(depositResult.TxHash, ValidatorAssignmentParams{
	//	Assignments: []Assignment{
	//		{
	//			UID:            "deposit_test",
	//			ValidatorCount: len(depositResult.Validators),
	//		},
	//	},
	//})
	//fmt.Println(assignmentResult, code, msg, err)
	//pp.Println(assignmentResult)
	//
	//detailsResult, code, msg, err := stakingService.GetUserDetails("deposit_test")
	//require.Nil(t, err)
	//require.Equal(t, uint(0), code)
	//require.NotNil(t, detailsResult)
	//require.Len(t, detailsResult.Validators, len(depositResult.Validators))
	//require.Equal(t, "ok", msg)
	//
	//deleteResult, code, msg, err := stakingService.DeleteUser("deposit_test")
	//require.Nil(t, err)
	//require.Equal(t, uint(4000202), code)
	//require.Nil(t, deleteResult)
	//require.Equal(t, "Broker User's validators are not all exited yet", msg)
}

func TestTransact(t *testing.T) {
	intToBytes32 := func(value *big.Int) [32]byte {
		var bytes32Value [32]byte
		copy(bytes32Value[:], value.Bytes())
		return bytes32Value
	}
	fmt.Printf("bytes32 value: %#x\n", intToBytes32(big.NewInt(116)))
}
