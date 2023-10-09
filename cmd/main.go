package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/liviudnicoara/egopay/internal/accounts"
	"github.com/liviudnicoara/egopay/internal/bills"
)

const ALCHEMY_TEST_URL = "https://eth-sepolia.g.alchemy.com/v2/w0uQJ2Oimfqh8H-ibuWNhP7dhv7cjeqs"
const ACCOUNT = "0xF97b8a8f9bA083b1FA5b05Cb956Ca0eca02fBD58"

func main() {
	client, err := ethclient.DialContext(context.Background(), ALCHEMY_TEST_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// // Get the balance of an account
	// account := common.HexToAddress(ACCOUNT)
	// balance, err := client.BalanceAt(context.Background(), account, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ethBalance := convertWEItoETH(balance)
	// fmt.Printf("Account balance: %s\n", ethBalance.Text('f', 10)) // 25893180161173005034

	// // Get the latest known block
	// block, err := client.BlockByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Latest block: %d\n", block.Number().Uint64())

	// pvk, pbk, account, err := getWallet("test")
	// fmt.Println(pvk)
	// fmt.Println(pbk)
	// fmt.Println(account)
	// fmt.Println(err)

	ksPath := "./.storage/accounts"
	as := accounts.NewAccountService(accounts.NewAccountRepository(ksPath), client)

	// storagePath := "./.storage/users"
	// us := users.NewUserService(users.NewUserRepository(storagePath), as)

	// user, err := us.Register("test", "pass")
	// ad, err := us.AddAccount(uuid.MustParse("77e34f97-cea3-4278-9570-20dc309d3a51"), "pass")
	// b, err := as.GetBalance(context.Background(), "0xCf9a951E338A3663804b5499706dc50A79AE908A")
	// fmt.Println(b.Float64())
	// fmt.Println(err)

	// ts := transfers.NewTransferService(as, client)
	// amount := big.NewFloat(0.01)
	// err = ts.Transfer(context.Background(), "", "0xCf9a951E338A3663804b5499706dc50A79AE908A", amount, "Ceparola123!")
	// fmt.Println(err)

	// b, err = as.GetBalance(context.Background(), "0xCf9a951E338A3663804b5499706dc50A79AE908A")
	// fmt.Println(b)
	// fmt.Println(err)

	// Deploy contract
	// account, err := as.GetAccount("0xCf9a951E338A3663804b5499706dc50A79AE908A", "Ceparola123!")

	bs := bills.NewBillService(as, client)
	address, tx, err := bs.Split(context.Background(), "0xCf9a951E338A3663804b5499706dc50A79AE908A", 100, "pass")

	fmt.Println(address)
	fmt.Println(tx)

}
