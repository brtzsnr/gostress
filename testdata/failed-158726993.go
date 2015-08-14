// gostress -seed 158726993 -want -9
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 4, 9, ); got != -9 {
fmt.Printf("f1_ssa(5, 6, 2, ) = %v, wanted -9\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint8, a3 uint, ) int64 {
switch {} // prevent inlining
v1 := int64(1) // int64
_ = v1
v2 := (a1 << (a3 >> (uint16(2) << 1)) - a1 >> a3 + a1 << 3 ^ a1 & a1 * a1 >> a3 | a1) << ((1 + uint64(1 ^ 1) << (uint64(2) << 3)) >> ((2 >> a2) >> 3 - (1 | 3) >> (2 & 1))) // int
_ = v2
v1 *= (v1 >> (a3 - a3) | v1 >> 1 & v1 | v1 ^ (v1 << a3) << (2 << a3) - v1 >> a3 | v1 * (v1 ^ v1) | v1 << 0 ^ v1 ^ v1 * v1) >> (3 | 0 * 1 * 0 * 1 ^ 0 & 0) // int64
a2 += a2 * a2 * a2 | 1 - a2 - a2 & a2 | a2 + a2 >> (uint32(1) >> a2) ^ (a2 - a2) ^ (a2 | a2 | a2 | a2 | a2 + a2) >> (1 - 1 | 2 - 0 ^ 2 + 3) ^ a2 << 2 // uint8
v2++
v1++
v2 -= ((a1 >> a2 & a1) + v2 * v2 ^ int(1) << 0 ^ (a1 << 0) << (2 + 1) ^ v2) // int
v2 += v2 << ((uint16(0 - 0 + 2 - 1) << (uint16(0 | 2) << a3)) << (2 + 1 | uint64(1) << 3 - uint64(0 + 3) << (1 & 3)) ^ (uint16(0) << (a3 ^ a3) - 1 - 3 ^ 0)) // int
v2 += int(v1 + v1) * (int(1) * a1 * a1) << (1 + 2) // int
v2 -= a1 // int
a2--
a1 *= a1 - (v2 & 0 - a1 + a1 >> 2 | a1 + v2 | a1) >> (2 * 2 - uint64(a3) - 0) // int
v3 := v1 // int64
_ = v3
a3 -= ((a3 | a3 * a3 | a3 + a3 + a3 & a3 >> a3) & a3) << (a2 >> a2 - a2) // uint
a1 -= a1 | a1 * (v2 & v2 | v2 ^ (a1 & v2) << (uint32(2) >> a2)) << ((uint64(2) >> a3 + 1 * 3) >> 3) // int
v1 = 0 + v3 + v3 ^ v1 ^ v3 ^ v1 + v3 * v3 << 1 | v1 // int64
v4 := ((v2 | v2 * 0 + a1) >> (2 + 1 | 0 + 1 + (1 | 1))) ^ a1 | int(uint64(2) << 0) * v2 ^ v2 // int
_ = v4
v2 *= ((int(uint(a1)) ^ a1 >> 0 ^ a1 & v4 - v2) >> ((0 * 0 | 0) >> (a3 + a3 * a3))) << ((a3 << 1 ^ a3 ^ a3 + a3 >> 3 ^ a3) << 0) // int
return ((v1 ^ v1) << 1 * (v3 - v3) << 1 & 0 ^ v3 - v1 - v1) * v3 << ((uint64(3 ^ 0 - 3) << ((a2 + a2) >> (a2 - a2))) >> 2)
}
