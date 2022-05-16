package main

import (
	"fmt"
	"io/ioutil"
	m "math"
	"net/http"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Hello world!")

	beyondHello()
}

func beyondHello() {
	var x int
	x = 3
	// "Short" declarations use := to infer the type, declare
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn Go!"

	s2 := `A "raw" string literal
can include line breaks.` // Same string type.

	g := 'Σ'

	f := 3.14195
	c := 3 + 4i

	var u uint = 7
	var pi float32 = 22. / 7

	n := byte('\n')

	var a4 [4]int
	a5 := [...]int{6, 0, 1, 0} // An array initialized with a fixed size

	a4_cpy := a4
	a4_cpy[0] = 25
	fmt.Println(a4_cpy[0] == a4[0]) // false

	// Slices have dynamic size. Arrays and slices each have advantages
	// but use cases for slices are much more common.
	s3 := []int{4, 5, 9}    // Compare to a5. No ellipsis here. // ellipsis省略号 // 没有省略号所以是个slice
	s4 := make([]int, 4)    // Allocates slice of 4 ints, initialized to all 0.
	var d2 [][]float64      // Declaration only, nothing allocated here.
	bs := []byte("a slice") // Type conversion syntax.

	// Slices (as well as maps and channels) have reference semantics.
	s3_cpy := s3                    // Both variables point to the same instance.
	s3_cpy[0] = 0                   // Which means both are updated.
	fmt.Println(s3_cpy[0] == s3[0]) // true  // 很神奇，slice，map和channels的=是指针

	// Because they are dynamic, slices can be appended to on-demand.
	// To append elements to a slice, the built-in append() function is used.
	// First argument is a slice to which we are appending. Commonly,
	// the array variable is updated in place, as in example below.
	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	p, q := learnMemory() // Declares p, q to be type pointer to int.
	fmt.Println(*p, *q)   // * follows a pointer. This prints two ints.

	// Maps are a dynamically growable associative array type, like the
	// hash or dictionary types of some other languages.
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a5, s4, bs

	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is how you write to a file, by the way")
	file.Close()

	// Output of course counts as using a variable.
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl() // Back in the flow.
}

// Go is fully garbage collected. It has pointers but no pointer arithmetic.
// You can make a mistake with a nil pointer, but not by incrementing a pointer.
// Unlike in C/Cpp taking and returning an address of a local variable is also safe.
func learnMemory() (p, q *int) {
	p = new(int) // Built-in function new allocates memory.  // 这个的作用？
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFlowControl() {
	if true {
		fmt.Println("told ya")
	}

	if false {
		//Pout.
	} else {
		//Gloat
	}

	x := 42.0
	switch x {
	case 0:
	case 1, 2: // Can have multiple matches on one case
	case 42:
	// Cases don't "fall through".
	/*
	   There is a `fallthrough` keyword however, see:
	     https://github.com/golang/go/wiki/Switch#fall-through
	*/
	case 43:
		// Unreached.
	default:
		// Default case is optional.
	}

	// Type switch allows switching on the type of something instead of value
	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println(c, "is a string")
	case int64:
		fmt.Printf("%d is an int64\n", c)
	default:
		// all other cases
	}

	for x := 0; x < 3; x++ {
		fmt.Println("iteration", x)
	}

	// For is the only loop statement in Go, but it has alternate forms.
	for { // Infinite loop.
		break    // Just kidding.
		continue // Unreached.
	}

	// You can use range to iterate over an array, a slice, a string, a map, or a channel.
	// range returns one (channel) or two values (array, slice, string and map).
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Println("key=%s, value=%d\n", key, value)
	}
	// If you only need the value, use the underscore as the key
	for _, name := range []string{"Bob", "Bill", "Joe"} { // 虽然是数组但返回也是key，value的格式
		fmt.Printf("Hello, %s\n", name)
	}

	// As with for, := in an if statement means to declare and assign
	// y first, then test y > x.
	if y := expensiveComputation(); y > x {
		x = y
	}

	// 函数式编程
	xBig := func() bool {
		return x > 10000
	}
	x = 99999
	fmt.Println("xBig:", xBig())
	x = 1.3e3
	fmt.Println("xBig:", xBig())

	// What's more is function literals may be defined and called inline,
	// acting as an argument to function, as long as:
	// a) function literal is called immediately (),
	// b) result type matches expected type of argument.
	fmt.Println("Add + double two numbers:",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // Called with args 10 and 2
	// => Add + double two numbers: 24

	goto love
love:
	learnFunctionFactory() // func returning func is fun(3)(3)
	learnDefer()           // A quick detour to an important keyword.
	learnInterfaces()      // Good stuff coming up!
}

func learnFunctionFactory() {
	// Next two are equivalent, with second being more practical
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))
}

// Decorators are common in other languages. Same can be done in Go
// with function literals that accept arguments.
// 妙r，有点像haskell的单一参数的函数
// 之前用go过一遍的设计模式的Factory模式没有用到，用上的话会巧妙
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintln("%s %s %s", before, mystring, after)
	}
}

func learnDefer() (ok bool) {
	// A defer statement pushes a function call onto a list. The list of saved
	// calls is executed AFTER the surrounding function returns.
	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because") // defer是LIFO实现
	// Defer is commonly used to close a file, so the function closing the
	// file stays close to the function opening the file.
	return true
}

// Define Stringer as an interface type with one method, String.
type Stringer interface {
	String() string
}

//
type pair struct {
	x, y int
}

func (p pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	p := pair{3, 4}
	fmt.Println(p.String())
	var i Stringer
	i = p

	fmt.Println(i.String())

	fmt.Println(p)
	fmt.Println(i)

	learnVariadicParams("great", "learning", "here!")
}

// Functions can have variadic parameters.
func learnVariadicParams(myStrings ...interface{}) {
	// Iterate each value of the variadic.
	// The underbar here is ignoring the index argument of the array.
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// Pass variadic value as a variadic parameter.
	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// ", ok" idiom used to tell if something worked or not.
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok { // ok will be false because 1 is not in the map.
		fmt.Println("no one there")
	} else {
		fmt.Print(x) // x would be the value, if it were in the map.
	}
	// An error value communicates not just "ok" but more about the problem.
	if _, err := strconv.Atoi("non-int"); err != nil {
		// prints 'strconv.ParseInt: parsing "non-int": invalid syntax'
		fmt.Println(err)
	}
	// We'll revisit interfaces a little later. Meanwhile,
	learnConcurrency()
}

func inc(i int, c chan int) {
	c <- i + 1
}

func learnConcurrency() {

	c := make(chan int)

	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)

	fmt.Println(<-c, <-c, <-c)

	cs := make(chan string)
	ccs := make(chan chan string) // A channel of string channels.
	go func() { c <- 84 }()       // Start a new goroutine just to send a value.
	go func() { cs <- "wordy" }()

	// Select has syntax like a switch statement but each case involves
	// a channel operation. It selects a case at random out of the cases
	// that are ready to communicate.
	select {
	case i := <-c:
		fmt.Printf("it's a %T", i)
	case <-cs:
		fmt.Println("it's a string")
	case <-ccs:
		fmt.Println("didn't happen.")
	}
	// At this point a value was taken from either c or cs. One of the two
	// goroutines started above has completed, the other will remain blocked.

	learnWebProgramming() // Go does it. You want to do it too.
}

// A single function from package http starts a web server.
func learnWebProgramming() {

	// First parameter of ListenAndServe is TCP address to listen to.
	// Second parameter is an interface, specifically http.Handler.
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // don't ignore errors
	}()

	requestServer()
}

func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You Learned Go in Y minutes!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`", string(body))
}
