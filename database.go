package nobreak

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

func init() {
	syncc = make(chan struct{}, 1)
	syncc <- struct{}{}
}

var (
	syncc chan struct{}
)

const (
	dbSchema = `
		CREATE TABLE cache_version (
			key TEXT PRIMARY KEY
		);
		CREATE TABLE cache_store (
			key TEXT PRIMARY KEY,
			header BLOB NOT NULL,
			body BLOB NOT NULL,
			created_at TEXT NOT NULL
		);
		INSERT INTO cache_version (key) VALUES ('0.0.0');
	`
)

// Save stores the http response.
func Save(db *sqlx.DB, key string, res *http.Response) ([]byte, error) {
	if v, err := getCache(db, key); err == nil {
		return v.Body, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//	log.Printf("b [%s]", string(b))
	h, err := json.Marshal(res.Header)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	o := <-syncc
	_, err = db.Exec(`
		INSERT INTO cache_store
		(key, header, body, created_at)
		VALUES (?, ?, ?, ?)
	`, key, h, b, now.Format(time.RFC3339))
	syncc <- o
	return b, err
}

// Cache type.
type Cache struct {
	Key       string `db:"key"`
	Header    []byte `db:"header"`
	Body      []byte `db:"body"`
	CreatedAt string `db:"created_at"`
}

func useCache(db *sqlx.DB, key string, w http.ResponseWriter, r *http.Request) error {
	v, err := getCache(db, key)
	if err != nil {
		return err
	}
	if _, err := w.Write(v.Body); err != nil {
		return err
	}
	var h http.Header
	if err := json.Unmarshal(v.Header, &h); err != nil {
		return err
	}
	if err := copyHeader(w.Header(), h); err != nil {
		return err
	}
	w.Header().Set("No-Break", "Use-Cache")
	// 	log.Printf("header : [%v]", w.Header())
	return nil
}

func getCache(db *sqlx.DB, key string) (Cache, error) {
	var v Cache
	o := <-syncc
	err := db.Get(&v, `
		SELECT key, header, body, created_at
		FROM cache_store
		WHERE key = ?`, key)
	syncc <- o
	return v, err
}
