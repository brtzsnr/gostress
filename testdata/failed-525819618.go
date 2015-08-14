// gostress -seed 525819618 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 4); got != 0 {
fmt.Printf("f1_ssa(5, 9, 9, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a3 int, ) int {
switch {} // prevent inlining
a1 -= a1 + (a1 - a1) << (2 & 1) & a1 // uint
v2 := (a3 & a3 << a1 - (a3 + a3) >> (1 << a1)) >> 2 // int
return v2
}
