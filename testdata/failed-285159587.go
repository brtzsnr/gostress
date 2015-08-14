// gostress -seed 285159587 -want -128
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(); got != -128 {
fmt.Printf("f1_ssa() = %v, wanted -128\n", got)
os.Exit(1)
}
}
func f1_ssa() int8 {
switch {} // prevent inlining
v1 := uint8(1 * 1) // uint8
_ = v1
v2 := 3 | (int64(3) >> 0 - 0 & 2 ^ 0 - 2 | int64(2) >> 1) // int64
_ = v2
v1++
v2 *= ((v2 ^ v2) >> 3 + (v2 & v2) << (3 - 1)) << (0 & 3 | 0 | 2) ^ v2 // int64
v3 := 2 & 0 * 1 & uint64(3) << 1 | 1 ^ 1 + 2 | uint64(2 & 2) >> (v1 << 2) | uint64(2 * 3) >> 0 // uint64
_ = v3
return (2 ^ int8(3) >> 2 - int8(3) >> v3) << v3
}
