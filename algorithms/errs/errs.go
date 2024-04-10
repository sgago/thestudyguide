package errs

import "errors"

var Empty error = errors.New("the collection is empty")
var NotFound error = errors.New("not found")
