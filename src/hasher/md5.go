package hasher

import (
	"crypto/md5"
	"fmt"
	"io"
)

type Md5 struct{}

func (m Md5) Encode(a string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, a)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
