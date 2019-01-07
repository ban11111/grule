package grule

import "reflect"

type ruler map[string]*rule

func (rs *ruler) add(r rule) {
	if len(*rs) <= 0 {
		*rs = make(map[string]*rule)
	}
	if (*rs)[r.Name] != nil {
		panic("rule: " + r.Name + " already exist, please use update instead.")
	}
	r.getCmp()
	(*rs)[r.Name] = &r
}

func (rs *ruler) addSub(name, passOrFail string, r rule) {
	if (*rs)[name] == nil {
		panic("rule: " + r.Name + " doesn't exist, please add it first.")
	}
	if passOrFail == "pass" {
		(*rs)[name].Pass = &r
	} else {
		(*rs)[name].Fail = &r
	}
}

func (rs *ruler) adds(r []rule) {
	if len(*rs) <= 0 {
		*rs = make(map[string]*rule)
	}
	for i := range r {
		if (*rs)[r[i].Name] != nil {
			panic("rule: " + r[i].Name + "already exist, please use update instead.")
		}
		r[i].getCmp()
		(*rs)[r[i].Name] = &r[i]
	}
}

type rule struct {
	Name         string      `json:"name"`
	ValueConfig  interface{} `json:"value_config"`
	Comparator   string      `json:"comparator"`
	DataReceived interface{} `json:"data_received"`
	Pass         *rule       `json:"pass"`
	Fail         *rule       `json:"fail"`
	cmp          int
	result       *bool
}

func (r *rule) compare() (pass *bool) {
	pass, _ = compareByTypes(r.ValueConfig, r.DataReceived, r.cmp)
	if pass != nil && *pass && r.Pass != nil {
		pass = r.Pass.compare()
	} else if pass != nil && !*pass && r.Fail != nil {
		pass = r.Fail.compare()
	}
	r.result = pass
	return pass
}

const (
	empty     = iota
	eq
	neq
	gt
	gte
	lt
	lte
	exists
	nexists
	regex
	matches
	contains
	ncontains
)

func (r *rule) getCmp() {
	var x = map[string]int{"eq": eq}
	r.cmp = x[r.Comparator]
}

func compareByTypes(ruleData, cmpData interface{}, cmp int) (resulted *bool, err error) {
	var result bool
	switch cmp {
	case empty:
		panic("no shit")
	case eq:
		result = ObjectsAreEqualValues(ruleData, cmpData)
		resulted = &result
	case neq:
		result = !ObjectsAreEqualValues(ruleData, cmpData)
		resulted = &result
	case gt:

	}
	return
}

// ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
func ObjectsAreEqualValues(expected, actual interface{}) bool {
	if objectsAreEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}

	return false
}

func objectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	return reflect.DeepEqual(expected, actual)
}