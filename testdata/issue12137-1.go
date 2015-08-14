package main
import "fmt"
func main() {
fmt.Println(f1_ssa())
}
func f1_ssa() uint64 {
switch {} // prevent inlining
v1 := uint64(3) // uint64
return v1 * v1 - (v1 & v1) & v1
}
