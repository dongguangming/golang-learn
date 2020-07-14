package main

import "fmt"

const MAX int = 5

func change(val int) {  
    val = 88
}

func changeValue(val *int) {  
    *val = 66
}

func main() {
  var city = [MAX]string{ "北京", "上海","广州","深圳","濮阳"}
   var i int
   var othercity [MAX]*string;

   for  i = 0; i < MAX; i++ {
      othercity[i] = &city[i] /* 字符串地址赋值给指针数组 */
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("指针数组：索引:%d 值:%s 值的内存地址:%d\n", i, *othercity[i] , othercity[i] )
   }
   
   var age int
   var ptr *int
   var pptr **int

   age = 99

   /* 指针 ptr 地址 */
   ptr = &age
   /* 指向指针 ptr 地址 */
   pptr = &ptr

   /* 获取 pptr 的值 */
   fmt.Printf("变量 age = %d\n", age )
   fmt.Printf("指针变量 *ptr = %d，内存地址是：%d\n", *ptr, ptr )
   fmt.Printf("指向指针的指针变量 **pptr = %d，内存地址是：%d\n", *pptr, pptr)
   
    //函数调用
   change(age)
   fmt.Printf("变量age更改后的值 = %d\n", age )
   
   //函数调用
   changeValue(ptr)
   fmt.Printf("变量age更改后的值 = %d\n", age )

   //通过new函数创建指针
    size := new(int)
    fmt.Printf("size的默认值= %d, 类型是： %T, 地址是： %v\n", *size, size, size)
    *size = 99
    fmt.Println("更改后新的值是：", *size)
}