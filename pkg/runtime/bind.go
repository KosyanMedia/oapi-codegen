package runtime

// Binder is the interface implemented by types that can be bound to a query string or a parameter string
// The input can be assumed to be a valid string.  If you define a Bind method you are responsible for all
// data being completely bound to the type.
//
// By convention, to approximate the behavior of Bind functions themselves,
// Binder implements Bind("") as a no-op.
type Binder interface {
	Bind(src string) error
}
