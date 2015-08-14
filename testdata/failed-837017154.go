// gostress -seed 837017154 -want 7
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 7, 5, ); got != 7 {
fmt.Printf("f1_ssa(4, 7, 5, ) = %v, wanted 7\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint64, a3 int64, ) int {
switch {} // prevent inlining
a3 += a1 // int64
a2 -= (a2 >> 2 & a2 - a2) + a2 | a2 // uint64
v1 := uint8(3) // uint8
_ = v1
v1++
a2 += (3 & a2 * (a2 >> v1) << (3 + 0)) // uint64
v2 := int8(0) << a2 + 0 ^ 0 & 0 * 1 + int8(2 * 3 & 2 ^ 0 & 0) >> a2 // int8
_ = v2
v3 := a1 // int64
_ = v3
a1 -= a1 // int64
v1 *= (uint8(a2 | a2) | v1 | v1 | v1 * v1 | v1 * v1) << ((v1 & v1 - 2 ^ v1) << 3 * (v1 + 1) * v1 + v1) // uint8
return int(0) >> v1 - 0 * 3 + 2 * 2 ^ 0 * 3 | int(1) << v1 ^ 2 ^ int(3) >> (3 & uint(3) << 1 - (1 | 1))
}
