package models

import (
	"github.com/sabouaram/transaction/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	_, err := sql.Clientdb.Exec("DELETE FROM users")
	if err != nil {
		log.Fatalf("Unable to clean the users table")
	}
}

func TestCreateUser(t *testing.T) {

	user := User{
		Id:      "03eb9399-e526-431f-812f-2fda01659000",
		Name:    "Test User",
		Balance: 1000,
	}
	err := user.CreateUser()

	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	user := User{
		Id:      "03eb9399-e526-431f-812f-2fda01659011",
		Name:    "Test User",
		Balance: float64(1000),
	}
	user.CreateUser()
	err := user.GetUser()
	assert.NoError(t, err)
	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, float64(1000), user.Balance)
}

func TestUpdateUser(t *testing.T) {
	user := User{
		Id:      "03eb9399-e526-431f-812f-2fda01659222",
		Name:    "Test User",
		Balance: float64(1000),
	}
	user.CreateUser()

	user.Name = "Updated User"
	user.Balance = float64(500)
	err := user.UpdateUser()
	assert.NoError(t, err)
	user.GetUser()
	assert.Equal(t, "Updated User", user.Name)
	assert.Equal(t, float64(500), user.Balance)
}

func TestDeleteUser(t *testing.T) {
	user := User{
		Id:      "03eb9399-e526-431f-812f-2fda01659999",
		Name:    "Test User",
		Balance: float64(1000),
	}
	user.CreateUser()
	err := user.DeleteUser()
	assert.NoError(t, err)
	err = user.GetUser()
	assert.Error(t, err)
}
