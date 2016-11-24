# Generic linear algebra poc

This is a proof-of-concept based on the gonum/matrix project, 
using a tiny portion of the code to play with making it support
all native go number types.

The overall idea is: Take the existing codebase, mark float64 values 
by typedeffing them as `TYPE`, and then use regular ast rewrite tools
to generate concrete versions of the code base for each type.

The end result is:

- Runnable & testable "template" code
- Checked-in, magic-hidden-from-end-users matrix code for all types

## Interesting files

    // Slightly modified version of gonum matrix.go, marking float64 as TYPE
    base/matrix/matrix.go
    
    // Code rewriter
    gen/main.go
    
    // Generated output
    numXX/matrixXX/gen_matrix.go
    
## Run the code generation

    ./generate.sh
    

## License

Gonum code as follows: https://github.com/gonum/license
