package main

type Account struct {
	ObjectType string `json:"docType"`
	Name       string `json:"Name"`     // 姓名
	Balance    int64  `json:"Balance"`  // 余额
	EntityID   string `json:"EntityID"` // EntityID

	History []string `json:"History"`
	Historys []HistoryItem // 当前acc的历史记录
}

type HistoryItem struct {
	TxId    string
	Account Account
}
