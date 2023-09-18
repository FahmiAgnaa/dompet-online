package usecase

import (
	"Kelompok-2/dompet-online/mock/repository_mock"
	"Kelompok-2/dompet-online/mock/usecase_mock"
	"Kelompok-2/dompet-online/model"
	"Kelompok-2/dompet-online/model/dto/req"
	"Kelompok-2/dompet-online/model/dto/resp"
	"Kelompok-2/dompet-online/util/common"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TransactionUsecaseTestSuite struct {
	suite.Suite
	transactionRepoMock *repository_mock.TransactionRepositoryMock
	walletUsecaseMock   *usecase_mock.WalletUseCaseMock
	transactionUseCase  TransactionUseCase
}

func (suite *TransactionUsecaseTestSuite) SetupTest() {
	suite.transactionRepoMock = new(repository_mock.TransactionRepositoryMock)
	suite.transactionUseCase = NewTransactionUseCase(suite.transactionRepoMock, suite.walletUsecaseMock)
}

func TestTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(WalletUsecaseTestSuite))
}

func (suite *TransactionUsecaseTestSuite) TestGetHistoryTransactions_Success() {

	mockData := []resp.GetTransactionsResponse{{
		Id:            "1",
		Destination:   "",
		Amount:        0,
		Description:   "",
		CreateAt:      time.Time{},
		User:          resp.User{},
		Wallet:        resp.Wallet{},
		PaymentMethod: resp.PaymentMethod{},
	}}
	userId := "1"
	suite.transactionRepoMock.On("FindAll", userId).Return(mockData, nil)
	result, err := suite.transactionUseCase.GetHistoryTransactions(userId)
	suite.transactionRepoMock.AssertExpectations(suite.T())
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *TransactionUsecaseTestSuite) TestGetHistoryTransactions_Fail() {

	mockData := []resp.GetTransactionsResponse{{
		Id:          "",
		Destination: "",
		Amount:      0,
		Description: "",
		CreateAt:    time.Time{},
		User: resp.User{
			UserName: "",
		},
		Wallet:        resp.Wallet{},
		PaymentMethod: resp.PaymentMethod{},
	}}
	userId := "1"
	suite.transactionRepoMock.On("FindAll", userId).Return(nil, errors.New("Simulated failed"))
	result, err := suite.transactionUseCase.GetHistoryTransactions(userId)
	suite.transactionRepoMock.AssertExpectations(suite.T())
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *TransactionUsecaseTestSuite) TestTopUp_Success(payload req.TopUpRequest) {
	payload = req.TopUpRequest{
		UserId:          "1",
		WalletID:        "1",
		Amount:          100.0,
		PaymentMethodId: "3",
	}
	mockData := model.Transactions{
		Id:              common.GenerateID(),
		UserId:          payload.UserId,
		SourceWalletID:  payload.WalletID,
		Destination:     "TopUp",
		Amount:          payload.Amount,
		Description:     "TopUp",
		PaymentMethodID: payload.PaymentMethodId,
		CreateAt:        time.Now(),
	}
	suite.transactionRepoMock.On("CreateTransaction", mock.Anything).Return(mockData, nil)
	suite.walletUsecaseMock.On("UpdateWalletBalance", mock.Anything, mock.Anything).Return(nil)
	result, err := suite.transactionUseCase.TopUp(payload)
	suite.transactionRepoMock.AssertExpectations(suite.T())
	suite.walletUsecaseMock.AssertExpectations(suite.T())
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *TransactionUsecaseTestSuite) TestTopUp_Fail() {

}

func (suite *TransactionUsecaseTestSuite) TestTransfer_Success() {

}

func (suite *TransactionUsecaseTestSuite) TestTransfer_Fail() {

}
