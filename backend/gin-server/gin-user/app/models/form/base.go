// Package form
// Date        : 2023/3/9 16:25
// Version     : 1.0.0
// Author      : 代码小学生王木木
// Email       : 18574945291@163.com
// Description :
package form

type PaginateForm struct {
	PageSize int64 `json:"pageSize" form:"pageSize" binding:"-"`
	PageNum  int64 `json:"pageNum" form:"pageNum" binding:"-"`
}
