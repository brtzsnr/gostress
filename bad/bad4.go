// gostress -seed 22363
package main
import "fmt"
var b2i = map[bool]int{true:1}
func main() {
fmt.Println(f1_ssa(0, 0, ))
}
func f1_ssa(a1 int64, a2 uint, ) uint64 {
switch {} // prevent inlining
a1 += int64(2) // int64
a1--
if (a2 * (0 ^ a2) > 1) || ((a2 & a2) - a2 >= 3) {
a2 -= (a2 << 2) ^ uint(b2i[false]) - uint(0) ^ a2 // uint
v1 := int8(1 - 0 - (1 & 2)) // int8
_ = v1
a1 *= a1 * a1 >> (uint32(3) >> 3) >> (3 * 3 * (0 ^ 1)) >> 1 // int64
v1--
}
v2 := a2 // uint
_ = v2
return 3
}
