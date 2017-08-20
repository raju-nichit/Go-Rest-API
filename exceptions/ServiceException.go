package exceptions

//ServiceError --
type ServiceError struct {
	Message string
}

func (serviceError *ServiceError) Error() string {
	return serviceError.Message
}
