package main

import "fmt"

func main() {
    var city = []string{ "北京", "上海","广州","深圳","濮阳","郑洲"}
    
    fmt.Println(city)
	fmt.Println(city[1:4])

	// missing low index implies 0
	fmt.Println(city[:3])
	// [2 3 5]

	// missing high index implies len(s)
	fmt.Println(city[4:])
	
	fmt.Println("准备用make方式创建切片")
	cities := make([]string, 3)
	cities[0] = "郑洲"
	cities[1] = "濮阳"
	cities[2] = "安阳"
	fmt.Printf("%q\n", cities)
	
	//附加多值到切片
	other_cities := []string{"南京"}
	
	other_cities = append(other_cities,  cities...)
	
	fmt.Println(other_cities)
	
	//拷贝切片
    fmt.Println("准备用copy方式创建切片")
    copycities := make([]string, len(other_cities))
    copy(copycities, other_cities)
    fmt.Println("copycities:", copycities)
    
	
	fmt.Println(len(cities))
	countries := make([]string, 42)
	fmt.Println(len(countries))
	
	//零片
	var temp []int
    fmt.Println(temp, len(temp), cap(temp))
    if temp == nil {
        fmt.Println("nil!")
    }

}