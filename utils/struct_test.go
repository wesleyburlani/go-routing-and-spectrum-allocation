package utils

import "testing"

func TestCopy(t *testing.T) {
	type TestStruct struct {
		Id string
	}

	object := &TestStruct{Id: "A"}
	object2 := &TestStruct{}
	Copy(object, object2)

	if object.Id != object2.Id {
		t.Errorf("Expected %s, got %s", object.Id, object2.Id)
	}
}
