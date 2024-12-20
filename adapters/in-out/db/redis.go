package db

import (
	"context"
	"encoding/json"
	"fmt"
	dtos "task-scheduler/adapters/DTOs"
	"task-scheduler/app/entities"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
	ctx context.Context
}

func NewRedisRepository() *RedisRepository {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0, 
    })


	return &RedisRepository{
		client: rdb,
		ctx: context.Background(),
	}
}

func (r *RedisRepository) Save(task *entities.Task) error{
	
	dto := dtos.TaskDTO{
		Entity: task,
	}
	
	data, err := dto.ToJson()
	

	if err != nil {
		return err
	}
	
	err = r.client.ZAdd(r.ctx, "tasks", 
		redis.Z{
			Member: data,
			Score: float64(task.Exp_time.Unix()),
		}).Err()
	
	
	fmt.Println("Task saved to redis")
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) SaveRecovery(task *entities.Task) error{
	dto := dtos.TaskDTO{
		Entity: task,
	}
	
	data, err := dto.ToJson()

	if err != nil {
		return err
	}
	
	err = r.client.SAdd(r.ctx, "tasks_recovery", data).Err()
	
	
	fmt.Println("Task saved to recovery")
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) GetFirst() (*entities.Task, error){
	fmt.Println("Task retrieved from redis")

	task_byte := r.client.ZPopMin(r.ctx, "tasks", 1).Val()

	if len(task_byte) == 0 {
		return nil, nil
	}
	
	task_json := []byte(task_byte[0].Member.(string))

	dto := dtos.TaskDTO{
		JSON: &dtos.TaskJSON{},
	}
	
	unm_err := json.Unmarshal(task_json, &dto.JSON)

	if unm_err != nil {
		return nil, unm_err
	}

	entity, err := dto.ToEntity()

	if err != nil {
		return nil, err
	}
	
	return entity, nil
}

func (r *RedisRepository) RemoveRecovery(task *entities.Task) error{
	fmt.Println("Task removed from redis")

	dto := dtos.TaskDTO{
		Entity: task,
	}
	
	data, err := dto.ToJson()
	if err != nil {
		return err
	}

	return r.client.SRem(r.ctx, "tasks_recovery", data).Err()
}