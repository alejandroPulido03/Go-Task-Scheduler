# Go Task Scheduler
An In-Memory Task Scheduler with recovery capacity and secondary storage in Redis.

This project is inspired by a [https://engineering.rappi.com/planificaci%C3%B3n-de-tareas-scheduler-4250961fa944](post) of a task-scheduler design proposed by the Rappi Tech Team. However, this project has a few differences, especially in the use case, which forces the use of non-spot machines (by the nature of in-memory storage).

Also, as a personal decision, the main storage is implemented as an in-memory Treap, a modern data structure that will be explained later.

## Introduction
In this document, we will talk about the technical details of this task scheduler, the trade-offs and decisions taken and the next steps in this project.

## Taking requirements
The functional requirement used by the construction of this software was the following:

1. As a client, I want to send an HTTP request description with a fixed time when I want to execute that request

Another functional requirements used (not all implemented yet) were:
1. As a client, I want to send an HTTP request description with a time schema when I want to execute that request
2. As a client, I want to specify a WebHook to get the confirmation of the task

To simulate the standards of a real product, the following quality requirements were taken:
1. As a client, I want a margin error in the time of execution of +/- 1 sec.
2. As a client, in the scenario of a downgrade, I want the scheduled tasks to be executed when the system is up again

## Architecture

![image](https://github.com/user-attachments/assets/0ab22798-dbe5-43b9-9819-3ca48f908ecd)


Because of the simplicity of the number of tasks that the scheduler was to do, It was decided to use a monolithic architecture, with three main responsibilities (traduced in services):
- Task record service: A service to listen to incoming requests, validate the business rules, and call the repository
- Repository service: A service to communicate with the storage of the data to save, retrieve, and delete tasks
- Worker service: A service to pool the time of execution of requests and to execute this request in the specified moment

Also, the system had many different inputs and outputs of data, so it was decided to use a hexagonal architecture, that allows the use of an abstraction model called ports and adapters, this method enables keeping the core logic separated from the in/out of the data, also allows having many in/out data connections in an organized way.

The storage technology decisions were made for a system with low latency needs, so the tasks are saved in an in-memory data model that will be explained later. This main storage has a replica in a relational Postgres database because of the simplicity of the data relations and how robust this technology is. 

Also, as the gross of the requests, the main storage could not maintain all the data, so we needed a secondary storage, with more flexibility in latency, for this secondary storage, the decision was to use Redis. Redis could be used as a non-relational database, with the flexibility of the use of structures to improve the time of retrieving and saving tasks. We use a sorted set structure because we have a sort key in the data, the time of execution, so we can get in a fast way some registers using the time of execution. Because Redis is also an in-ram database, we need to generate a replica of the data.

Finally, we decided to use asynchronous communication for the sending of the confirmation messages to the webhook, because a client could receive a lot of confirmation messages and the endpoint of the webhook could be down, so we decided to use the message queue rabbitMQ for this requirement.

## Implementation details

### Task record service
The unique interesting detail of this service is the must to confirm that the time of the scheduled request was being at least one minute after the actual time, this because the worker retrieves the next minute's tasks to execute it, and the arrival of a task in the actual minute will not be considered (this use case is very rare in a real-use so it's not worth losing sleep over.)








