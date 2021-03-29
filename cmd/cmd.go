package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ninjahome/ninja-go/node"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/wallet"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path/filepath"
)

const (
	DefaultBaseDir = ".ninja"
	ConfFileName   = "config.json"
	MainNet        = "main"
	TestNet        = "test"
)

type StoreCfg map[string]*CfgPerNetwork

type CfgPerNetwork struct {
	Name string            `json:"name"`
	PCfg *node.Config      `json:"node"`
	UCfg *utils.Config     `json:"utils"`
	WCfg *wallet.Config    `json:"wallet"`
	RCfg *websocket.Config `json:"websocket"`
	CCfg *contact.Config   `json:"contact"`
}

func (sc StoreCfg) DebugPrint() {
	for _, c := range sc {
		fmt.Println(c)
	}
}

func (c CfgPerNetwork) String() string {
	s := fmt.Sprintf("\n<<<===================System[%s] Config===========================", c.Name)
	s += c.PCfg.String()
	s += c.UCfg.String()
	s += c.WCfg.String()
	s += c.RCfg.String()
	s += c.CCfg.String()
	s += fmt.Sprintf("\n======================================================================>>>")
	return s
}

var param struct {
	baseDir     string
	servicePort *int16
	password    string
	forTest     bool
}

func init() {

	InitCmd.Flags().StringVarP(&param.baseDir, "baseDir", "d", DefaultBaseDir,
		"ninja init -d [DIR]")

	walletCreateCmd.Flags().StringVarP(&param.password, "password", "p", "",
		"ninja wallet create -p|--password [PASSWORD]")
	walletCreateCmd.Flags().BoolVarP(&param.forTest, "test", "t", false, "ninja wallet create -t")

	WalletCmd.AddCommand(walletCreateCmd)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "ninja init -d [DIR]",
	Long:  `TODO::.`,
	Run:   initNode,
	//Args:  cobra.MinimumNArgs(1),
}

func initNode(_ *cobra.Command, _ []string) {
	dir := utils.BaseUsrDir(param.baseDir)
	if utils.FileExists(dir) {
		panic("duplicate init operation! please save the old config or use -baseDir for new node config")
	}
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		panic(err)
	}
	if err := initDefault(dir); err != nil {
		panic(err)
	}
}

func initDefault(baseDir string) error {
	conf := make(StoreCfg)
	mainConf := &CfgPerNetwork{
		Name: MainNet,
		PCfg: node.DefaultConfig(true, baseDir),
		UCfg: &utils.Config{
			LogLevel: zerolog.ErrorLevel,
		},
		WCfg: wallet.DefaultConfig(true, baseDir),
		RCfg: websocket.DefaultConfig(true, baseDir),
		CCfg: contact.DefaultConfig(true, baseDir),
	}
	conf[MainNet] = mainConf

	testConf := &CfgPerNetwork{
		Name: TestNet,
		PCfg: node.DefaultConfig(false, baseDir),
		UCfg: &utils.Config{
			LogLevel: zerolog.DebugLevel,
		},
		WCfg: wallet.DefaultConfig(false, baseDir),
		RCfg: websocket.DefaultConfig(false, baseDir),
		CCfg: contact.DefaultConfig(false, baseDir),
	}

	conf[TestNet] = testConf

	bts, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		panic(err)
	}
	path := filepath.Join(baseDir, string(filepath.Separator), ConfFileName)
	if err := os.WriteFile(path, bts, 0644); err != nil {
		panic(err)
	}
	return nil
}

var WalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "ninja wallet",
	Long:  `TODO::.`,
	Run:   walletAction,
}

func walletAction(c *cobra.Command, _ []string) {
	_ = c.Usage()
}

var walletCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "ninja wallet create -p [PASSWORD] -t",
	Long:  `TODO::.`,
	Run:   createAcc,
	Args:  cobra.MinimumNArgs(1),
}

func createAcc(_ *cobra.Command, _ []string) {
	if param.password == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		param.password = string(pw)
	}

	dir := utils.BaseUsrDir(param.baseDir)
	var walletDir = ""
	if param.forTest {
		walletDir = filepath.Join(dir, string(filepath.Separator), wallet.TestKeyStoreScheme)
	} else {
		walletDir = filepath.Join(dir, string(filepath.Separator), wallet.KeyStoreScheme)
	}

	wallet.InitConfig(&wallet.Config{Dir: walletDir})
	if err := wallet.Inst().CreateNewKey(param.password); err != nil {
		panic(err)
	}
	fmt.Println("create success!")
}
