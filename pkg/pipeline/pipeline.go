package pipeline

func Generate(values ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			out <- value
		}
	}()
	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			out <- value * value
		}
	}()
	return out
}

func FilterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			if value%2 == 0 {
				out <- value
			}
		}
	}()
	return out
}
