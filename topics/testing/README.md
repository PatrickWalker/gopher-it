
  

# Testing

  

## Hello World

  

Let's say we have  file called sum.go
```go
package main

func Sum(x int, y int) int {
    return x + y
}

func main() {
    Sum(5, 5)
}
```
then in another file sum_test.go (and can be another package)
```go
package main
//std lib yo
import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
    // this is how we fail the test
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
```

Testing is a big thing in Go. You may be used to needing external libraries to perform your unit testing (things like JUnit and NUnit) but in Go it's part of the [standard lib](https://golang.org/pkg/testing/). 

  ### Naming

So there's a bit to unpack here so let's get started.

 - The file you created was called _test.go. That's important that's how
   go knows it's a test file and so doesn't put it in your production distribution
- The function is called TestXyz again that's a convention for a unit test. *func TestXxx(\*testing.T)*  is the standard the testing.T object is pretty important as it orchestrates (and fails the tests)
- It's in the same package which allows you to test unexported functions but you can perform more 'black box' testing by testing from another package if you want. That has the benefit of making you act like a consumer and is [recommended by many](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c#.q88391hne)

If you're new to Unit Testing never fear the goal is really to try and isolate your code and logic from as many dependencies as you can and test that small unit. That should mean the tests are fast to run and have less reason to change. Unit tests are traditionally thought of as the cheapest form of automation testing. They aren't perfect but they act as a real comfort blanket when making changes to your code and should be the solid foundation on which you build with quality.

## Assert your authority

Go test does not include an assertion library or package. If you're used to Junit and Nunit this will feel weird. You're used to doing Assert.IsEmpty or the like in your tests and so it can be a bit weird initially in Go but there's reasonings behind it

 - most assertion functions are easily reproduced if you need but writing go code gives you lots more flexibility
 - execution can continue if you use t.Error vs ending immediately with t.Fatal 

If assertions are a deal breaker to you then there are options. [Testify](https://github.com/stretchr/testify) is the most popular without doubt. 

Personal recommendation would be try and deal without them for as long as you can. Write a few helper functions if you need them but try and keep it as vanilla as possible

# Exercise 1

 
> Stop with the the Fizzbuzz mate. NEVER

> I did fizzbuzz wrong. Write some tests to show the error of my ways

> Then fix the code and run again

  
* FizzBuzz wrong  [source][fbs] / [playground][fbp]
* Exercise template: [source][t1s] / [playground][t1p]

* Example solution: [source][s1s] / [playground][s1p]

  

## Test Tables

 You probably noticed the solution from the last exercise feels a bit janky. There's a lot of repeated stuff there. Could we make it better? Step in test tables or data driven testing 

```go
package main

import "testing"

func TestFizzTable(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected string
	}{
		{"one","1","1"},
		{"zero","0","0"},
		{"3=fizz","3","fizz"},
		{"5=buzz","5","buzz"},
		{"15=fizzbuzz","15","0"},
		
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		out := fizzbuzz(test.input)
		if out != test.expected {
			// this is how we fail the test
			t.Errorf("Output was incorrect, got: %v, want: %v.", out, test.expected)
		}
		})
	}
}	
```
[playground](https://play.golang.org/p/h1fqjmrxahs)
t.Run is us calling a sub test. That's why we have the name too this allows us to get quicker feedback when something fails


## Are you mocking me?
No. I mean I'm really not because mocking isn't something that's part of the test lib in Go. 

### What is mocking?

[MOQ](https://github.com/Moq/moq4/wiki/Quickstart) and Moqito in Java are good examples. What happens is you specify you're creating an implementation of an interface and use 'setup' to put in listeners or behaviour you want for your test...

**So shouldn't Go have it?** I mean it certainly would be expected but the argument go has is....use interfaces. 

The fringe benefit here being if you use dependency injection your actual production code should (in theory) be easier to change in future. Also the way Go applies interfaces implicitly means we can even wrap external objects and dependencies in interfaces that make sense to us [Worked Example](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)

Again if this causes you headaches testify has you covered.

## Examples
 Little easter egg here but go test can help ensure your documentation examples are correct. Check out this [this calhoun article](https://www.calhoun.io/how-to-test-with-go/)  



## Coverage

Test coverage can be a useful if not imperfect metric. Understanding how much of your code has unit test coverage can show where you may want to put some attention to.
  
 Again this is built in and you can see the output of this in your IDE as well if you want. 
  
  Run go test -cover to see the coverage of the fizzbuzz code
  ``` go
W-D-M-PW:fixbuzz patrick$ go test -cover
coverage: 100.0% of statements

```

This is great 100% is amazing but what if it was 50% how would you know which lines weren't covered? HTML report got you

``` go
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 
```
Open the html file and you're covered or your friendly IDE can do it for you

## Race Conditions
Go test execution has the support for detecting race conditions. That said a lot of the toolchain does but if you want to be a little more confident with your tests run the following
  
  ```go
  test -race ./...
```
Issues here can be annoying to fix but the detector makes it a lot easier than getting a phone call at 2 in the morning that production is acting funny :) 
[More reading](https://blog.golang.org/race-detector)

## Benchmarks

Performance testing or benchmarking is also part of the same library. This is amazing. Think you might be able to make a function more performant? Well wonder no more as you can check
  

```go
package main

func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}

```

```go
package main

import "testing"

func benchmarkFib(i int, b *testing.B) {
		// b.N is a value that gives 
        for n := 0; n < b.N; n++ {
                Fib(i)
        }
}

//couldn't we use test driven tests for this? we could but as we don't check output we wouldn't save much
func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
//test name is BencharkXxxx not TestXxxxx
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
//its a different argument as well
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
```


The output of the above looks like 
```
TW-D-M-PW:bench patrick$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkFib1-8         1000000000               2.26 ns/op
BenchmarkFib2-8         200000000                6.53 ns/op
BenchmarkFib3-8         100000000               11.3 ns/op
BenchmarkFib10-8         5000000               389 ns/op
BenchmarkFib20-8           30000             48129 ns/op
BenchmarkFib40-8               2         721651723 ns/op
PASS
ok      _/Users/patrick/code/go-102/topics/testing/bench        12.048s
```

*one thing* - Unless you can guarantee that the machine is pretty standardly utilized ie - same usage profile between runs then these benchmarks are hard to compare run by runs so don't panic if tomorrow it's slower. Running it more frequently helps you understand your true values

Performance testing at that level is a science this is more a trend indicator

[Dave Cheney](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) has written blog posts on this and done some amazing videos.

[fbs]: exercises/fixbuzz/fizzbuzz.go

[fbp]: https://play.golang.org/p/GtIAcIEs2TP

[t1s]: exercises/fixbuzz/fizzbuzztmp_test.go

[t1p]:https://play.golang.org/p/I7NkuidFWBQ

[s1s]: exercises/fixbuzz/fizzbuzz_test.go

[s1p]: https://play.golang.org/p/g9t8625QBVm

