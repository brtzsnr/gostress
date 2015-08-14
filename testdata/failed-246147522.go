// gostress -seed 246147522 -want -1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 4, 4, 0, ); got != -1 {
fmt.Printf("f1_ssa(0, 4, 5, 3, ) = %v, wanted -1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 int64, a3 int8, a4 uint, ) int64 {
switch {} // prevent inlining
a1++
a1++
a3 *= a3 ^ a3 >> 3 + a3 << 0 | (a3 >> 3) << (2 | 2) + (a3 >> a4 & 3 ^ a3) >> (uint32(1) << a4 * uint32(1) >> 2 | 0) ^ a3 + a3 * a3 + a3 | a3 + a3 | a3 // int8
v1 := a4 // uint
_ = v1
v2 := int64(0 - (3 | 3 - 1) & 0) // int64
_ = v2
v3 := 3 * 2 & 0 + 1 & 2 ^ 2 * 2 ^ 0 | 0 ^ 3 ^ 2 | 3 & uint64(0) >> 0 + 1 * uint64(2) << (1 & uint64(1) << 0) * 0 * 0 & 2 & 2 & (1 | 0) // uint64
_ = v3
a3 -= a3 // int8
a4 -= v1 // uint
v3 += v3 & (uint64(a2 ^ v2) - v3 >> v1 ^ v3 & v3 + v3) >> (((v1 & a4) | a4 ^ a4 * a4 - a4) >> (v1 ^ v1 & a4 * v1 + 3 + 2)) // uint64
a1 -= (v2 * a2 ^ a1 + a1 | a1 + a2 & a1 | (v2 | a2) << 2) >> 1 // int64
a4 -= ((a4 ^ v1 | v1) ^ a4 & v1 - v1 ^ v1 * v1 - v1 + a4) << v1 & (0 | v1 | v1 & a4 >> 3) + v1 // uint
v1 *= v1 | v1 & v1 + (a4 * v1 - v1 + (v1 >> v3) << (3 ^ 2) + a4 >> 1 + v1 >> v3 * uint(a3)) // uint
a2++
v1--
v1 *= (a4 ^ a4 ^ 2 | a4 + a4 | v1 | v1) << v3 & 2 & 2 | a4 * v1 ^ v1 | v1 | a4 ^ a4 // uint
a1++
a4 -= ((a4 - 1 - v1) << (3 ^ 0 | 2) & a4 | v1 - a4) >> ((2 - (uint8(2) << v3) >> 1 + 1) << (a4 - v1 - v1 ^ v1 + v1 - v1 | v1)) // uint
v3 = v3 << (uint8(2) << (uint32(2) >> v1) + 2 | 3 | 3) ^ (v3 << (2 + 2) & v3) << ((3 + 0) - uint32(2 - 2) >> (v3 - v3)) ^ v3 * 3 + v3 & v3 * v3 & v3 | v3 // uint64
a4++
return a2 ^ v2 >> a4 | a2 + a2 * a1 ^ a2 & a2 >> 3 * a2 | v2 | 1 * v2 & a1 + v2 | int64(v3 << 2) ^ a2 | v2 & v2 & v2 | a1 & v2 | a2 ^ v2 * v2 ^ v2
}
