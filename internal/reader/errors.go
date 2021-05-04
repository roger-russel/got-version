package reader

import "errors"

var ErrDataIsNotFromAPIPE error = errors.New("this version only accept data from pipe")
