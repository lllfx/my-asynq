// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskDao is the data access object for table task.
type TaskDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TaskColumns // columns contains all the column names of Table for convenient usage.
}

// TaskColumns defines and stores column names for table task.
type TaskColumns struct {
	Id          string // id
	TaskType    string // task Type
	TaskPayload string // task Payload
	TaskId      string // task id
	TaskInfo    string // task info
	TaskStatus  string // task status
	ErrMsg      string // err msg
	CreateAt    string // Created Time
	UpdateAt    string // Updated Time
}

// taskColumns holds the columns for table task.
var taskColumns = TaskColumns{
	Id:          "id",
	TaskType:    "task_type",
	TaskPayload: "task_payload",
	TaskId:      "task_id",
	TaskInfo:    "task_info",
	TaskStatus:  "task_status",
	ErrMsg:      "err_msg",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}

// NewTaskDao creates and returns a new DAO object for table data access.
func NewTaskDao() *TaskDao {
	return &TaskDao{
		group:   "default",
		table:   "task",
		columns: taskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskDao) Columns() TaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
