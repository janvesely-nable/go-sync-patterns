# Concurrency Patterns in Go

## Go Concurrency 

While you can use most languages to implement concurrent code, Go stands out with its built-in support for concurrency making it easier.
This article is an overview of the concurrent support in Go accompanied using my own examples.
If you do not like my examples, before heading away, you might want to check the references section at the bottom, 
where I gathered various resources on the topic. 

## Basics

When writing concurrent code in Go, you make use of goroutines, which are directly built-in to the language. 
These are mapped on the OS level threads by the Go runtime and hence goroutines can be viewed as green threads (https://en.wikipedia.org/wiki/Green_thread). 
Goroutines are typically used with channels and the select statement. 
Additional synchronisation tools are in the standard [sync](https://pkg.go.dev/sync) package and [golang.org/x/sync](https://pkg.go.dev/golang.org/x/sync) package. 

### Concurrency and Parallelism

These terms are often confused hence Rob Pike from the Go team dedicated a whole lecture on the [topic](https://www.youtube.com/watch?v=oV9rvDllKEg). 
In brief, writing concurrent code using goroutines can make the code capable of running in parallel. 
The actual parallel execution can happen when the machine it runs on has multiple processors. 
Writing code using goroutine and channels allows to implement the [Actor model](https://en.wikipedia.org/wiki/Actor_model) and structure the code in that fashion, 
which may be seen as a benefit of its own.

## Main Patterns

### 1 WaitGroup

WaitGroup allows the simplest synchronisation between goroutines. It is typically used to await the completion of a go routine or a set of goroutines.
Here is an [example](./01-waitgroup/main.go).

### 2 ErrGroup

ErrGroup is a WaitGroup that allows the main goroutine to detect that any of the sub-goroutines had an error. So rather than adding a channel for the errors,
you can just use ErrGroup that does it for you. The ErrGroup can even handle context. Here is a simple example. (/.02-err-group/main.go).

### 3 Await Result

A channel facilitates a simple synchronisation between goroutines. It is typically used to transfer some values from a task back to the main goroutine. 
The main goroutine waits by reading from the channel. When the task is finished, it will close the channel and the main goroutine will be unblocked. 
Here is an [example](./03-await-result/main.go).

### 4 Await Result with Cancellation Context

Suppose we need to remain responsive to two different scenarios.
 - task completes and reports results as in the previous case
 - there is a cancellation request and the main goroutine should stop waiting when context is cancelled

Here is an [example](./04-await-result-context/main.go).

Note: If you wanted the sub-task to be terminated, the context would also have to be passed in to the task goroutine 
and a select statement would be used to determine whether the next iteration of the computation should run or whether the computation should be aborted.

### 5 Use a Mutex to protect updates to data by multiple goroutines

Communicating by sharing data is not recommended. However, Mutex is still available in the standard library, and it is occasionally needed.
So here is an [example](./05-mutex/main.go), which also shows perhaps a lesser known fact about Golang's Mutex: Mutex is not clever enough to know
when a lock is already held by a goroutine appendValue(value int) method.

### 6 Use RWMutex when you have readers and writers

When shared data is used by distinct writers and distinct readers, performance may be improved if a Mutex is replaced with and RWMutex. 
The RWMutex is smart enough to ensure an attempted writer Lock() eventually becomes available by blocking any subsequent reader RLock() calls. 
Here is an [example](./06-rw-mutex/main.go).

### 7 Once

If you want to ensure a certain function is only called once, you can create a wrapper sync.Once object 
and invoke the method through that wrapper. Here is an [example](./07-once/main.go).

### 8 Producers and Consumers

Go channel can be used to connect consumers and producer processes. In fact, it is possible to have multiple producers and multiple consumers using the same channel. 
When a channel is used in this way, specifying a buffer size is necessary to utilise all producers and consumers. 
See an [example](./08-producer-consumer/main.go).

### 9 sync.Map

Standard Go maps are not safe for concurrent use. Rather than creating a struct with an embedded Mutex, 
just use sync.Map. However, when it is necessary to co-ordinate data outside its content as well,
it is recommended to use the standard map with separate locking and co-ordination.

### 10 sync.Pool

A set of temporary objects that may be individually saved or retrieved in a way of caching.

### 11 ErrGroup

Allows controlling multiple go routines simplifying the handling of Context (i.e. cancellation) and errors coming
from the go routines. 

### 12 SingleFlight

Allows suppressing duplicate longer-running function calls such as when request for a token is already in progress. 

## References

* Synchronisation packages supported by the Go team
  * [sync](https://pkg.go.dev/sync) package
  * [golang.org/x/sync](https://pkg.go.dev/golang.org/x/sync) package.

* Lecture on concurrency vs parallelism by Rob Pike https://www.youtube.com/watch?v=oV9rvDllKEg

* Concurrency patterns by Brian Mills in GopherCon: https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view

* Lecture on Advanced Concurrency techniques from GopherCon https://www.youtube.com/watch?v=y2zc9gvIMPM

* Interesting blog on concurrency patterns https://blogtitle.github.io/categories/concurrency/

* Udemy course on working with concurrency https://swmsp.udemy.com/course/working-with-concurrency-in-go-golang/learn/lecture/32032356?start=0#overview


