package day4

type BoardElement struct {
	// Number of this ``BoardElement``.
	// Since this is mostly used for equality comparisons, we'll just keep this a ``string``.
	value string
	// Whether this ``BoardElement``'s `value` was already called out.
	marked bool
}

type Board [5][5]BoardElement

type WinEvent struct {
	// The ``Board`` at the moment it won.
	board Board
	// The number which made the board win.
	lastCalledOut int
}
