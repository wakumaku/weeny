package hasher

import "testing"

func TestMd5Hash(t *testing.T) {

	h := Md5{}

	result, err := h.Encode("hello")
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
	expected := "5d41402abc4b2a76b9719d911017c592"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
