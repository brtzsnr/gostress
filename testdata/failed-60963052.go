// gostress -seed 60963052 -want -54043193515180032
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 5, ); got != -54043193515180032 {
fmt.Printf("f1_ssa(9, 5, ) = %v, wanted -54043193515180032\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint, ) int64 {
switch {} // prevent inlining
a1 *= ((a1 & a1) >> (a2 + a2)) << 3 + a1 // uint8
a2 *= a2 // uint
a1--
v1 := int64(2) // int64
_ = v1
v1++
v1--
v1 = v1 & v1 // int64
v1 *= v1 // int64
v2 := int64(int(0) << 0 ^ int(3) >> a2 & 1 ^ 1) + v1 - v1 << a2 ^ 1 & v1 & v1 // int64
_ = v2
a1--
return ((v2 * v1 * 3 - 0) | 0) << a2
}
