package securebytes

import (
	"bytes"
	"encoding/asn1"
	"encoding/gob"
	"encoding/json"
)

// Serializer is an interface which allows to choose between GOBSerializer and JSONSerializer
type Serializer interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

// GOBSerializer uses encoding/gob to encode and decode data
type GOBSerializer struct{}

// Marshal data with gob serializer
func (g GOBSerializer) Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(v)
	return buf.Bytes(), err
}

// Unmarshal data with gob serializer
func (g GOBSerializer) Unmarshal(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}

// JSONSerializer uses encoding/json to encode and decode data
type JSONSerializer struct{}

// Marshal data with json serializer
func (j JSONSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal data with json serializer
func (j JSONSerializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// ASN1Serializer uses encoding/asn1 to encode and decode data
type ASN1Serializer struct{}

// Marshal data with asn1 serializer
func (a ASN1Serializer) Marshal(v interface{}) ([]byte, error) {
	return asn1.Marshal(v)
}

// Unmarshal data with asn1 serializer
func (a ASN1Serializer) Unmarshal(data []byte, v interface{}) error {
	_, err := asn1.Unmarshal(data, v)
	return err
}
