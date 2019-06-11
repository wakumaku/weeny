package hasher

import (
	"github.com/speps/go-hashids"
)

type Hashids struct{}

func (enc Hashids) Encode(a string) (string, error) {
	hd := hashids.NewData()
	hd.Salt = a
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	return h.Encode([]int{1, 2, 3, 4, 5})
}
