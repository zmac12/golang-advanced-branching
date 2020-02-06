package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

func init() {

	inventory = []vehicle{
		bike{"FTR 1200", "Indian"},
		bike{"Iron 1200", "Harley"},
		car{"Sonata", "Hyundai", "Sedan"},
		car{"SantaFe", "Hyundai", "SUV"},
		car{"Civic", "Honda",  "Hatchback"},
		car{"A5", "Audi", "Coupe"},
		car{"Mazda6", "Mazda", "Sedan"},
		car{"CRV", "Honda", "SUV"},
		car{"Camry", "Toyota", "Sedan"},
		truck{"F-150", "Ford",  "Truck" , },
		truck{"RAM1500", "Dodge", "Truck"} }
}

func main() {

	//Generate ratings for the different vehicles

	//Print ratings for the different vehicles
}

 
