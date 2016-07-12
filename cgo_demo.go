package main

/*
int square(int x) {
    return x*x;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Printf("%v\n", C.square(2))
}
