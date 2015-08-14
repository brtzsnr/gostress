// gostress -seed 152693686 -want 8
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(); got != 8 {
fmt.Printf("f1_ssa(2, 5, 8, ) = %v, wanted 8\n", got)
os.Exit(1)
}
}
func f1_ssa() uint64 {
switch {} // prevent inlining
return 1 * uint64(3 + 0) << (1 & 3) | 3 ^ 2 & 0 | 1 * 0 & 0
}
