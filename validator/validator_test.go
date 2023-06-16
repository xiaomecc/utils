package validator_test

import (
	"fmt"
	"testing"

	"github.com/xiaomecc/utils/validator"
)

func TestRunning(t *testing.T) {
	valida := validator.NewValidate()
	valida.Binder()

	type SystemDepartmentCreateReq struct {
		Name    string `json:"name" validate:"required,min=2,max=6"`
		Type    int64  `json:"type" validate:"eq=1|eq=2"`
		OwnerId int64  `json:"owner_id" validate:"required"`
		Remark  string `json:"remark,min=1,len=255"`
		Mobile  string `json:"mobile" validate:"mobile880"`
	}

	var AddGroup = SystemDepartmentCreateReq{
		Name:    "ffff",
		Type:    1,
		OwnerId: 99,
		Remark:  "qaqaqa",
		Mobile:  "1812598624",
	}

	var AddGroupErrInfo = map[string]string{
		"Name.min":          "a",
		"Name.max":          "b",
		"Type.eq=1|eq=2":    "c",
		"ParentId.required": "d",
		"RoleId.required":   "e",
		"OwnerId.required":  "f",
		"Remark.min":        "g",
		"Remark.len":        "h",
		"Mobile.mobile880":  "mobile error",
	}

	err := valida.Validate(AddGroupErrInfo, AddGroup)
	fmt.Println(err)
}
