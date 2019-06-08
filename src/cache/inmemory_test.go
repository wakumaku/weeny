package cache

import (
	"testing"
)

func TestStoreAndRetrievesAKey(t *testing.T) {

	c := NewInMemory()

	err := c.Save("id1", "value1")
	if err != nil {
		t.Fatal("error saving the value")
	}

	value, err := c.Retrieve("id1")
	if err != nil {
		t.Fatal("error retrieving the value")
	}

	if value != "value1" {
		t.Fatal("Unexpected value stored")
	}

}

func TestRetrieveFailsWhenAnItemIsNotFound(t *testing.T) {

	c := NewInMemory()

	err := c.Save("id1", "value1")
	if err != nil {
		t.Fatal("error saving the value")
	}

	if _, err := c.Retrieve("id2"); err == nil {
		t.Fatal("an error was expected")
	}

}
