package validator_test

import (
	"fmt"
	"testing"

	"github.com/xiaomecc/utils/validator"
)

func TestRunning(t *testing.T) {
	valida := validator.NewValidate()

	type SystemDepartmentCreateReq struct {
		Name    string `json:"name" validate:"required,min=2,max=6"`
		Type    int64  `json:"type" validate:"eq=1|eq=2"`
		OwnerId int64  `json:"owner_id validate:"required"`
		Remark  string `json:"remark,min=1,len=255"`
	}

	var AddGroup = SystemDepartmentCreateReq{
		Name:    "ffffffa",
		Type:    1,
		OwnerId: 99,
		Remark:  "qaqaqa",
	}

	var AddGroupErrInfo = map[string]string{
		"Name.min":          "名字错误3",
		"Name.max":          "名字错误5",
		"Type.eq=1|eq=2":    "类型错误",
		"ParentId.required": "父级ID错误",
		"RoleId.required":   "角色错误",
		"OwnerId.required":  "角色错误",
		"Remark.min":        "备注错误",
		"Remark.len":        "备注错误",
	}

	err := valida.Validate(AddGroupErrInfo, AddGroup)
	fmt.Println(err)
}
