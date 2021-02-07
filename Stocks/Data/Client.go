package Data

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const connectionString string = "user=postgres password=123 dbname=trades sslmode=disable"

//Создаем таблицу сделок
var schema = `
CREATE TABLE trade (
    id bigint,
    price real,
    quantity real,
	buyerId bigint,
	sellerId bigint,
	market bigint
);`


//Для простоты считаем только по тем кто что то покупал, по формуле: сумма покупки - сумма продажи
var accountsQuery = `
SELECT t1.debet - coalesce(t2.credit, 0) as account, t1.buyerid as accountid FROM 
(SELECT SUM(price * quantity) as debet, buyerid FROM trade GROUP BY buyerid) as t1
LEFT  JOIN (
	SELECT SUM(price * quantity) as credit, sellerId FROM trade GROUP BY sellerId
) as t2 ON t1.buyerid = t2.sellerid`

type Client struct {
	db *sqlx.DB
}

//Подключение к бд и получения клиента для запросов
func GetSqlClient() Client {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	//Запрос, проверяющий что таблица существует
	_, isExist := db.Query("SELECT * FROM trade limit 1")

	if isExist != nil{
		_, err = db.Exec(schema)
		if err != nil{
			log.Fatal("Не удалось применить схему ", err)
		}
	}

	return Client{db}
}

//Вставка батча записей
func (client Client) BulkInsert(trades []*Trade){
	query := generateInsertQuery(trades)
	_, err := client.db.Exec(query)
	if err != nil{
		log.Fatal("Не удалось вставить записи ", err)
	}
}

//Запрос на получение баланса
func (client Client) GetAccountsBalance() []Account{
	accounts := make([]Account, 0, 100)
	err := client.db.Select(&accounts, accountsQuery)
	if err != nil{
		log.Fatal("Не удалось посчитать баланс ", err)
	}
	return  accounts
}

//Генерация запроса на вставку большого количества значений
//TODO переделать на COPY
func generateInsertQuery(trades []*Trade) string {
	query := "INSERT INTO trade VALUES "
	for _, t := range trades{
		query += fmt.Sprintf(
			"(%d, %f, %f, %d, %d, %d),",t.Id, t.Price, t.Quantity, t.BuyerId, t.SellerId, t.Market)
	}
	length := len(query)
	query = query[:length-1]
	return  query
}