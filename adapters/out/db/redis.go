package db

import (
	"context"
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
        DB:       0,  // use default DB
    })
	return &RedisRepository{
		client: rdb,
		ctx: context.Background(),
	}
}

func (r *RedisRepository) Save(task entities.Task) error{
	dto := dtos.TaskDTO{
		Entity: task,
	}
	
	data, err := dto.ToJson()

	if err != nil {
		return err
	}
	
	err = r.client.SAdd(r.ctx, task.Exp_time.String(), data).Err()
	
	
	fmt.Println("Task saved to redis")
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) GetByTime(time string) (entities.Task, error){
	fmt.Println("Task fetched by time from redis")
	return entities.Task{}, nil
}

func (r *RedisRepository) GetByInterval(start string, end string) (entities.Task, error){
	fmt.Println("Task fetched by interval from redis")
	return entities.Task{}, nil
}