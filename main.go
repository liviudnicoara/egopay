package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/liviudnicoara/egopay/accounts"
	contracts "github.com/liviudnicoara/egopay/contracts"
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
	account, err := as.GetAccount("0xCf9a951E338A3663804b5499706dc50A79AE908A", "pass")
	if err != nil {
		log.Fatal(err)
	}

	nounce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(account.PrivateKey.PublicKey))
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, chainID)
	auth.Nonce = big.NewInt(int64(nounce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)

	payers := []common.Address{crypto.PubkeyToAddress(account.PrivateKey.PublicKey)}
	amount := big.NewInt(100)
	address, tx, instance, err := contracts.DeploySplitBill(auth, client, payers, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address)
	fmt.Println(tx.Hash().Hex())
	_ = instance

}

// func createNewWallet() (string, string, string, error) {
// 	pvk, err := crypto.GenerateKey()
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	binaryPVK := crypto.FromECDSA(pvk)
// 	encodedPVK := hexutil.Encode(binaryPVK)

// 	binaryPBK := crypto.FromECDSAPub(&pvk.PublicKey)
// 	encodedPBK := hexutil.Encode(binaryPBK)

// 	address := crypto.PubkeyToAddress(pvk.PublicKey).Hex()

// 	return encodedPVK, encodedPBK, address, nil
// }

// func createNewKeyStoreWallet(password string) error {
// 	ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
// 	_, err := ks.NewAccount(password)

// 	return err
// }

// func getWallet(password string) (string, string, string, error) {
// 	f, err := os.ReadFile("./wallet/UTC--2023-09-28T12-16-40.596555300Z--ca1d3dd253cade17ee7af14cb3a0c9c7cd49a509")
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	ks, err := keystore.DecryptKey(f, password)

// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	binaryPVK := crypto.FromECDSA(ks.PrivateKey)
// 	encodedPVK := hexutil.Encode(binaryPVK)

// 	binaryPBK := crypto.FromECDSAPub(&ks.PrivateKey.PublicKey)
// 	encodedPBK := hexutil.Encode(binaryPBK)

// 	address := crypto.PubkeyToAddress(ks.PrivateKey.PublicKey).Hex()

// 	return encodedPVK, encodedPBK, address, nil
// }
