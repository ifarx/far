package util

type Codec interface {
	Unmarshal(bys []byte, in interface{}) error
	Marshal(in interface{}) ([]byte, error)
}
