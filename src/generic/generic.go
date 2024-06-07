package main

func main() {

}

func sum[T any](x, y T) T {
	return x
}

func sum2[T int | float32 | float64](x, y T) T {
	return x + y
}

func sum3[T comparable](x, y T) T {
	if (x == nil) != (y == nil) {
		return x
	}
	return y
}
