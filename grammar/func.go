package grammar

import "fmt"

//func (p mytype) funcname(q int) (r,s int) { return 0,0 }
//函数可以定义用于特定的类型mytype,这类函数更加通俗的称呼是 method。
//这部分称 作 receiver 而它是可选的。

func factorial1(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * factorial1(x-1)
	}
}

//naming return value
func factorial2(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * factorial2(x-1)
	}
	return
}

//defer demo
func deferDemo() {
	fmt.Printf("\ninside of deferDemo ...\n")
	fmt.Printf("defer cmds FILO \n")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("defer cmd: print i = %d\n", i)
	}
	fmt.Print("leaving deferDemo func \n")
}

//deferDemo2 should return 1
func deferDemo2() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

func deferDemo3() {
	fmt.Printf("inside of deferDemo3\n")

	fmt.Printf("Begin defer func...\n")
	defer func() {
		fmt.Printf("defer func(){}() //() \n ")
	}()
	defer func(x int) {
		fmt.Printf("defer func(x int){}(%d) //() \n ", x)
	}(5)
	fmt.Printf("End defer func...\n")
}

func varParamDemo(arg ...int) {
	fmt.Printf("\ninside of varParamDemo \nargs is: ")
	for i, n := range arg {
		fmt.Printf("i=%d->%d,", i, n)
	}
	fmt.Printf("\n")
}

func anonymousDemo() {
	fmt.Printf("\ninside of anonymousDemo \n")
	afunc := func() {
		println("i am a anonymous func")
	}
	afunc()

	fmt.Printf("afunc type is: %T\n", afunc)

}

func printX(x int) {
	fmt.Printf("printX: %v\n", x)
}

func callbackDemo(y int, f func(int)) {
	fmt.Printf("\ncallback define: func callbackDemo(y int, f func(int)) \n")
	fmt.Printf("before f(y)\n")
	f(y)
}

func throwsPanicDemo(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

func DemoFuncs() {
	fmt.Printf("\n-------------------------------------------\n")
	fmt.Printf("Run DemoFunc ...\n\n")

	fmt.Printf("os package *File.Write: \n")
	fmt.Printf("func (file *File) Write(b []byte) (n int, err error) \n")

	v := factorial1(5)
	fmt.Printf("v := factorial1(5) v=%d \n", v)

	v = factorial2(5)
	fmt.Printf("v =  factorial2(5) v=%d \n", v)

	deferDemo()

	v2 := deferDemo2()
	fmt.Printf("v2 := deferDemo2(), v2 = %d //should not v2 = 0\n", v2)

	deferDemo3()

	varParamDemo(2, 4, 6, 8, 10)

	anonymousDemo()

	callbackDemo(100, printX)
}
