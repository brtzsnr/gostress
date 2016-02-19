// gostress -seed 26194
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(1, false, 0, ))
}
func f1_ssa(a1 int8, a2 bool, a3 uint64, ) int {
switch {} // prevent inlining
v1 := ((a3 == 0) && (a2 || a2)) || a2 && (a2 || a2) // bool
_ = v1
a1 *= a1 // int8
a1 -= (3 ^ a1 - (a1 & a1) * a1) + (a1 * (a1 + (a1 * a1))) // int8
if v1 {
v2 := uint(2) // uint
_ = v2
if (((int64(1) << a3) != (int64(1) >> 1)) && (a2 == (v1 && a2))) || a2 {
v2--
v2 = v2 | uint((a3 << 3) - (a3 + a3)) // uint
a1 -= a1 // int8
v2++
if v1 {
if (v2 == v2) && (((1 != 3) && (1 == 3)) || (a1 != (a1 * a1))) {
a1 += a1 + (a1 >> (3 * 2) << (a3 - a3 << (a3 | a3))) // int8
a3 = a3 << (0 + 1 & (1 | 0) & (3 - (uint16(3) >> v2))) // uint64
v2 += v2 | v2 // uint
a1 *= a1 // int8
}
v3 := a3 // uint64
_ = v3
v2--
if (a2 != a2) && a2 && ((0 + 0 - (uint8(1) << a3)) == 0) {
a1 += a1 // int8
v1 = (1 + (3 | (0 * 0))) != ((int(0 | 2) >> (v2 >> 0)) ^ ((2 & 2) + (3 + 2))) // bool
if ((a3 << 1 << (v2 + v2)) == ((v3 >> v2) | (v3 >> 3))) && (v3 == (a3 * (v3 << v3))) {
if ((a1 * a1) | (a1 & a1) ^ (a1 << 1 >> (2 & 1))) != (a1 << 2 << (1 & 2) * (a1 ^ a1 << (v2 ^ v2))) {
if ((a1 + a1 * a1) == (a1 - a1)) && a2 {
v3 = uint64(3) // uint64
a3 -= v3 // uint64
a1++
if (2 != (int64(1) >> v2) != (v1 && a2 && v1)) || (((uint8(0) >> 2) != (uint8(1) << 1)) || a2) {
v4 := (uint(1) >> 1) + v2 >> 1 * v2 // uint
_ = v4
v5 := v3 * (a3 | a3 | v3 << (3 << a3 * 0)) // uint64
_ = v5
a1 -= int8((int64(1) << v5) | (3 * 2) << ((v3 & a3) ^ (v3 >> v2))) // int8
v3 -= a3 ^ v3 ^ (v5 + (a3 - v3)) >> (uint32(1 * 0) << (v3 ^ v5) * 2) // uint64
v5 = a3 // uint64
}
a1 *= a1 // int8
v2 -= (v2 & v2 & (v2 - v2 - v2)) + v2 // uint
if a2 {
a2 = v1 // bool
v2 += (v2 * v2 << (a3 + v3) & (v2 & (v2 >> v2))) | ((v2 * v2) - v2 + uint(a1 + a1)) // uint
v6 := 3 // int
_ = v6
a3 = a3 - a3 << ((2 * (3 - 1)) - ((uint16(3) >> v2) + (3 + 0))) // uint64
v6++
return v6
}
}
v7 := 2 | ((2 & 1) + (1 * 2)) | (1 + (int(0) << 2) | (3 & (int(0) << a3))) // int
_ = v7
a3--
a3 = a3 ^ (v3 ^ a3 - v3) | a3 // uint64
}
}
}
v8 := a3 << (1 + 1) << 3 * (v3 * ((a3 << 3) - (v3 & v3))) // uint64
_ = v8
}
a1++
}
a3 *= a3 | (a3 + (a3 * a3)) + (a3 - a3 | (a3 & a3) << a3) // uint64
}
return (0 - 1 + (3 - 2) * (int(3) << (uint8(1) << 2))) | ((int(2) << (1 + 0)) ^ 0)
}
