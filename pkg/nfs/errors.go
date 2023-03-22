package nfs

import "errors"

var (
	ErrorInvalidShareFormat     error = errors.New("share has invalid format")
	ErrorInvalidDirectoryFormat error = errors.New("directoy has invalid format")
)
