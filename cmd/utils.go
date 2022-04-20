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