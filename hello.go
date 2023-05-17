// program starts executing at the main package.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	// * Say "Hello somebody"
	/*
		pl("What is your name?")
		   	reader := bufio.NewReader(os.Stdin)
		   	// Stop reading util get a new line or type the ENTER key of keyboard
		   	name, err := reader.ReadString('\n')
		   	if err == nil {
		   		pl("Hello", name)
		   	} else {
		   		log.Fatal(err)
		   	}
	*/

	// * Variables
	/*
		// var name type
		var vName string = "hwangblood"
		var v1, v2 = 1.2, 3.4
		var v3 = "hello"
		v4 := 2.4
		v4 = 5.4
	*/

	// * Data Types
	/*
		pl(reflect.TypeOf(25))      // int
		pl(reflect.TypeOf(3.14))    // float64
		pl(reflect.TypeOf(true))    // bool
		pl(reflect.TypeOf("hello")) // string
		pl(reflect.TypeOf('👋'))     // int32
	*/

	// * Casting
	/*
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
	*/

	// * If Conditional
	/*
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
	*/

	// * Strings
	/*
		sV1 := "A world"
		replacer := strings.NewReplacer("A", "Replaced")
		sV2 := replacer.Replace((sV1))
		pl(sV2)

		pl("Length:", len("Have a good day!"))
		pl("Contains good:", strings.Contains("good night", "good"))
		pl("o index:", strings.Index("looking", "o"))
		pl("Replace:", strings.Replace("course", "o", "0000", -1))

		sV3 := "\nSome Words\n" // more like: \t \" \\
		sV3 = strings.TrimSpace(sV3)
		pl(sV3)

		pl("Split:", strings.Split("a-b-c-d", "-"))
		pl("Lower:", strings.ToLower("ABCdef"))
		pl("Upper:", strings.ToUpper("xyZ"))
		pl("Prefix:", strings.HasPrefix("tacocat", "taco"))
		pl("Suffix:", strings.HasSuffix("tacocat", "cat"))
	*/

	// * Runes
	/*
		rStr := "abcd"
		pl("Rune Count:", utf8.RuneCountInString(rStr))
	*/

	// * Printf
	/*
		for i, runeVal := range "play" {
			fmt.Printf("%d - %#U - %c\n", i, runeVal, runeVal)
		}

		fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A',
			3.14, true, 1, 1)

		// Float formatting
		fmt.Printf("%9f\n", 3.14)      // Width 9
		fmt.Printf("%.2f\n", 3.141592) // Decimal precision 2
		fmt.Printf("%9.f\n", 3.141592) // Width 9 no precision

		// Sprintf returns a formatted string instead of printing
		sp1 := fmt.Sprintf("%9.f\n", 3.141592)
		pl(sp1)

	*/

	// * Time
	/*
		now := time.Now()
		pl(now.Year(), now.Month(), now.Day())
		pl(now.Hour(), now.Minute(), now.Second())
	*/

	// * Math
	/*
		pl("5 + 4 =", 5+4)
		pl("5 - 4 =", 5-4)
		pl("5 * 4 =", 5*4)
		pl("5 / 4 =", 5/4)
		pl("5 % 4 =", 5%4)
		pl("10.0 / 3 =", 10.0/3)

		// Generate a random number

		seedSecs := time.Now().Unix()
		rand.Seed(seedSecs)          // ! Deprecated
		randNum := rand.Intn(50) + 1 // 1- 50
		pl("Random:", randNum)

		// There are many math functions
		pl("math.Abs(-10) =", math.Abs(-10))
		pl("math.Pow(4, 2) ", math.Pow(4, 2))
		pl("math.Sqrt(16) =", math.Sqrt(16))
		pl("math.Cbrt(8) =", math.Cbrt(8))
		pl("math.Ceil(4.4) =", math.Ceil(4.4))
		pl("math.Floor(4.4) =", math.Floor(4.4))
		pl("math.Round(4.4) =", math.Round(4.4))
		pl("math.Log2(8) =", math.Log2(8))
		pl("math.Log10(100) =", math.Log10(100))

		// Ge the log of e to the power of 2
		pl("math.Log(math.Exp(2)) =", math.Log(math.Exp(2)))
		pl("math.Max(5, 4) =", math.Max(5, 4))
		pl("math.Max(4, 5) =", math.Max(4, 5))

		// Convewrt 90 degrees to radians
		r90 := 90 * math.Pi / 180
		d90 := r90 * (180 / math.Pi)
		fmt.Printf("%.2f radians = %.2f degreesn\n", r90, d90)

		pl("Sin(90) =", math.Sin(r90))
		// There are also functions for Cos, Tan, Acos, Asin
		// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
		// Htpot
	*/

	// * For Loop
	/*
		// for initialization; condition;
		for x := 5; x >= 1; x-- {
			// for x := 1; x <= 5; x++ {
			pl(x)
		}
	*/

	// * While Loop
	/*
		fx := 0
		for fx < 5 {
			pl(fx)
			fx++
		}

		seedSecs := time.Now().Unix()
		rand.Seed(seedSecs)
		randNum := rand.Intn(10) + 1

		for true {
			fmt.Print("Guess a number between 0 and 10 : ")
			reader := bufio.NewReader(os.Stdin)
			guess, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			guess = strings.TrimSpace(guess)
			iGuess, err := strconv.Atoi(guess)
			if err != nil {
				log.Fatal(err)
			}
			if iGuess > randNum {
				pl("Pick a Lower Value")
			} else if iGuess < randNum {
				pl("Pick a Higher Value")
			} else {
				pl("You guessed it")
				break
			}
		}
		pl("Random Number is :", randNum)
	*/

	// * Range
	/*
		aNums := []int{1, 2, 3, 4, 5, 6}
		for _, num := range aNums {
			pl(num)
		}
	*/

	// * Arrays
	/*
		var arr1 [5]int
		arr1[0] = 1
		arr2 := [5]int{1, 2, 3, 4, 5}
		pl("index 0:", arr2[0])
		pl("arr length:", len(arr2))
		for i := 0; i < len(arr2); i++ {
			pl(arr2[i])
		}
		for i, v := range arr2 {
			fmt.Printf("index %d: %d\n", i, v)
		}
		arr3 := [2][2]int{
			{1, 2},
			{3},
		}
		for i := 0; i < len(arr3); i++ {
			for j := 0; j < len(arr3[i]); j++ {
				pl(arr3[i][j])
			}
		}
		aStr1 := "abcde"
		rArr1 := []rune(aStr1)
		for i, r := range rArr1 {
			fmt.Printf("Rune %d: %d\n", i, r)
		}
		byteArr := []byte{'a', 'b', 'c', 'd', 'e'}
		bStr := string(byteArr[:])
		pl("byte array is a string:", bStr)
	*/
	// * Slices
	/*
		// var name []datatype
		sl1 := make([]string, 6)
		sl1[0] = "Society"
		sl1[1] = "of"
		sl1[2] = "the"
		sl1[3] = "Simulated"
		sl1[4] = "Universe"
		pl("Slice 1 length:", len(sl1))
		for i, v := range sl1 {
			fmt.Printf("index %d: %s\n", i, v)
		}

		sliArr := [5]int{1, 2, 3, 4, 5}
		sl3 := sliArr[:2] // [0:2]
		pl("sl3[:3]:", sl3[:3])
		pl("sliArr[2:]", sliArr[2:])
		sliArr[0] = 10
		pl("sl3:", sl3)
		sl3[0] = 8
		pl("sliArr:", sliArr)

		sl3 = append(sl3, 12)
		pl("sl3:", sl3)

		sl4 := make([]string, 6)
		pl("sl4:", sl4)
		pl("sl4[0]:", sl4[0]) // "", empty string
	*/

	pl(getSum(2, 5))

	twice, pow := getTwo(5)
	pl(twice, pow)

	pl(getQuotient(5, 0))
	pl(getQuotient(5, 4))

	pl(getSum2(1, 2, 3, 4))
	pl(getSum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	pl("array sum", getArrSum([]int{2, 4, 6, 8, 10, 12, 14}))

	num1 := 9
	pl("num1 before func:", num1)
	changeVarInFunc(num1)
	pl("num1 after func:", num1)

	v4 := 10
	var p4 *int = &v4 // store pointer
	pl("p4 address:", p4)
	pl("p4 value:", *p4)
	*p4 = 14
	pl("p4 value:", *p4)
	pl("v4 before func:", v4)
	changeVarInFunc2(&v4)
	pl("v4 after func:", v4)

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl("pArr:", pArr)
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average: %.3f\n", getAverage(iSlice...))

	fileIO()
	fileIO2()
}

