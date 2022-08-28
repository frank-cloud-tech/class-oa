package errs

import (
	"net/http"
)

var (
	// 4XX 默认都是warning
	ErrBadRequest   = newCustomWebError(http.StatusBadRequest, 4000, "bad request", "参数错误", "warning")
	ErrUnauthorized = newCustomWebError(http.StatusUnauthorized, 4001, "Unauthorized", "用户未认证", "warning")
	ErrForbidden    = newCustomWebError(http.StatusForbidden, 4003, "forbidden", "无权限", "warning")
	ErrNotFound     = newCustomWebError(http.StatusNotFound, 4004, "not found", "请求的资源不存在", "warning")
	ErrConflict     = newCustomWebError(http.StatusConflict, 4009, "conflict", "已存在", "warning")
	// 5XX 默认都是error
	ErrInternalServerError = newCustomWebError(http.StatusInternalServerError, 5000, "Internal Server Error", "内部错误", "error")

	ErrGetPodLog        = newCustomWebError(http.StatusNotFound, 4010, "error in opening stream", "无法获取容器标准输出，请检查容器是否启动", "warning")
	ErrPipelineResult   = newCustomWebError(http.StatusInternalServerError, 5000, "get pipeline status failed", "流水线失败", "error")
	ErrGetContainerType = newCustomWebError(http.StatusNotFound, 4011, "can not get container type", "容器规格不存在", "warning")

	ErrAccessDenied = newCustomWebError(http.StatusForbidden, 4013, "access denied", "无权限", "warning")

	ErrMigrateCheckFailed = newCustomWebError(http.StatusBadRequest, 4014, "bad request", "参数错误", "warning")
)
