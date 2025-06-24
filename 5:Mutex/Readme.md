# Mutex

## Critical Section
* When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time. This section of code that modifies shared resources is called critical section.

## Race Condition
* the output of the program depends on the sequence of execution of Goroutines is called race condition

### What is Mutex ?
- A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time to prevent race conditions from happening.