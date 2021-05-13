package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) SaveAcc(acc Account) (string, error) {
	eventID := "eventAddAcc"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将acc对象序列化成为字节数组
	b, err := json.Marshal(acc)
	if err != nil {
		return "", fmt.Errorf("指定的acc对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addAcc", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) FindAccInfoByEntityID(entityID string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryAccInfoByEntityID", Args: [][]byte{[]byte(entityID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindAccByCertNoAndName(certNo, name string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryAccByCertNoAndName", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) ModifyAcc(acc Account) (string, error) {

	eventID := "eventModifyAcc"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将acc对象序列化成为字节数组
	b, err := json.Marshal(acc)
	if err != nil {
		return "", fmt.Errorf("指定的acc对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateAcc", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) DelAcc(entityID string) (string, error) {

	eventID := "eventDelAcc"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "delAcc", Args: [][]byte{[]byte(entityID), []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) TransferAcc(id1 string, id2 string, amount int64) (string, error) {

	if amount < 0 {
		return "", fmt.Errorf("不能转账为负数")
	}

	eventID := "eventTransferAcc"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	var negativeAmount = amount * (-1)

	acc1 := Account{
		Name:     "",
		Balance:  negativeAmount,
		EntityID: id1,
		History:  []string{"To " + id2 + " Amount: " + strconv.FormatInt(amount, 10) + ";"},
	}

	// 将acc1对象序列化成为字节数组
	b, err := json.Marshal(acc1)
	if err != nil {
		return "", fmt.Errorf("指定的acc对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "transferAcc", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	acc2 := Account{
		Name:     "",
		Balance:  amount,
		EntityID: id2,
		History:  []string{"From " + id1 + " Amount: " + strconv.FormatInt(amount, 10) + ";"},
	}

	// 将acc2对象序列化成为字节数组
	b, err = json.Marshal(acc2)
	if err != nil {
		return "", fmt.Errorf("指定的acc对象序列化时发生错误")
	}

	req = channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "transferAcc", Args: [][]byte{b, []byte(eventID)}}
	respone, err = t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}
