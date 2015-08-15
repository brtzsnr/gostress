// gostress -seed 2986720012 -want 97330151817216
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 0, ); got != 97330151817216 {
fmt.Printf("f1_ssa(2, 0, ) = %v, wanted 97330151817216\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint8, ) int64 {
switch {} // prevent inlining
a2 += a2 ^ a2 ^ a2 | a2 * a2 - a2 * a2 + a2 + a2 >> 0 - a2 << (1 + 3) - a2 * (a2 + a2 ^ a2) >> (1 * 0) & a2 ^ a2 // uint8
a1 = (a1 | a1 * a1 ^ a1 & a1 * a1 >> 2 ^ (a1 << (1 * 3) ^ a1) << (uint32(0) >> (0 & 2) & 0)) << (3 + uint(3 + 1) >> (3 | 2) ^ uint(0 | 2) << (1 * 0) * 2 | 3 + 0) // int64
a1 *= a1 & a1 ^ ((a1 ^ a1) << (uint16(1 | 3) << (2 - 0))) << ((uint16(2) << (3 ^ 1)) << (2 ^ 3 ^ 3)) | a1 << (uint64(2) >> 2 | uint64(0) >> a2 & 3 - 2) * a1 // int64
a1--
a2 += a2 | a2 << 0 - a2 - a2 << 1 & a2 & a2 ^ a2 >> 1 & a2 & a2 * a2 >> 1 * a2 ^ a2 * a2 - (a2 * a2) >> uint16(1) - (a2 >> 1) >> ((a2 - 2) & a2 - a2 * a2) // uint8
a1 -= ((a1 + a1) >> (a2 >> a2) + a1 + a1 ^ a1 * a1) | a1 // int64
v1 := 2 // int
_ = v1
a1 = a1 << (3 * (3 & 1) + 3 ^ 1) | a1 | a1 | a1 << a2 ^ (a1 ^ a1) << 3 - a1 & (a1 ^ a1 & a1 & a1 >> 3) // int64
v1 = (v1 + v1 & v1 ^ v1 >> 2 + v1 - v1 * v1 - v1) // int
v1 += (((v1 - v1) << (uint(1) << 3) + v1 ^ v1 - v1) << (((uint(2) << 1) >> 1) << a2) + ((v1 + 3) << a2 ^ v1 >> 0 | v1)) // int
return a1 ^ a1 - a1 ^ a1 * a1 - 3 & a1 & a1 - a1 | a1 - a1 << (0 + 2 ^ 2 & uint32(0) << 1) * a1 * 0 ^ a1 - a1 ^ a1 ^ (a1 + a1 & a1 >> a2) >> 0
}
