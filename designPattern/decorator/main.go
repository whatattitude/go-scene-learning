package decorator

import "fmt"
func main()  {
    var beverage Beverage
    // 买了一杯咖啡
    beverage = Coffee{description:"houseBlend"}
    // 给咖啡加上 Mocha
    beverage = Mocha{beverage:beverage, description:"Mocha"}
    // 给咖啡加上 Whip
    beverage = Whip{beverage:beverage, description:"whip"}
    // 最后计算 Coffee 的价格
    fmt.Println(beverage.getDescription(), ", cost is ", beverage.cost())
}