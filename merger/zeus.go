package main

import (
	"fmt"
	"strconv"
)

type ZeusCombination struct {
	background   int
	skinTone     int
	weaponName   int
	weaponColour int
	hairColour   int
	hairStyle    int
	beardStyle   int
	armourName   int
	armourColour int
	eyeMarkings  int
}

var ZeusPathToImages string = "../characters/CHAR1"
var ZeusPathToStore string = "../charFinal/zeus"
var ZeusRareTraitsToAllocate = 1
var ZeusLayers int = 10

func generateZeusCombinations(combinationCount int, rareTraitAllocation, rareBackgroundAllocation [][]int) []ZeusCombination {
	combinations := []ZeusCombination{}

	for i := 0; i < combinationCount; i++ {
		thisCombo := ZeusCombination{
			background:   generateNumber(normalBackgroundOptions) + 5,
			skinTone:     generateNumber(normalVariationOptions),
			weaponName:   generateNumber(normalVariationOptions),
			weaponColour: generateNumber(normalVariationOptions),
			hairColour:   generateNumber(normalVariationOptions),
			hairStyle:    generateNumber(normalVariationOptions),
			beardStyle:   generateNumber(normalVariationOptions),
			armourName:   generateNumber(normalVariationOptions),
			armourColour: generateNumber(normalVariationOptions),
			eyeMarkings:  generateNumber(normalVariationOptions),
		}

		checkComboExists := CheckIfZeusCombinationExists(combinations, thisCombo)
		if checkComboExists == "notexists" {
			combinations = append(combinations, thisCombo)
		} else if checkComboExists == "exists" {
			// retry!!
			fmt.Printf("Got duplicate on index '%v', retrying...\n", i)
			i--
		}
	}

	// Set rare trait allocations
	noArmourAllocations := rareTraitAllocation[0]
	for i := 0; i < rareTraitCount; i++ {
		fmt.Printf("%v: Armour: Updating trait for image [%v] to be rare trait.\n", i, noArmourAllocations[i])
		combinations[noArmourAllocations[i]].armourName = 888
	}
	for i := 0; i < rareImagesBackground; i++ {
		background := i + 1 // background files are numbered from 1-11 so add one.
		backgroundBatch := rareBackgroundAllocation[i]
		for j := 0; j < len(backgroundBatch); j++ {
			fmt.Printf("Rare background #%v: Updating background for image [%v].\n", background, backgroundBatch[j])
			combinations[backgroundBatch[j]].background = background
		}
	}

	return combinations
}

func CheckIfZeusCombinationExists(combinations []ZeusCombination, thisCombo ZeusCombination) string {
	for i := 0; i < len(combinations); i++ {
		if combinations[i] == thisCombo {
			return "exists"
		}
	}
	return "notexists"
}

var ZeusConvertConfigToStringArray = func(combination ZeusCombination) ([]string, map[string]interface{}) {
	armourColour := convertToAlphabet(combination.armourColour)
	skinTone := strconv.Itoa(combination.skinTone)
	hairColour := convertToAlphabet(combination.hairColour)
	armourName := strconv.Itoa(combination.armourName)
	background := strconv.Itoa(combination.background)
	if armourName == "888" {
		armourName = "None"
		armourColour = "None"
	}

	config := []string{
		background, // 0. background
		hairColour, // 1. backhair
		skinTone,   // 2. basebody
		fmt.Sprintf("%v/%v", armourName, armourColour), // 3. armourset
		skinTone, // 4. basehead
		fmt.Sprintf("%v/%v", combination.eyeMarkings, hairColour),                                               // 5. eyesbrows
		fmt.Sprintf("%v/%v", combination.hairStyle, hairColour),                                                 // 6. hairheadwear
		fmt.Sprintf("%v/%v", combination.beardStyle, hairColour),                                                // 7. beardmouth
		fmt.Sprintf("%v/%v", strconv.Itoa(combination.weaponName), convertToAlphabet(combination.weaponColour)), // 8. weaponaura
		skinTone, // 9. basehand
	}

	weaponName, weaponColour := getZeusWeaponMetadata(combination.weaponColour, combination.weaponName)
	if armourName != "None" {
		armourName, armourColour = getZeusArmourMetadata(combination.armourColour, combination.armourName)
	}
	configMetadata := map[string]interface{}{
		"God":           "Zeus",
		"Pantheon":      "Greek",
		"Skin Tone":     getSkinTone(skinTone),
		"Weapon Name":   weaponName,
		"Weapon Colour": weaponColour,
		"Hair Colour":   getZeusHairColour(combination.hairColour),
		"Crown / Hair":  getZeusHairStyle(combination.hairStyle),
		"Beard Style":   getZeusBeardStyle(combination.beardStyle),
		"Armour Name":   armourName,
		"Armour Colour": armourColour,
		"Eye Markings":  getZeusEyeMarkings(combination.eyeMarkings),
		"Background":    getBackgroundName(combination.background),
	}
	return config, configMetadata
}

