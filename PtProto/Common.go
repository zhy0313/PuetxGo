package PtProto

/**
 * Title:common组件
 * User: iuoon
 * Date: 2016-9-23
 * Version: 1.0
 */

import ("github.com/golang/protobuf/proto"
	"fmt"
	"math"
        )


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PlayerRaw struct {
	Username         *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Pwd              *string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Email            *string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayerRaw) Reset()         { *m = PlayerRaw{} }
func (m *PlayerRaw) String() string { return proto.CompactTextString(m) }
func (*PlayerRaw) ProtoMessage()    {}

func (m *PlayerRaw) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *PlayerRaw) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *PlayerRaw) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

type GameSrvInfo struct {
	Index            *uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status           *uint32 `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Lastts           *uint32 `protobuf:"varint,4,opt,name=lastts" json:"lastts,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameSrvInfo) Reset()         { *m = GameSrvInfo{} }
func (m *GameSrvInfo) String() string { return proto.CompactTextString(m) }
func (*GameSrvInfo) ProtoMessage()    {}

func (m *GameSrvInfo) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *GameSrvInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GameSrvInfo) GetStatus() uint32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *GameSrvInfo) GetLastts() uint32 {
	if m != nil && m.Lastts != nil {
		return *m.Lastts
	}
	return 0
}

// 登录游戏相关数据.
type LoginRspInfo struct {
	Token            *string        `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Host             *string        `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Port             *uint32        `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Srvs             []*GameSrvInfo `protobuf:"bytes,4,rep,name=srvs" json:"srvs,omitempty"`
	Version1         *string        `protobuf:"bytes,5,opt,name=version1" json:"version1,omitempty"`
	Version2         *string        `protobuf:"bytes,6,opt,name=version2" json:"version2,omitempty"`
	Version3         *string        `protobuf:"bytes,7,opt,name=version3" json:"version3,omitempty"`
	Download         *string        `protobuf:"bytes,8,opt,name=download" json:"download,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *LoginRspInfo) Reset()         { *m = LoginRspInfo{} }
func (m *LoginRspInfo) String() string { return proto.CompactTextString(m) }
func (*LoginRspInfo) ProtoMessage()    {}

func (m *LoginRspInfo) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *LoginRspInfo) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *LoginRspInfo) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *LoginRspInfo) GetSrvs() []*GameSrvInfo {
	if m != nil {
		return m.Srvs
	}
	return nil
}

func (m *LoginRspInfo) GetVersion1() string {
	if m != nil && m.Version1 != nil {
		return *m.Version1
	}
	return ""
}

func (m *LoginRspInfo) GetVersion2() string {
	if m != nil && m.Version2 != nil {
		return *m.Version2
	}
	return ""
}

func (m *LoginRspInfo) GetVersion3() string {
	if m != nil && m.Version3 != nil {
		return *m.Version3
	}
	return ""
}

func (m *LoginRspInfo) GetDownload() string {
	if m != nil && m.Download != nil {
		return *m.Download
	}
	return ""
}

