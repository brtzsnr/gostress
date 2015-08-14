// gostress -seed 305707100 -want 1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, 0, 2, 7, 6, ); got != 1 {
fmt.Printf("f1_ssa(5, 0, 2, 7, 6, ) = %v, wanted 1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint64, a3 int64, a4 int, a5 int64, ) int {
switch {} // prevent inlining
a3 += (a3 + a3 - (a5 * a3) << (a2 - a2) ^ a5 >> 1 * 1 ^ a3 | a3) >> (a1 + (a1 | a1) << (a2 - 1) & a1) // int64
a2 -= ((a2 >> 3) << 2) >> (3 ^ 0 << 0) + (a2 * a2 + a2) >> (1 + 1 - 0 + 1) - a2 // uint64
a4 *= a4 ^ a4 | a4 * a4 ^ a4 * 0 + a4 - a4 * a4 - a4 - a4 // int
a1 -= a1 + a1 - a1 | a1 * a1 ^ a1 & (a1 >> 0) & uint(a2 ^ a2) << (0 >> (a1 & a1 | a1)) // uint
a4++
a5 = a5 ^ a3 << a2 ^ a5 ^ 1 * a3 & a5 & a3 | a3 - a3 ^ a3 << a1 * a5 >> 3 - ((int64(3) << a1) >> (0 ^ 1 ^ 1)) // int64
a3 *= a5 // int64
a4 -= (a4 & a4 >> a1) >> 0 | a4 & a4 >> 0 + (a4 & a4) << (a1 & a1) * a4 + (a4 & a4) // int
a5++
return a4
}
