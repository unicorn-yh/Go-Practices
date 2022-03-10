package main
import "fmt"
func Hello(){
	fmt.Println("Hello!")
}
func Hi(){
	fmt.Println("Hi!")
}
func AllGreetings(){
	Hello()
	Hi()
}
func main(){
	AllGreetings()
	fmt.Println("in greeting.go")
}
