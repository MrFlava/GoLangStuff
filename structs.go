package main

import "fmt"

func main(){
	mybill := newBill("my bill")
	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg", 4.50)
	mybill.addItem("perepichka", 4.50)
	mybill.addItem("coffee", 4.50)
	mybill.updateTip(10)
	fmt.Println(mybill.format())
}