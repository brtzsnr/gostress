// gostress -seed 24577
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(0, false, 0, ))
}
func f1_ssa(a1 int8, a2 bool, a3 int8, ) bool {
switch {} // prevent inlining
return ((a1 >> 3 * a1) != a1) && (a2 || (0 == (2 + 3)))
}
