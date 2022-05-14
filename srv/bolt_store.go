package srv

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	bolt "go.etcd.io/bbolt"
)

var BucketName = []byte("updatemgr")

type BoltStore struct {
	DB *bolt.DB
}

func InitBoltStore(path string) *BoltStore {
	db, err := bolt.Open(path, 0644, nil)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to open bolt database file")
	}
	bs := BoltStore{db}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(BucketName)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to create bucket in boltdb")
		}
		return nil
	})
	return &bs
}

func (bs *BoltStore) addOnline(o Online) {
	bs.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		v := bucket.Get([]byte(o.Name))
		host := Host{}
		if v == nil {
			host = Host{
				Name:   o.Name,
				Online: o,
			}
		} else {
			host = *UnmarshallHost(v)
			host.Online = o
		}
		err := bucket.Put([]byte(o.Name), host.Marshall())

		if err != nil {
			log.Error().
				Str("host", o.Name).
				Err(err).
				Msg("Could not store online info for host")
			return err
		}
		return nil
	})
}

func (bs *BoltStore) addUpdatesAvailable(u UpdatesAvailable) {
	bs.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		v := bucket.Get([]byte(u.Name))
		host := Host{}
		if v == nil {
			host = Host{
				Name:             u.Name,
				UpdatesAvailable: u,
			}
		} else {
			host = *UnmarshallHost(v)
			host.UpdatesAvailable = u
		}
		err := bucket.Put([]byte(u.Name), host.Marshall())

		if err != nil {
			log.Error().
				Str("host", u.Name).
				Err(err).
				Msg("Could not store updates available info for host")
			return err
		}
		return nil
	})
}

func (bs *BoltStore) addRebootRequired(r RebootRequired) {
	bs.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		v := bucket.Get([]byte(r.Name))
		host := Host{}
		if v == nil {
			host = Host{
				Name:           r.Name,
				RebootRequired: r,
			}
		} else {
			host = *UnmarshallHost(v)
			host.RebootRequired = r
		}
		err := bucket.Put([]byte(r.Name), host.Marshall())

		if err != nil {
			log.Error().
				Str("host", r.Name).
				Err(err).
				Msg("Could not store reboot required info for host")
			return err
		}
		return nil
	})
}

func (bs *BoltStore) getHost(hostname string) (Host, error) {
	host := Host{}
	err := bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		r := bucket.Get([]byte(hostname))
		if r != nil {
			return errors.New("No much hostname")
		}
		host = *UnmarshallHost(r)
		return nil
	})
	return host, err
}

func (bs *BoltStore) getHosts() []Host {
	var hosts []Host = make([]Host, 0)
	bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			current := *UnmarshallHost(v)
			hosts = append(hosts, current)
		}
		return nil
	})
	return hosts
}

func (bs *BoltStore) daHosts() map[string]Host {
	var hosts map[string]Host = make(map[string]Host, 0)
	bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			current := *UnmarshallHost(v)
			hosts[current.Name] = current
		}
		return nil
	})
	return hosts
}

func (bs *BoltStore) getRebootRequired(hostname string) (bool, error) {
	host := Host{}
	err := bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		r := bucket.Get([]byte(hostname))
		if r != nil {
			return errors.New("No much hostname")
		}
		host = *UnmarshallHost(r)
		return nil
	})

	return host.RebootRequired.RebootRequired, err
}

func (bs *BoltStore) getUpdatesAvailable(hostname string) ([]string, error) {
	host := Host{}
	err := bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		r := bucket.Get([]byte(hostname))
		if r != nil {
			return errors.New("No much hostname")
		}
		host = *UnmarshallHost(r)
		return nil
	})

	return host.UpdatesAvailable.Packages, err
}

func (bs *BoltStore) purge(ttl int) {
	oldest := time.Now().Unix() - int64(ttl)

	bs.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			current := *UnmarshallHost(v)
			if current.Online.TimeStamp >= oldest {
				bucket.Delete(k)
			}
		}
		return nil
	})
}
