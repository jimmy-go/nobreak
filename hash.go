package nobreak

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
)

// HashRequest returns url + method and body content hashed.
func HashRequest(scheme string, r *http.Request) (*bytes.Buffer, string, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, "", err
	}
	s := fmt.Sprintf("%s-%s://%s/%s[body:%s]", r.Method, scheme, r.Host, r.URL.RawPath, buf.String())
	log.Printf("HashRequest : s [%s] [%#v]", s, *r)
	key, err := Checksum(s)
	if err != nil {
		return nil, "", err
	}
	return buf, key, nil
}

var (
	mdfive = md5.New()
)

// Checksum generates a md5 checksum from s.
func Checksum(s string) (string, error) {
	if _, err := io.WriteString(mdfive, s); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", mdfive.Sum(nil)), nil
}
