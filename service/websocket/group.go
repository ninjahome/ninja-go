package websocket

import (
	"errors"
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb"
	"sync"
)


type GroupMember struct {
	MemberId 	string `json:"member_id"`
	NickName    string `json:"nick_name"`
}

func (gm *GroupMember)String() string  {
	if gm == nil{
		return "Group Member struct is nil"
	}

	return fmt.Sprintf("Member Id: %s\tNickName: %s\r\n",gm.MemberId,gm.NickName)
}


type GroupDesc struct {
	GroupName string			`json:"group_name"`
	GroupId   string 			`json:"group_id"`
	Owner     string			`json:"owner"`
	Members   []*GroupMember 	`json:"members"`
}



func (gd *GroupDesc)dup() *GroupDesc  {
	g:=&GroupDesc{
		GroupName: gd.GroupName,
		GroupId: gd.GroupId,
		Owner: gd.Owner,
	}

	for i:=0;i<len(gd.Members);i++{
		g.Members = append(g.Members,&GroupMember{
			MemberId: gd.Members[i].MemberId,
			NickName: gd.Members[i].NickName,
		})
	}

	return g

}

func (gd *GroupDesc)String() string  {
	if gd == nil{
		return "group describe struct is nil"
	}

	msg:=fmt.Sprintf("Group Name: %s\r\nGroup Id: %s\r\nGroup Owner: %s\r\n",
		gd.GroupName,gd.GroupId,gd.Owner)

	for _,gm:=range gd.Members{
		msg += gm.String()
	}

	return msg
}



type GroupStore struct {
	storeLock *sync.RWMutex
	groupLock map[string]*sync.RWMutex
	groupList map[string]*GroupDesc
	dblock *sync.Mutex
	db *leveldb.DB
}

func NewGroupStore(db *leveldb.DB) *GroupStore {
	return &GroupStore{
		storeLock: &sync.RWMutex{},
		groupLock: make(map[string]*sync.RWMutex),
		groupList: make(map[string]*GroupDesc),
		dblock: &sync.Mutex{},
		db: db,
	}
}

func (gs *GroupStore)_getGroupLocker(groupId string) *sync.RWMutex  {
	gs.storeLock.RLock()
	defer gs.storeLock.RUnlock()
	if l,ok:=gs.groupLock[groupId];!ok{
		return nil
	}else{
		return l
	}
}

func (gs *GroupStore)getGroupLocker(groupId string) *sync.RWMutex  {
	if l:=gs._getGroupLocker(groupId);l!=nil{
		return l
	}

	gs.storeLock.Lock()
	defer gs.storeLock.Unlock()

	if l,ok:=gs.groupLock[groupId];ok{
		return l
	}

	l:=&sync.RWMutex{}

	gs.groupLock[groupId] = l

	return l
}

func (gs *GroupStore)DelLocker(groupId string)   {
	gs.storeLock.Lock()
	defer gs.storeLock.Unlock()

	if _,ok:=gs.groupLock[groupId];ok{
		delete(gs.groupLock,groupId)
	}
}



func (gs *GroupStore)GetGroup(groupId string) *GroupDesc  {
	locker:=gs.getGroupLocker(groupId)
	locker.RLock()
	defer locker.RUnlock()

	if g,ok:=gs.groupList[groupId];ok{
		return g.dup()
	}

	return nil
}


func (gs *GroupStore)AddGroup(groupId,groupName,groupOwner string, MemberIds, nickNames []string) {
	locker:=gs.getGroupLocker(groupId)
	locker.Lock()
	defer locker.Unlock()

	var (
		g *GroupDesc
		ok bool
	)

	if g,ok=gs.groupList[groupId];!ok{
		g = &GroupDesc{}
	}

	g.GroupId = groupId
	g.GroupName = groupName
	g.Owner = groupOwner

	lm:=len(MemberIds)
	ln:=len(nickNames)

	if lm > ln{
		lm = ln
	}

	for i:=0;i<lm;i++{
		member:=&GroupMember{
			MemberId: MemberIds[i],
			NickName: nickNames[i],
		}
		g.Members = append(g.Members,member)
	}

	gs.groupList[groupId] = g

}

