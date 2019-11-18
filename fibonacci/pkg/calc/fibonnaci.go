package calc

// Fibonacci structure for holding fibonacci features.
type Fibonacci struct{}

// Calculate method on structure for calculating results.
func (f Fibonacci) Calculate(steps int) (result int) {
	if steps == 1 || steps == 2 {
		result = 1
	} else if steps > 2 {
		result = f.Calculate(steps-1) + f.Calculate(steps-2)
	}
	return result
}
