package _48

import (
	"errors"
	"fmt"
)

func panickingBad(a, b int) {
	// Why is this potentially bad?
	if a+b > 8 {
		panic("This is a panic")
	}
}

/* For a few reasons
- We have to have separate logic to recover via recover(), different from typical error handling
- Panics generally have to be converted to an error, or we risk having the panic bubble up
- Unline an error, there are no "typed" panics. A panic is a panic
*/

var tooLargeError = errors.New("This is an error from a large value")

func notPanicking(a, b int) error {
	if a+b > 8 {
		return tooLargeError
	}
	return nil
}

func handleNotPanicking() {
	err := notPanicking(8, 8)
	if errors.Is(err, tooLargeError) {
		// Much better than panicking
		fmt.Println("Handle the error")
	} else {
		fmt.Println("Handle something else")
	}
}

func loadDependency() {
	panic("Could not load dependency")
}

func goodPanic() {
	//In general, there are two situations in which we should panic:
	// - We fail to load a dependency which our application relies on (for example, loading Authz)
	// - There's a programmer error that is not application/business specific, including
	// a nil pointer dereference, re-registering a dependency, etc

	// Defer a function to recover from any panics in main or functions it calls
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) // r contains the value passed to panic
		}
	}()

	loadDependency()
}

// A good thing about panics, stack traces (order of operations from bottom to top, with the top frame being where
// the error occurs)
/*

panic: runtime error: index out of range [0] with length 0

goroutine 62893 [running]:
github.com/justworkshr/cdms/internal/icp.parseAddresses({0xc003b8b880, 0x1, 0xc003c3c960?}, {0xc003a06e70, 0x1, 0x0?})
    /build/internal/icp/addresses.go:133 +0x448
github.com/justworkshr/cdms/internal/icp.(*Repo).FindAllAddresses(0xc000ca9bd0, {0x24e61b0, 0xc003c3c960}, {0xc00317a340?, 0x1000000?, 0xc003860f38?})
    /build/internal/icp/addresses.go:95 +0x988
github.com/justworkshr/cdms/internal/gateway.Service.BatchGetAddresses.func2()
    /build/internal/gateway/addresses.go:304 +0x6b
golang.org/x/sync/errgroup.(*Group).Go.func1()
    /go/pkg/mod/golang.org/x/sync@v0.19.0/errgroup/errgroup.go:93 +0x50
created by golang.org/x/sync/errgroup.(*Group).Go in goroutine 7243
    /go/pkg/mod/golang.org/x/sync@v0.19.0/errgroup/errgroup.go:78 +0x95

*/
