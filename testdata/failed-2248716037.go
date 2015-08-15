// gostress -seed 2248716037 -want 45035996273704960
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 3, ); got != 45035996273704960 {
fmt.Printf("f1_ssa(4, 3, ) = %v, wanted 45035996273704960\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint64, ) uint64 {
switch {} // prevent inlining
v1 := (uint8(3) << (3 | 3) & uint8(0) >> (0 | 3)) >> (1 + 2 - 0 ^ 3 * 2 ^ 3) - 0 | uint8(3) >> 3 + (2 & 3) & uint8(1) << a2 ^ (uint8(1) >> 3) * uint8(0) << 2 & uint8(2) >> 2 - 1 - 0 * 2 + 0 // uint8
_ = v1
a1--
v1++
a1 *= int64(0 ^ int(0) >> (uint16(0 + 2) >> (uint16(3) >> 2) * 1 + 0 + 0 ^ 3)) // int64
a1 += a1 ^ (a1 | a1 ^ a1) << (a2 ^ a2 + a2) - a1 >> 2 & a1 & (a1 >> 0) >> (v1 * v1) | (a1 ^ a1 + a1 | a1 + (a1 * a1) << (v1 - v1)) * int64(a2 << a2 ^ a2 - a2) // int64
v1 = v1 << 1 & v1 + (v1 * v1 | v1 * v1) & v1 << (2 ^ 3 & 2 * 3 - 1) | v1 // uint8
v2 := a1 & ((a1 << 0) >> (a2 & a2)) & ((a1 - a1) >> (v1 * v1) + a1 * a1 ^ int64(0) << 0) << 3 // int64
_ = v2
v2 = v2 - (a1 - a1 >> 1 - (a1 - v2)) >> (0 + 2 & 0 | (3 + 3) >> (1 | 2)) ^ (((v2 << 1) >> v1 * a1) >> ((a2 + a2) - a2 * a2 ^ a2 & a2)) // int64
v2 += int64((int8(1) >> (uint8(2 | 1) * v1 + v1 ^ v1 ^ v1)) >> (2 >> (v1 & v1) & 1 >> a2 - 0 | (1 ^ 1 * 3))) // int64
v2--
return ((a2 & a2 >> 2 * a2 | a2 & a2 + a2) << (v1 & v1 + v1 ^ (v1 >> 2) >> uint32(3)) + a2 ^ a2 | a2 + a2 * a2 & a2 * a2 + a2 * a2 - a2) << (a2 & (1 | a2) << (v1 | v1) | a2 & a2 << (uint(0 + 1) >> v1 + uint(3) >> (2 ^ 2)))
}
