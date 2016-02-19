// gostress -seed 29723
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(0, 3, 0, 2, 2, 3, ))
}
func f1_ssa(a1 int, a2 int8, a3 int64, a4 uint, a5 uint64, a6 uint8, ) bool {
switch {} // prevent inlining
if true {
a3 = (a3 << (a6 + a6 - (a6 & a6))) | (a3 - (a3 | 1) * (a3 ^ a3 << (3 * 3))) // int64
a1 *= a1 - (a1 ^ (a1 >> 2)) >> (((0 << a5) >> a4) - (0 * (3 & 3))) // int
a1 += a1 * a1 * (a1 + a1) >> (a6 << a6) << a4 // int
a1 -= 3 // int
a4--
}
a4 += a4 ^ (a4 ^ (a4 & a4)) | ((a4 << a6 >> 1) + (a4 - (a4 | a4))) // uint
if (a6 != (a6 ^ a6 & (a6 & a6))) || ((a4 == a4) && false) {
v1 := a2 | a2 // int8
_ = v1
v1--
v2 := a5 // uint64
_ = v2
a4 = (a4 << 3 & (a4 + a4)) | (2 & (a4 ^ a4)) << (a4 - (0 - a4 >> (a4 << v2))) // uint
v3 := a4 // uint
_ = v3
a6 -= (a6 >> v3 << (v3 ^ v3) * (a6 + a6 + uint8(a3))) ^ (a6 & a6) // uint8
}
a1 += a1 // int
a3 = (a3 << 2) ^ 1 | a3 - (a3 * 2 & a3 << (a6 + a6 >> (2 | 2))) // int64
a4 += (a4 >> 1 * uint(a5)) - 1 & a4 // uint
a1 -= a1 << a5 >> a5 << a4 >> (0 + 1 | (1 & 1) ^ 3) // int
v4 := a6 & (a6 | (a6 | a6)) & (a6 + a6 + a6 * 1) // uint8
_ = v4
v5 := true // bool
_ = v5
return (a2 != (a2 ^ a2)) || (v5 || v5 && (a5 != a5)) || (v5 || v5 || v5 && v5)
}
