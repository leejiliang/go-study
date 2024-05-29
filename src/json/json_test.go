package json

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func TestJsonSerialize(t *testing.T) {
	var u = User{
		Name:   "张三",
		Age:    18,
		Gender: "fame",
	}
	uNew := new(User)
	if marshal, err := json.Marshal(u); err == nil {
		t.Log(string(marshal))
		if err := json.Unmarshal(marshal, uNew); err == nil {
			t.Logf("new User is : %v", uNew)
		}
	}

}
