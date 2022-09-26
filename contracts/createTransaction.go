package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"golang_standard/album"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetContract() (*album.Album, *bind.TransactOpts) {
	// address of etherum env
	infura := "https://goerli.infura.io/v3/0c84ae418d3042a490f3f8f46bd823c4"
	client, err := ethclient.Dial(infura)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := album.NewAlbum(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	auth := GetAccountAuth(client, os.Getenv("PRIVATE_KEY"))
	return conn, auth
}

func GetAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("ChainID", chainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
