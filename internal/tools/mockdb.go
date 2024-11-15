package tools

import(
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"HelloWorld":{
		AuthToken: "100",
		Username: "HelloWorld",
	},
	"HelloMars":{
		AuthToken: "200",
		Username: "HelloMars",
	},
}

var mockBalanceDetail = map[string]BalanceDetails{
	"HelloWorld":{
		Balance: 1000,
		Username: "HelloWorld",
	},
	"HelloMars":{
		Balance: 2000,
		Username: "HelloMars",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetBankBalance(username string) *BalanceDetails{
	time.Sleep(time.Second * 1)
	var clientData = BalanceDetails{}
	clientData, ok := mockBalanceDetail[username]
	if !ok{
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error  {
	return nil
}