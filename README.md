# Concurrency Patterns in Go

## Go Concurrency 

While you can use most languages to implement concurrent code, Go stands out with its built-in support for concurrency. This article is an overview of the concurrent support in Go using examples.

## Basics

When writing concurrent code in Go, you make use of goroutines, which are directly built-in to the language. These are mapped on the OS level threads by the Go runtime and hence go routines can be viewed as green threads (https://en.wikipedia.org/wiki/Green_thread). Goroutines typically are used with channels and the select statement. Additional synchronisation tools are in the standard [sync](https://pkg.go.dev/sync) package and [golang.org/x/sync](https://pkg.go.dev/golang.org/x/sync) package. 

### Concurrency and Parallelism

These terms are often confused hence Rob Pike from the Go team dedicated a whole lecture on the [topic](https://www.youtube.com/watch?v=oV9rvDllKEg). In brief, writing concurrent code using goroutines can make the code capable of running in parallel, provided that the machine it runs on has multiple processors. Writing code using goroutine and channels allows to implement the [Actor model](https://en.wikipedia.org/wiki/Actor_model) and structure the code in that fashion, which may be seen as a benefit of its own, and also giving the code an opportunity to utilise parallel capabilities of the system it is running on.

## Example Patterns

### 01 WaitGroup

WaitGroup allows the simplest synchronisation between go routines. It is typically used to await the completion of a go routine or a set of go routines. Here is an [example](./01-waitgroup/).
