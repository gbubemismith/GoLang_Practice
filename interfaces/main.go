package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c *Chicken) PrepareDish() {
	fmt.Println("Preparing dish Chicken!!!")
}

func (c *Salad) PrepareDish() {
	fmt.Println("Preparing dish salad")
}

func main() {

}
