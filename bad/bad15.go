// gostress -seed 32491
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(0, 1, 0, ))
}
func f1_ssa(a1 int8, a2 int64, a3 uint8, ) int {
switch {} // prevent inlining
v1 := 1 - ((3 * 1 & (uint(3) << 3)) ^ (uint(0) << (3 + 0))) // uint
_ = v1
if (int(a3 * a3) << ((a3 >> 0) | (a3 >> v1))) == (int(2) << (1 + 2)) {
a3--
v1 *= v1 // uint
if ((v1 & (v1 + v1)) == (v1 >> (uint16(0) << 1))) && true {
a2 = a2 ^ (int64(1) >> 0) - ((a2 << v1 >> (1 - 1)) + a2) // int64
v1 -= v1 + v1 >> (uint32(1 * 3) >> 3) * (v1 & (v1 >> 1) & (v1 >> (2 - 0))) // uint
v2 := v1 // uint
_ = v2
v3 := v2 // uint
_ = v3
v4 := (int(3) << 2) - (0 * 0) - 0 >> ((2 << (1 ^ 2)) - 1) // int
_ = v4
if true {
a3--
a2 -= a2 | ((a2 << v3) - a2 & ((a2 & a2) - (a2 & a2))) // int64
if uint8(int8(2) << 0 * a1) == a3 {
a2 += a2 >> (2 + 1) // int64
a3++
a3 = a3 + uint8(b2i(a1 == a1)) ^ (a3 - (a3 ^ a3 | a3)) // uint8
a2 = (a2 << v1 >> 0) | ((a2 ^ a2 << (v1 * v1)) | a2) // int64
v2 += v1 // uint
a3 += a3 >> (v1 - (v1 - 1)) & a3 // uint8
}
}
v5 := a3 + a3 | (a3 * a3) // uint8
_ = v5
a2++
a1--
a1 = (a1 >> (3 & 2)) - (a1 + a1) ^ a1 // int8
}
v1 += v1 // uint
a3 *= a3 // uint8
a3--
if (v1 & ((v1 & v1) + (v1 - v1))) == v1 {
a3 += a3 * ((a3 | a3 * (a3 + a3)) | (a3 ^ (a3 ^ a3))) // uint8
if (a2 >= 3) && ((a2 == int64(b2i(true))) && ((1 == 0) || false)) {
v1 *= v1 // uint
a2--
a1++
v6 := ((a3 * a3) - (a3 >> 1) * uint8(b2i(v1 != v1))) ^ a3 // uint8
_ = v6
a1 += a1 - (0 ^ a1 & (a1 - a1) >> (0 ^ 0 + 2)) // int8
v7 := a2 // int64
_ = v7
a1++
if false {
v8 := v1 // uint
_ = v8
a1++
v6++
a2 *= a2 - (v7 >> 3) // int64
v8--
return 3
}
}
v1--
a3 *= a3 & a3 // uint8
a2 += a2 // int64
v9 := (v1 << a3) - (v1 << v1) << ((uint32(3) >> v1) ^ 0) << 2 // uint
_ = v9
v10 := true // bool
_ = v10
a3--
}
v11 := v1 // uint
_ = v11
v12 := v11 // uint
_ = v12
v13 := a1 << 2 * (a1 - a1 - int8(v11)) * ((a1 * a1) | (a1 >> 1) | ((a1 << 1) ^ a1)) // int8
_ = v13
}
a2 += a2 // int64
a2 += (a2 & (a2 * a2) * ((a2 >> 2) + a2)) - (a2 ^ a2 >> (uint32(0) << 2) * (a2 * (a2 + a2))) // int64
a2 += (a2 - a2 << 0) + (a2 ^ a2 & (a2 + a2)) & a2 // int64
a2--
return 1
}
