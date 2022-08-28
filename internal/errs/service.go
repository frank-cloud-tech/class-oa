package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrParseJson        = errors.New("jsonParseError")
	ErrParseNginxConfig = errors.New("nginxConfigParseError")
	ErrParseRegex       = errors.New("regexParseError")

	ErrServiceMeshBindingNotFound = fmt.Errorf("no binding info found")
)
