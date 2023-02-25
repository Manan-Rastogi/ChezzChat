package models

import (
	"testing"

	
	"github.com/stretchr/testify/require"
)

var TestDB = Init()

func TestCreateNewUser(t *testing.T) {
	// helpers.Logger.Info("CreateNewUser Test Started!")

	user1 := Users{
		Name: "Manan Rastogi",
		Email: "manan.rastogi@chezz.com",
		Password: "MananIsGreat",
	}

	err := CreateNewUser(TestDB, user1)
	require.NoError(t, err)
}
