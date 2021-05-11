package repository

import (
	"api-grpc/authentication/models"
	"api-grpc/db"
	"fmt"
	"log"
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

func getUser(id bson.ObjectId) *models.User {
	user := models.User{
		Id:       id,
		Name:     "TEST",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	return &user
}

func TestUsersRepositorySaveSuccess(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	user := getUser(bson.NewObjectId())
	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)
}

func TestUsersRepositoryGetIdSuccess(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	user := getUser(bson.NewObjectId())

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Id.Hex(), found.Id.Hex())
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Password, found.Password)
	assert.Equal(t, user.Email, found.Email)
}

func TestUsersRepositoryGetEmail(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	user := getUser(bson.NewObjectId())

	repo := NewUsersRepository(conn)
	err = repo.Save(user)
	assert.NoError(t, err)

	resp, err := repo.GetByEmail(user.Email)
	assert.NoError(t, err)

	assert.Equal(t, resp.Email, user.Email)
	assert.Equal(t, resp.Name, user.Name)
	assert.Equal(t, resp.Id, user.Id)
	assert.Equal(t, resp.Password, user.Password)
}

func TestUsersRepositoryGetIdNotFound(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	r := NewUsersRepository(conn)
	found, err := r.GetById(bson.NewObjectId().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
	assert.Nil(t, found)
}

func TestUsersRepositoryDeleteId(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)

	r := NewUsersRepository(conn)
	user := getUser(bson.NewObjectId())

	err = r.Save(user)
	assert.NoError(t, err)

	err = r.Delete(user.Id.Hex())
	assert.NoError(t, err)

	_, err = r.GetById(user.Id.Hex())
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
}
