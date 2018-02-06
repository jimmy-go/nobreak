package nobreak

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// HashRequest returns url + method and body content hashed.
func HashRequest(host string, r *http.Request) (*bytes.Buffer, string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, "", err
	}
	targetURL := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, r.RequestURI)

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, "", err
	}
	s := fmt.Sprintf("%s-%s[body:%s]", r.Method, targetURL, buf.String())
	return buf, s, nil
	// FIXME; checksum fails for root.

	//	key, err := Checksum(s)
	//	log.Printf("key [%s] pre key [%s]", key, s)
	//	if err != nil {
	//		return nil, "", err
	//	}
	//	return buf, key, nil
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
