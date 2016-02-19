// gostress -seed 10410
package main
import "fmt"
var b2i = map[bool]int{true:1}
func main() {
fmt.Println(f1_ssa(0, 1, 3, ))
}
func f1_ssa(a1 uint64, a2 uint, a3 uint, ) bool {
switch {} // prevent inlining
v1 := (a3 - a2 << (a1 | 1 * (a1 & a1))) != (a3 ^ ((a2 << 1) - (a2 | a3))) // bool
_ = v1
a3++
a1 += a1 // uint64
v2 := a1 != ((a1 * a1) | a1 << (0 & 2)) // bool
_ = v2
v3 := int64(0) // int64
_ = v3
v4 := 3 // int
_ = v4
v5 := int8(1 ^ 2) // int8
_ = v5
v5 = (v5 * v5 >> a2) ^ (v5 | v5 - v5) ^ v5 // int8
if ((3 - 2) != (1 - 2)) && ((v4 | v4) != (v4 * v4)) && ((v4 - (v4 & v4)) == v4) {
a2++
v3 *= (v3 >> (a1 * a1)) | (v3 | v3 & (v3 | v3)) & v3 // int64
v5 *= v5 >> (1 * ((0 | 1) << 3)) // int8
v4 = v4 >> (2 - 1 ^ (3 & 2) | (2 * (uint32(2) << 0))) // int
v5 += (v5 ^ v5 << (uint16(0) << a1)) | (v5 | v5 - v5) | (v5 - v5) // int8
v6 := a1 // uint64
_ = v6
v5 = v5 // int8
v7 := (v4 >> (uint32(0) >> 1 & (uint32(1) >> a2))) | ((v4 * v4) + (v4 | v4) | v4) // int
_ = v7
v8 := v3 & (v3 * v3 * v3) & v3 // int64
_ = v8
}
a2 = (a3 << (a1 + a1 | (a1 ^ a1))) - (a3 & (a2 ^ (a2 & a2))) // uint
v4--
return v1 == (v2 && v2 && v2) == ((a1 - (a1 - a1)) == ((a1 * 1) - (a1 & a1)))
}
