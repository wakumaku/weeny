package application

import (
	"errors"
	"testing"
)

// cacheTester implements cache.Cache
type cacheTester struct {
	funcSave     func(key, value string) error
	funcRetrieve func(key string) (string, error)
}

func (c *cacheTester) Save(key, value string) error {
	if c.funcSave != nil {
		return c.funcSave(key, value)
	}
	return nil
}
func (c *cacheTester) Retrieve(key string) (string, error) {
	if c.funcRetrieve != nil {
		return c.funcRetrieve(key)
	}
	return "", nil
}

// hasherTester implements hasher.Hash
type hasherTester struct {
	funcEncode func(string) (string, error)
}

func (h *hasherTester) Encode(v string) (string, error) {
	if h.funcEncode != nil {
		return h.funcEncode(v)
	}
	return "", nil
}

func TestSaveOk(t *testing.T) {

	c := &cacheTester{}

	expectedHash := "1234567890ASDFGHJKL"
	h := &hasherTester{
		funcEncode: func(v string) (string, error) {
			return expectedHash, nil
		},
	}

	app := New(c, h)

	hash, err := app.Save("hello")
	if err != nil {
		t.Fatalf("Unexpected error, got: %s", err)
	}

	if hash != expectedHash {
		t.Fatalf("Unexpected hash, got: %s, expected: %s", hash, expectedHash)
	}
}

func TestRetrieveOk(t *testing.T) {

	expectedCachedValue := "hello"
	c := &cacheTester{
		funcRetrieve: func(k string) (string, error) {
			return expectedCachedValue, nil
		},
	}

	h := &hasherTester{}

	app := New(c, h)

	cachedValue, err := app.Get("1234567890ASDFGHJKL")
	if err != nil {
		t.Fatalf("Unexpected error, got: %s", err)
	}

	if cachedValue != expectedCachedValue {
		t.Fatalf("Unexpected hash, got: %s, expected: %s", cachedValue, expectedCachedValue)
	}
}

func TestSaveError(t *testing.T) {

	expectedCachedError := errors.New("error caching the value")
	c := &cacheTester{
		funcSave: func(k, v string) error {
			return expectedCachedError
		},
	}

	h := &hasherTester{}

	app := New(c, h)

	cachedValue, err := app.Save("hello")
	if err != expectedCachedError {
		t.Fatalf("Unexpected error, got: %s", err)
	}

	if cachedValue != "" {
		t.Fatalf("Unexpected hash, got: %s, expected: <empty string>", cachedValue)
	}
}
func TestSaveEncoderError(t *testing.T) {

	c := &cacheTester{}

	expectedEncodeError := errors.New("error encoding the value")
	h := &hasherTester{
		funcEncode: func(v string) (string, error) {
			return "", expectedEncodeError
		},
	}

	app := New(c, h)

	cachedValue, err := app.Save("hello")
	if err != expectedEncodeError {
		t.Fatalf("Unexpected error, got: %s", err)
	}

	if cachedValue != "" {
		t.Fatalf("Unexpected hash, got: %s, expected: <empty string>", cachedValue)
	}
}
