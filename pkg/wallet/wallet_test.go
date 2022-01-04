package wallet_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/umedjon-programm/wallet/pkg/wallet"
)

func TestService_Repeat_success(t *testing.T) {
	svc:=&wallet.Service{}
	account,err:=svc.RegisterAccount("+992920000001")
	if account==nil{
		t.Error(err)
		return
	}
	err=svc.Deposit(account.ID,10000000)
	if err!=nil{
		t.Error(err)
		return
	}
	payment,err:=svc.Pay(account.ID,1000000,"food")
	if payment==nil{
		t.Error(err)
		return
	}
	payment,err=svc.Repeat(payment.ID)
	if payment==nil{
		t.Error(err)
		return
	}
}
func TestService_Repeat_fail(t *testing.T) {
	svc:=&wallet.Service{}
	account,err:=svc.RegisterAccount("+992920000001")
	if account==nil{
		t.Error(err)
		return
	}
	err=svc.Deposit(account.ID,10000000)
	if err!=nil{
		t.Error(err)
		return
	}
	payment,err:=svc.Pay(account.ID,10000000,"food")
	if payment==nil{
		t.Error(err)
		return
	}
	_,err=svc.Repeat(uuid.New().String())
	if err==nil{
		t.Error("must return error")
		return
	}
}
func TestSevice_Reject_nil(t *testing.T) {
	svc:=&wallet.Service{}
	account,err:=svc.RegisterAccount("+992920000001")
	if account==nil{
		t.Error(err)
		return
	}
	err=svc.Deposit(account.ID,10000000)
	if err!=nil{
		t.Error(err)
		return
	}
	payment,err:=svc.Pay(account.ID,10000000,"food")
	if payment==nil{
		t.Error(err)
		return
	}
	err=svc.Reject(payment.ID)
	if err!=nil{
		t.Error(err)
		return
	}
}
func TestSevice_Reject_fail(t *testing.T) {
	svc:=&wallet.Service{}
	account,err:=svc.RegisterAccount("+992920000001")
	if account==nil{
		t.Error(err)
		return
	}
	err=svc.Deposit(account.ID,10000000)
	if err!=nil{
		t.Error(err)
		return
	}
	payment,err:=svc.Pay(account.ID,10000000,"food")
	if payment==nil{
		t.Error(err)
		return
	}
	err=svc.Reject(uuid.New().String())
	if err==nil{
		t.Error(err)
		return
	}
}
func Test_Favorite_success(t *testing.T) {
	svc:=&wallet.Service{}
	account,err:=svc.RegisterAccount("+992920000001")
	if account==nil{
		t.Error(err)
		return
	}
	err=svc.Deposit(account.ID,10000000)
	if err!=nil{
		t.Error(err)
		return
	}
	payment,err:=svc.Pay(account.ID,1000000,"food")
	if payment==nil{
		t.Error(err)
		return
	}
	favorite, err:=svc.FavoritePayment(payment.ID, "Barq")
	if favorite==nil{
		t.Error(err)
		return
	}
	payment,err=svc.PayFromFavorite(favorite.ID)
	if payment==nil{
		t.Error(err)
		return
	}
}
func Test_Favorite_fail(t *testing.T) {
	svc:=&wallet.Service{}
	account,err:=svc.RegisterAccount("+992920000001")
	if account==nil{
		t.Error(err)
		return
	}
	err=svc.Deposit(account.ID,10000000)
	if err!=nil{
		t.Error(err)
		return
	}
	payment,err:=svc.Pay(account.ID,1000000,"food")
	if payment==nil{
		t.Error(err)
		return
	}
	favorite, err:=svc.FavoritePayment(payment.ID, "Barq")
	if favorite==nil{
		t.Error(err)
		return
	}
	_,err=svc.PayFromFavorite(uuid.NewString())
	if err==nil{
		t.Error(err)
		return
	}
}