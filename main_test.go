package securebytes

import (
	"reflect"
	"testing"
)

type testStruct struct {
	UserID  int
	Message string
}

var secret = testStruct{123123123, "secret"}

func encryptDecrypt(t *testing.T, sb *SecureBytes) {
	b64, err := sb.EncryptToBase64(secret)
	if err != nil {
		t.Error(err)
	}
	var result testStruct
	err = sb.DecryptBase64(b64, &result)
	if err != nil {
		t.Error(err)
	}
	t.Log(b64)
	if !reflect.DeepEqual(result, secret) {
		t.Log(result)
		t.Error("source and decoded data don't match")
	}
}

func TestEncryptDecryptJSON(t *testing.T) {
	sb := New(nil, JSONSerializer{})
	encryptDecrypt(t, sb)
}

func TestEncryptDecryptGOB(t *testing.T) {
	sb := New(nil, GOBSerializer{})
	encryptDecrypt(t, sb)
}

func TestEncryptDecryptASN1(t *testing.T) {
	sb := New(nil, ASN1Serializer{})
	encryptDecrypt(t, sb)
}

func BenchmarkSecureBytesJSON(b *testing.B) {
	var b64 string
	var result testStruct
	sb := New(nil, JSONSerializer{})
	for i := 0; i < b.N; i++ {
		b64, _ = sb.EncryptToBase64(secret)
		sb.DecryptBase64(b64, &result)
	}
}

func BenchmarkSecureBytesGOB(b *testing.B) {
	var b64 string
	var result testStruct
	sb := New(nil, GOBSerializer{})
	for i := 0; i < b.N; i++ {
		b64, _ = sb.EncryptToBase64(secret)
		sb.DecryptBase64(b64, &result)
	}
}

func BenchmarkSecureBytesASN1(b *testing.B) {
	var b64 string
	var result testStruct
	sb := New(nil, ASN1Serializer{})
	for i := 0; i < b.N; i++ {
		b64, _ = sb.EncryptToBase64(secret)
		sb.DecryptBase64(b64, &result)
	}
}
