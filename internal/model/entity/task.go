// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id          uint        `json:"id"          orm:"id"           ` // task ID
	TaskType    string      `json:"taskType"    orm:"task_type"    ` // task Type
	TaskPayload string      `json:"taskPayload" orm:"task_payload" ` // task Payload
	TaskId      string      `json:"taskId"      orm:"task_id"      ` // task id
	TaskInfo    string      `json:"taskInfo"    orm:"task_info"    ` // task info
	TaskStatus  int         `json:"taskStatus"  orm:"task_status"  ` // task status
	ErrMsg      string      `json:"errMsg"      orm:"err_msg"      ` // err msg
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"    ` // Created Time
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"    ` // Updated Time
}
