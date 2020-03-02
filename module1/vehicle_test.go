package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// Task 1: Define vehicle interface
func TestVehicleInterfaceIsDefined(t *testing.T) {

	didFindAnInterface, didFindTheInterface := checkInterface("vehicle")

	if !didFindAnInterface || !didFindTheInterface {
		t.Error("Did not define an interface named `vehicle`")
	}

}

// Task 2: Define 3 structs
func TestMainStructsAreDefined(t *testing.T) {

	// Check for car struct
	didFindAStruct, didFindTheStruct := checkStruct("car")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `car`")
	}

	// Check for truck struct
	didFindAStruct, didFindTheStruct = checkStruct("truck")

	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `truck`")
	}

	// Check for bike struct
	didFindAStruct, didFindTheStruct = checkStruct("bike")

	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `bike`")
	}
}

// Task 3: Define car and truck fields
func TestCarTruckFields(t *testing.T) {

	// car fields
	if !checkStructProperties("car", "model", "string") {
		t.Error("Did not define `model` field in `car` with the proper type")
	}
	if !checkStructProperties("car", "make", "string") {
		t.Error("Did not define `make` field in `car` with the proper type")
	}
	if !checkStructProperties("car", "typeVehicle", "string") {
		t.Error("Did not define `typeVehicle` field in `car` with the proper type")
	}

	// truck fields
	if !checkStructProperties("truck", "model", "string") {
		t.Error("Did not define `model` field in `truck` with the proper type")
	}
	if !checkStructProperties("truck", "make", "string") {
		t.Error("Did not define `make` field in `truck` with the proper type")
	}
	if !checkStructProperties("truck", "typeVehicle", "string") {
		t.Error("Did not define `typeVehicle` field in `truck` with the proper type")
	}
}

// Task 4: Define bike fields
func TestBikeFields(t *testing.T) {
	// bike fields
	if !checkStructProperties("bike", "model", "string") {
		t.Error("Did not define `model` field in `bike` with the proper type")
	}
	if !checkStructProperties("bike", "make", "string") {
		t.Error("Did not define `make` field in `bike` with the proper type")
	}
}

// Task 5: Define Values struct
func TestValuesStructIsDefined(t *testing.T) {

	// Check for car struct
	didFindAStruct, didFindTheStruct := checkStruct("Values")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `Values`")
	}
}

// Task 6: Define Model struct
func TestModelStructIsDefined(t *testing.T) {

	// Check for car struct
	didFindAStruct, didFindTheStruct := checkStruct("Model")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `Model`")
	}
}

// Task 7: Add Values fields
func TestValuesFields(t *testing.T) {
	// Values fields
	if !checkStructProperties("Values", "Models", "[]Model") {
		t.Error("Did not define `Models` field in `Values` with the proper type")
	}

}

// Task 8: Add Model fields
func TestModelFields(t *testing.T) {
	// Model fields
	if !checkStructProperties("Model", "Name", "string") {
		t.Error("Did not define `Name` field in `Model` with the proper type")
	}

	if !checkStructProperties("Model", "Feedback", "[]string") {
		t.Error("Did not define `Feedback` field in `Model` with the proper type")
	}
}

// Task 9: Define feedback struct and values
func TestFeedbackResultStructIsDefinedAndHasValues(t *testing.T) {

	// Check for feedback struct
	didFindAStruct, didFindTheStruct := checkStruct("feedbackResult")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named `feedbackResult`")
	}

	// feedbackResult fields
	if !checkStructProperties("feedbackResult", "feedbackTotal", "int") {
		t.Error("Did not define `feedbackTotal` field in `feedbackResult` with the proper type")
	}
	if !checkStructProperties("feedbackResult", "feedbackPositive", "int") {
		t.Error("Did not define `feedbackPositive` field in `feedbackResult` with the proper type")
	}
	if !checkStructProperties("feedbackResult", "feedbackNegative", "int") {
		t.Error("Did not define `feedbackNegative` field in `feedbackResult` with the proper type")
	}
	if !checkStructProperties("feedbackResult", "feedbackNeutral", "int") {
		t.Error("Did not define `feedbackNeutral` field in `feedbackResult` with the proper type")
	}

}

// Task 10: Define variables
func TestVariables(t *testing.T) {

	// vehicleResult map
	if !checkMap("vehicleResult", "string", "feedbackResult") {
		t.Error("Did not declare `vehicleResult` with the proper type.")
	}

	// inventory slice
	if !checkSlice("inventory", "vehicle") {
		t.Error("Did not declare `inventory` slice of type 'vehicle' ")
	}

}

