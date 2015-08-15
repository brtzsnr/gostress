// gostress -seed 842010851 -want 3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(); got != 3 {
fmt.Printf("f1_ssa() = %v, wanted 3\n", got)
os.Exit(1)
}
}
func f1_ssa() int64 {
switch {} // prevent inlining
v1 := (3 + 1 ^ 3 * 2 & 0 ^ 2 & uint8(0) >> 1 | 1) << (uint64(2 + 1) << (0 & 0) - 3 - 0 | 3) - 0 // uint8
_ = v1
v1 = (v1 & v1 >> 3 * v1 | (v1 * v1) << (0 ^ 2) ^ v1) // uint8
v1 += v1 // uint8
v1--
v1 = (v1 >> (0 * 3) - (v1 * v1) + (v1 - v1) + v1 & v1 & v1 ^ v1) >> (uint(1 | 1) >> 0 | 3 * 0 ^ 1 ^ uint(3) << 0 ^ 0 & 3 ^ 0 | 1 + (1 + 3 - 2 - 2 ^ 3)) // uint8
v1 = v1 // uint8
return 2 - (int64(3) >> v1) << v1 | 0 + 3 | ((2 ^ 0 & 3) * int64((3 ^ 0)) >> (2 >> 1)) >> (3 | 1 | 0 & 1 ^ 1 * 1 >> 3 + 0 >> 3 | 0)
}
