A list of programs on which dev.ssa go compiler performs poorly.
To see the generated code run `GOSSAFUNC=f1_ssa go build bad1.go`
and then check the output or the generated ssa.html file.

# Bad
* bad13.go - many redundant branches
* bad15.go - redundant comparison at b17

# Fixed
* bad0.go - extra store at v17
* bad2.go -
* bad4.go - extra store at v8
* bad6.go - missed constant folding v153
* bad7.go - redundant comparison at b10
* bad8.go - redundant comparison at b12
* bad9.go - missed comparison folding v33
* bad11.go - redundant comparison at b2
* bad12.go - redundant comparison at b2
* bad14.go - redundant comparison at b2
* bad16.go - redundant comparison at b6
* bad17.go - many redundant branches
* bad18.go - many redundant branches
