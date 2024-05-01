// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table task for DAO operations like Where/Data.
type Task struct {
	g.Meta      `orm:"table:task, do:true"`
	Id          interface{} // task ID
	TaskType    interface{} // task Type
	TaskPayload interface{} // task Payload
	TaskId      interface{} // task id
	TaskInfo    interface{} // task info
	TaskStatus  interface{} // task status
	ErrMsg      interface{} // err msg
	CreateAt    *gtime.Time // Created Time
	UpdateAt    *gtime.Time // Updated Time
}
