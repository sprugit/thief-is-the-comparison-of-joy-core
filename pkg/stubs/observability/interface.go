package observability

import "errors"

type SubmitLog func(string) error

type Observer struct {
	Message    string
	LoggerFunc SubmitLog
}

func (observer *Observer) Submit() (err error) {
	if observer == nil {
		return errors.New("Observer is nil")
	}
	err = observer.LoggerFunc(observer.Message)
	return err
}
