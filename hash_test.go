package nobreak

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	table := []struct {
		Purpose, Input, Exp string
	}{
		{"1. OK: empty", ``, `d41d8cd98f00b204e9800998ecf8427e`},
		{"2. OK: hello", `hello`, `5d41402abc4b2a76b9719d911017c592`},
	}
	for _, x := range table {
		actual, err := Checksum(x.Input)
		assert.Nil(t, err)
		assert.EqualValues(t, x.Exp, actual, x.Purpose)
	}
}

func TestHashRequest(t *testing.T) {
	table := []struct {
		Purpose, Method, URL, Body, ExpKey string
	}{
		{"1. OK: new body", "POST", "http://localhost:8080/hello", `{"msg":"my body"}`,
			`3954e94e0e0aec969024b0d1e073fba3`,
		},
		{"2. OK: same body", "POST", "http://localhost:8080/hello/2", `{"msg":"my body"}`,
			`dcbfc170b18d0b8db5ca58537fad87a8`,
		},
		{"3. OK: same body GET", "GET", "http://localhost:8080/hello/2", `{"msg":"my body"}`,
			`e0e39956f19c55a3cf89d5d598bbf7f5`,
		},
		{"4. OK: new body", "GET", "http://localhost:8080/hello", `{"msg":"my body1"}`,
			`0cdaea41c9cfa3d10a01c595f0015cdd`,
		},
	}
	for _, x := range table {
		//		req, err := http.NewRequest(x.Method, x.URL, bytes.NewBufferString(x.Body))
		u, err := url.Parse(x.URL)
		assert.Nil(t, err, x.Purpose)

		body := bytes.NewBufferString(x.Body)
		req := &http.Request{
			RequestURI: u.RequestURI(),
			Method:     x.Method,
			URL:        u,
			Body:       ioutil.NopCloser(body),
		}

		buf, key, err := HashRequest(x.URL, req)
		assert.Nil(t, err, x.Purpose)
		assert.EqualValues(t, x.ExpKey, key, x.Purpose)
		assert.EqualValues(t, x.Body, buf.String(), x.Purpose)
	}
}
