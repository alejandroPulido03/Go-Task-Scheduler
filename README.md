# Go Task Scheduler
An In-Memory Task Scheduler with recovery capacity and secondary storage in Redis.

This project is inspired by a [https://engineering.rappi.com/planificaci%C3%B3n-de-tareas-scheduler-4250961fa944](post) of a task-scheduler design proposed by the Rappi Tech Team. However, this project has a few differences, especially in the use case, which forces the use of non-spot machines (by the nature of in-memory storage).

Also, as a personal decision, the main storage is implemented as an in-memory Treap, a modern data structure that will be explained later.

## Introduction
In this document, we will talk about the technical details of this task scheduler, the trade-offs and decisions taken and the next steps in this project.

## Taking requirements
The functional requirement used by the construction of this software was the following:

1. As a client, I want to send an HTTP request description with a fixed time when I want to execute that request

Another functional requirements used (not all implemented yet)
