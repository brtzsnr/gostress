// gostress -seed 12650
package main
import "fmt"
var b2i = map[bool]int{true:1}
func main() {
fmt.Println(f1_ssa(1, ))
}
func f1_ssa(a1 uint64, ) bool {
switch {} // prevent inlining
a1 += a1 & (a1 << 0) << (3 | (3 * 2)) << uint64(b2i[(uint(0) >> 0) == (0 | 3)]) // uint64
v1 := 3 - ((int64(2) << 3) | (0 + 1) * (2 | (3 ^ 1))) // int64
_ = v1
a1 *= a1 // uint64
a1--
v2 := true || (0 != 1) || ((false == true) && (false || true)) && ((v1 - v1 - (v1 >> 3)) != (v1 + v1 ^ (v1 << 3))) // bool
_ = v2
return (3 + 1) == (uint8(2) >> 3 & (uint8(0) >> 0) << a1)
}
