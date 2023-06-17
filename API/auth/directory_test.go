package auth

import "testing"

func TestCreateDirectory(t *testing.T) {

	err := CreateDirectory()
	if err == nil {
		t.Errorf("Test Failed")
	}
}

func TestConfirmDirectoryExistence(t *testing.T) {
	b, _ := ConfirmDirectoryExistence()
	if b != false {
		t.Errorf("Test Failed")
	}
}
