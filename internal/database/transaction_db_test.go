package database

import (
	"database/sql"
	"testing"

	"github.com.br/limaantonio/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	accountDB     *AccountDB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE accounts (id TEXT, client_id TEXT, balance REAL, created_at date)")
	db.Exec("CREATE TABLE clients (id TEXT, name TEXT, email TEXT, created_at date)")
	db.Exec("CREATE TABLE transactions (id TEXT, account_id_from TEXT, account_id_to TEXT, amount REAL, created_at date)")
	s.accountDB = NewAccountDB(db)
	s.transactionDB = NewTransactionDB(db)
	s.client, _ = entity.NewClient("Antonio", "a.com")
	s.client2, _ = entity.NewClient("Antonio", "a.com")
	s.accountFrom = entity.NewAccount(s.client)
	s.accountTo = entity.NewAccount(s.client2)

	s.transactionDB = NewTransactionDB(db)

}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 10)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
