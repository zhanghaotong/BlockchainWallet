/**
 * @auther Zhang Haotong
 *
 */
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"strconv"
)

const DOC_TYPE = "accObj"

// 保存acc
// args: account
func PutAcc(stub shim.ChaincodeStubInterface, acc Account) ([]byte, bool) {

	acc.ObjectType = DOC_TYPE
	b, err := json.Marshal(acc)
	if err != nil {
		return nil, false
	}

	// 保存acc状态
	err = stub.PutState(acc.EntityID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

// 根据EntityID查询信息状态
// args: entityID
func GetAccInfo(stub shim.ChaincodeStubInterface, entityID string) (Account, bool)  {
	var acc Account
	// 根据EntityID查询信息状态
	b, err := stub.GetState(entityID)
	if err != nil {
		return acc, false
	}

	if b == nil {
		return acc, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &acc)
	if err != nil {
		return acc, false
	}

	// 返回结果
	return acc, true
}

// 添加信息
// args: accountObject
// EntityID为 key, Account 为 value
func (t *AccountChaincode) addAcc(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var acc Account
	err := json.Unmarshal([]byte(args[0]), &acc)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: EntityID必须唯一
	_, exist := GetAccInfo(stub, acc.EntityID)
	if exist {
		return shim.Error("要添加的EntityID已存在")
	}

	acc.History = append(acc.History, "Account created, initial balance = " + strconv.FormatInt(acc.Balance,10) + ";" )

	_, bl := PutAcc(stub, acc)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根据EntityID查询详情（溯源）
// args: entityID
func (t *AccountChaincode) queryAccInfoByEntityID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据EntityID查询acc状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据EntityID查询信息失败")
	}

	if b == nil {
		return shim.Error("根据EntityID没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var acc Account
	err = json.Unmarshal(b, &acc)
	if err != nil {
		return  shim.Error("反序列化acc信息失败")
	}

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(acc.EntityID)
	if err != nil {
		return shim.Error("根据指定的EntityID查询对应的历史变更数据失败")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []HistoryItem
	var hisAcc Account
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取acc的历史变更数据失败")
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisAcc)

		if hisData.Value == nil {
			var empty Account
			historyItem.Account = empty
		}else {
			historyItem.Account = hisAcc
		}

		historys = append(historys, historyItem)

	}

	acc.Historys = historys

	// 返回
	result, err := json.Marshal(acc)
	if err != nil {
		return shim.Error("序列化acc信息时发生错误")
	}
	return shim.Success(result)
}

// 根据EntityID更新信息
// args: accountObject
func (t *AccountChaincode) updateAcc(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var info Account
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return  shim.Error("反序列化acc信息失败")
	}

	// 根据EntityID查询信息
	result, bl := GetAccInfo(stub, info.EntityID)
	if !bl{
		return shim.Error("根据Entity查询信息时发生错误")
	}

	result.Name = info.Name
	result.Balance = info.Balance
	result.EntityID = info.EntityID
	
	_, bl = PutAcc(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

// 根据EntityID删除信息
// args: entityID
func (t *AccountChaincode) delAcc(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	/*var acc Account
	result, bl := GetAccInfo(stub, info.EntityID)
	err := json.Unmarshal(result, &acc)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息删除成功"))
}

// 根据EntityID转账
// args: accountObject
func (t *AccountChaincode) transferAcc(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var info Account
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return  shim.Error("反序列化acc信息失败")
	}

	// 根据EntityID查询信息
	result, bl := GetAccInfo(stub, info.EntityID)
	if !bl{
		return shim.Error("根据EntityID查询信息时发生错误")
	}

	result.Balance = result.Balance + info.Balance
	
	if result.Balance < 0 {
		return shim.Error("转账金额大于余额，发生错误")
	}
	
	for _, s := range info.History {
		result.History = append(result.History,s)
	}

	_, bl = PutAcc(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}