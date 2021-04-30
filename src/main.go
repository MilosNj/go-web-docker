// main.go

package main
import (
	"strconv" // strconv implementira konverzije iz i u string oblike osnovnih tipova podataka
	"github.com/astaxie/beego" // beego je Go web framework. Inspiraciju su uzeli od Flask-a
)

func main() {
	/*
	  Ovo će se podudarati sa rutama kao što su:
	  /sum/3/5
	  /product/6/23
	*/
	beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	// Preuzmi vrednosti parametara sa rute
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	// Setuj vrednosti za korišćenje u templejtu
	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.TplName = "result.html"

	// Izvrši kalkulacije u zavisnosti od 'operation' parametra u ruti
	switch operation {
	case "sum":
		c.Data["result"] = add(num1, num2)
	case "product":
		c.Data["result"] = multiply(num1, num2)
	default:
		c.TplName = "invalid-route.html"
	}
}

func add(n1, n2 int) int {
	return n1 + n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}
