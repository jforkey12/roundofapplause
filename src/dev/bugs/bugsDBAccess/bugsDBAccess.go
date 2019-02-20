package bugsDBAccess

import (
	m "dev/bugs/models"
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
	Collection *mgo.Collection
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
	indexes := []string{"createdBy", "device"}

	mgo.Session, mgo.Collection, err = db.Init("bugs", uniqueIndexes, indexes)
	return err

}

func (mgo *DbService) InsertBug(bug m.Bug) (m.Bug, error) {

	if mgo.Collection == nil {
		return m.Bug{}, errors.New("DB not initialized")
	}

	err := mgo.Collection.Insert(bug)
	if err != nil {
		return bug, err
	}
	return bug, nil
}

func (mgo *DbService) ReplaceBug(bug m.Bug) (m.Bug, error) {

	if mgo.Collection == nil {
		return m.Bug{}, errors.New("DB not initialized")
	}

	err := mgo.Collection.Update(bson.M{"id": bug.ID}, bug)
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

	if mgo.Collection == nil {
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
		err := mgo.Collection.Find(nil).All(&bugs)
		return bugs, err
	}
	err := mgo.Collection.Find(conditions).All(&bugs)
	return bugs, err
}

func (mgo DbService) GetBugByID(id string) (m.Bug, error) {
	bug := m.Bug{}

	if mgo.Collection == nil {
		return bug, errors.New("DB not initialized")
	}

	err := mgo.Collection.Find(bson.M{"id": id}).One(&bug)
	if err != nil {
		return bug, err
	}
	return bug, nil
}

func (mgo DbService) DeleteBug(id string) error {
	if mgo.Collection == nil {
		return errors.New("DB not initialized")
	}
	err := mgo.Collection.Remove(bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
