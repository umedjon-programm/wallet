package wallet

import (
	"errors"

	"github.com/google/uuid"
	"github.com/umedjon-programm/wallet/pkg/types"
)
type Service struct{
	nextAccountID int64
	accounts []*types.Account
	payments []*types.Payment
}

var ErrAccountNotFound=errors.New("account not found")
var ErrPaymentNotFound=errors.New("payment not found")
var ErrPhoneRegistered=errors.New("phone already registered")
var ErrAmountMustBePositive=errors.New("amount must be greater than 0")
var ErrNotEnoughBalance=errors.New("not enough balance")
func (s *Service) RegisterAccount(phone types.Phone)(*types.Account,error){
	for _,account:=range s.accounts{
		if account.Phone==phone{
			return nil,ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account:=&types.Account{
		ID: s.nextAccountID,
		Phone: phone,
		Balance: 0,
	}
	s.accounts=append(s.accounts,account )
	return account,nil

}
func (s *Service) Deposit(accountID int64, amount types.Money) error{
	if amount<=0{
		return ErrAmountMustBePositive
	}
	account, err:=s.FindAccountByID(accountID)
	if account==nil{
		return err
	} 
	account.Balance+=amount
	return nil
}
func (s *Service) Pay(accountID int64, amount types.Money,category types.PaymentCategory)(*types.Payment,error){
	if amount<=0{
		return nil, ErrAmountMustBePositive
	}
	account,err:=s.FindAccountByID(accountID)
	if account==nil{
		return nil,err
	}
	if account.Balance<amount{
		return nil, ErrNotEnoughBalance
	}
	account.Balance-=amount
	paymentID:=uuid.New().String()
	payment:=&types.Payment{
		ID: paymentID,
		AccountID: accountID,
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}
	s.payments=append(s.payments, payment)
	return payment, nil
}

func(s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
		var account *types.Account
	for _,acc:=range s.accounts{
		if accountID==acc.ID{
			account=acc
			return account,nil
		}
	} 
	return nil,ErrAccountNotFound
	
}
func(s *Service) FindPaymentByID(paymentID string)(*types.Payment,error){
	var payment *types.Payment
	for _,pay:=range s.payments{
		if paymentID==pay.ID{
			payment=pay
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}

func (s *Service) Repeat(paymentID string)(*types.Payment, error){
	payment,err:=s.FindPaymentByID(paymentID)
	if payment==nil{
		return nil, err
	}
	account,err:=s.FindAccountByID(payment.AccountID)
	if account==nil{
		return nil, err
	}
	return s.Pay(account.ID,payment.Amount,payment.Category)
}