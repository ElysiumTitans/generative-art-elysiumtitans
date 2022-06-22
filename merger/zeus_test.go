package main

import (
	"testing"
)

var testComboDoesExist = ZeusCombination{
	skinTone:     3,
	weaponName:   4,
	weaponColour: 5,
	hairColour:   5,
	hairStyle:    4,
	beardStyle:   5,
	armourName:   1,
	armourColour: 5,
	eyeMarkings:  3,
}

var testComboDoesNotExist = ZeusCombination{
	skinTone:     2,
	weaponName:   6,
	weaponColour: 1,
	hairColour:   1,
	hairStyle:    5,
	beardStyle:   3,
	armourName:   2,
	armourColour: 4,
	eyeMarkings:  4,
}

var testCombinations = []ZeusCombination{
	{skinTone: 4, weaponName: 4, weaponColour: 4, hairColour: 2, hairStyle: 1, beardStyle: 3, armourName: 1, armourColour: 4, eyeMarkings: 2},
	{skinTone: 4, weaponName: 3, weaponColour: 1, hairColour: 3, hairStyle: 5, beardStyle: 5, armourName: 5, armourColour: 1, eyeMarkings: 3},
	{skinTone: 2, weaponName: 4, weaponColour: 4, hairColour: 3, hairStyle: 2, beardStyle: 3, armourName: 5, armourColour: 4, eyeMarkings: 3},
	{skinTone: 2, weaponName: 4, weaponColour: 2, hairColour: 4, hairStyle: 3, beardStyle: 4, armourName: 2, armourColour: 3, eyeMarkings: 1},
	{skinTone: 2, weaponName: 4, weaponColour: 4, hairColour: 2, hairStyle: 4, beardStyle: 4, armourName: 5, armourColour: 1, eyeMarkings: 3},
	{skinTone: 2, weaponName: 5, weaponColour: 3, hairColour: 3, hairStyle: 3, beardStyle: 5, armourName: 2, armourColour: 5, eyeMarkings: 3},
	{skinTone: 4, weaponName: 1, weaponColour: 4, hairColour: 5, hairStyle: 5, beardStyle: 2, armourName: 4, armourColour: 1, eyeMarkings: 3},
	{skinTone: 2, weaponName: 5, weaponColour: 2, hairColour: 4, hairStyle: 1, beardStyle: 3, armourName: 6, armourColour: 3, eyeMarkings: 5},
	{skinTone: 2, weaponName: 6, weaponColour: 3, hairColour: 2, hairStyle: 2, beardStyle: 3, armourName: 3, armourColour: 2, eyeMarkings: 4},
	{skinTone: 2, weaponName: 1, weaponColour: 2, hairColour: 4, hairStyle: 1, beardStyle: 4, armourName: 3, armourColour: 2, eyeMarkings: 3},
	{skinTone: 5, weaponName: 2, weaponColour: 3, hairColour: 3, hairStyle: 4, beardStyle: 4, armourName: 3, armourColour: 5, eyeMarkings: 4},
	{skinTone: 2, weaponName: 2, weaponColour: 3, hairColour: 3, hairStyle: 5, beardStyle: 3, armourName: 6, armourColour: 1, eyeMarkings: 3},
	{skinTone: 2, weaponName: 6, weaponColour: 3, hairColour: 5, hairStyle: 3, beardStyle: 2, armourName: 3, armourColour: 1, eyeMarkings: 1},
	{skinTone: 3, weaponName: 4, weaponColour: 3, hairColour: 1, hairStyle: 4, beardStyle: 4, armourName: 3, armourColour: 5, eyeMarkings: 1},
	{skinTone: 2, weaponName: 6, weaponColour: 2, hairColour: 3, hairStyle: 1, beardStyle: 4, armourName: 1, armourColour: 3, eyeMarkings: 5},
	{skinTone: 5, weaponName: 1, weaponColour: 2, hairColour: 2, hairStyle: 1, beardStyle: 3, armourName: 5, armourColour: 1, eyeMarkings: 4},
	{skinTone: 3, weaponName: 4, weaponColour: 5, hairColour: 5, hairStyle: 4, beardStyle: 5, armourName: 1, armourColour: 5, eyeMarkings: 3},
	{skinTone: 5, weaponName: 5, weaponColour: 3, hairColour: 1, hairStyle: 5, beardStyle: 2, armourName: 2, armourColour: 4, eyeMarkings: 3},
	{skinTone: 1, weaponName: 3, weaponColour: 4, hairColour: 2, hairStyle: 2, beardStyle: 1, armourName: 4, armourColour: 2, eyeMarkings: 1},
}

func Test_CheckCombinationDoesExist(t *testing.T) {
	checkComboExists := CheckIfZeusCombinationExists(testCombinations, testComboDoesExist)
	if checkComboExists != "exists" {
		t.Error("expected", "exists", "got", "notexists")
	}
}

func Test_CheckCombinationDoesNotExist(t *testing.T) {
	checkComboExists := CheckIfZeusCombinationExists(testCombinations, testComboDoesNotExist)
	if checkComboExists != "notexists" {
		t.Error("expected", "notexists", "got", "exists")
	}
}
