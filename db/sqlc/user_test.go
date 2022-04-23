package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/tornvallalexander/go-backend-template/utils"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomPassword(8))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Username:       utils.RandomUsername(),
		HashedPassword: hashedPassword,
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	require.NotZero(t, user.Username)
	require.NotZero(t, user.HashedPassword)
	require.NotZero(t, user.Email)

	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.Email, arg.Email)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	_, err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

// Not entirely applicable in this case, but this is how we would
// test a list functionality
func TestListUsers(t *testing.T) {
	// var lastUser User
	for i := 0; i < 10; i++ {
		// lastUser = createRandomUser(t)
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	/*
		for _, user := range users {
			require.NotEmpty(t, user)
			require.Equal(t, user.Username, lastUser.Username)
		}
	*/
}
