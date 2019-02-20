package dbadapter

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/globalsign/mgo"
)

type DBAdapter struct {
	Session *mgo.Session
}

var dba *DBAdapter
var once sync.Once

func getDBAdapter() *DBAdapter {
	once.Do(func() {

		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{"localhost:27017"},
			Timeout:  10 * time.Second,
			Username: "admin1",
			Password: "qwerty123456!",
			Database: "applause",
		}

		session, _ := mgo.DialWithInfo(mongoDBDialInfo)
		dba = &DBAdapter{Session: session}
	})
	return dba
}

func Init(collection string, uniqueInds []string, inds []string) (*mgo.Session, *mgo.Collection, error) {
	dba = getDBAdapter()
	var err error

	if dba.Session == nil {
		err = errors.New("Unable to connect to the db, Ensure db is running")
		return nil, nil, err
	}

	session := dba.Session.Copy()

	session.SetMode(mgo.Monotonic, true)

	col := session.DB("applause").C(collection)
	for _, index := range uniqueInds {
		var s []string
		s = append(s, index)
		index := mgo.Index{
			Key:    s,
			Unique: true,
		}
		err := col.EnsureIndex(index)

		if err != nil {
			os.Exit(1)
		}
	}

	for _, index := range inds {
		var s []string
		s = append(s, index)
		index := mgo.Index{
			Key:    s,
			Unique: false,
		}
		err := col.EnsureIndex(index)

		if err != nil {
			os.Exit(1)
		}
	}

	return session, col, err
}
