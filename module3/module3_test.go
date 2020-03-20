package module3

import (
	"testing"
)

// Task 1: Create showRating() function
func TestShowRatingFuncIsDefined(t *testing.T) {
	if !checkFunc("showRating", "model") {
		t.Error("Create function showRating() with one parameter `model` of type `string` and no return value")
	}

}

// Task 2: ratingFound := false
func TestVarIsDefined(t *testing.T) {
	if !checkAssignedValue(funcBlock, "ratingFound:=false") {
		t.Error("Variable `ratingFound` assigned to bool value `false` is not found")
	}
}

// Task 3: Check `for m , r:= range vehicleResult`
func TestForRangeFeedback(t *testing.T) {

	if !checkForStmt(funcBlock, "m", "r", "vehicleResult") {
		t.Error("Could not find the for statement with key `m` and value `r` and `vehicleResult` as range")
	}
}

// Task 4: Check `if m == model`
func TestIfStmt(t *testing.T) {
	if !checkIfStmt(forBlock.Body, "m==model") {
		t.Error("If statment is either not defined or condition is not correct.")
	}
}


// Task 4: Statements under if
func TestStmts(t *testing.T) {

	if !checkStmts(ifBlock, "fmt.Printf") {
		t.Error("If statment is either not defined or condition is not correct.")
	}
	if !checkStmts(ifBlock, "ratingFound=true") {
		t.Error("rating not found is either not defined or condition is not correct.")
	}

}

// Task 6: Check `if !ratingFound`
func TestIfRatingFound(t *testing.T) {
	if !checkIfStmt(funcBlock, "!ratingFound") {
		t.Error("If statment is rating not found.")
	}

	if !checkStmts(ifBlock, "fmt.Printf") {
		t.Error("If statment is either not defined or condition is not correct.")
	}


}

// Task 7: Create carDetails method
func TestCarDetailsMethod(t *testing.T) {
	if !checkMethod("carDetails", "c *car") {
		t.Error("Create method `carDetails()` which has a receiver of type `*car` named `c`")
	}

}

// Task 8: Create bikeDetails method
func TestBikeDetailsMethod(t *testing.T) {
	if !checkMethod("bikeDetails", "b *bike") {
		t.Error("Create method `bikeDetails()` which has a receiver of type `*bike`named `b`")
	}

}

// Task 9: Create truckDetails method
func TestTruckDetailsMethod(t *testing.T) {
	if !checkMethod("truckDetails", "t *truck") {
		t.Error("Create method `truckDetails()` which has a receiver of type `*truck` named `t`")
	}

}

// Task 10: Check `for _, veh:= range inventory`
func TestForRangeInventory(t *testing.T) {
	if !checkForWithinMain("main", "_", "veh", "inventory") {
		t.Error("Could not find the for statement with no key and `v` as value and `f.Models` as range")
	}
}

// Task 11: Check for `switch`
func TestSwitchType(t *testing.T) {
	if !checkSwitchType(mainForStmt, "v:=veh.(type)") {
		t.Error("Could not find the switch statement with no initialization and no tag")
	}

}