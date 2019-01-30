package securebytes

import (
	"bytes"
	"testing"

	"github.com/gorilla/securecookie"
)

var secret = "my secret string which is 44 charasters long"

func encryptDecrypt(t *testing.T, sb *SecureBytes) {
	b64, err := sb.EncryptToBase64(secret)
	if err != nil {
		t.Error(err)
	}
	var result string
	err = sb.DecryptBase64(b64, &result)
	if err != nil {
		t.Error(err)
	}
	if result != secret {
		t.Log(b64)
		t.Log(result)
		t.Error("source and decoded data don't match")
	}
}

func TestEncryptDecryptJSON(t *testing.T) {
	sb := New(nil)
	encryptDecrypt(t, sb)
}

func TestEncryptDecryptGOB(t *testing.T) {
	sb := New(nil)
	sb.Serializer = GOBSerializer{}
	encryptDecrypt(t, sb)
}

func BenchmarkSecureBytesJSON(b *testing.B) {
	var b64 string
	sb := New(nil)
	for i := 0; i < b.N; i++ {
		b64, _ = sb.EncryptToBase64(secret)
		sb.DecryptBase64(b64, nil)
	}
}

func BenchmarkSecureBytesGOB(b *testing.B) {
	var b64 string
	sb := New(nil)
	sb.Serializer = GOBSerializer{}
	for i := 0; i < b.N; i++ {
		b64, _ = sb.EncryptToBase64(secret)
		sb.DecryptBase64(b64, nil)
	}
}

func BenchmarkSecureCookieJSON(b *testing.B) {
	var b64 string
	hashKey := bytes.Repeat([]byte("H"), 32)
	blockKey := bytes.Repeat([]byte("B"), 24)
	var sc = securecookie.New(hashKey, blockKey)
	sc.SetSerializer(securecookie.JSONEncoder{})
	for i := 0; i < b.N; i++ {
		b64, _ = sc.Encode("", secret)
		sc.Decode("", b64, nil)
	}
}
func BenchmarkSecureCookieGOB(b *testing.B) {
	var b64 string
	hashKey := bytes.Repeat([]byte("H"), 32)
	blockKey := bytes.Repeat([]byte("B"), 24)
	var sc = securecookie.New(hashKey, blockKey)
	for i := 0; i < b.N; i++ {
		b64, _ = sc.Encode("", secret)
		sc.Decode("", b64, nil)
	}
}
