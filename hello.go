//設置名稱，命名空間
//如果程式碼是應用程式用 main 命名
//如果是設計函式庫，就自己命名
package main

//import 是引入套件
//fmt 是用於命令列環境的格式化輸出入(formatted input/output)
import "fmt"

//主函式，應用程式的起始點或進入點
//每個 Go 應用程式只會有一個主函式
//主函式寫法是固定的
func main() {
	sayHello()
}

func sayHello() {
	fmt.Println("Hello World")
}
