package db

import (
	"testing"
)

func Test__NewM(t *testing.T) {
	in := NewM()

	if in.isNew != true {
		t.Error("should set isNew to true")
	}
}

func Test__M_IsNew_and_Store(t *testing.T) {
	m := NewM()
	if !m.IsNew() {
		t.Error("should return true when model is new")
	}

	if m.Store(); m.IsNew() {
		t.Error("should return false when model is not new")
	}
}

type testModel struct {
	*M
	ID int64
	A  int
}

func newTestModel() *testModel {
	m := new(testModel)
	m.M = NewM()
	return m
}

func (m *testModel) PK() (string, interface{}) {
	return "id", m.ID
}

func (m *testModel) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"id": m.ID,
		"a":  m.A,
	}
}

func Test__Div(t *testing.T) {
	ms := make([]Model, 0, 65535)
	cp := cap(ms)
	newMsExp, exMsExp := make([]Model, 0, 1+cp%2), make([]Model, 0, cp%2)

	for id := 1; id <= cp; id++ {
		m := &testModel{NewM(), int64(id), id}
		ms = append(ms, m)
		if id%2 == 0 {
			ms[len(ms)-1].Store()
			exMsExp = append(exMsExp, m)
		} else {
			newMsExp = append(newMsExp, m)
		}
	}

	newMs, exMs := NewAndExisting(ms)

	for _, ms := range [][][]Model{{newMs, newMsExp}, {exMs, exMsExp}} {
		act, exp := ms[0], ms[1]
		if len(act) != len(exp) {
			t.Errorf("expected %d models, but got %d", len(exp), len(act))
		}

		for i := range exp {
			if act[i] != exp[i] {
				t.Errorf("expected model %p, but got %p", exp[i], act[i])
			}
		}
	}
}

func Test__PKs(t *testing.T) {
	ln := 65535
	ms, pksExp := make([]Model, 0, ln), make([]interface{}, 0, ln)

	for id := 1; id <= ln; id++ {
		m := &testModel{NewM(), int64(id), id}
		ms = append(ms, m)
		_, pkv := m.PK()
		pksExp = append(pksExp, pkv)
		m.Store()
	}

	pks := PKs(ms)

	for i := range pksExp {
		if pks[i].(int64) != pksExp[i].(int64) {
			t.Errorf("expected pk %d, but got %d", pksExp[i].(int64),
				pks[i].(int64))
		}
	}
}
