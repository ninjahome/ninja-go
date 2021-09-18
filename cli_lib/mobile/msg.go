package chatLib

func RecoverMsgFromDb(msgCount int, time int64) {

}

func RecoverGroupMsg(msgCount int, groupId string, time int64) {

}

func RecoverMsg(msgCount int, to string, time int64) {

}

//if startTime == endTime, just delete one msg
//if startTime == endTime == 0, delete all chat msg
//if startTime < endTime, error
//id means userId or groupId
//if id == "", clear all msg between startTime and endTime
func ClearMsg(startTime, endTime int64, id string) error {
	return nil
}


//one hour = 3600
//one day = 86400
//one week = 604800
//one month = 2592000
//3 month = 7776000
//half year = 15552000
//one year = 31104000
//-1 for ever
func SaveTimeInterval(seconds int64)  {
	_inst.chatStore.SetHistoryInterval(seconds)
}

func DeleteUser(addr string) error {
	return nil
}


