package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Metadata map[string]interface{}

type Details struct {
	id       int
	ipfslink string
	policyId string
}

// must be above 0
var maxCountRareBackgrounds = 4
var rareTraitCount = 20
var rareImagesBackground = 5
var rareBackgroundCount = 20
var normalBackgroundOptions = 6
var normalVariationOptions = 5

func generateNumber(nums int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(nums) + 1
}

func main() {
	// Useful if you are have multiple characters in the same set e.g. zeus starts at count 0-399, athena starts at 400-799
	startCountingFrom := 0
	numToCreate := 40
	fmt.Printf("Allocating %v sets of %v rare traits...\n", ZeusRareTraitsToAllocate, rareTraitCount)
	rareTraitAllocation := allocateRareTraits(numToCreate, rareTraitCount, ZeusRareTraitsToAllocate)
	fmt.Printf("Allocating one set of %v rare backgrounds...\n", rareBackgroundCount)
	rareBackgroundAllocation := allocateRareTraits(numToCreate, rareBackgroundCount, 1)

	// Triple check for duplicates in rare traits - it is super important to get this right.
	for i := 0; i < ZeusRareTraitsToAllocate; i++ {
		dupCount(rareTraitAllocation[i])
	}
	dupCount(rareBackgroundAllocation[0])

	// Split background allocations into batches of x(what is specified for maxCountRareBackgrounds) per image
	rareBackgroundAllocationBatched := batchActions(rareBackgroundAllocation[0], maxCountRareBackgrounds)
	combinations := generateZeusCombinations(numToCreate, rareTraitAllocation, rareBackgroundAllocationBatched)
	buildPrep(ZeusPathToImages, ZeusPathToStore, ZeusLayers, ZeusConvertConfigToStringArray, combinations, startCountingFrom)

}

func batchActions(a []int, c int) [][]int {
	r := (len(a) + c - 1) / c
	b := make([][]int, r)
	lo, hi := 0, c
	for i := range b {
		if hi > len(a) {
			hi = len(a)
		}
		b[i] = a[lo:hi:hi]
		lo, hi = hi, hi+c
	}
	return b
}

func allocateRareTraits(numToCreate, rareTraitCount, ZeusRareTraitsToAllocate int) [][]int {
	var traitAllocations [][]int
	for i := 0; i < ZeusRareTraitsToAllocate; i++ {
		if ZeusRareTraitsToAllocate > 1 {
			fmt.Printf("Generating trait allocations for trait #%v\n", i)
		}
		var imagesWithAllocatedTraits []int
		for j := 0; j < rareTraitCount; j++ {
			setSelection := true
			imageToAllocate := generateNumber(numToCreate - 1)
			for _, alreadyAllocated := range imagesWithAllocatedTraits {
				if alreadyAllocated == imageToAllocate {
					// try again, this image already had this rare trait allocated to it
					fmt.Printf("Double trait allocation for image %v for slot %v, trying again...\n", alreadyAllocated, j)
					setSelection = false
					j--
				}
			}
			if setSelection {
				imagesWithAllocatedTraits = append(imagesWithAllocatedTraits, imageToAllocate)
			}
		}
		traitAllocations = append(traitAllocations, imagesWithAllocatedTraits)
	}
	return traitAllocations
}

func buildPrep(pathToImages, pathToStore string, layers int, convertConfigToStringArray func(ZeusCombination) ([]string, map[string]interface{}), combinations []ZeusCombination, startCountingFrom int) {
	// Replace with your policyId
	policyId := "xxxx"
	count := startCountingFrom
	for i := 0; i < len(combinations); i++ {

		pngId := fmt.Sprintf("%06d", count)
		count++
		displayName := fmt.Sprintf("Character #%v", pngId)
		fmt.Print("\nDisplay Name: ", displayName)
		fmt.Printf("\nCombination %v: %+v", i, combinations[i])
		config, configMetadata := convertConfigToStringArray(combinations[i])
		metadata := buildMetadata(policyId, pngId, displayName, configMetadata)
		buildFinal(layers, pathToImages, pathToStore, pngId, config, metadata)
	}
}

