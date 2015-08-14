// gostress -seed 846226224 -want -3223696
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 0, 8, 8, ); got != -3223696 {
fmt.Printf("f1_ssa(7, 0, 8, 8, ) = %v, wanted -3223696\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, a3 uint8, a4 uint, ) int {
switch {} // prevent inlining
v1 := 2 + 2 * 0 | 2 | 1 | 0 - 3 + 3 + 2 | int(0) >> a4 + int(1) >> a4 ^ 0 * (int(0) >> (1 + 2)) // int
_ = v1
a1--
v1--
v1 = v1 // int
v1 = v1 // int
v1 *= (v1 - v1 | v1 - v1 ^ v1) ^ (v1 * v1 + v1 + v1) // int
v2 := a1 // int8
_ = v2
a3 -= a3 + a3 * (a2 ^ a3) >> 0 // uint8
return v1 ^ v1 - v1 * v1 ^ v1 << a4 * v1 | v1 & v1 - v1 & v1 | v1
}
