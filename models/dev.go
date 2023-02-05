package models

type TestCases struct {
	Title    string
	Args     []interface{}
	Expected []interface{}
	Message  string
}
