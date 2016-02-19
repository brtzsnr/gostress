// gostress -seed 21636
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(3, true, 1, true, ))
}
func f1_ssa(a1 int, a2 bool, a3 uint, a4 bool, ) bool {
switch {} // prevent inlining
a3 -= a3 // uint
v1 := int8(2 & 1) // int8
_ = v1
if (a3 << 0 >> 3) != a3 {
if a2 {
a2 = a2 // bool
a1 = a1 ^ ((a1 << 0) ^ b2i(a2) & (a1 >> a3 & (a1 + a1))) // int
a1 += a1 // int
return 1 != ((2 * 2) ^ 1 & ((1 * 1) - (1 + 0)))
}
a3++
v1 *= (v1 & v1) + (v1 + v1) | ((v1 << a3) + (2 + v1)) & ((v1 - v1 * (v1 - v1)) ^ v1) // int8
v2 := a1 + (a1 ^ a1 << 2) | a1 // int
_ = v2
v1 -= v1 // int8
v1++
v3 := v1 // int8
_ = v3
v4 := int64(0 ^ 2) // int64
_ = v4
v5 := v4 // int64
_ = v5
v6 := a1 - a1 - (v2 * a1) + (a1 | v2 | a1) + (a1 ^ a1 + (v2 << 2) | (v2 | v2 + a1)) // int
_ = v6
}
if v1 == v1 {
a4 = 3 == ((uint8(1) >> a3) ^ (3 + 3) - (uint8(2) >> 3 >> (3 | 1))) // bool
v7 := a3 >> ((uint32(3) << a3) ^ 3 * (3 - 1 - (0 ^ 0))) // uint
_ = v7
v1 -= v1 - (v1 >> v7) >> 2 // int8
v1 *= v1 // int8
v7 *= v7 // uint
v1 *= (v1 >> v7 & v1) - int8(3 + 3 | (uint64(0) << 3)) // int8
v8 := uint64(3) // uint64
_ = v8
v1 += (v1 >> (0 ^ (3 - 3))) | v1 // int8
a1 = int(v8) ^ (a1 | a1) * (a1 + a1 >> (v8 << 0)) >> (uint32(3) >> (0 + 2) * 0) // int
}
if 0 != (0 - 2) {
v1++
v1 -= v1 - (v1 * (v1 | v1)) & (v1 + ((v1 * v1) ^ v1)) // int8
v9 := uint64(2) // uint64
_ = v9
v1++
v9 -= v9 >> (2 ^ 0) // uint64
a1 += a1 | a1 // int
}
v10 := int64(0) // int64
_ = v10
return a4
}
