package value

// Saver saves values in constant time intervals.
type Saver interface {
	Save(int)
}
