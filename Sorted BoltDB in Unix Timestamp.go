package main

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func main() {

// Open a new DB
	db, err := bolt.Open("db", 0600, nil)
	if err != nil {
	    log.Fatal(err)
	}

// Update DB
	db.Update(func(tx *bolt.Tx) error {
		log.Println("Writing 1 key/value pairs to 'Bucket'")
		var curTime int64 = time.Now().Unix()
			if err := put(tx, "Bucket", curTime, "Add any value"); err != nil {
				return err
			}
		return nil
	})
		log.Println("\nCommitting...\n")
		if err != nil {
			log.Fatal()
		}
		log.Println("Committed data:)\n")

// Read DB
 db.View(func(tx *bolt.Tx) error {
		log.Println("Reading data for 'Bucket':")
		log.Println("    Key            Value")
		c := tx.Bucket([]byte("Bucket")).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			t := int64(binary.BigEndian.Uint64(k)) 
			  log.Printf("• %d » %s", t, v)
		}
		log.Println("")
			return nil
	})
}

// Put function 
func put(tx *bolt.Tx, bucket string, curTime int64, value string) error { 
	k := make([]byte, 8) 
	binary.BigEndian.PutUint64(k, uint64(curTime)) 
	b, err := tx.CreateBucketIfNotExists([]byte(bucket)) 
          if err != nil {
            return err
          }	
	if err := b.Put(k, []byte(value)); err != nil {
		return err
	}
	return nil
}

