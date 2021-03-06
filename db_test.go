package boltdb

import (
	"io/ioutil"

	"github.com/boltdb/bolt"
	"github.com/facebookgo/ensure"
	// "log"
	"path"
	"testing"
)

func TestString(t *testing.T) {
	db := newBoltDB(t)
	defer db.Close()

	bucket, err := db.Bucket([]byte("0"))
	ensure.Nil(t, err)

	err = bucket.Set([]byte("version"), []byte("1"))
	ensure.Nil(t, err)

	val, err := bucket.Get([]byte("version"))
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, []byte("1"))

	elemType, err := bucket.TypeOf([]byte("version"))
	ensure.Nil(t, err)
	ensure.DeepEqual(t, elemType, STRING)
}

func newBoltDB(t *testing.T) *DB {
	dir, err := ioutil.TempDir("", "bolt")
	ensure.Nil(t, err)

	dbpath := path.Join(dir, "bolt.db")
	// log.Println("dbpath:", dbpath)
	opt := &Options{
		ReadOnly: false,
	}
	db, err := Open(dbpath, 0644, opt)
	ensure.Nil(t, err)

	return db
}

func scan(db *bolt.DB, bucket []byte, t *testing.T) {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucket).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t.Logf("%s  %s\n", k, v)
		}
		return nil
	})
}
