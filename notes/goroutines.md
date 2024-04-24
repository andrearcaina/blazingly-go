# Goroutines (Concurrency in Go)

First we have to understand what concurrency is.

## Concurrency

- Concurrency is the ability to run multiple programs at the same time.
- It is not the same as parallelism or parallel processing.
  - Concurrency is about dealing with lots of things at once.
  - Parallelism is about doing lots of things at once.
- Concurrency is a way to structure a program by breaking it into pieces that can be executed independently.
- It's like having multiple tasks running at the same time of a single core processor.
  - An example of concurrency is a web server that can process multiple requests at the same time.
  - Another example is a chat application that can receive messages from multiple users at the same time.
  - Another example of concurrency in real life is a human being who can walk and talk at the same time.

A program can run concurrently but may or may not run in parallel. This is because concurrency is about dealing with lots of things at once, but it doesn't necessarily mean that a program is running multiple tasks at the same time.  A program running in parallel is always running concurrently (due to the fact that it is running multiple tasks at the same time).

## Goroutines

- Goroutines are functions or methods that run concurrently with other functions or methods.
- Goroutines can be thought of as lightweight threads and are managed by the Go runtime.
- Goroutines are useful for performing tasks independently and concurrently.

## Mutex

- A mutex is a lock that allows only one goroutine to access a shared resource at a time.
- A mutex is used to prevent mutual exclusion, which is when two or more goroutines try to access the same shared resource at the same time.
- A mutex is also used to prevent a race condition, which is when the output of a program depends on the order in which the goroutines are executed.

### Read-Write Mutex

- A read-write mutex is a lock that allows multiple goroutines to read a shared resource at the same time, but only one goroutine to write to the shared resource at a time.
- A read-write mutex is useful when the shared resource is read more often than it is written to.
- A read-write mutex is used to prevent the **Readers/Writers Problem**, which is when multiple goroutines try to read and write to a shared resource at the same time.

## WaitGroup

- A WaitGroup is used to wait for a collection of goroutines to finish executing.
- A WaitGroup is useful for coordinating the execution of multiple goroutines.
- A WaitGroup is used to prevent a program from exiting before all the goroutines have finished executing.

