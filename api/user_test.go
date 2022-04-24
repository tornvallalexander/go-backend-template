package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/tornvallalexander/go-backend-template/db/mock"
	db "github.com/tornvallalexander/go-backend-template/db/sqlc"
	"github.com/tornvallalexander/go-backend-template/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserAPI(t *testing.T) {
	user, _ := randomUser(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetUser(gomock.Any(), gomock.Eq(user.Username)).
		Times(1).
		Return(user, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/users/%s", user.Username)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchUser(t, recorder.Body, user)
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = utils.RandomPassword(8)
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:          utils.RandomUsername(),
		HashedPassword:    hashedPassword,
		Email:             utils.RandomEmail(),
		PasswordChangedAt: utils.RandomDate(),
		CreatedAt:         utils.RandomDate(),
	}

	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
