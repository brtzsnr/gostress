// gostress -seed 265812795 -want -1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, 5, ); got != -1 {
fmt.Printf("f1_ssa(5, 5, ) = %v, wanted -1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint64, ) int64 {
switch {} // prevent inlining
v1 := (int8(a1 ^ a1) | int8(3 | 0) >> 1 | 0 | 2 - 3 - 3 & int8(1) >> a2) << (0 | 0 & 3 ^ 2 * uint16(a1 + a1) & uint16(a2) >> a1 & 2 & 0 ^ 0) // int8
_ = v1
a2 = a2 // uint64
v2 := (3 * 0 | 3 * 1) + (uint8(1) >> a2) >> (2 & 3) ^ uint8(3) >> (uint8(2) >> 3) ^ 3 // uint8
_ = v2
a2--
a1 *= a1 // uint
v3 := ((v1 + v1) >> a2 ^ v1 >> a1 + v1 + v1 >> (1 ^ 3) * v1 - v1 & v1) // int8
_ = v3
v4 := v1 ^ v1 | v3 & v3 * v1 * v1 + v3 - v3 * v3 ^ v1 | v1 // int8
_ = v4
a1 -= a1 - 3 + a1 * 0 | a1 - a1 & a1 & a1 & a1 ^ a1 & ((a1 >> v2) & uint(v2 + v2)) // uint
a2 += a2 & a2 | a2 - a2 - a2 | uint64(1 ^ 2) ^ a2 + a2 | a2 // uint64
return (0 | 1 - 0 | int64(1) << 1 + 2 * 2 - int64(3) << v2) >> (v2 + v2 | v2 - v2 * v2 | uint8(1) * v2 << a2)
}
