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
	//result, code, msg, err := stakingService.ManualDepositBrokerUser("0x0361f7813b412e8cbd026b930022d3d728f7d6690b430881d193b89190758b47")
	//fmt.Println("--result:", result)
	//fmt.Println("--code:", code)
	//fmt.Println("--msg:", msg)
	//fmt.Println("--err:", err)

	depositResult, code, msg, err := stakingService.DepositData(DepositDataRequestParams{
		Uid:               "mkx",
		Quantity:          1,
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
	fmt.Println("---unsignedTransaction:", depositResult.Ethereum.UnsignedTransaction)

	//sendTransactionResult, code, msg, err := stakingService.SendTransaction(SendTransactionRequestParams{
	//	SignedTransaction: "0xb9035802f90354827e7e0a050a830173b694ec98ac89df4f62fa6fc79ef67309e1b436342dc78901bc16d674ec800000b902e46225d68c000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001e000000000000000000000000000000000000000000000000000000000000002a07400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000030a90d5626dc35029f52fafd15f084a3dc1b33fc0289e8f85af8f187cdfcedea05f09334bd89d2e12b8e1a802534aef0af000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200100000000000000000000002b3779a253db55b98eced3ef427992740c17db17000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000060a9190ba5aa168e0c6078ecf519d40825f5b3030f1762bb2c297e3195f96f4f7d04a4bfd7eabc65d13a1065dea33013ea1977574cd227ba00b66a2e3004fa74b04284b7119c80b584858465b79d34683d7615578407b21d2dc0ffd9c0a4bf07ca00000000000000000000000000000000000000000000000000000000000000018a07182d3aaf9d700895ba0fc0120419829c98e3f4ecf883e28dda42d831e70ac001a0cf6300a6c6863d5e76f13371bb80adf459a841aed376478391ec3edfeac40152a06b2691638d0a8c16474d22771587474a7b32abeee74102dac57f10268c3651f8",
	//	ChainName:         "ethereum",
	//})
	//fmt.Println(code, msg, err)
	//fmt.Println("txHash:", sendTransactionResult.TxHash)

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