// Task 11: Checking var initialization under func init
func TestInitializeVars(t *testing.T) {

	if vehicleResult == nil {
		t.Error("Did not initialize `vehicleResult` using make statement within `init` function")
	} 

	if inventory == nil {
		t.Error("Did not uncomment `inventory' slice within 'init` function")
	}
	
}



// ------------------------------------- Compute functions -------------------------------

//Function for reading the file
func readFile() *ast.File {

	_, currentFile, _, _ := runtime.Caller(1)

	src, err := ioutil.ReadFile(path.Join(path.Dir(currentFile), "vehicle.go"))

	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)

	if err != nil {
		panic(err)
	}
	return f
}

// Function for checking map
func checkMap(mapName, mapKey, mapValue string) bool {

	foundMapName, foundMapKeyValue := false, false

	f := readFile()

	var findMap *ast.MapType

	//var vehicleResult map[string]feedbackResult

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.ValueSpec" {
					ValueSpec := spec.(*ast.ValueSpec)
					if reflect.TypeOf(ValueSpec.Type).String() == "*ast.MapType" {
						for _, name := range ValueSpec.Names {
							if name.String() == mapName {
								foundMapName = true
								findMap = ValueSpec.Type.(*ast.MapType)
								break
							}
						}
					}
				}
			}
		}
	}

	if foundMapName {
		//fmt.Println(findMap.Key, findMap.Value)
		if findMap.Key.(*ast.Ident).String() == mapKey && findMap.Value.(*ast.Ident).String() == mapValue {
			foundMapKeyValue = true
		}
	}

	return foundMapName && foundMapKeyValue
}

// Function for checking slice
func checkSlice(sliceName, sliceType string) bool {

	foundSliceName, foundSliceType := false, false

	f := readFile()
	var findSlice *ast.ArrayType
	//	var

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				//	fmt.Println(spec, reflect.TypeOf(spec).String())
				if reflect.TypeOf(spec).String() == "*ast.ValueSpec" {
					ValueSpec := spec.(*ast.ValueSpec)
					if reflect.TypeOf(ValueSpec.Type).String() == "*ast.ArrayType" {
						for _, arr := range ValueSpec.Names {
							if arr.Name == sliceName {
								foundSliceName = true
								findSlice = ValueSpec.Type.(*ast.ArrayType)
								break
							}
						}
					}
				}
			}
		}
	}

	if foundSliceName  && findSlice.Elt.(*ast.Ident).String() == sliceType {
		foundSliceType = true

	}

	return foundSliceName && foundSliceType
}

// Function for checking said interface
func checkInterface(interfaceName string) (bool, bool) {

	var foundAInterface bool
	var foundVehicleInterface bool

	f := readFile()

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.TypeSpec" {
					typeSpec := spec.(*ast.TypeSpec)
					if reflect.TypeOf(typeSpec.Type).String() == "*ast.InterfaceType" {
						foundAInterface = true
						if typeSpec.Name.Name == interfaceName {
							foundVehicleInterface = true
						}
					}
				}
			}
		}
	}
	return foundAInterface, foundVehicleInterface
}

// Function for checking said struct
func checkStruct(structName string) (bool, bool) {

	var foundAStruct bool
	var foundNamedStruct bool

	f := readFile()

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.TypeSpec" {
					typeSpec := spec.(*ast.TypeSpec)
					if reflect.TypeOf(typeSpec.Type).String() == "*ast.StructType" {
						foundAStruct = true
						if typeSpec.Name.Name == structName {
							foundNamedStruct = true
						}
					}
				}
			}
		}
	}
	return foundAStruct, foundNamedStruct
}

// Function for checking struct fields
func checkStructProperties(structName, fieldName, fieldType string) bool {
	var targetStruct *ast.TypeSpec

	f := readFile()

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.TypeSpec" {
					typeSpec := spec.(*ast.TypeSpec)
					if reflect.TypeOf(typeSpec.Type).String() == "*ast.StructType" {
						if typeSpec.Name.Name == structName {
							targetStruct = typeSpec
							break
						}
					}
				}
			}
		}
	}

	if targetStruct == nil {
		return false
	}

	targetStructType := targetStruct.Type.(*ast.StructType)
	for _, field := range targetStructType.Fields.List {
		for _, name := range field.Names {
			if name.Name == fieldName {
				switch reflect.TypeOf(field.Type).String() {
				case "*ast.Ident":
					fType := field.Type.(*ast.Ident)
					return fType.Name == fieldType

				case "*ast.ArrayType":
					if !strings.Contains(fieldType, "[]") {
						return false
					}
					aType := field.Type.(*ast.ArrayType)
					elt := aType.Elt.(*ast.Ident)
					return elt.Name == strings.ReplaceAll(fieldType, "[]", "")
				}
			}
		}
	}
	return false
}