func getZeusHairColour(hairColour int) string {
	switch hairColour {
	case 1:
		return "Platinum"
	case 2:
		return "Soft Blond"
	case 3:
		return "Charcoal"
	case 4:
		return "Brown"
	case 5:
		return "Chestnut Brown"
	}

	fmt.Printf("\nERROR: HAIRCOLOUR: No match for conversion of `%v` - returning empty string.\n", hairColour)
	return ""
}

func getZeusHairStyle(hairStyle int) string {
	switch hairStyle {
	case 1:
		return "Champion of Olympus"
	case 2:
		return "Poseidon's Gift"
	case 3:
		return "Aetos of God"
	case 4:
		return "Lord of Thunder"
	case 5:
		return "Councillor of The Heavens"
	}

	fmt.Printf("\nERROR: HAIRSTYLE: No match for conversion of `%v` - returning empty string.\n", hairStyle)
	return ""
}

func getZeusBeardStyle(beardStyle int) string {
	switch beardStyle {
	case 1:
		return "Double Braid"
	case 2:
		return "Single Braid"
	case 3:
		return "Unkempt"
	case 4:
		return "Full"
	case 5:
		return "Trim"
	}

	fmt.Printf("\nERROR: BEARDSTYLE: No match for conversion of `%v` - returning empty string.\n", beardStyle)
	return ""
}

func getZeusEyeMarkings(eyeMarkings int) string {
	switch eyeMarkings {
	case 1:
		return "Head of Olympus"
	case 2:
		return "Hephaestus' Gift"
	case 3:
		return "Traveler of Worlds"
	case 4:
		return "Titan Slayer"
	case 5:
		return "Athena's Blessing"
	}
	fmt.Printf("\nERROR: EYEMARKINGS: No match for conversion of `%v` - returning empty string.\n", eyeMarkings)
	return ""
}

func getZeusWeaponMetadata(weaponColour, weapon int) (string, string) {
	weaponName := ""
	switch weapon {
	case 1:
		weaponName = "The Axe of Kratos"
	case 2:
		weaponName = "Thunder"
	case 3:
		weaponName = "The Caduceus"
	case 4:
		weaponName = "Lightning"
	case 5:
		weaponName = "Poseidons Trident"
	}
	if weaponName == "" {
		fmt.Printf("\nERROR: WEAPON: No match for conversion of `%v` - returning empty string.\n", weaponName)
	}
	weaponColourName := fmt.Sprintf("%s - Variation #%v", weaponName, weaponColour)
	return weaponName, weaponColourName
}

func getZeusArmourMetadata(armourColour, armour int) (string, string) {
	armourName := ""
	switch armour {
	case 1:
		armourName = "Aetos Dios"
	case 2:
		armourName = "The Dwarf's Gift"
	case 3:
		armourName = "Storm"
	case 4:
		armourName = "Final Order of Hephaestus"
	case 5:
		armourName = "Slayer of Titans"
	}
	if armourName == "" {
		fmt.Printf("\nERROR: ARMOUR: No match for conversion of `%v` - returning empty string.\n", armourName)
	}
	armourColourName := fmt.Sprintf("%s - Variation #%v", armourName, armourColour)
	return armourName, armourColourName
}