func buildFinal(layers int, pathToImages, pathToStore, pngId string, config []string, metadataJson map[string]interface{}) {
	// Layers correspond to the number assigned to each body part so they should be assembled in this order -> 1-9
	// Layer 0 is always the background
	pathList := []string{}
	for i := 0; i < layers; i++ {
		if strings.Contains(config[i], "None") {
			continue
		}
		path := fmt.Sprintf("%s/%v/%s.png", pathToImages, i, config[i])
		if i == 0 {
			path = fmt.Sprintf("../characters/backgrounds/%s.png", config[i])
		}
		pathList = append(pathList, path)
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	var cmd *exec.Cmd
	whereToSaveImage := fmt.Sprintf("%s/pngs/%v.png", pathToStore, pngId)

	if len(pathList) == 5 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4])
	} else if len(pathList) == 6 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4], pathList[5])
	} else if len(pathList) == 7 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4], pathList[5], pathList[6])
	} else if len(pathList) == 8 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4], pathList[5], pathList[6], pathList[7])
	} else if len(pathList) == 9 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4], pathList[5], pathList[6], pathList[7], pathList[8])
	} else if len(pathList) == 10 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4], pathList[5], pathList[6], pathList[7], pathList[8], pathList[9])
	} else if len(pathList) == 11 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4], pathList[5], pathList[6], pathList[7], pathList[8], pathList[9], pathList[10])
	} else if len(pathList) == 12 {
		cmd = exec.Command("python3", "image.py", whereToSaveImage, pathList[0], pathList[1], pathList[2], pathList[3], pathList[4], pathList[5], pathList[6], pathList[7], pathList[8], pathList[9], pathList[10], pathList[11])
	} else {
		fmt.Print(len(pathList), pathList)
		fmt.Print("\nOH NO WE DIDNT GET A MATCHING LENGTH YOU BETTER CHECK THAT OUT.\nYou'll probably need to add an entry to match the number of layers your image has.")
	}
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println(out.String())

	// NFTMakerPro expects the metadata in a file with the prefix `.metadata`
	metadataFileName := fmt.Sprintf("%s/metadata/%v.metadata", pathToStore, pngId)
	// // Create json file for metadata
	metadata, err := os.Create(metadataFileName)
	if err != nil {
		log.Fatalf("failed to create metadata file `%s.metadata`: %s", pngId, err)
	}
	file, err := json.MarshalIndent(metadataJson, "", "	")
	if err != nil {
		log.Fatalf("failed to parse metadataJson: %s", err)
	}
	err = os.WriteFile(metadataFileName, file, 0644)
	if err != nil {
		log.Fatalf("failed to write metadata to file: %s", err)
	}

	defer metadata.Close()
}

func buildMetadata(policyId, pngId, displayName string, configMetadata map[string]interface{}) Metadata {
	idName := fmt.Sprintf("NFTCHAR1%s", pngId)
	ipfslinkPlaceholder := "<ipfs_link>"
	var files = make([]map[string]interface{}, 1)
	files[0] = map[string]interface{}{
		"name":      displayName,
		"mediaType": "image/png",
		"src":       ipfslinkPlaceholder,
	}
	descriptionPlaceholder := "Project description."
	return map[string]interface{}{
		"721": map[string]interface{}{
			policyId: map[string]interface{}{
				idName: map[string]interface{}{
					"name":        displayName,
					"image":       ipfslinkPlaceholder,
					"mediaType":   "image/png",
					"seriesName":  "Series Name",
					"description": descriptionPlaceholder,
					"files":       files,
					"attributes":  configMetadata,
					"links": map[string]interface{}{
						"Discord": "https://discord.gg/xxx",
						"Twitter": "https://twitter.com/xxx",
						"Website": "https://xxx.xx/",
					},
				},
			},
			"version": "1.0",
		},
	}
}

func convertToAlphabet(num int) string {
	switch num {
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	case 5:
		return "E"
	case 6:
		return "None"
	}

	fmt.Printf("\nERROR: No match for conversion of `%v` - returning `C` by default.\n", num)
	return "C"
}

func getBackgroundName(background int) string {
	switch background {
	case 1:
		return "White"
	case 2:
		return "Black"
	case 3:
		return "Magenta"
	case 4:
		return "Turquoise"
	case 5:
		return "Cream"
	case 6:
		return "Green"
	case 7:
		return "Yellow"
	case 8:
		return "Pink"
	case 9:
		return "Blue"
	case 10:
		return "Purple"
	case 11:
		return "Grey"
	}

	fmt.Printf("\nERROR: BACKGROUNDS: No match for conversion of `%v` - returning empty string.\n", background)
	return ""
}

func getSkinTone(skinTone string) string {
	switch skinTone {
	case "1":
		return "One"
	case "2":
		return "Two"
	case "3":
		return "Three"
	case "4":
		return "Four"
	case "5":
		return "Five"
	}

	fmt.Printf("\nERROR: SKINTONE: No match for conversion of `%v` - returning empty string.\n", skinTone)
	return ""
}

func dupCount(list []int) {

	duplicate_frequency := make(map[int]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
			fmt.Printf("DUPLICATE ON ITEM %v.\n", item)
			panic("Cannot have duplicated for rare traits.")
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}

}
