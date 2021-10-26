package erratum

import "errors"

func Use(ro ResourceOpener, input string) (e error) {
	var resource Resource
	var err error
	defer func() {
		rec := recover()
		if rec != nil {

			if errorType, isFrobError := rec.(FrobError); isFrobError {
				resource.Defrob(errorType.defrobTag)
			}
			resource.Close()
			e = errors.New("meh")
		}
	}()

	for {
		resource, err = ro()

		if err == nil {
			break
		}
		if _, isTransientError := err.(TransientError); !isTransientError {
			return errors.New("too awesome")
		}
	}
	resource.Frob(input)
	resource.Close()
	return nil
}
