package models

// ResponseModel --
type ResponseModel struct {
	Error   string
	Message string
	Status  int
	Object  interface{}
}
