package nobreak

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// Handler _
func Handler(c *Config) http.HandlerFunc {
	client := &http.Client{
		Timeout: time.Duration(c.Timeout) * time.Millisecond,
	}
	log.Printf("db [%s]", c.Database)
	db, err := sqlx.Connect("sqlite3", c.Database)
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec(dbSchema); err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(c.Host)
		if err != nil {
			log.Printf("Parse : err [%s]", err)
			return
		}

		targetURL := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, r.RequestURI)

		buf, key, err := HashRequest(u.Scheme, r)
		if err != nil {
			log.Printf("NewRequest : err [%s]", err)
			return
		}
		log.Printf("Handler : target url [%s] key [%s]", targetURL, key)

		req, err := http.NewRequest(r.Method, targetURL, buf)
		if err != nil {
			log.Printf("NewRequest : err [%s]", err)
			return
		}
		if err := copyHeader(req.Header, r.Header); err != nil {
			log.Printf("copyHeader : err [%s]", err)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			if err := useCache(db, key, w, req); err != nil {
				log.Printf("useCache : err [%s] key [%s]", err, key)
			}
			return
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("Bad response code [%d] key [%s]", resp.StatusCode, key)
			if err := useCache(db, key, w, req); err != nil {
				log.Printf("useCache : err [%s] key [%s]", err, key)
			}
			return
		}

		b, err := Save(db, key, resp)
		if err != nil {
			log.Printf("Save : err [%s]", err)
			return
		}
		if _, err := w.Write(b); err != nil {
			log.Printf("write : err [%s]", err)
		}
	}
}

func copyHeader(dst, src http.Header) error {
	for k, v := range src {
		if strings.ToLower(k) == strings.ToLower("Accept-Encoding") {
			// TODO; add gzip support.
			continue
		}
		dst.Set(k, v[0])
	}
	return nil
}
