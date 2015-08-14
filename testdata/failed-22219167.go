// gostress -seed 22219167 -want -3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 6, 1, 3, ); got != -3 {
fmt.Printf("f1_ssa(1, 6, 1, 3, ) = %v, wanted -3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint, a3 uint64, a4 int, ) int8 {
switch {} // prevent inlining
a2--
a2 *= (a2 << 3) << a3 - uint(a4 >> a3) & 2 + uint(a1) * a2 & a2 & 0 | a2 // uint
v1 := int8(3) << (a1 | (a1 << a1) & a1 + a1 - a1 << a3) // int8
_ = v1
v2 := v1 // int8
_ = v2
v2 = ((v1 ^ v1 - 2 - int8(0) >> (a1 - 1)) >> (a2 << (a2 - a2 << a2))) << (a2 << (3 ^ 0) | a2 >> a1 - a2) // int8
a3++
a2--
v3 := a4 * a4 // int
_ = v3
a1 *= a1 // uint8
a3 += a3 // uint64
return v2 & v1 >> 2 - v1 << 0 ^ v2 * v2
}
