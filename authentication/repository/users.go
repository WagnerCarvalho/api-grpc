package repository

import (
	"api-grpc/authentication/models"
	"api-grpc/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UserRepository interface {
	Save(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Update(user *models.User) error
	Delete(id string) error
}

type userRepository struct {
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) UserRepository {
	return &userRepository{c: conn.DB().C(UsersCollection)}
}

func (r *userRepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

func (r *userRepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

func (r *userRepository) GetByEmail(email string) (user *models.User, err error) {
	r.c.Find(bson.M{"email": email}).One(&user)
	return
}

func (r *userRepository) GetAll() (users []*models.User, err error) {
	r.c.Find(bson.M{}).One(&users)
	return
}

func (r *userRepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

func (r *userRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

func (r *userRepository) DeleteAll() error {
	return r.c.DropCollection()
}
