// gostress -seed 29199
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(2, 0, 0, 0, 2, 2, true, ))
}
func f1_ssa(a1 uint64, a2 int64, a3 int8, a4 uint, a5 uint8, a6 int64, a7 bool, ) int64 {
switch {} // prevent inlining
if a7 {
v1 := (a1 ^ a1 + a1 & (a1 >> (2 * 1))) - a1 // uint64
_ = v1
a4 *= a4 // uint
a6 += a2 // int64
v2 := (a2 == (a6 >> 0)) || (a7 && a7 && a7) || ((v1 != (v1 - v1)) && a7) // bool
_ = v2
v2 = v1 != (v1 + v1 | (a1 - v1) | (v1 >> (0 + 1))) // bool
}
v3 := a2 // int64
_ = v3
v4 := (0 - 0 ^ (int(2) >> a1) << 2) + (int(2) >> (a1 & a1 << 3)) // int
_ = v4
v4 *= (v4 ^ v4 >> (a5 - a5) >> (uint16(2) << (a5 | a5))) + (v4 + v4 | (v4 - v4) * (v4 | (v4 << a1))) // int
a1 *= a1 & a1 // uint64
return a2 & (a2 & ((a6 * v3) + (v3 ^ 2)))
}
