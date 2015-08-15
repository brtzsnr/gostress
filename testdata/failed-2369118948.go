// gostress -seed 2369118948 -want 256
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 8, ); got != 256 {
fmt.Printf("f1_ssa(7, 8, ) = %v, wanted 256\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, ) int64 {
switch {} // prevent inlining
a2 = a2 // uint8
v1 := (a2 >> 2 - a2 | a2 >> a2 ^ 2 ^ a2 & a2) + a2 * a2 - (a2 + a2) << (0 ^ 3 | 2 & 3) + a2 << 3 & 3 + a2 // uint8
_ = v1
v1 *= v1 & ((v1 | v1 * v1 >> 3) << (uint32(0) >> 1 | 2 & 1) ^ uint8(3) >> 3) << (3 + 3 ^ 3 ^ 3 + 2 * 0 - 0 & 2) // uint8
a1++
a1 += a1 * a1 ^ (a1 << 2) << (3 - 2 + 3) - (a1 & a1) * a1 // int8
a1 -= a1 + a1 ^ a1 | a1 << 0 & 0 & int8(1) << (v1 | v1 - v1 - v1 * v1) + a1 ^ (a1 >> 3 + a1 ^ a1 + int8(3) >> (2 ^ 3)) // int8
v1 += (v1 << (v1 ^ v1) | v1 & v1 << v1 - v1 ^ v1) // uint8
v2 := ((2 * 2 ^ uint64(0) << v1) >> 1 + 2 - 2 & (uint64(1) << 2)) >> a2 ^ (uint64(2) << 2 | 2 & 3 - 1 + 3 & 2 - 2 * 2 | 2 * 3 ^ 0 + 2) << 0 // uint64
_ = v2
v2 = v2 // uint64
return (((int64(1) >> a2) - int64(0) >> 2 | 1 | 0) << a2) >> 0 ^ (0 ^ 0)
}
