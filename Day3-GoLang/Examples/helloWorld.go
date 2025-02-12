package main

import "fmt"

type student struct {
	Name string
	Rgno float64
	Dept string
}

func main() {
	// message := "hello world!"
	// fmt.Printf("%s", message)
	st := student{Name: "Student1", Rgno: 12.2, Dept: "CS"}
	fmt.Println("Name: ", st.Name, "\nRegestration No: ", st.Rgno, "\nDepartment: ", st.Dept)
}

/*
func ifelseDemo() {
	var n1, n2 int
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	if n1 > n2 {
		fmt.Println("n1 is larger than n2")
	} else if n1 == n2 {
		fmt.Println("n1 is equal to n2")
	} else {
		fmt.Println("n1 is smaller to n2")
	}
}
func forThreeVarDemo() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)
}
func forCondiDemo() {
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}

func infiniteFor() {
	sum := 0
	for {
		sum++ // repeated forever
	}
	//fmt.Println(sum)
}

func forPythonStyle() {
	strings := []string{"hello", "world", "GoLang", "NIE"}
	for _, s := range strings {
		fmt.Println(s)
	}
}
*/
