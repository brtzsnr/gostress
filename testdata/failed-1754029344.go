// gostress -seed 1754029344 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(); got != 0 {
fmt.Printf("f1_ssa() = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa() int64 {
switch {} // prevent inlining
v1 := uint64(3) // uint64
_ = v1
v1 = v1 // uint64
v1 *= uint64(int(2 - 1 & 1) << (v1 + v1 * v1) | 3 - int(0) << 3 * 2) // uint64
v1 -= v1 // uint64
v1 = v1 | v1 << 1 ^ v1 + v1 & (v1 << v1) << (3 & 1) ^ v1 << (uint8(1) << 3 - 0 * 1) // uint64
v2 := (uint(3 - 2) << (1 - 1) + 2 ^ 2 ^ uint(1) >> 1) | ((3 + uint(1) >> 3) >> (uint16(3) << 3 & uint16(1) << v1)) >> (2 + v1 - v1 & v1 ^ (v1 + v1)) // uint
_ = v2
v2 = v2 // uint
return ((int64(1) | 1) << (v2 * v2 ^ v2)) << (uint8(0) * uint8(3) << 1) & (0 & 0 & 0 & 3 | int64(1 - 3) >> (v1 & v1))
}
