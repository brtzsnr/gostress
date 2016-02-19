// gostress -seed 26591
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(1, true, 2, 1, ))
}
func f1_ssa(a1 uint64, a2 bool, a3 int, a4 uint8, ) int {
switch {} // prevent inlining
if (3 != 0) && ((a1 * a1 & (a1 | a1)) != (a1 >> (a1 << a1))) {
a1--
v1 := 3 | ((uint(3) << 1) | (3 - 0) - (uint(0) >> a1 & 0)) // uint
_ = v1
a4 *= (a4 & a4 & a4 & (a4 & a4)) ^ (a4 & (a4 >> v1) << (a4 + (a4 >> a4))) // uint8
}
if a2 || ((uint(1) << a4) != (2 + 1)) && (a2 || (a2 && a2) || (a2 || (1 == 2))) {
v2 := int64(3 ^ 3) >> (0 + 1 ^ 0 | 2) // int64
_ = v2
a4 += a4 // uint8
v3 := (uint(3) << (uint32(3) >> (1 * a4))) != ((2 & 3 & (1 + 3)) + (3 | 3 * (2 - 2))) // bool
_ = v3
v4 := (uint(2) >> (a1 | a1 + (a1 & a1))) - (3 + 0 | (0 * 0) * (uint(1) << (uint16(2) << a1))) // uint
_ = v4
a3 *= (a3 & a3) ^ (a3 >> a4 << (uint16(3) << v4)) << ((uint32(1) << 3) | (2 | 0) + (2 & 1)) // int
a1 -= a1 & (uint64(int8(3) << a1) | a1) // uint64
v2--
v2 *= v2 // int64
a3 -= (a3 << (0 + 2) >> (a4 << (uint32(2) >> 2))) ^ (a3 + a3 & (a3 & a3) * 1) // int
}
a4 = a4 // uint8
a4--
a3--
a1 *= a1 | ((a1 * a1) - (a1 << 1)) - (a1 >> (a4 - (a4 | a4))) // uint64
a1 += a1 * ((a1 + 0 << (3 + 3)) ^ (a1 * a1 << (1 ^ 2))) // uint64
a4 -= (a4 & (a4 >> a1)) + (a4 >> 1 >> 2) | (a4 << 3) // uint8
v5 := int64(1 | (0 + 1 | (1 - 1)) ^ (1 * 0)) // int64
_ = v5
v6 := a2 || a2 // bool
_ = v6
return (a3 - a3 ^ (a3 >> 0) * (a3 * a3)) | (a3 | a3 - a3 * ((a3 << 2) ^ a3))
}
