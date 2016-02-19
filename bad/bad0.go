// gostress -seed 4188
package main
import "fmt"
var b2i = map[bool]int{true:1}
func main() {
fmt.Println(f1_ssa(0, 1, 1, 3, 0, 2, 0, 1, 1, 2, ))
}
func f1_ssa(a1 int8, a2 uint8, a3 uint64, a4 int64, a5 int8, a6 uint64, a7 uint, a8 int, a9 uint64, a10 uint, ) uint8 {
switch {} // prevent inlining
a3 = ((a3 * a6) - a3 >> 1) | (a6 ^ (a6 + a3) >> a10) // uint64
v1 := a8 | int(1) // int
_ = v1
if (true || (false || false) && (a4 == (a4 + 1))) || ((a2 ^ a2 >> (a10 - a7)) != (a2 * a2 * (a2 | a2))) {
v2 := a7 // uint
_ = v2
v1 += a8 - ((v1 * v1) ^ (v1 >> a2)) << (((0 << 2) << (a6 >> a2)) | 2) // int
v3 := (true == (a2 != a2)) && (((a4 + a4) != (a4 ^ a4)) && false) // bool
_ = v3
a5--
v4 := (v2 & a10) - v2 * (a10 ^ v2) * (a10 + v2 - (v2 ^ v2) - (v2 << (0 | 0))) // uint
_ = v4
v5 := a2 & (a2 >> (1 & 0)) << 0 // uint8
_ = v5
a2 += a2 // uint8
a4 += 1 + a4 // int64
a6 *= (a3 << 1) - a6 // uint64
v6 := (v5 ^ a2 * (v5 >> v5)) ^ (a2 << a2) << (uint16(2 ^ (2 ^ 3)) >> (a2 * v5 >> (uint(1) << 1))) // uint8
_ = v6
}
a10 -= (a7 - a10 + uint(a6) << ((2 * 0) | 1)) ^ a7 // uint
v1 = v1 // int
a4++
v1++
a7 += a7 // uint
v1 *= (v1 << a2) - (v1 ^ v1) | v1 + (v1 << (a3 ^ (a9 | a3))) // int
return a2
}
