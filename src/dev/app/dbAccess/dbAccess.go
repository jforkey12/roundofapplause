package dbAccess

import (
	m "dev/app/models"
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
	Session        *mgo.Session
	userCollection *mgo.Collection
	bugsCollection *mgo.Collection
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
	buInds := []string{"id"}
	bInds := []string{"createdBy", "device"}

	uuInds := []string{"id"}
	uInds := []string{"firstName", "lastName", "country"}

	mgo.Session, mgo.bugsCollection, err = db.Init("bugs", buInds, bInds)
	mgo.Session, mgo.userCollection, err = db.Init("users", uuInds, uInds)

	return err

}

func (mgo *DbService) InsertBug(bug m.Bug) (m.Bug, error) {

	if mgo.bugsCollection == nil {
		return m.Bug{}, errors.New("DB not initialized")
	}

	err := mgo.bugsCollection.Insert(bug)
	if err != nil {
		return bug, err
	}
	return bug, nil
}

func (mgo *DbService) ReplaceBug(bug m.Bug) (m.Bug, error) {

	if mgo.bugsCollection == nil {
		return m.Bug{}, errors.New("DB not initialized")
	}

	err := mgo.bugsCollection.Update(bson.M{"id": bug.ID}, bug)
	if err != nil {
		return bug, err
	}
	return bug, err
}

func (mgo DbService) GetBugs(testerIds []int, devices []string) ([]m.Bug, error) {
	bugs := []m.Bug{}
	var dParams bson.M
	var tParams bson.M
	conditions := bson.M{}

	if mgo.bugsCollection == nil {
		return bugs, errors.New("DB not initialized")
	}
	if len(devices) > 0 {
		dParams = bson.M{"device": bson.M{"$in": devices}}
	}
	if len(testerIds) > 0 {
		tParams = bson.M{"createdby": bson.M{"$in": testerIds}}
	}
	if len(devices) > 0 && len(testerIds) > 0 {
		conditions =
			bson.M{"$and": []bson.M{
				dParams,
				tParams,
			},
			}
	} else if len(devices) > 0 {
		conditions = dParams
	} else if len(testerIds) > 0 {
		conditions = tParams
	} else {
		err := mgo.bugsCollection.Find(nil).All(&bugs)
		return bugs, err
	}
	err := mgo.bugsCollection.Find(conditions).All(&bugs)
	return bugs, err
}

func (mgo DbService) GetBugByID(id string) (m.Bug, error) {
	bug := m.Bug{}

	if mgo.bugsCollection == nil {
		return bug, errors.New("DB not initialized")
	}

	err := mgo.bugsCollection.Find(bson.M{"id": id}).One(&bug)
	if err != nil {
		return bug, err
	}
	return bug, nil
}

func (mgo DbService) DeleteBug(id string) error {
	if mgo.bugsCollection == nil {
		return errors.New("DB not initialized")
	}
	err := mgo.bugsCollection.Remove(bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
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
