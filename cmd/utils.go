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