// * Functions
// func funcName(parameters) returnType {BODY}
func getSum(x int, y int) int {
	return x + y
}

// * Return Multiple
func getTwo(x int) (int, int) {
	return x * 2, x * x
}

// * Function Errors
func getQuotient(x float64, y float64) (answer float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

// * Varadic Functions
// it's able to recive an unknown number of parameters
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// * Passing Arrays
func getArrSum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func changeVarInFunc(x int) int {
	x += x
	return x
}

// * Pointers
func changeVarInFunc2(xPtr *int) {
	*xPtr += *xPtr
}

// * Pass Array Pointers
func dblArrVals(arr *[4]int) {
	for i := 0; i < 4; i++ {
		arr[i] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, v := range nums {
		sum += v
	}
	return sum / numSize

}

// * File IO
func fileIO() {
	// write file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, s := range sPrimeArr {
		_, err := f.WriteString(s + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	// read file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}

// Append to file
/*
Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified

O_RDONLY : open the file read-only
O_WRONLY : open the file write-only
O_RDWR   : open the file read-write

These can be or'ed

O_APPEND : append data to the file when writing
O_CREATE : create a new file if none exists
O_EXCL   : used with O_CREATE, file must not exist
O_SYNC   : open for synchronous I/O
O_TRUNC  : truncate regular writable file when opened
*/
func fileIO2() {
	// Check if file exists
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("File doesn't exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("append line\n"); err != nil {
			log.Fatal(err)

		}
	}

}

// * Command Line

// * Packages / Modules
// * Maps
// * Generics
// * Constraints
// * Structs
// * Composition
// * Defined types
// * Associate Methods
// * Protecting Data
// * Getter / Setter
// * Encapsulation
// * Interfaces
// * Concurrency / GoRoutines
// * Sleep
// * Channels
// * Mutex / Lock
// * Closures
// * Passing Functions
// * Recursion
// * Regular Expressions
// * Automated Testing
// * Web app
// * Templates / HTML
// * Installation
