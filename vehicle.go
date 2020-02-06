package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type vehicle interface {
}

type car struct {
	model       string
	make        string
	typeVehicle string
}

type truck struct {
	model       string
	make        string
	typeVehicle string
}

type bike struct {
	model string
	make  string
}



//Values Array from the Json File
type Values struct {
	Models []Model `json:"values"`
}

//Model feedback of the vehicles
type Model struct {
	Name      string   `json:"model"`
	Feedbacks []string `json:"feedbacks"`
}


const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

type feedbackResult struct {
	feedbackTotal    int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral  int
}

type rating float32
var vehicleResult map[string]feedbackResult
var inventory []vehicle


func init() {

	vehicleResult = make(map[string]feedbackResult)

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
	generateRating()

	//Print ratings for the different vehicles
	for _, veh:= range inventory {
		switch v := veh.(type) {
			case car:
				v.carDetails()
			case bike:
				v.bikeDetails()
			case truck:
				v.truckDetails()
			default:
				fmt.Printf("Are you sure this Vehicle Type exists")
		}
	}

}

 
func showRating(model string) {
	ratingFound := false
	for m , r:= range vehicleResult {
		if m == model {
		fmt.Printf("Total Ratings:%v\t\tPositive:%v\tNegative:%v\tNeutral:%v" , r.feedbackTotal , r.feedbackPositive , r.feedbackNegative , r.feedbackNeutral)
			ratingFound = true
		} 
	}
	if !ratingFound {
		fmt.Printf("No Ratings for this vehicle" )
	}

}  

func (c *car)carDetails() {
	fmt.Printf("\nCar: \t%v:%v\t\t",c.make,c.model)
	showRating(c.model)
}

func (b *bike) bikeDetails() {
	fmt.Printf("\nBike: \t%v:%v\t\t",b.make,b.model)
	showRating(b.model)
}

func (t *truck) truckDetails() {
	fmt.Printf("\nTruck: \t%v:%v\t\t",t.make,t.model)
	showRating(t.model)
}

func readJSONFile() Values{
	jsonFile, err := os.Open("feedback.json")

	if err != nil {
		log.Fatal("File not found")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var content Values

	json.Unmarshal(byteValue, &content)
	
	return content
}

func generateRating() {

	f := readJSONFile()
	
	for _, v := range f.Models {
		var vehResult feedbackResult
		var vehRating rating
		for _, msg := range v.Feedbacks {
		
			if text := strings.Split(msg, " ") ; len(text) >= 5 {
				vehRating = 5.0
				vehResult.feedbackTotal++
				for _, word := range text {
					switch s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r"); s {
						case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
							vehRating += extraPositive
						case "help", "helpful", "thanks", "thank you", "happy":
							vehRating += positive
						case "not helpful", "sad", "angry", "improve", "annoy":
							vehRating += negative
						case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
							vehRating += extraNegative
						}
					
				}
				switch {
					case vehRating > 8.0:
						vehResult.feedbackPositive++
					case vehRating >= 4.0 && vehRating <= 8.0:
						vehResult.feedbackNeutral++
					case vehRating < 4.0:
						vehResult.feedbackNegative++
					}
				}
			}
		vehicleResult[v.Name] = vehResult
	}


}
