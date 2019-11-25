package errors

//EventError ...
type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	//ErrDateBusy ...
	ErrDateBusy = EventError("This Date is used")
)
