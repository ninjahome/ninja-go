package main

import (
	"fmt"
	"github.com/ninjahome/bls-wallet/bls"
	"github.com/ninjahome/ninja-go/extra_node/cmd"
	"github.com/ninjahome/ninja-go/extra_node/config"
	"github.com/ninjahome/ninja-go/extra_node/ethwallet"
	"github.com/ninjahome/ninja-go/extra_node/webserver"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
)

const (
	PidFileName = ".pid"
	Version     = "0.0.1"
)

var param = struct {
	version bool
	passwd  string
}{}

var rootCmd = &cobra.Command{
	Use: "exnode",

	Short: "extra node for service node",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "exnode -v")
	flags.StringVarP(&param.passwd, "password", "p", "", "password for open wallet")

	rootCmd.AddCommand(cmd.WalletCmd)
	rootCmd.AddCommand(cmd.InitCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func mainRun(_ *cobra.Command, _ []string) {

	if param.version {
		fmt.Println(Version)
		return
	}
	if param.passwd == "" {
		fmt.Println("please input password")
		return
	}

	c := config.GetExtraConfig()
	if c == nil {
		fmt.Println("please initial exnode")
		return
	}

	w, err := ethwallet.LoadWallet(c.GetWalletFile())
	if err != nil {
		panic(err)
	}
	if err = w.Open(param.passwd); err != nil {
		panic(err)
	}

	bls.Init(bls.BLS12_381)
	//bls.SetETHmode(bls.EthModeDraft07)

	go webserver.StartWebDaemon(w)

	waitShutdownSignal()
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>extra node start at pid(%s)<<<<<<<<<<\n", pid)

	h, _ := config.GetExtraHome()
	pidfile := path.Join(h, PidFileName)

	if err := ioutil.WriteFile(pidfile, []byte(pid), 0644); err != nil {
		fmt.Print("failed to write running pid", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)

	sig := <-sigCh

	webserver.StopWebDaemon()

	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
