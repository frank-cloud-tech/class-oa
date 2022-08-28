package errs

import "github.com/frank-cloud-tech/class-oa/pkg/web"

func newCustomWebError(httpCode, code int, message, userMessage, level string) web.WebError {
	return web.WebError{
		Response: web.Response{
			Code:        code,
			Message:     message,
			UserMessage: userMessage,
			Level:       level,
		},
		HttpCode: httpCode,
	}
}
