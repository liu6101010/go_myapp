package libs

import "fmt"

var lib_param_a int = 999
var lib_param_b string = "paramb"

//public
func Show() {
	fmt.Println("######## lib1 called")
}

func GetParams() (a int, b string) {
	a = lib_param_a
	b = lib_param_b
	return
}

func init() {
	fmt.Println("######## lib1 used")
}
