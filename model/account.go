package model

type Account struct {
	Id               int64  `gorm:"column:id;primaryKey"`
	Title            string `gorm:"title"`
	Account          string `gorm:"account"`
	TokenAddress     string `gorm:"token_address"`
	TokenName        string `gorm:"token_name"`
	BalanceThreshold string `gorm:"balance_threshold"`
	NotifyWho        string `gorm:"notify_who"`
	Rpc              string `gorm:"rpc"`
	Remark           string `gorm:"remark"`
}
