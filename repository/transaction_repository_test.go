package repository

import (
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/resp"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    TransactionRepository
}

func (suite *TransactionRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSQL = mock
	suite.repo = NewTransactionRepository(suite.mockDb)
}

func TestTransactionRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepoTestSuite))
}

func (suite *TransactionRepoTestSuite) TestCreateTransaction_Success() {
	mockData := model.Transactions{
		Id:              "1",
		SourceWalletID:  "001",
		UserId:          "1",
		PaymentMethodID: "dana",
		Destination:     "pln",
		Amount:          50000,
		Description:     "tagihan listrik",
		CreateAt:        time.Time{},
	}
	rows := sqlmock.NewRows([]string{"id", "Source_wallet_ID", "user_Id", "Payment_Method_id", "destination", "amount", "description", "created_at"})
	rows.AddRow(mockData.Id, mockData.SourceWalletID, mockData.UserId, mockData.PaymentMethodID, mockData.Destination, mockData.Amount, mockData.Description, mockData.CreateAt)
	suite.mockSQL.ExpectQuery("INSERT INTO transactions").
		WithArgs(mockData.SourceWalletID, mockData.UserId, mockData.PaymentMethodID, mockData.Destination, mockData.Amount, mockData.Description, mockData.CreateAt).
		WillReturnRows(rows)
	result, err := suite.repo.CreateTransaction(mockData)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *TransactionRepoTestSuite) TestCreateTransaction_Fail() {
	mockData := model.Transactions{
		Id:              "1",
		SourceWalletID:  "001",
		UserId:          "1",
		PaymentMethodID: "dana",
		Destination:     "pln",
		Amount:          50000,
		Description:     "tagihan listrik",
		CreateAt:        time.Time{},
	}

	// TODO create transaction failed
	rows := sqlmock.NewRows([]string{"id", "source_wallet_ID", "user_Id", "Payment_Method_id", "destination", "amount", "description", "created_at"})
	rows.AddRow(mockData.Id, mockData.SourceWalletID, mockData.UserId, mockData.PaymentMethodID, mockData.Destination, mockData.Amount, mockData.Description, mockData.CreateAt)
	suite.mockSQL.ExpectExec("INSERT INTO transactions").
		WithArgs(mockData.Id, mockData.SourceWalletID, mockData.UserId, mockData.PaymentMethodID, mockData.Destination, mockData.Amount, mockData.Description, mockData.CreateAt).
		WillReturnError(errors.New("Insertion failed"))
	result, err := suite.repo.CreateTransaction(mockData)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.NoError(suite.T(), suite.mockSQL.ExpectationsWereMet())
}

func (suite *TransactionRepoTestSuite) TestFindAll_success() {
	mockData := []resp.GetTransactionsResponse{{
		Id:          "1",
		Destination: "1",
		Amount:      0,
		Description: "1",
		CreateAt:    time.Time{},
		User: resp.User{
			UserName: "sai",
		},
		Wallet: resp.Wallet{
			RekeningUser: "",
			Balance:      0,
		},
		PaymentMethod: resp.PaymentMethod{
			Name:        "",
			Description: "",
		},
	}}

	rows := sqlmock.NewRows([]string{"id", "detination", "amount", "description", "create_at",
		"user_username", "wallet_rekening_user", "wallet_balance", "payment_method_name", "payment_method_description"})
	for _, t := range mockData {
		rows.AddRow(t.Id, t.Destination, t.Amount, t.Description, t.CreateAt, t.User.UserName, t.Wallet.RekeningUser, t.Wallet.Balance, t.PaymentMethod.Name, t.PaymentMethod.Description)
	}
	expectedQuery := `
       SELECT
           t.id,
           t.destination,
           t.amount,
           t.description,
           t.created_at,
           u.user_name,
           w.rekening_user,
           w.balance,
           p.name,
           p.description
       FROM transactions AS t
       JOIN users AS u ON t.user_id = u.id
		JOIN wallets AS w ON t.source_wallet_id = w.id
		JOIN payment_method AS p ON t.payment_method_id = p.id;`

	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnRows(rows)
	_, got := suite.repo.FindAll("1")
	assert.Nil(suite.T(), got)
}
