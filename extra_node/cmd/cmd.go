package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var param = struct {
	walletForce bool
}{}

var WalletCmd = &cobra.Command{
	Use: "wallet",

	Short: "wallet command group",

	Long: `usage description::TODO::`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please use sub command")
	},
}

var walletCreateCmd = &cobra.Command{
	Use: "create",

	Short: "create wallet",

	Long: `usage description::TODO::`,

	Run: createWallet,
}

func init() {
	walletCreateCmd.Flags().BoolVarP(&param.walletForce, "force", "f", false, "force to create wallet whatever wallet exists")
	WalletCmd.AddCommand(walletCreateCmd)
}

func createWallet(cmd *cobra.Command, args []string) {

}
