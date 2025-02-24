package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/Zulhaidir/microservice/mantap/db/mock"
	db "github.com/Zulhaidir/microservice/mantap/db/sqlc"
	"github.com/Zulhaidir/microservice/mantap/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAccountGetAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	/* build stubs */
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	/* start test and send request */
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	/* check response */
	require.Equal(t, http.StatusOK, recorder.Code)

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
