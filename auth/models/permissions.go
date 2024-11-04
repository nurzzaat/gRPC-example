package models

import "github.com/nurzzaat/gRPC-example/common"

type Role struct {
	Name        int
	Permissions []string
}

const (
	AllowFAQRead   = "faq.read"
	AllowFAQManage = "faq.update"

	AllowSupportRead   = "support.read"
	AllowSupportManage = "support.update"
)

var Roles = []Role{
	{Name: common.ADMIN, Permissions: []string{
		AllowFAQManage, AllowFAQRead, AllowSupportManage, AllowSupportRead,
	}},

	{Name: common.USER, Permissions: []string{
		AllowFAQRead, AllowSupportRead,
	}},
}
