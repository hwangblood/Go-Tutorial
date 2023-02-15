// program starts executing at the main package.
package main

import (
	"fmt"
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
}
