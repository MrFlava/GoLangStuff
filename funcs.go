package main
import (
	"fmt"
	"math"
)

func saySomething(n string) {
	fmt.Printf("Good monin' %v \n", n)
}

func sayBye(n string){
	fmt.Printf("by mah %v \n", n)
}

func cyclNames(n []string, f func(string)) {
	for _, v :=range n {
		f(v)
	}

}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main(){
	// saySomething("mario")
	// saySomething("Arthur")
	// sayBye("mario")


	cyclNames([]string{"cloud", "slay", "bagget"}, saySomething)	
	cyclNames([]string{"cloud", "slay", "bagget"}, sayBye)	
	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 i %0.3f and circle 2 is %0.3f", a1, a2)
}