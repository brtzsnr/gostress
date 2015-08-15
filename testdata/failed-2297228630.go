// gostress -seed 2297228630 -want 709396502639
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, ); got != 709396502639 {
fmt.Printf("f1_ssa(9, ) = %v, wanted 709396502639\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, ) int {
switch {} // prevent inlining
a1--
v1 := 2 // int
_ = v1
v1++
v1 *= ((v1 | v1 * v1) >> (1 | 0 * 2)) << (uint64((3 * 1)) << ((a1 ^ a1) << (uint8(1) << 3))) ^ v1 | v1 * v1 | v1 >> a1 | v1 & v1 << 2 - v1 >> 2 & v1 * v1 >> 0 - v1 // int
v1 = (((v1 * v1) >> (1 - 1)) << (uint32(1) >> 1 | 2) ^ v1 << (1 * 1 ^ 0 | 0) | v1) << (1 ^ 3) // int
v1 += (v1 ^ v1) - v1 << a1 | v1 ^ v1 | v1 ^ v1 + v1 * v1 & (v1 + v1) | v1 & v1 * 2 + v1 & v1 * v1 | v1 >> 3 ^ v1 - v1 ^ v1 | 3 | 3 + (v1 + v1 & v1 | 3 * v1 & v1 - v1 | v1 - v1 - v1 | v1 + v1 << 0 * v1) // int
v1 -= ((v1 + v1 + v1 | v1 + v1 & v1) << 1 + v1 * v1 & v1 - v1 & v1 ^ v1 & int(3 - 0 ^ 3 & 2)) // int
a1--
return v1 + (v1 >> 0) >> (a1 | a1) | v1 ^ v1 - v1 + ((v1 - v1) << (1 - 1) + (v1 & v1) >> (3 ^ 2)) >> (1 ^ 1 * 0 & 3 ^ 0)
}
