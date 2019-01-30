package securebytes

import (
	"encoding/gob"
	"encoding/json"
	"io"
)

// Serializer is an interface which allows to choose between GOBSerializer and JSONSerializer
type Serializer interface {
	Encode(io.Writer, interface{}) error
	Decode(io.Reader, interface{}) error
}

// GOBSerializer uses encoding/gob to encode and decode data
type GOBSerializer struct{}

// Encode data with gob serializer
func (g GOBSerializer) Encode(w io.Writer, e interface{}) error {
	return gob.NewEncoder(w).Encode(e)
}

// Decode data with gob serializer
func (g GOBSerializer) Decode(r io.Reader, e interface{}) error {
	return gob.NewDecoder(r).Decode(e)
}

// JSONSerializer uses encoding/json to encode and decode data
type JSONSerializer struct{}

// Encode data with json serializer
func (j JSONSerializer) Encode(w io.Writer, e interface{}) error {
	return json.NewEncoder(w).Encode(e)
}

// Decode data with json serializer
func (j JSONSerializer) Decode(r io.Reader, e interface{}) error {
	return json.NewDecoder(r).Decode(e)
}
