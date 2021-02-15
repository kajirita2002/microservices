package repository

import (
	"fmt"
	"log"
	"microservices/authentication/models"
	"microservices/db"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestUsersRepositorySave(t *testing.T) {
	// 前提処理
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUserRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())

	assert.NoError(t, err)
	assert.NotNil(t, found)

}

func TestUsersRepositoryGetById(t *testing.T) {
	// 前提処理
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUserRepository(conn)
	err = r.Save(user)

	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())

	assert.NoError(t, err)

	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

	found, err = r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, mgo.ErrNotFound, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)

}

func TestUsersRepositoryGetByEmail(t *testing.T) {
	// 前提処理
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUserRepository(conn)
	err = r.Save(user)

	assert.NoError(t, err)

	found, err := r.GetByEmail(user.Email)

	assert.NoError(t, err)

	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)

	found, err = r.GetByEmail("")
	assert.Error(t, mgo.ErrNotFound, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)

}

func TestUsersRepositoryUpdate(t *testing.T) {
	// 前提処理
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUserRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	user.Name = "UPDATE"
	err = r.Update(user)
	assert.NoError(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, "UPDATE", found.Name)
}

func TestUsersRepositoryDelete(t *testing.T) {
	// 前提処理
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()
	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUserRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)


	err = r.Delete(user.Id.Hex())
	assert.NoError(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}