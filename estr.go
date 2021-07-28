package estr

import (
	"github.com/speps/go-hashids"
)

const (
	defaultSalt      = "estr"
	defaultMinLength = 10
)

var hashID *hashids.HashID

func Init(salt string, minLength uint8) error {
	hd := hashids.NewData()
	if salt == "" {
		salt = defaultSalt
	}
	hd.Salt = salt

	if minLength == 0 {
		minLength = defaultMinLength
	}

	hd.MinLength = int(minLength)
	var err error
	if hashID, err = hashids.NewWithData(hd); err != nil {
		return err
	}

	return nil
}

func Encode(src string) (string, error) {
	var ns []int64
	for _, r := range src {
		ns = append(ns, int64(r))
	}
	return hashID.EncodeInt64(ns)
}

func Decode(src string) (string, error) {
	var ns []rune
	dst, err := hashID.DecodeInt64WithError(src)
	if err != nil {
		return "", err
	}

	for _, i := range dst {
		ns = append(ns, rune(i))
	}

	return string(ns), nil
}
