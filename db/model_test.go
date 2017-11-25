package db

import (
	"testing"
)

func Test__NewIN(t *testing.T) {
	in := NewIN()

	if in.isNew != true {
		t.Error("should set isNew to true")
	}
}

func Test__IN_IsNew(t *testing.T) {
	in := NewIN()

	if !in.IsNew() {
		t.Error("should return true when model is new")
	}

	in.isNew = false
	if in.IsNew() {
		t.Error("should return false when model is not new")
	}
}

type testModel struct {
	*IN
	ID int64
}

func newTestModel() Model {
	m := new(testModel)
	m.IN = NewIN()
	return m
}

func (m *testModel) PK() (string, interface{}) {
	return "id", m.ID
}

func Test__Model(t *testing.T) {
	var mi = newTestModel()
	m := mi.(*testModel)

	m.ID = 5
	if id, v := mi.PK(); id != "id" {
		t.Error("should return primary key name")
	} else if v.(int64) != 5 {
		t.Error("should return primary key value")
	}

	if !mi.IsNew() {
		t.Error("should return true when model is new")
	}

	m.isNew = false
	if mi.IsNew() {
		t.Error("should return false when model is not new")
	}
}
