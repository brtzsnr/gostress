// gostress -seed 3095522295 -want 381
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(); got != 381 {
fmt.Printf("f1_ssa() = %v, wanted 381\n", got)
os.Exit(1)
}
}
func f1_ssa() uint {
switch {} // prevent inlining
v1 := (1 ^ (1 + 0) ^ (uint(1) << 1) << (0 ^ 2) * 0 ^ 3 | 0 + 1 - 0 - ((uint(2) >> 2) << (1 | 1)) << (uint8(2) >> 3 + 1 & 2)) // uint
_ = v1
v1 += (v1 >> 0 - uint(2) << 3 | v1 + v1 & v1 - v1 | (3 ^ v1 ^ v1 << v1) >> (uint32(2 - 2) >> v1) | v1) << 0 // uint
v2 := (2 ^ 3 | uint64(1 ^ 3) << (uint8(3) >> v1)) << (2 >> (v1 * v1) ^ (1 | 1) - (2 - 3) >> (2 * 0) & 3 >> 0) & uint64(0 ^ 2 * 0) << (((1 << 1) << (uint32(3) >> 0)) >> (1 | 0)) ^ uint64(v1 ^ v1 | v1) - 3 & 3 * 2 - 3 & 2 // uint64
_ = v2
v3 := (v1 * v1 - v1 - 3 + v1) >> (v2 | v2 & v2 ^ 3 - uint64(v1) >> (uint64(3) << 0)) - uint(2) * v1 // uint
_ = v3
v1 *= uint(2 - 2 ^ 2 * 2) | (v3 ^ v1 & v1) >> (0 << (v2 >> 2)) & v1 ^ v1 & v3 - v1 * v3 ^ uint(1) - v1 | ((v3 * v3) >> 0) >> 1 * v1 + v3 // uint
v1 = ((v1 >> 1) << (1 + 2) ^ v1 | v1 & v3 + (v3 - v1)) >> (v3 >> (1 | (uint16(3) >> v3) << (v1 << 1) * 1)) // uint
v3 = v1 // uint
return (v3 - v3) >> 3 | v1 | v3 + (v3 >> v2) ^ v1 ^ v3 + v3 - v1 & (v3 | v3) + v3 + v1 >> 3 ^ v1 + v1 + v3 ^ v1 - v1 ^ v3 | (v3 | v1) << (v2 << v2) - (v1 - v1) >> 0
}
