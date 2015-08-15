// gostress -seed 1952624067 -want 4096
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 9, 1, ); got != 4096 {
fmt.Printf("f1_ssa(8, 9, 1, ) = %v, wanted 4096\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint8, a3 uint64, ) int {
switch {} // prevent inlining
a1 = a1 << 2 // uint
v1 := ((0 ^ 3) ^ 2 + 2 & 0) - 2 + 0 - 0 + 0 - 1 & int8(0) << a1 + 2 - 2 + 3 ^ 3 ^ int8(2) << (2 & 0) + int8(0) >> (a1 | a1 + a1) * int8(3 | 1 - 2) >> (3 | a3 ^ a3 >> a1) // int8
_ = v1
v1 -= (v1 & (v1 * v1) & v1 | v1) << a2 // int8
v2 := 1 ^ int64(0) >> a3 * 1 + int64(2) >> 1 | 0 | int64(2 ^ 0) >> (3 ^ 1) | 1 * (2 ^ int64(0) >> a2 | int64(1) >> a1) << (a2 - a2 | a2) | 2 ^ 1 | 2 | 3 | 2 & 3 & 3 & 1 // int64
_ = v2
v3 := (v2 & v2 & v2 - 1) - v2 ^ v2 & v2 | v2 - v2 ^ v2 & v2 ^ v2 - v2 >> (3 * 2 - 3 * uint32(v2)) + v2 & v2 * v2 << a1 & (v2 ^ v2) >> (1 & 0) * v2 - v2 >> (3 & 1) // int64
_ = v3
a2 += a2 // uint8
a2 -= (a2 - a2 ^ a2 << 2) << (2 | 0 + 1 ^ 2) ^ (a2 >> a3) | a2 | a2 ^ a2 & a2 - a2 + ((a2 | a2) >> (a1 >> a3 ^ a1 >> 0) & a2 | a2 ^ a2 >> a2) >> (((3 << 3) >> (a2 * a2)) >> ((2 * uint16(3) >> 1) >> a3)) // uint8
a1++
a3 *= a3 & a3 + a3 + a3 + a3 * a3 * a3 * (a3 & a3) | a3 - 3 - a3 * a3 | a3 + ((a3 & a3 + uint64(3) << a1) << (3 + 1 << a1) ^ ((a3 * a3) >> (a3 | a3))) // uint64
a1++
return (int(1) << (a3 * a3 - a3 | a3 + a3)) >> (0 & 1 + 1 | 1 & 2 * 3 ^ 3 * ((3 + 3) >> 2) >> (a3 + a3 ^ a3 & a3)) * ((int(3 - 0 ^ 2) << (3 | 1 & 2 + 1)) << (a1 >> 2)) >> ((a3 | (a3 - a3) >> a3) >> 0)
}
