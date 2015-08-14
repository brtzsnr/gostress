// gostress -seed 158765451 -want 2305843009213693945
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 2, 6, ); got != 2305843009213693945 {
fmt.Printf("f1_ssa(4, 2, 6, ) = %v, wanted 2305843009213693945\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint64, a3 int64, ) uint64 {
switch {} // prevent inlining
a1 -= (a1 & a1 | a1 + a1 - 0 + uint(2)) << (0 + 1 & 2) // uint
a1 -= a1 | a1 // uint
a3 -= a3 >> 1 // int64
a3 -= 3 ^ a3 & a3 >> 0 ^ a3 >> a1 - a3 & a3 * a3 & a3 // int64
a3 += a3 & int64(1) + a3 + a3 ^ a3 ^ int64(0) >> (2 | 3 & 1 * 3 + 3 | 1 - 1) // int64
a3 = (a3 >> uint(3 - 2)) - ((a3 >> a2) << (a2 + a2 ^ a2 | a2)) >> (a1 >> a1 - a1 >> a2 - a1 & a1 * a1 & a1) // int64
a3 = a3 >> (a2 * (a2 - a2 | a2) << (a1 ^ a1 | a1)) // int64
a1 = (a1 - a1 << (0 - 0) - ((a1 >> 3) >> (a2 * a2))) // uint
a2 = a2 - a2 & a2 - a2 + a2 ^ a2 // uint64
a2 = a2 | (a2 * a2 & a2 - a2 + a2 + a2 - a2 << 2) // uint64
return a2 + a2 >> (uint16(0) << a2 - 2 & 2 & uint16(1) >> a2 ^ 3)
}
