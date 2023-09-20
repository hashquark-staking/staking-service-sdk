package staking_service

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestDepositProcess(t *testing.T) {
	setupTest()
	stakingService, err := NewStakingService(STAKING_SERVICE_BASE_URL, PRIVATE_KEY, ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	depositResult, code, msg, err := stakingService.DepositData(DepositDataRequestParams{
		Uid:               "mkx",
		Quantity:          5,
		WithdrawalAddress: "0x2B3779A253dB55B98eCED3EF427992740C17db17",
	})
	fmt.Println("finish time:", time.Now())
	fmt.Println(*depositResult, code, msg, err)
	fmt.Println("period:", depositResult.Period)
	Pubkeys := make([]string, len(depositResult.Ethereum.DepositData))
	withdrawalCredentials := make([]string, len(depositResult.Ethereum.DepositData))
	signatures := make([]string, len(depositResult.Ethereum.DepositData))
	depositDataRoots := make([]string, len(depositResult.Ethereum.DepositData))
	if err == nil {
		for i, list := range depositResult.Ethereum.DepositData {
			Pubkeys[i] = `"` + list.Pubkey + `"`
			withdrawalCredentials[i] = `"` + list.WithdrawalCredential + `"`
			signatures[i] = `"` + list.Signature + `"`
			depositDataRoots[i] = `"` + list.DepositDataRoot + `"`
		}
	}
	fmt.Println("---Pubkeys:", "["+strings.Join(Pubkeys, ",")+"]")
	fmt.Println("---withdrawalCredentials:", "["+strings.Join(withdrawalCredentials, ",")+"]")
	fmt.Println("---signatures:", "["+strings.Join(signatures, ",")+"]")
	fmt.Println("---depositDataRoots:", "["+strings.Join(depositDataRoots, ",")+"]")

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
	fmt.Printf("bytes32 value: %#x\n", intToBytes32(big.NewInt(116))) // 0x7400000000000000000000000000000000000000000000000000000000000000
}