type RoleCommonInfo struct {
	Id              *int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name            *string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Sex             *int32      `protobuf:"varint,3,opt,name=sex" json:"sex,omitempty"`
	VocationId      *int32      `protobuf:"varint,4,opt,name=vocationId" json:"vocationId,omitempty"`
	Vip             *int32      `protobuf:"varint,5,opt,name=vip" json:"vip,omitempty"`
	Expr            *int64      `protobuf:"varint,6,opt,name=expr" json:"expr,omitempty"`
	Money           *int64      `protobuf:"varint,7,opt,name=money" json:"money,omitempty"`
	Gold            *int64      `protobuf:"varint,8,opt,name=gold" json:"gold,omitempty"`
	Power           *int64      `protobuf:"varint,9,opt,name=power" json:"power,omitempty"`
	Powerts         *int64      `protobuf:"varint,10,opt,name=powerts" json:"powerts,omitempty"`
	Honor           *int64      `protobuf:"varint,11,opt,name=honor" json:"honor,omitempty"`
	ArmyPay         *int64      `protobuf:"varint,12,opt,name=armyPay" json:"armyPay,omitempty"`
	Wings           *PbItemInfo `protobuf:"bytes,13,opt,name=wings" json:"wings,omitempty"`
	Weapon          *PbItemInfo `protobuf:"bytes,14,opt,name=weapon" json:"weapon,omitempty"`
	Coat            *PbItemInfo `protobuf:"bytes,15,opt,name=coat" json:"coat,omitempty"`
	Cloakd          *PbItemInfo `protobuf:"bytes,16,opt,name=cloakd" json:"cloakd,omitempty"`
	Trouser         *PbItemInfo `protobuf:"bytes,17,opt,name=trouser" json:"trouser,omitempty"`
	Armguard        *PbItemInfo `protobuf:"bytes,18,opt,name=armguard" json:"armguard,omitempty"`
	Shoes           *PbItemInfo `protobuf:"bytes,19,opt,name=shoes" json:"shoes,omitempty"`
	FightCardDistri *int32      `protobuf:"varint,20,opt,name=fightCardDistri" json:"fightCardDistri,omitempty"`
	BuyPowerCnt     *int32      `protobuf:"varint,21,opt,name=buyPowerCnt" json:"buyPowerCnt,omitempty"`
	BuyPowerTs      *int64      `protobuf:"varint,22,opt,name=buyPowerTs" json:"buyPowerTs,omitempty"`
	GetFreePowerTs  *int64      `protobuf:"varint,23,opt,name=getFreePowerTs" json:"getFreePowerTs,omitempty"`
	// 好友功能使用
	FightValue       *int32 `protobuf:"varint,24,opt,name=fightValue" json:"fightValue,omitempty"`
	Online           *int32 `protobuf:"varint,25,opt,name=online" json:"online,omitempty"`
	Send             *int32 `protobuf:"varint,26,opt,name=send" json:"send,omitempty"`
	Recv             *int32 `protobuf:"varint,27,opt,name=recv" json:"recv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RoleCommonInfo) Reset()         { *m = RoleCommonInfo{} }
func (m *RoleCommonInfo) String() string { return proto.CompactTextString(m) }
func (*RoleCommonInfo) ProtoMessage()    {}

func (m *RoleCommonInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *RoleCommonInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RoleCommonInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *RoleCommonInfo) GetVocationId() int32 {
	if m != nil && m.VocationId != nil {
		return *m.VocationId
	}
	return 0
}

func (m *RoleCommonInfo) GetVip() int32 {
	if m != nil && m.Vip != nil {
		return *m.Vip
	}
	return 0
}

func (m *RoleCommonInfo) GetExpr() int64 {
	if m != nil && m.Expr != nil {
		return *m.Expr
	}
	return 0
}

func (m *RoleCommonInfo) GetMoney() int64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *RoleCommonInfo) GetGold() int64 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *RoleCommonInfo) GetPower() int64 {
	if m != nil && m.Power != nil {
		return *m.Power
	}
	return 0
}

func (m *RoleCommonInfo) GetPowerts() int64 {
	if m != nil && m.Powerts != nil {
		return *m.Powerts
	}
	return 0
}

func (m *RoleCommonInfo) GetHonor() int64 {
	if m != nil && m.Honor != nil {
		return *m.Honor
	}
	return 0
}

func (m *RoleCommonInfo) GetArmyPay() int64 {
	if m != nil && m.ArmyPay != nil {
		return *m.ArmyPay
	}
	return 0
}

func (m *RoleCommonInfo) GetWings() *PbItemInfo {
	if m != nil {
		return m.Wings
	}
	return nil
}

func (m *RoleCommonInfo) GetWeapon() *PbItemInfo {
	if m != nil {
		return m.Weapon
	}
	return nil
}

func (m *RoleCommonInfo) GetCoat() *PbItemInfo {
	if m != nil {
		return m.Coat
	}
	return nil
}

func (m *RoleCommonInfo) GetCloakd() *PbItemInfo {
	if m != nil {
		return m.Cloakd
	}
	return nil
}

func (m *RoleCommonInfo) GetTrouser() *PbItemInfo {
	if m != nil {
		return m.Trouser
	}
	return nil
}

func (m *RoleCommonInfo) GetArmguard() *PbItemInfo {
	if m != nil {
		return m.Armguard
	}
	return nil
}

func (m *RoleCommonInfo) GetShoes() *PbItemInfo {
	if m != nil {
		return m.Shoes
	}
	return nil
}

func (m *RoleCommonInfo) GetFightCardDistri() int32 {
	if m != nil && m.FightCardDistri != nil {
		return *m.FightCardDistri
	}
	return 0
}

func (m *RoleCommonInfo) GetBuyPowerCnt() int32 {
	if m != nil && m.BuyPowerCnt != nil {
		return *m.BuyPowerCnt
	}
	return 0
}

func (m *RoleCommonInfo) GetBuyPowerTs() int64 {
	if m != nil && m.BuyPowerTs != nil {
		return *m.BuyPowerTs
	}
	return 0
}

func (m *RoleCommonInfo) GetGetFreePowerTs() int64 {
	if m != nil && m.GetFreePowerTs != nil {
		return *m.GetFreePowerTs
	}
	return 0
}

func (m *RoleCommonInfo) GetFightValue() int32 {
	if m != nil && m.FightValue != nil {
		return *m.FightValue
	}
	return 0
}

func (m *RoleCommonInfo) GetOnline() int32 {
	if m != nil && m.Online != nil {
		return *m.Online
	}
	return 0
}

func (m *RoleCommonInfo) GetSend() int32 {
	if m != nil && m.Send != nil {
		return *m.Send
	}
	return 0
}

func (m *RoleCommonInfo) GetRecv() int32 {
	if m != nil && m.Recv != nil {
		return *m.Recv
	}
	return 0
}

type RoleOtherInfo struct {
	UnreadMailCmt    *int32 `protobuf:"varint,1,opt,name=unreadMailCmt" json:"unreadMailCmt,omitempty"`
	ChapterId        *int32 `protobuf:"varint,2,opt,name=chapterId" json:"chapterId,omitempty"`
	LevelId          *int32 `protobuf:"varint,3,opt,name=levelId" json:"levelId,omitempty"`
	PointId          *int32 `protobuf:"varint,4,opt,name=pointId" json:"pointId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RoleOtherInfo) Reset()         { *m = RoleOtherInfo{} }
func (m *RoleOtherInfo) String() string { return proto.CompactTextString(m) }
func (*RoleOtherInfo) ProtoMessage()    {}

func (m *RoleOtherInfo) GetUnreadMailCmt() int32 {
	if m != nil && m.UnreadMailCmt != nil {
		return *m.UnreadMailCmt
	}
	return 0
}

func (m *RoleOtherInfo) GetChapterId() int32 {
	if m != nil && m.ChapterId != nil {
		return *m.ChapterId
	}
	return 0
}

func (m *RoleOtherInfo) GetLevelId() int32 {
	if m != nil && m.LevelId != nil {
		return *m.LevelId
	}
	return 0
}

func (m *RoleOtherInfo) GetPointId() int32 {
	if m != nil && m.PointId != nil {
		return *m.PointId
	}
	return 0
}

// 道具信息
type PbItemInfo struct {
	Id               *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Cnt              *int32 `protobuf:"varint,2,opt,name=cnt" json:"cnt,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PbItemInfo) Reset()         { *m = PbItemInfo{} }
func (m *PbItemInfo) String() string { return proto.CompactTextString(m) }
func (*PbItemInfo) ProtoMessage()    {}

func (m *PbItemInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PbItemInfo) GetCnt() int32 {
	if m != nil && m.Cnt != nil {
		return *m.Cnt
	}
	return 0
}

// 背包单元格信息
type PbBagCellInfo struct {
	Indx             *int32      `protobuf:"varint,1,opt,name=indx" json:"indx,omitempty"`
	Item             *PbItemInfo `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PbBagCellInfo) Reset()         { *m = PbBagCellInfo{} }
func (m *PbBagCellInfo) String() string { return proto.CompactTextString(m) }
func (*PbBagCellInfo) ProtoMessage()    {}

func (m *PbBagCellInfo) GetIndx() int32 {
	if m != nil && m.Indx != nil {
		return *m.Indx
	}
	return 0
}

func (m *PbBagCellInfo) GetItem() *PbItemInfo {
	if m != nil {
		return m.Item
	}
	return nil
}

// 背包信息
type PbBagInfo struct {
	BuyCount         *int32           `protobuf:"varint,1,opt,name=buyCount" json:"buyCount,omitempty"`
	Cells            []*PbBagCellInfo `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PbBagInfo) Reset()         { *m = PbBagInfo{} }
func (m *PbBagInfo) String() string { return proto.CompactTextString(m) }
func (*PbBagInfo) ProtoMessage()    {}

func (m *PbBagInfo) GetBuyCount() int32 {
	if m != nil && m.BuyCount != nil {
		return *m.BuyCount
	}
	return 0
}

func (m *PbBagInfo) GetCells() []*PbBagCellInfo {
	if m != nil {
		return m.Cells
	}
	return nil
}

// 卡包信息
type PbCardBagInfo struct {
	Indx             *int32        `protobuf:"varint,1,opt,name=indx" json:"indx,omitempty"`
	Name             *string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Cards            []*PbItemInfo `protobuf:"bytes,3,rep,name=cards" json:"cards,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PbCardBagInfo) Reset()         { *m = PbCardBagInfo{} }
func (m *PbCardBagInfo) String() string { return proto.CompactTextString(m) }
func (*PbCardBagInfo) ProtoMessage()    {}

func (m *PbCardBagInfo) GetIndx() int32 {
	if m != nil && m.Indx != nil {
		return *m.Indx
	}
	return 0
}

func (m *PbCardBagInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PbCardBagInfo) GetCards() []*PbItemInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

// 战斗卡包信息
type PbFightCardBagInfo struct {
	Cards            []*PbCardBagInfo `protobuf:"bytes,1,rep,name=cards" json:"cards,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PbFightCardBagInfo) Reset()         { *m = PbFightCardBagInfo{} }
func (m *PbFightCardBagInfo) String() string { return proto.CompactTextString(m) }
func (*PbFightCardBagInfo) ProtoMessage()    {}

func (m *PbFightCardBagInfo) GetCards() []*PbCardBagInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

// 请求附带信息
// 掉线重连后可以设置token继续操作，token保存一个小时
type RequestInfo struct {
	Token            *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}

func (m *RequestInfo) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type PbCellInfo struct {
	Indx             *int32      `protobuf:"varint,1,opt,name=indx" json:"indx,omitempty"`
	Item             *PbItemInfo `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PbCellInfo) Reset()         { *m = PbCellInfo{} }
func (m *PbCellInfo) String() string { return proto.CompactTextString(m) }
func (*PbCellInfo) ProtoMessage()    {}

func (m *PbCellInfo) GetIndx() int32 {
	if m != nil && m.Indx != nil {
		return *m.Indx
	}
	return 0
}

func (m *PbCellInfo) GetItem() *PbItemInfo {
	if m != nil {
		return m.Item
	}
	return nil
}

// 元宝	40001
// 金币	40002
// 体力	40003
// 荣誉	40004
// 个人贡献	40005
// 军团贡献	40006
// 回廊币	40007
type RoleChangeInfo struct {
	Items            []*PbCellInfo `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *RoleChangeInfo) Reset()         { *m = RoleChangeInfo{} }
func (m *RoleChangeInfo) String() string { return proto.CompactTextString(m) }
func (*RoleChangeInfo) ProtoMessage()    {}

func (m *RoleChangeInfo) GetItems() []*PbCellInfo {
	if m != nil {
		return m.Items
	}
	return nil
}

// Jie
type PbBuyCardInfo struct {
	ExtractCardType  *int32 `protobuf:"varint,1,opt,name=extractCardType" json:"extractCardType,omitempty"`
	FreeCnt          *int32 `protobuf:"varint,2,opt,name=freeCnt" json:"freeCnt,omitempty"`
	Ts               *int64 `protobuf:"varint,3,opt,name=ts" json:"ts,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PbBuyCardInfo) Reset()         { *m = PbBuyCardInfo{} }
func (m *PbBuyCardInfo) String() string { return proto.CompactTextString(m) }
func (*PbBuyCardInfo) ProtoMessage()    {}

func (m *PbBuyCardInfo) GetExtractCardType() int32 {
	if m != nil && m.ExtractCardType != nil {
		return *m.ExtractCardType
	}
	return 0
}

func (m *PbBuyCardInfo) GetFreeCnt() int32 {
	if m != nil && m.FreeCnt != nil {
		return *m.FreeCnt
	}
	return 0
}

func (m *PbBuyCardInfo) GetTs() int64 {
	if m != nil && m.Ts != nil {
		return *m.Ts
	}
	return 0
}

// 响应通用信息
type RespondInfo struct {
	Obtain           *RoleChangeInfo `protobuf:"bytes,1,opt,name=obtain" json:"obtain,omitempty"`
	Consume          *RoleChangeInfo `protobuf:"bytes,2,opt,name=consume" json:"consume,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RespondInfo) Reset()         { *m = RespondInfo{} }
func (m *RespondInfo) String() string { return proto.CompactTextString(m) }
func (*RespondInfo) ProtoMessage()    {}

func (m *RespondInfo) GetObtain() *RoleChangeInfo {
	if m != nil {
		return m.Obtain
	}
	return nil
}

func (m *RespondInfo) GetConsume() *RoleChangeInfo {
	if m != nil {
		return m.Consume
	}
	return nil
}
