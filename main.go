package main

import (
	"encoding/json"
	"fmt"
	"github.com/ninjahome/ninja-go/cmd"
	"github.com/ninjahome/ninja-go/node"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/fdlimit"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/ninjahome/ninja-go/wallet"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

type SysParam struct {
	version  bool
	baseDir  string
	network  string
	password string
	keyAddr  string
	wsIP     string
	wsPort   int16
}

const (
	PidFileName = "pid"
	Version     = "0.1.0"
)

var (
	param = &SysParam{}
)

var rootCmd = &cobra.Command{
	Use: "ninja",

	Short: "ninja is a peer to peer chat system",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "ninja -v")

	flags.StringVarP(&param.network, "network",
		"n", cmd.MainNet,
		"ninja -n|--network ["+cmd.MainNet+"|"+cmd.TestNet+"] default is "+cmd.MainNet+".")

	flags.StringVarP(&param.password, "password",
		"p", "", "ninja -p [PASSWORD OF SELECTED KEY]")

	flags.StringVarP(&param.keyAddr, "key",
		"k", "", "ninja -k [ADDRESS OF KEY]")

	flags.StringVarP(&param.baseDir, "dir",
		"d", cmd.DefaultBaseDir, "chord -d [BASIC DIRECTORY]")

	flags.StringVar(&param.wsIP, "ws.ip", "",
		"ninja --ws.ip [IP]")
	flags.Int16Var(&param.wsPort, "ws.port", -1,
		"ninja --ws.port [Port]")

	rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(cmd.WalletCmd)
	rootCmd.AddCommand(cmd.DebugCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func initNinjaConfig() (err error) {

	conf := make(cmd.StoreCfg)
	dir := utils.BaseUsrDir(param.baseDir)
	confPath := filepath.Join(dir, string(filepath.Separator), cmd.ConfFileName)
	bts, e := os.ReadFile(confPath)
	if e != nil {
		return e
	}

	if err = json.Unmarshal(bts, &conf); err != nil {
		return
	}

	result, ok := conf[param.network]
	if !ok {
		err = fmt.Errorf("failed to find node config")
		return
	}
	if param.wsPort != -1 {
		result.RCfg.WsPort = param.wsPort
	}
	if param.wsIP != "" {
		result.RCfg.WsIP = param.wsIP
	}
	fmt.Println(result.String())

	wallet.InitConfig(result.WCfg)
	node.InitConfig(result.PCfg)
	utils.InitConfig(result.UCfg)
	websocket.InitConfig(result.RCfg)
	//TODO:: configure contact service dynamically
	contact.InitConfig(result.CCfg)

	return
}

func initWalletKey() error {
	var pwd = param.password
	if pwd == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}
		pwd = string(pw)
	}

	if err := wallet.Inst().Active(pwd, param.keyAddr); err != nil {
		return err
	}
	return nil
}

func initSystem() error {

	if err := os.Setenv("GODEBUG", "netdns=go"); err != nil {
		return err
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(int64(time.Now().Nanosecond()))
	limit, err := fdlimit.Maximum()
	if err != nil {
		return fmt.Errorf("failed to retrieve file descriptor allowance:%s", err)
	}
	_, err = fdlimit.Raise(uint64(limit))
	if err != nil {
		return fmt.Errorf("failed to raise file descriptor allowance:%s", err)
	}
	return nil
}

func mainRun(_ *cobra.Command, _ []string) {

	if param.version {
		fmt.Println(Version)
		return
	}

	if err := initSystem(); err != nil {
		panic(err)
	}

	if err := initNinjaConfig(); err != nil {
		panic(err)
	}

	if err := initWalletKey(); err != nil {
		panic(err)
	}



	if err := node.Inst().Start(); err != nil {
		panic(err)
	}
	thread.NewThreadWithName(cmd.ThreadName, cmd.StartCmdRpc).Run()


	waitShutdownSignal()
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>ninja node start at pid(%s)<<<<<<<<<<\n", pid)
	path := filepath.Join(utils.BaseUsrDir(param.baseDir), string(filepath.Separator), PidFileName)
	if err := ioutil.WriteFile(path, []byte(pid), 0644); err != nil {
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
	node.Inst().ShutDown()
	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
