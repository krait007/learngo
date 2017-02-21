package grammar

import "fmt"

func switchExample(i int) {
	fmt.Printf("switch i=%d result is: ", i)
	switch i {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}
	fmt.Printf("\n")

}

//demo flow control
//if else, switch case ..., for, goto break continue fallthrough
func DemoFlow() {
	fmt.Printf("-------------------------------------------\n")
	fmt.Printf("Run DemoFlow ...\n\n")
	len := 5
	if len > 10 {
		fmt.Printf("len > 10 is true \n")
	} else {
		fmt.Printf("len > 10 is false \n")
	}

	if len = 100; len > 10 {
		fmt.Printf("len = 100; len > 10 is true \n")
	} else {
		fmt.Printf("len = 100; len > 10 is false \n")
	}

	fmt.Printf("example for switch and fallthrough \n")
	switchExample(0)
	switchExample(1)
	switchExample(2)
	switchExample(3)
	switchExample(4)
	switchExample(5)
	switchExample(6)
	switchExample(7)

	fmt.Printf("example for: without expression after switch keyword \n")
	num := 5
	switch {
	case 0 <= num && num <= 3:
		fmt.Printf("num is 0-3")
	case 4 <= num && num <= 6:
		fmt.Printf("num is 4-6")
	case 7 <= num && num <= 9:
		fmt.Printf("num is 7-9")
	}
	fmt.Print("\n")

	////////////////////////////////////////
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum = 0
	for {
		sum++
		if sum > 100 {
			break // break somelabel
		}
	}

	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

}

func DemoArray() {
	var arr1 [3]int
	var arr2 = [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}
	arr4 := [...]int{7, 8, 9, 10}

	aa1 := [2][2]int{[2]int{1, 2}, [2]int{3, 4}}
	aa2 := [2][2]int{[...]int{1, 2}, [...]int{3, 4}}
	aa3 := [2][2]int{{1, 2}, {3, 4}}

	fmt.Printf("-------------------------------------------\n")
	fmt.Printf("Run DemoAraray ...\n\n")
	fmt.Printf("var arr1 [3]int \n")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[%d]=%d ", i, arr1[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("var arr2 = [3]int{1, 2, 3} \n")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d]=%d ", i, arr2[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("arr3 := [3]int{4, 5, 6} \n")
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3[%d]=%d ", i, arr3[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("arr4 := [...]int{7, 8, 9, 10} \n")
	fmt.Printf("arr4 len is %d \n", len(arr4))
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("arr4[%d]=%d ", i, arr4[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("aa1 := [2][2]int{[2]int{1, 2}, [2]int{3, 4}} \n")
	for i := 0; i < len(aa1); i++ {
		for j := 0; j < len(aa1[i]); j++ {
			fmt.Printf("aa1[%d][%d]=%d ", i, j, aa1[i][j])
		}
	}
	fmt.Printf("\n\n")

	fmt.Printf("aa2 := [2][2]int{[...]int{1, 2}, [...]int{3, 4}} \n")
	for i := 0; i < len(aa2); i++ {
		for j := 0; j < len(aa2[i]); j++ {
			fmt.Printf("aa2[%d][%d]=%d ", i, j, aa2[i][j])
		}
	}
	fmt.Printf("\n\n")

	fmt.Printf("aa3 := [2][2]int{{1, 2}, {3, 4}} \n")
	for i := 0; i < len(aa3); i++ {
		for j := 0; j < len(aa3[i]); j++ {
			fmt.Printf("aa3[%d][%d]=%d ", i, j, aa3[i][j])
		}
	}
	fmt.Printf("\n\n")
}

func DemoSlice() {

}

func DemoMap() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // last , can not delete
	}

	fmt.Printf("\n-------------------------------------------\n")
	fmt.Printf("Run DemoMap ...\n\n")
	fmt.Printf("Define a map format: \nmap[<from type>]<to type> \nmake(map[string]int) \nmake(map[string] int, init_size ) \n\n")

	fmt.Printf("add a month: monthdays[\"Undecim\"] = 30 \n")
	monthdays["Undecim"] = 30

	fmt.Printf("update month days: monthdays[\"Feb\"] = 29  \n")
	monthdays["Feb"] = 29

	v, ok := monthdays["Jan"]
	fmt.Printf("v, ok := monthdays[\"Jan\"]\nv=%d, ok=%t\n", v, ok)

	delete(monthdays, "Mar")
	fmt.Printf("delete(monthdays, \"Mar\")\n")
	v, ok = monthdays["Mar"]
	fmt.Printf("v, ok := monthdays[\"Mar\"]\nv=%d, ok=%t\n", v, ok)

}
