package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@.com", client.Email)

}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)

}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@.com")
	err := client.Update("John Doe Update", "j@.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@.com", client.Email)

}

func TestUpdateClientWithArgsInvalid(t *testing.T) {
	client, _ := NewClient("John Doe Update", "j@.com")
	err := client.Update("", "j@.com")
	assert.Error(t, err, "name is required")

}

func TestCreateAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, client.ID, account.Client.ID)

}
