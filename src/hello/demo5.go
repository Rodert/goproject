package main

import (
	"fmt"
)

// type student struct {
// 	name string
// 	age  int
// }

// func main() {
// 	m := make(map[string]*student)
// 	stus := []student{
// 		{name: "pprof.cn", age: 18},
// 		{name: "测试", age: 23},
// 		{name: "博客", age: 28},
// 	}

// 	for _, stu := range stus {
// 		m[stu.name] = &stu
// 	}
// 	for k, v := range m {
// 		fmt.Println(k, "=>", v.name)
// 	}
// }

// ---

//Student 学生
// type Student struct {
// 	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
// 	Gender string //json序列化是默认使用字段名作为key
// 	name   string //私有不能被json包访问
// }

// func main() {
// 	s1 := Student{
// 		ID:     1,
// 		Gender: "女",
// 		name:   "pprof",
// 	}
// 	data, err := json.Marshal(s1)
// 	if err != nil {
// 		fmt.Println("json marshal failed!")
// 		return
// 	}
// 	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}
// }

// ---

//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
	age    int
	City   string
}

// func main() {
// 	s1 := Student{
// 		ID:     1,
// 		Gender: "女",
// 		name:   "pprof",
// 		age:    18,
// 		City:   "China",
// 	}
// 	data, err := json.Marshal(s1)
// 	if err != nil {
// 		fmt.Println("json marshal failed!")
// 		return
// 	}
// 	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}

// 	var a int = 10

// 	if n := "abc"; a < 20 {
// 		fmt.Println("a 小于 20")
// 		fmt.Println(n[1])
// 	}

// }

func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
