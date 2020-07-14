package main

import "fmt"

type User struct {
     name   string
     age    int
     gender string
     city   string
 }

// 方法定义
 func (u *User) describe() {
     fmt.Printf("%v 今年 %v岁了\n", u.name, u.age)
 }
 //指针参数
 func (u *User) setAge(age int) {
     u.age = age
 }
 //值参数
 func (u User) setName(name string) {
     u.name = name
 }
 
func main() {
   fmt.Println("结构体开始了")

   user := User{name: "董广明", age: 31, gender: "man", city: "南京"}

   fmt.Println(user)
   
   user.describe()
   user.setAge(99)
   fmt.Println(user.age)
   user.setName("dongguangming")
   fmt.Println(user.name)
   user.describe()


   
   //简写
   u := User{"董广明", 31,  "man",  "南京"}

   fmt.Println(u)
   
    //通过指针访问
   u1 := &User{"董广明", 31,  "man",  "南京"}

   fmt.Println(u1.name)
   
   emp := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "董",
        lastName:  "广明",
        age:       31,
        salary:    5000,
    }

    fmt.Println("员工", emp)
    
}