package utils

import "testing"

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("zhangsan")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	validationToken, err := ValidationToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(validationToken)
}
