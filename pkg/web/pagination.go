package web

type PagedResult struct {
	Items    interface{} `json:"items"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

//NewPagedResult 校验结果是否为空，并创建分页类型的返回格式
func NewPagedResult(items interface{}, total, page, pageSize int) PagedResult {
	return PagedResult{
		Items:    EnsureResultNotNull(items),
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}
