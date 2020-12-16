package handler

import (
	"testing"

	"github.com/lz1998/ecust_library/model/admin"
)

func TestGenerateJwtTokenString(t *testing.T) {
	token, err := GenerateJwtTokenString(&admin.EcustAdmin{
		ID:       1,
		Username: "123",
		Password: "123",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", token)
	}
}
