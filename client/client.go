/*
*
@author: Administrator
@Date: 2022-12-01-0001
@Note
*
*/
package client

import (
	"fmt"
	"go_test/log"
	"go_test/network"
	"go_test/network/router"
	"go_test/network/tcp"
	"go_test/proto/pb"
	"time"

	"google.golang.org/protobuf/proto"
)

type ClientStruct struct {
	conn    network.IConn
	IsClose chan bool
}

func CreateClient(addr string) {
	route := router.RouterMgr
	client := &ClientStruct{
		IsClose: make(chan bool, 1),
	}
	defer close(client.IsClose)
	tcpConn := tcp.NewClient(tcp.NewPacket(), addr, network.WithMax(10, time.Now().UnixNano()/1e6, nil), route, client.Stop).Connect()
	if tcpConn == nil {
		client.IsClose <- true
	} else {
		client.conn = tcpConn
		fmt.Println("链接成功!!!!")
		// 注册回调
		client.addRouter(route)
		// 发送登录协议
		client.conn.Send(&pb.Role_Login_C2S{
			Iggid: "222211112316",
		})
	}
	<-client.IsClose
	log.AppLogger.Info("客户端关闭")
}

func (*ClientStruct) getRoleData(conn network.IConn, data []byte) {
	// 登录返回
	// conn.Send(&pb.Hero_CostPropComposite_C2S{Heroid: 6001}) // 合成武将
	conn.Send(&pb.Role_Gmcmd_C2S{Cmd: 1001, Args: "[{\"id\" : 3401, \"num\" : 10}]"}) // gm添加物品
	// conn.Send(&pb.Hero_UseExpProp_C2S{Heroid: 6001, Itemid: 1981, Num: 1})            // 武将使用经验书
	// conn.Send(&pb.Hero_UseExpProp_C2S{Heroid: 600101, Itemid: 1981, Num: 1}) // 武将升级
	// conn.Send(&pb.Role_Gmcmd_C2S{Cmd: 1005, Args: "{\"id\" : 0}"}) // gm添加武将
	// conn.Send(&pb.Role_Gmcmd_C2S{Cmd: 1007, Args: "{\"id\" : 600201,\"num\" : 5000000000}"}) // gm增加武将经验
	// conn.Send(&pb.Role_Gmcmd_C2S{Cmd: 1008, Args: "{\"id\" : 0,\"level\" : 20}"}) // gm武将达到指定等级

	// conn.Send(&pb.Role_Gmcmd_C2S{Cmd: 1009, Args: "[{\"id\" : 1110101, \"num\" : 10},{\"id\" : 1120101, \"num\" : 10}]"}) // gm添加士兵
	// conn.Send(&pb.Role_Gmcmd_C2S{Cmd: 1011, Args: "{\"areaid\" : 1, \"id\" : 1020101,\"pos\" : {\"x\":1,\"y\":1}}"}) // gm添加建筑
	// ----------------部队相关
	// conn.Send(&pb.Army_ListInfo_C2S{}) // 获取部队简略信息
	// conn.Send(&pb.Army_Create_C2S{Armyinfo: &pb.ArmyCreateInfo{ // 创建部队
	// 	Heros: []int32{6001},
	// 	Soldiers: []*pb.ArmyCreateSoldierInfo{
	// 		&pb.ArmyCreateSoldierInfo{Id: 1110101, Num: 99},
	// 		&pb.ArmyCreateSoldierInfo{Id: 1120101, Num: 1},
	// 	},
	// }})
	// conn.Send(&pb.Army_DisbandArmy_C2S{Armyid: 1}) //解散部队
	// conn.Send(&pb.Mail_SendPrivate_C2S{Playerid: 1, Title: "111", Content: "2222"})
	// conn.Send(&pb.Military_GetInfo_C2S{Areaid: 1, Buildingid: 1})
	// conn.Send(&pb.Military_SoldierTrain_C2S{Areaid: 1, Buildingid: 28, Soldierid: 1110101, Trainnum: 100, IsBuy: false}) // 训练士兵
	conn.Send(&pb.Military_Upgrade_C2S{Areaid: 1, Buildingid: 28, Soldierid: 1110101, Trainnum: 100, IsBuy: false}) // 晋升士兵
	// conn.Send(&pb.Military_UseSpeedProp_C2S{Areaid: 1, Buildingid: 28, Itemid: 3401, Itemnum: 1}) // 使用加速道具
	// conn.Send(&pb.Military_BuyTimeFinish_C2S{Areaid: 1, Buildingid: 28}) // 立即完成
	// conn.Send(&pb.Military_CancelProduction_C2S{Areaid: 1, Buildingid: 1}) // 取消生产
	// conn.Send(&pb.Military_CollectSoldier_C2S{Areaid: 1, Buildingid: 28}) // 收取士兵

}

func (server *ClientStruct) Stop(conn network.IConn) {
	server.IsClose <- true
}

func (client *ClientStruct) addRouter(router router.IRouter) {
	router.AddRouter(&pb.Role_Login_S2C{}, client.getRoleData)
	router.AddRouter(&pb.Military_Upgrade_S2C{}, client.militaryUpgrade)
}

func (*ClientStruct) militaryUpgrade(conn network.IConn, data []byte) {
	msg := &pb.Military_Upgrade_S2C{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("解析错误")
	}
}
