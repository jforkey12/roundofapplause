package usersDBAccess

import (
	m "dev/users/models"
	db "dev/utils/dbadapter"
	"errors"
	"os"
	"sync"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var once sync.Once
var instance m.DbInterface

type DbService struct {
	Session    *mgo.Session
	userCollection *mgo.userCollection
}

func GetMgoService() m.DbInterface {
	once.Do(func() {
		instance = &DbService{}

		err := instance.InitSession()
		if err != nil {
			os.Exit(1)
		}
	})

	return instance
}

func (mgo *DbService) InitSession() (err error) {
	uniqueIndexes := []string{"id"}
	indexes := []string{"firstName", "lastName", "country"}
	mgo.Session, mgo.userCollection, err = db.Init("users", uniqueIndexes, indexes)
	return err

}

func (mgo *DbService) InsertUser(user m.User) (m.User, error) {

	if mgo.userCollection == nil {
		return m.User{}, errors.New("DB not initialized")
	}

	err := mgo.userCollection.Insert(user)
	return user, err
}

func (mgo *DbService) ReplaceUser(user m.User) (m.User, error) {

	if mgo.userCollection == nil {
		return m.User{}, errors.New("DB not initialized")
	}

	err := mgo.userCollection.Update(bson.M{"id": user.ID}, user)
	return user, err
}

func (mgo DbService) GetUsers(countries []string, devices []string) ([]m.User, error) {
	users := []m.User{}
	var devParams bson.M
	var cParams bson.M
	conditions := bson.M{}

	if mgo.userCollection == nil {
		return users, errors.New("DB not initialized")
	}
	if len(devices) > 0 {
		devParams = bson.M{"devices": bson.M{"$in": devices}}
	}
	if len(countries) > 0 {
		cParams = bson.M{"country": bson.M{"$in": countries}}
	}
	if len(devices) > 0 && len(countries) > 0 {
		conditions =
			bson.M{"$and": []bson.M{
				devParams,
				cParams,
			},
			}
	} else if len(devices) > 0 {
		conditions = devParams
	} else if len(countries) > 0 {
		conditions = cParams
	} else {
		err := mgo.userCollection.Find(nil).All(&users)
		return users, err
	}
	err := mgo.userCollection.Find(conditions).All(&users)

	return users, err
}

func (mgo DbService) GetUserByID(id string) (m.User, error) {
	user := m.User{}

	if mgo.userCollection == nil {
		return user, errors.New("DB not initialized")
	}

	err := mgo.userCollection.Find(bson.M{"id": id}).One(&user)
	return user, err
}

func (mgo DbService) DeleteUser(id string) error {
	if mgo.userCollection == nil {
		return errors.New("DB not initialized")
	}
	err := mgo.userCollection.Remove(bson.M{"id": id})

	return err
}
