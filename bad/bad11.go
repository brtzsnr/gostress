// gostress -seed 298
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(1, 1, 3, 2, 1, true, 3, 1, ))
}
func f1_ssa(a1 uint8, a2 int64, a3 int, a4 int64, a5 int8, a6 bool, a7 int, a8 uint, ) uint {
switch {} // prevent inlining
a1 += (a1 * a1 >> (3 | 2)) - (3 * a1 * (a1 | a1)) + (uint8(a4 * a2) << 0) // uint8
a4++
v1 := ((a8 >> 0) != (a8 << 2)) || a6 && a6 // bool
_ = v1
a2 -= a2 // int64
a1 += a1 // uint8
a1 -= ((a1 & a1) - (a1 - a1) >> (a1 | a1 << (uint16(1) << 2))) | a1 // uint8
a3 += (a3 & a3 * int(a8)) | a3 * 0 // int
a8++
a7 = a3 | (b2i(v1) | (a3 | a3) & a7) // int
v2 := (a4 & ((a2 >> 3) | (a4 >> 1))) ^ a2 // int64
_ = v2
return a8 * (a8 ^ a8) >> (2 & ((3 & 3) << a8))
}
