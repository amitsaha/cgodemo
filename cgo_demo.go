package main

/*
int myfunc(int x) {
    return x*x;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Printf("%v\n", C.myfunc(2))
}
