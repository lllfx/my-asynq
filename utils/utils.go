package utils

import (
	"context"
	"encoding/json"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
	"my-asynq/internal/dao"
	"my-asynq/internal/model/entity"
)

// TaskStatus 任务状态
type TaskStatus int

const (
	TASK_STATUS_INVALID   TaskStatus = 0
	TASK_STATUS_UNSTART   TaskStatus = 1
	TASK_STATUS_WAITING   TaskStatus = 2
	TASK_STATUS_RUNNING   TaskStatus = 3
	TASK_STATUS_SUCCESS   TaskStatus = 4
	TASK_STATUS_FAILED    TaskStatus = 5
	TASK_STATUS_STOPED    TaskStatus = 6
	TASK_STATUS_DELETE    TaskStatus = 7
	TASK_STATUS_EXPORTING TaskStatus = 8
)

func GetRedisClientOpt(ctx context.Context) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{Addr: g.Cfg().MustGet(ctx, "redis.addr", "").String(),
		Password: g.Cfg().MustGet(ctx, "redis.password", "").String()}
}

func enqueueTask(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	client := asynq.NewClient(GetRedisClientOpt(ctx))
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(client)
	taskInfo, err := client.Enqueue(task, opts...)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return taskInfo, nil
}
func EnqueueTask(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	//先写表
	var resTaskInfo *asynq.TaskInfo
	err := dao.Task.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := tx.Ctx(ctx).Model(dao.Task.Table()).Data(&g.Map{
			dao.Task.Columns().TaskType:    task.Type(),
			dao.Task.Columns().TaskPayload: string(task.Payload()),
			dao.Task.Columns().TaskInfo:    "",
			dao.Task.Columns().TaskId:      "",
			dao.Task.Columns().TaskStatus:  int(TASK_STATUS_INVALID),
		}).FieldsEx(dao.Task.Columns().Id).InsertAndGetId()
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		taskInfo, err := enqueueTask(ctx, task, opts...)
		if err != nil {
			g.Log().Error(ctx, err)
			_, err1 := tx.Ctx(ctx).Model(dao.Task.Table()).Data(&g.Map{
				dao.Task.Columns().ErrMsg:     err.Error(),
				dao.Task.Columns().TaskStatus: int(TASK_STATUS_INVALID),
			}).Where(dao.Task.Columns().Id, id).Update()
			if err1 != nil {
				g.Log().Error(ctx, err1)
				return err1
			}
			//保留这个错误记录
		} else {
			resTaskInfo = taskInfo
			taskInfoByte, err := json.Marshal(taskInfo)
			if err != nil {
				g.Log().Error(ctx, err)
			}
			taskInfoStr := ""
			if taskInfoByte != nil {
				taskInfoStr = string(taskInfoByte)
			}
			_, err1 := tx.Ctx(ctx).Model(dao.Task.Table()).Data(&g.Map{
				dao.Task.Columns().TaskInfo:   taskInfoStr,
				dao.Task.Columns().TaskId:     taskInfo.ID,
				dao.Task.Columns().TaskStatus: int(TASK_STATUS_WAITING),
			}).Where(dao.Task.Columns().Id, id).Update()
			if err1 != nil {
				g.Log().Error(ctx, err1)
				return err1
			}
		}
		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return resTaskInfo, nil
}

type TaskHandler struct {
	handler func(ctx context.Context, t *asynq.Task) error
}

func NewTaskHandler(handler func(ctx context.Context, t *asynq.Task) error) *TaskHandler {
	return &TaskHandler{
		handler: handler,
	}
}

func (p *TaskHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	//处理
	taskId := t.ResultWriter().TaskID()
	taskEntity := new(entity.Task)
	err := dao.Task.Ctx(ctx).Where(dao.Task.Columns().TaskId, taskId).Scan(&taskEntity)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	_, err = dao.Task.Ctx(ctx).Data(g.Map{
		dao.Task.Columns().TaskStatus: int(TASK_STATUS_RUNNING),
	}).Where(dao.Task.Columns().TaskId, taskId).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	err = p.handler(ctx, t)
	if err != nil {
		g.Log().Error(ctx, err)
		_, err1 := dao.Task.Ctx(ctx).Data(g.Map{
			dao.Task.Columns().TaskStatus: int(TASK_STATUS_FAILED),
			dao.Task.Columns().ErrMsg:     err.Error(),
		}).Where(dao.Task.Columns().TaskId, taskId).Update()
		if err1 != nil {
			g.Log().Error(ctx, err1)
		}
		return err
	}
	_, err = dao.Task.Ctx(ctx).Data(g.Map{
		dao.Task.Columns().TaskStatus: int(TASK_STATUS_SUCCESS),
	}).Where(dao.Task.Columns().TaskId, taskId).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
