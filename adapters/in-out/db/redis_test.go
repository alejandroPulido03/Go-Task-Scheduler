package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	dtos "task-scheduler/adapters/DTOs"

	"task-scheduler/app/entities"
)

func TestRedisRepository_Save(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	repo := &RedisRepository{
		client:   rdb,
		ctx:      context.Background(),
	}

	task := &entities.Task{
		Uuid:    "task-1",
		Exp_time: time.Now(),
	}

	dto := dtos.TaskDTO{
		Entity: task,
	}
	data, _ := dto.ToJson()

	mock.ExpectZAdd("tasks", redis.Z{
		Member: data,
		Score:  float64(task.Exp_time.Unix()),
	}).SetVal(1)

	err := repo.Save(task)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisRepository_SaveRecovery(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	repo := &RedisRepository{
		client:   rdb,
		ctx:      context.Background(),
	}

	task := &entities.Task{
		Uuid:    "task-1",
		Exp_time: time.Now(),
	}

	dto := dtos.TaskDTO{
		Entity: task,
	}
	data, _ := dto.ToJson()

	mock.ExpectSAdd("tasks_recovery", data).SetVal(1)

	err := repo.SaveRecovery(task)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisRepository_GetFirst(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	repo := &RedisRepository{
		client:   rdb,
		ctx:      context.Background(),
	}


	task := &entities.Task{
		Uuid:    "task-1",
		Exp_time: time.Now(),
	}

	dto := dtos.TaskDTO{
		Entity: task,
	}
	data, _ := dto.ToJson()

	mock.ExpectZPopMin("tasks", 1).SetVal([]redis.Z{
		{
			Member: data,
			Score:  float64(task.Exp_time.Unix()),
		},
	})

	retrievedTask, err := repo.GetFirst()
	assert.NoError(t, err)
	assert.Equal(t, task.Uuid, retrievedTask.Uuid)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisRepository_RemoveRecovery(t *testing.T) {
	rdb, mock := redismock.NewClientMock()

	repo := &RedisRepository{
		client:   rdb,
		ctx:      context.Background(),
	}

	task := &entities.Task{
		Uuid:    "task-1",
		Exp_time: time.Now(),
	}

	dto := dtos.TaskDTO{
		Entity: task,
	}
	data, _ := dto.ToJson()

	mock.ExpectSRem("tasks_recovery", data).SetVal(1)
	

	err := repo.RemoveRecovery(task)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
