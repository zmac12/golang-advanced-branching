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

/*
func TestMainStructs(t *testing.T) {

	t.Run("VehicleStructs" , func(t *testing.T) {
		t.Run("struct=car" , TestCarStructIsDefined )
		t.Run("struct=truck" , TestTruckStructIsDefined )
		t.Run("struct=bike" , TestBikeStructIsDefined )
	})
}
*/

// Task 1: Define vehicle interface
func TestVehicleInterfaceIsDefined(t *testing.T) {

	didFindAnInterface, didFindTheInterface := checkInterface("vehicle")

	if !didFindAnInterface || !didFindTheInterface {

		t.Error("Did not define an interface named vehicle")

	}

}

// Task 2: Define 3 structs
func TestMainStructsAreDefined(t *testing.T) {

	// Check for car struct
	didFindAStruct, didFindTheStruct := checkStruct("car")
	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named car")
	}

	// Check for truck struct
	didFindAStruct, didFindTheStruct = checkStruct("truck")

	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named truck")
	}

	// Check for bike struct
	didFindAStruct, didFindTheStruct = checkStruct("bike")

	if !didFindAStruct || !didFindTheStruct {
		t.Error("Did not define a struct named bike")
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
