package main

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

// Values array for the feedback.json file
type Values struct {
	Models []Model `json:"values"`
}

// Model array for the feedback.json file
type Model struct {
	Name     string   `json:"model"`
	Feedback []string `json:"feedback"`
}

type feedbackResult struct {
	feedbackTotal    int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral  int
}

var (
	vehicleResult map[string]feedbackResult
	inventory     []vehicle
)

type rating float32

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
		car{"Civic", "Honda", "Hatchback"},
		car{"A5", "Audi", "Coupe"},
		car{"Mazda6", "Mazda", "Sedan"},
		car{"CRV", "Honda", "SUV"},
		car{"Camry", "Toyota", "Sedan"},
		truck{"F-150", "Ford", "Truck"},
		truck{"RAM1500", "Dodge", "Truck"}}
	vehicleResult = make(map[string]feedbackResult)

}

func main() {

	// Generate ratings for the different vehicles

	// Print ratings for the different vehicles
}

/*
func readJSONFile() Values {
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
*/
