package c

func main() {
	go func() {}() // want "a `goroutine` statement forbidden to use."
}
