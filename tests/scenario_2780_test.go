// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uaccount"
	"github.com/ucloud/ucloud-sdk-go/services/ubill"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario2780(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "2780",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Phone": 18800230967,
				"Email": "testsdk@ucloud.cn",
			}
		},
		Owners: []string{"sunny.zhang@ucloud.cn"},
		Title:  "UBill+UAccount API测试集",
		Steps: []*driver.Step{
			testStep2780GetProjectList01,
			testStep2780DescribeMemberList02,
			testStep2780CreateProject03,
			testStep2780InviteSubaccount04,
			testStep2780AddMemberToProject05,
			testStep2780FreezeMember06,
			testStep2780RemoveMemberFromProject07,
			testStep2780TerminateMember08,
			testStep2780TerminateCharacter09,
			testStep2780TerminateProject10,
			testStep2780GetBillDataFileUrl11,
		},
	})
}

var testStep2780GetProjectList01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UAccount")
		if err != nil {
			return nil, err
		}
		client := c.(*uaccount.UAccountClient)

		req := client.NewGetProjectListRequest()
		err = utils.SetRequest(req, map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		resp, err := client.GetProjectList(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "GetProjectListResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取帐号下的项目列表",
	FastFail:      false,
}

var testStep2780DescribeMemberList02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeMemberList")
		err = req.SetPayload(map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeMemberListResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取成员列表",
	FastFail:      false,
}

var testStep2780CreateProject03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UAccount")
		if err != nil {
			return nil, err
		}
		client := c.(*uaccount.UAccountClient)

		req := client.NewCreateProjectRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ProjectName": "test-api-project",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateProject(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ProjectId", step.Must(utils.GetValue(resp, "ProjectId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateProjectResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建项目",
	FastFail:      false,
}

var testStep2780InviteSubaccount04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("InviteSubaccount")
		err = req.SetPayload(map[string]interface{}{
			"UserPhone": step.Scenario.GetVar("Phone"),
			"UserName":  "testucloudsdk",
			"UserEmail": step.Scenario.GetVar("Email"),
			"IsFinance": "no",
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "InviteSubaccountResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "邀请子帐号成员",
	FastFail:      false,
}

var testStep2780AddMemberToProject05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("AddMemberToProject")
		err = req.SetPayload(map[string]interface{}{
			"ProjectId":   step.Scenario.GetVar("ProjectId"),
			"MemberEmail": step.Scenario.GetVar("Email"),
			"CharacterId": step.Scenario.GetVar("CharacterId"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "AddMemberToProjectResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "添加成员到项目",
	FastFail:      false,
}

var testStep2780FreezeMember06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("FreezeMember")
		err = req.SetPayload(map[string]interface{}{
			"MemberEmail": step.Scenario.GetVar("Email"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "FreezeMemberResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "冻结成员",
	FastFail:      false,
}

var testStep2780RemoveMemberFromProject07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("RemoveMemberFromProject")
		err = req.SetPayload(map[string]interface{}{
			"ProjectId":   step.Scenario.GetVar("ProjectId"),
			"MemberEmail": step.Scenario.GetVar("Email"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "RemoveMemberFromProjectResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "从项目中移除成员",
	FastFail:      false,
}

var testStep2780TerminateMember08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("TerminateMember")
		err = req.SetPayload(map[string]interface{}{
			"MemberEmail": step.Scenario.GetVar("Email"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "TerminateMemberResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除子账号",
	FastFail:      false,
}

var testStep2780TerminateCharacter09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("TerminateCharacter")
		err = req.SetPayload(map[string]interface{}{
			"CharacterId": step.Scenario.GetVar("CharacterId"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "TerminateCharacterResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除角色",
	FastFail:      false,
}

var testStep2780TerminateProject10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UAccount")
		if err != nil {
			return nil, err
		}
		client := c.(*uaccount.UAccountClient)

		req := client.NewTerminateProjectRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ProjectId": step.Scenario.GetVar("ProjectId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateProject(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "TerminateProjectResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除项目",
	FastFail:      false,
}

var testStep2780GetBillDataFileUrl11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UBill")
		if err != nil {
			return nil, err
		}
		client := c.(*ubill.UBillClient)

		req := client.NewGetBillDataFileUrlRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"BillType":   0,
			"BillPeriod": step.Must(functions.GetTimestamp(10)),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.GetBillDataFileUrl(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "GetBillDataFileUrlResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "生成账单数据文件下载的 url",
	FastFail:      false,
}