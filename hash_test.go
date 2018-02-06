package nobreak

import (
	"bytes"
	"net/http"
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
			`843e076e3b56edd4e730c1ae59f570f1`,
		},
		{"2. OK: same body", "POST", "http://localhost:8080/hello/2", `{"msg":"my body"}`,
			`decdd8a95e061316a08fdbd92c2ac4d5`,
		},
		{"3. OK: same body GET", "GET", "http://localhost:8080/hello/2", `{"msg":"my body"}`,
			`58cf18899120484eb5b8880d38408d3e`,
		},
	}
	for _, x := range table {
		req, err := http.NewRequest(x.Method, x.URL, bytes.NewBufferString(x.Body))
		assert.Nil(t, err, x.Purpose)

		buf, key, err := HashRequest("https", req)
		assert.Nil(t, err, x.Purpose)
		assert.EqualValues(t, x.ExpKey, key, x.Purpose)
		assert.EqualValues(t, x.Body, buf.String(), x.Purpose)
	}
}
