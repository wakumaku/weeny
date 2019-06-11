package hasher

import "testing"

func TestShortUUID(t *testing.T) {

	h := Hashids{}

	result, err := h.Encode("http://duckduckgo.com")
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
	expected := "bRFOuAsbUl"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
