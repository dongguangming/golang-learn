package main

 import (
     "fmt"
 )

func change(name *string) {
  *name = "dongguangming"
}

func print(items ...string) {
  for _, v := range items {
    fmt.Println(v)
  }
}


 func main() {
    fmt.Println("函数开始了")
    name := "董广明"
    fmt.Println(name)

    change(&name)
    fmt.Println(name)
    
    print("董广明", "dongguangming", "dmg")
    
    list := []string{"Hello", "World"}
    print(list...) // An array argument

    func () {
  fmt.Println("我是一个匿名函数")
} ()

 }
