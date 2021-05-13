/**
 * @auther Zhang Haotong
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kongyixueyuan.com/education/sdkInit"
	"github.com/kongyixueyuan.com/education/service"
	"github.com/kongyixueyuan.com/education/web"
	"github.com/kongyixueyuan.com/education/web/controller"
)

const (
	configFile  = "config.yaml"
	initialized = false
	EduCC       = "educc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/kongyixueyuan.com/education/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     EduCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/kongyixueyuan.com/education/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID: EduCC,
		Client:      channelClient,
	}

	acc := service.Account{
		Name:     "张三",
		Balance:  10,
		EntityID: "101",
	}

	acc2 := service.Account{
		Name:     "李四",
		Balance:  20,
		EntityID: "102",
	}

	acc3 := service.Account{
		Name:     "admin",
		Balance:  999999999,
		EntityID: "admin",
	}

	msg, err := serviceSetup.SaveAcc(acc)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	msg, err = serviceSetup.SaveAcc(acc2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	msg, err = serviceSetup.SaveAcc(acc3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	// 根据EntityID查询信息
	result, err := serviceSetup.FindAccInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var acc service.Account
		json.Unmarshal(result, &acc)
		fmt.Println("根据EntityID查询信息成功：")
		fmt.Println(acc)
	}

	// 修改/添加信息
	info := service.Account{
		Name:     "张三",
		Balance:  200,
		EntityID: "101",
	}
	msg, err = serviceSetup.ModifyAcc(info)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	//Transfer
	msg, err = serviceSetup.TransferAcc("101", "102", 10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	//Transfer
	msg, err = serviceSetup.TransferAcc("101", "102", 30)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	//Transfer
	msg, err = serviceSetup.TransferAcc("102", "101", 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	// 根据EntityID查询信息
	result, err = serviceSetup.FindAccInfoByEntityID("102")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var acc service.Account
		json.Unmarshal(result, &acc)
		fmt.Println("根据EntityID查询信息成功：")
		fmt.Println(acc)

	}

	// 根据EntityID查询信息
	result, err = serviceSetup.FindAccInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var acc service.Account
		json.Unmarshal(result, &acc)
		fmt.Println("根据EntityID查询信息成功：")
		fmt.Println(acc)

	}

	//===========================================//

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
