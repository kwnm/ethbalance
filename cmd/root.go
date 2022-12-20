package cmd

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ethbalance",
	Short: "A tool to check the balance of an Ethereum wallet",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial("https://mainnet.infura.io")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		account := common.HexToAddress(address)
		balance, err := client.BalanceAt(nil, account, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Balance for address %s: %s ETH\n", address, balance)
	},
}

var address string

func init() {
	rootCmd.Flags().StringVarP(&address, "address", "a", "", "Ethereum address to check balance for")
	rootCmd.MarkFlagRequired("address")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//
///*
//Copyright © 2022 NAME HERE <EMAIL ADDRESS>
//*/
//package cmd
//
//import (
//	"context"
//	"fmt"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"math"
//	"math/big"
//	"os"
//
//	"github.com/spf13/cobra"
//)
//
//var address string
//
//// Execute adds all child commands to the root command and sets flags appropriately.
//// This is called by main.main(). It only needs to happen once to the rootCmd.
//func Execute() {
//	err := rootCmd.Execute()
//	if err != nil {
//		os.Exit(1)
//	}
//}
//
//func init() {
//	// Here you will define your flags and configuration settings.
//	// Cobra supports persistent flags, which, if defined here,
//	// will be global for your application.
//
//	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.strait-etl.yaml)")
//
//	// Cobra also supports local flags, which will only run
//	// when this action is called directly.
//	rootCmd.AddCommand(ethbalanceCmd)
//}
//
//var ethbalanceCmd = &cobra.Command{
//
//	Use:   "ethbalance",
//	Short: "A tool to check the balance of an Ethereum wallet",
//	Run: func(cmd *cobra.Command, args []string) {
//		client, err := ethclient.Dial("https://mainnet.infura.io/v3/e73825dafe8940d5b523f25fe4891aac")
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		account := common.HexToAddress(address)
//		balance, err := client.BalanceAt(context.Background(), account, nil)
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//		fbalance := new(big.Float)
//		fbalance.SetString(balance.String())
//		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
//		fmt.Printf("Balance for address %s: %f ETH\n", address, ethValue)
//
//	},
//}
