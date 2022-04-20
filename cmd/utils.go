package cmd

import(
	"context"
	"io/ioutil"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
)

type Wallet struct{
	account types.account
	c *client.Client
}

func CreateNewWallet(RPCEndpoint string) Wallet {
	// create a new wallet using solana-go-sdk that provides a NewAccount() method in types
	newAccount := types.NewAccount()

	data := []byte(newAccount.PrivateKey) // convert the private key to byte array for storage
	err := ioutil.WriteFile("data", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return Wallet{
		newAccount,
		client.NewClient(RPCEndpoint),
	}
}

func ImportOldWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {
	// import a wallet with bytes slice private key
	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet {
		wallet, 
		client.NewClient(RPCEndpoint),
	}, nil
}

func GetBalance() (uint64, error){
	wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
	balance, err := wallet.c.GetBalance(
		context.TODO(), //request context
		wallet.account.PublicKey.ToBase58() // wallet to fetch balance for
	)
	if err != nil{
		return 0, nil
	}
	// The functin returns us the balance in lamports.
	// We can convert this to SOL by dividing with 1e9
	return balance, nil
}

func RequestAirdrop(amount uint64) (string, error){
	// request for SOL using RequestAirdrop()
	wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
	amount = amount * 1e9 // turning SOL into lamports
	txhash, err := wallet.c.RequestAirdrop(
		context.TODO() // request context
		wallet.account.PublicKey.ToBase58() // wallet address requesting airdrop
		amount, // amount of SOL in lamport
	)
	if err != nil{
		return "", err
	}
	return txhash, nil
}