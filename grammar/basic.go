package grammar

import "fmt"

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
	fmt.Printf("Define a map format: \nmap[<from type>]<to type> \nmake(map[string]int)  \n\n")

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
