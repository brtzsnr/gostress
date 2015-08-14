// gostress -seed 1586724092 -want 197
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 1, 9, 9, ); got != 197 {
fmt.Printf("f1_ssa(1, 1, 9, 9, ) = %v, wanted 197\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint, a3 int64, a4 uint8, ) int {
switch {} // prevent inlining
v1 := (int8(0) >> (2 & 2 + 2) - int8(3 | 1) >> (3 << 2) * 0) << (uint16(2) >> (1 + 0 ^ 3 * 0 & 1 + uint32(3) >> a1)) // int8
_ = v1
a1++
a1 = (a1 << (a2 + 3 + a2 & a2 | a2 | a2)) >> (((uint32(1) >> a4) >> (uint16(3) >> a4)) << ((a1 >> 3) >> (a2 - a2))) // uint64
v1 -= v1 >> (a4 * a4 - a4 >> 2 & a4) // int8
v1 *= v1 // int8
v1 *= (((v1 << 2) >> a4) << (3 << 0 - 2 & 3 - 1 & 2 - 1)) << (a1 | (a1 & 0) << a1 * a1) // int8
v1--
v1 *= (int8(a4) + 2 | v1) >> (a4 ^ a4 - a4 + a4) & (v1 ^ 1 & v1 ^ v1) << (3 | 2 | 0 - 2) ^ (v1 & v1 | v1) << (a4 ^ a4 * a4 & a4) & v1 - v1 & v1 & v1 * v1 // int8
a3 = a3 // int64
return int(a4 ^ 2) + (int(0) << 3) << (uint8(3) << a1) * (int(2) << 0) << (a2 + a2) | 2 | 3 | int(a1) * 3 - int(a3) + 3 | 1 + int(1) >> 2
}
