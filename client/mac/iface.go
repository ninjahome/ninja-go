package mac

import "C"

type MacApp struct {
	uiAPI C.UserInterfaceAPI
}

type UIAPI interface {
	Log(...interface{})
	ActionNotify(int, ...interface{})
	SysExit(err error)
}

//export initConf
func initConf(C.UserInterfaceAPI) *C.char {
	return nil
}
