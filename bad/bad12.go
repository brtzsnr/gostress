// gostress -seed 24791
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(2, false, ))
}
func f1_ssa(a1 uint8, a2 bool, ) uint8 {
switch {} // prevent inlining
a1 -= a1 // uint8
if a2 && (((2 - 1) == (0 - 1)) || ((a1 != a1) && a2)) {
a2 = a2 && (3 != (2 | 2)) && a2 // bool
a2 = a2 // bool
v1 := 2 & (2 ^ 3) // int
_ = v1
v1++
v1 += (v1 + v1 >> 3) - (v1 << (1 | 0)) | v1 // int
}
if a2 {
v2 := uint64((1 | 0 + (3 * 1) * (0 ^ (0 + 1))) | (2 | 1 & (3 - 2) * 3)) // uint64
_ = v2
v3 := uint(3) // uint
_ = v3
a1--
a1++
v4 := v2 // uint64
_ = v4
v5 := v4 >> ((uint32(3 & 2) >> (3 * 0)) + (0 + 1 * (0 & 0))) // uint64
_ = v5
v6 := v3 - (v3 >> (a1 >> 0)) << 3 // uint
_ = v6
}
a1 *= a1 ^ a1 - (a1 | a1 | (uint8(2) << 3)) ^ (a1 << 3 & (a1 | a1) & a1) // uint8
if (2 == (int64(1 - 1) << 0)) && (1 == ((1 & 1) ^ (1 * 3))) {
v7 := a1 + ((a1 >> 0) | (a1 * a1)) - a1 // uint8
_ = v7
v8 := (uint(3 + 3) << (0 ^ 2)) | (uint(3) >> (2 | 3)) ^ (uint(3) >> (uint64(3) << 1 * (3 ^ 0))) // uint
_ = v8
v8 -= v8 // uint
v7--
v7--
v8--
v7 *= v7 + v7 + (v7 >> ((uint16(2) >> 1) + (3 + 1))) // uint8
a2 = a2 && a2 // bool
v9 := a1 - (a1 ^ a1 << 3) - a1 // uint8
_ = v9
}
return a1
}
