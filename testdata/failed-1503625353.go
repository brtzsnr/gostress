// gostress -seed 1503625353 -want 256
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 3, 5, 1, ); got != 256 {
fmt.Printf("f1_ssa(7, 3, 5, 1, ) = %v, wanted 256\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int, a3 int8, a4 int8, ) int64 {
switch {} // prevent inlining
a1 = a2 // int
a3 += int8(2) << 3 // int8
v1 := uint64((3 ^ 1 & 0 ^ 0 - 0 & 1 ^ 0 ^ 1)) // uint64
_ = v1
a2--
a2--
a2 = (a1 | a2 >> (2 ^ 0) & a2 << 2) >> (((1 << (2 + 2)) << 3) << (2 - v1 | v1 + v1 ^ v1 | v1 & 2)) // int
v1 *= ((v1 >> 1) >> 3 ^ v1 & v1 | v1) // uint64
v2 := uint8(3 ^ 3) << ((3 << 0 ^ 0) >> 0 + 1 | 0 + 3 * 2 * 3 - 0 ^ 0 - 3) // uint8
_ = v2
a2 *= a2 & (a2 | a2) << (1 & 3) ^ a1 * a1 + a2 - a2 << 3 & a1 + a1 + a2 // int
return ((int64(1 * 0 + 1) << (2 & uint16(3) << 2)) >> v1) << v1
}
