package Data

type Account struct {
	Balance float64 `db:"account"`
	AccountId int64 `db:"accountid"`
}
