# Package

[Everything you need to know about Packages in Go - Medium](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)

# Import

```go
//  Import one package
import "fmt"

//  Factored import statement
import (
	"fmt"
    "time"
)
```

# Alias

```go
var pl = fmt.Println
```

# Comments

```go
// This is a one-line comment

/*
This is a
multi-line comment
*/
```

# Main

```go
func main(){
	pl("What is your name?")
}
```

# User Input

```go
reader := bufio.NewReader(os.Stdin)
```

# Error Handling

```go
name , error := reader.ReadString('\n')
```

# Blank Identifier

```go
// This a bad practice
name , _ := reader.ReadString('\n')
```

# Variables

```go
	var vName string = "hwangblood"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4
	v4 = 5.4
```

# Data Types

int, float64, bool, string, rune

0 , 0.0 , false, "" ,

```go
	pl(reflect.TypeOf(25))      // int
	pl(reflect.TypeOf(3.14))    // float64
	pl(reflect.TypeOf(true))    // bool
	pl(reflect.TypeOf("hello")) // string
	pl(reflect.TypeOf('👋'))     // int32
```

# Casting

```go
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2) // 1

	// convert string to int
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4)) // 50000000 <nil> int

	// convert int to string
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	// convert string to float
	cV7 := "3.14"
	cV8, err := strconv.ParseFloat(cV7, 64)
	if err == nil {
		pl(reflect.TypeOf(cV8)) // float64
	}

	// format print
	cV9 := fmt.Sprintf("cV9 = %f", 3.14)
	pl(cV9)
```

# If Conditional

Conditional Operators: > < >= <= == !=

Logical Operators: && || !

```go
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an important Birthday")
	}

	pl("!ture =", !true) // false
```

# Strings

functions for string

# Runes

https://stackoverflow.com/questions/19310700/what-is-a-rune/19311218#19311218

```go
utf8.RuneCountInString("abcd") // 4
```

# Printf

https://pkg.go.dev/fmt#Printf

Go has its own version of C's printf
%d : Integer
%c : Character
%f : Float
%t : Boolean
%s : String
%o : Base 8
%x : Base 16
%v : Guesses based on data type
%T : Type of supplied value

# Time

https://pkg.go.dev/time

```go
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
```

# Math

https://pkg.go.dev/math

+, -, \*, /, %

++, --

```go
5 + 4 = 9
5 - 4 = 1
5 * 4 = 20
5 / 4 = 1
5 % 4 = 1
10.0 / 3 = 3.3333333333333335
Random: 28
math.Abs(-10) = 10
math.Pow(4, 2)  16
math.Sqrt(16) = 4
math.Cbrt(8) = 2
math.Ceil(4.4) = 5
math.Floor(4.4) = 4
math.Round(4.4) = 4
math.Log2(8) = 3
math.Log10(100) = 2
math.Log(math.Exp(2)) = 2
math.Max(5, 4) = 5
math.Max(4, 5) = 5
1.57 radians = 90.00 degreesn
Sin(90) = 1
```

There are also functions for Cos, Tan, Acos, Asin,
Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos,
Htpot

# For Loop

```go
	for x := 5; x >= 1; x-- {
		// for x := 1; x <= 5; x++ {
		pl(x)
	}

	for x := 1; x <= 5; x++ {
		pl(x)
	}
```

# While Loop

```go
fx := 0
for fx < 5 {
	pl(fx)
	fx++
}

n := 1
for true {
	n++
	if n > 10 {
		break
	}
}
```
