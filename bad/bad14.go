// gostress -seed 31221
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(false, false, ))
}
func f1_ssa(a1 bool, a2 bool, ) uint64 {
switch {} // prevent inlining
if (a1 == (false || (3 != 0))) || a1 {
v1 := (uint64(2) << (uint(0) >> 0 >> (0 + 0))) ^ 3 // uint64
_ = v1
v1 *= v1 | v1 + (v1 | v1) - ((v1 << 3) | uint64(2)) ^ v1 // uint64
v1 *= v1 + v1 // uint64
v1 += v1 // uint64
}
v2 := uint(2) // uint
_ = v2
v3 := uint64(3) // uint64
_ = v3
v4 := int64(3 - 1) >> (2 + (0 | 0) & (0 + 3)) // int64
_ = v4
v2--
return v3 + v3 + (v3 + v3) ^ v3 << (2 | ((0 * 2) + (2 - 2)))
}