func (gs *GroupStore)DelGroup(groupId string) bool {
	locker:=gs.getGroupLocker(groupId)
	locker.Lock()
	defer locker.Unlock()

	if _,ok:=gs.groupList[groupId];ok{
		delete(gs.groupList,groupId)
		return true
	}

	return false
}

func (gs *GroupStore)UpdateGroupName(groupId,groupName string) error {
	locker:=gs.getGroupLocker(groupId)
	locker.Lock()
	defer locker.Unlock()

	if g,ok:=gs.groupList[groupId];ok{
		g.GroupName = groupName
		return nil
	}

	return errors.New("group not found")

}

func (gs *GroupStore)UpdateGroupOwner(groupId,groupOwner string) error {
	locker:=gs.getGroupLocker(groupId)
	locker.Lock()
	defer locker.Unlock()

	if g,ok:=gs.groupList[groupId];ok{
		g.Owner = groupOwner
		return nil
	}

	return errors.New("group not found")

}

func (gs *GroupStore)AddMember(groupId,memberId,nickName string) (newmember bool,err error)  {
	locker:=gs.getGroupLocker(groupId)
	locker.Lock()
	defer locker.Unlock()

	if g,ok:=gs.groupList[groupId];!ok{
		return false,errors.New("group not found")
	}else{

		for i:=0;i<len(g.Members);i++{
			if g.Members[i].MemberId == memberId{
				return false,errors.New("member id found")
			}
		}

		g.Members = append(g.Members,&GroupMember{MemberId: memberId,NickName: nickName})

		return true,nil
	}
}

func (gs *GroupStore)UpdateMember(groupId,memberId,nickName string) (newmember bool,err error)  {
	locker:=gs.getGroupLocker(groupId)
	locker.Lock()
	defer locker.Unlock()

	if g,ok:=gs.groupList[groupId];!ok{
		return false,errors.New("group not found")
	}else{
		for i:=0;i<len(g.Members);i++{
			if g.Members[i].MemberId == memberId{
				g.Members[i].NickName = nickName
				return false,nil
			}
		}

		g.Members = append(g.Members,&GroupMember{MemberId: memberId,NickName: nickName})

		return true,nil
	}
}

func (gs *GroupStore)MemberInGroup(groupId, memberId string) (bool,error)  {
	locker:=gs.getGroupLocker(groupId)
	locker.RLock()
	defer locker.RUnlock()

	if g,ok:=gs.groupList[groupId];!ok{
		return false, errors.New("group no found")
	}else{
		for _,gm:=range g.Members{
			if gm.MemberId == memberId{
				return true, nil
			}
		}

		return false, nil
	}

}

func (gs *GroupStore)DelMember(groupId, memberId string) (bool,error) {
	locker:=gs.getGroupLocker(groupId)
	locker.Lock()
	defer locker.Unlock()

	if g,ok:=gs.groupList[groupId];!ok{
		return false, errors.New("group no found")
	}else{
		idx:=-1
		for i:=0;i<len(g.Members);i++{
			if g.Members[i].MemberId == memberId{
				idx = i
				break
			}
		}
		if idx == -1{
			return false, errors.New("member not found")
		}

		last:=len(g.Members)-1

		if idx != last{
			g.Members[idx] = g.Members[last]
		}

		g.Members = g.Members[:last]

		return true, nil
	}
}

func (gs *GroupStore)ListGroup() []*GroupDesc  {

	gs.storeLock.RLock()
	defer gs.storeLock.Unlock()

	var gds []*GroupDesc

	for _, groupDesc:=range gs.groupList{
		gds = append(gds,groupDesc.dup())
	}

	return gds
}
