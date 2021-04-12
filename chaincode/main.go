package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
)

type AccountChaincode struct {

}

func (t *AccountChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *AccountChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addAcc"{
		return t.addAcc(stub, args)		// 添加信息
	}else if fun == "queryAccInfoByEntityID" {
		return t.queryAccInfoByEntityID(stub, args)	// 根据EntityID查询详情
	}else if fun == "updateAcc" {
		return t.updateAcc(stub, args)		// 根据EntityID更新信息
	}else if fun == "delAcc"{
		return t.delAcc(stub, args)	// 根据EntityID删除信息
	}else if fun == "transferAcc" {
		return t.transferAcc(stub, args) //转账
	}

	return shim.Error("指定的函数名称错误")
}

func main(){
	err := shim.Start(new(AccountChaincode))
	if err != nil{
		fmt.Printf("启动AccountChaincode时发生错误: %s", err)
	}
}
