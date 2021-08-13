package main

import (
	"fmt"
	"github.com/ninjahome/ninja-go/extra_node/cmd"
	"github.com/ninjahome/ninja-go/extra_node/config"
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

	rootCmd.AddCommand(cmd.WalletCmd)

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

	if err := config.InitSystem(); err != nil {
		panic(err)
	}

	go webserver.StartWebDaemon()

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
