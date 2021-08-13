package cmd

import (
	"fmt"
	"github.com/ninjahome/ninja-go/extra_node/config"
	"github.com/ninjahome/ninja-go/extra_node/ethwallet"
	"github.com/ninjahome/ninja-go/tools"
	"github.com/spf13/cobra"
	"os"
)

var param = struct {
	walletForce bool
	initForce   bool
	passwd      string
}{}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init exnode",
	Long:  "init exnode",
	Run:   initExNode,
}

var WalletCmd = &cobra.Command{
	Use: "wallet",

	Short: "wallet command group",

	Long: `usage description::TODO::`,

	Run: showWallet,
}

var walletCreateCmd = &cobra.Command{
	Use: "create",

	Short: "create wallet",

	Long: `usage description::TODO::`,

	Run: createWallet,
}

func init() {
	InitCmd.Flags().BoolVarP(&param.initForce, "force", "f", false, "force init exnode")
	walletCreateCmd.Flags().BoolVarP(&param.walletForce, "force", "f", false, "force to create wallet whatever wallet exists")
	walletCreateCmd.Flags().StringVarP(&param.passwd, "password", "p", "", "password for protect wallet")
	WalletCmd.AddCommand(walletCreateCmd)
}

func createWallet(cmd *cobra.Command, args []string) {
	var (
		c *config.Config
		err error
	)

	if param.passwd == "" {
		fmt.Println("please input password")
		return
	}

	if c,err = config.InitConfig(); err!=nil{
		fmt.Println("please initial exnode first")
		return
	}

	if !param.walletForce {
		if wfile := c.GetWalletFile(); tools.FileExists(wfile) {
			fmt.Println("wallet have been initialized")
			return
		}
	}

	if w, err := ethwallet.NewWallet(param.passwd); err != nil {
		panic(err)
	} else {
		if err = w.SaveToPath(c.GetWalletFile()); err != nil {
			panic(err)
		}
		fmt.Println("wallet create success")
		fmt.Println("wallet address is:", w.MainAddress().String())
	}

}

func initExNode(cmd *cobra.Command, args []string) {
	h, err := config.GetExtraHome()
	if err != nil {
		panic(err)
	}

	if !param.initForce {
		if be := tools.FileExists(h); be {
			fmt.Println("exnode have been initialized")
			return
		}
	}

	if err := os.MkdirAll(h, 0755); err != nil {
		panic(err)
		return
	}

	c := config.DefaultConfig()
	if err = c.Save(); err != nil {
		panic(err)
	}

}

func showWallet(cmd *cobra.Command, args []string) {
	var (
		c *config.Config
		err error
		w   ethwallet.Wallet
	)
	if c,err = config.InitConfig(); err != nil {
		fmt.Println("please initial exnode first")
		return
	}

	if w, err = ethwallet.LoadWallet(c.GetWalletFile()); err != nil {
		fmt.Println("load wallet failed", err.Error())
		return
	}

	fmt.Println("wallet address:", w.MainAddress().String())
	//todo...
	fmt.Println("Eth Balance:")

}
