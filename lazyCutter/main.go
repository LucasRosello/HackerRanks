package main

func main() {

}

func SimpleLazyCutter(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func OptimizedLazyCutter(n int) int {
	return (n * (n + 1) / 2) + 1
}
