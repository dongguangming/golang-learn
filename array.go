package main

import "fmt"

func main() {
    var city = [5]string{ "北京", "上海","广州","深圳","濮阳"}
    fmt.Println(city[0], city[1], city[2], city[3], city[4])
    fmt.Println(city)
    
    //简写版
    othercity := [5]string{ "北京", "上海","广州", "深圳","濮阳"}
    fmt.Printf("%q\n", othercity)
    
     //不同方式打印数组
    other_city := [...]string{ "北京", "上海","广州", "深圳","濮阳"}
    fmt.Println("不同方式打印数组")
    fmt.Println(other_city)
    fmt.Printf("%s\n", other_city)
    fmt.Printf("%q\n", other_city)
    
    var multi [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            multi[i][j] = i + j
        }
    }
    fmt.Printf("二维数组: ", multi)
}