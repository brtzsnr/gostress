// gostress -seed 19271
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(false, false))
}
func f1_ssa(a4 bool, a6 bool) bool {
switch {} // prevent inlining
return a6 || (a6 || (a6 || a4)) || (a6 || (a4 || a6 || (false || a6)))
}
