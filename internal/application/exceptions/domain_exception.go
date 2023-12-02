package exceptions

type DomainException struct {
	Message string
}

func NewDomainException(message string) *DomainException {
	return &DomainException{
		Message: message,
	}
}

func (de *DomainException) Error() string {
	return de.Message
}
