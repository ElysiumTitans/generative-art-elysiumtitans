
<h1 align="center">Elysium Titans Generative Art Script</h1>

This script is designed to created NFT images composed of a range of different layers. It also creates the metadata files with all of the characteristic information and stores them in a format that is easy to upload to NFTMakerPro. It requires a lot of customization to fit your project

# What it does:
- Creates a defined number of characters with the image layers you provide it
- Creates a matching metadata file per character following the CIP 721 standard
- You can add rarity to characteristics e.g. if you are creating 400 characters but only was 4 of them to have a certain background you can specify this in the script
- Ensures every character is unique
- Produces the images using a python package for the best quality

Instructions for set up follow:

## Preparing the images
There is a folder named `characters` - below is an example of the expected folder layout showcasing 2 image layers, the first image layer has 5 variations, the second has 5 variations and each variation has 5 different colours each

### Layer One
- characters/CHAR1/1/A.png    <--- body --- skintone 1
- characters/CHAR1/1/B.png    <--- body --- skintone 2
- characters/CHAR1/1/C.png    <--- body --- skintone 3
- characters/CHAR1/1/D.png    <--- body --- skintone 4
- characters/CHAR1/1/E.png    <--- body --- skintone 5

### Layer Two
- characters/CHAR1/2/1/A.png    <--- outfit one swimmer --- yellow
- characters/CHAR1/2/1/B.png    <--- outfit one swimmer --- red
- characters/CHAR1/2/1/C.png    <--- outfit one swimmer --- green
- characters/CHAR1/2/1/D.png    <--- outfit one swimmer --- blue
- characters/CHAR1/2/1/E.png    <--- outfit one swimmer --- black

- characters/CHAR1/2/2/A.png    <--- outfit two sailor --- yellow
- characters/CHAR1/2/2/B.png    <--- outfit two sailor --- red
- characters/CHAR1/2/2/C.png    <--- outfit two sailor --- green
- characters/CHAR1/2/2/D.png    <--- outfit two sailor --- blue
- characters/CHAR1/2/2/E.png    <--- outfit two sailor --- black

- characters/CHAR1/2/3/A.png    <--- outfit three pirate --- yellow
- characters/CHAR1/2/3/B.png    <--- outfit three pirate --- red
- characters/CHAR1/2/3/C.png    <--- outfit three pirate --- green
- characters/CHAR1/2/3/D.png    <--- outfit three pirate --- blue
- characters/CHAR1/2/3/E.png    <--- outfit three pirate --- black

- characters/CHAR1/2/4/A.png    <--- outfit four lecturer --- yellow
- characters/CHAR1/2/4/B.png    <--- outfit four lecturer --- red
- characters/CHAR1/2/4/C.png    <--- outfit four lecturer --- green
- characters/CHAR1/2/4/D.png    <--- outfit four lecturer --- blue
- characters/CHAR1/2/4/E.png    <--- outfit four lecturer --- black

- characters/CHAR1/2/5/A.png    <--- outfit five astronaut --- yellow
- characters/CHAR1/2/5/B.png    <--- outfit five astronaut --- red
- characters/CHAR1/2/5/C.png    <--- outfit five astronaut --- green
- characters/CHAR1/2/5/D.png    <--- outfit five astronaut --- blue
- characters/CHAR1/2/5/E.png    <--- outfit five astronaut --- black

### If you have backgrounds store them like so
- characters/backgrounds/1.png
- characters/backgrounds/2.png
- characters/backgrounds/3.png


The program expects the images in this format so you'll either have to stick to it or modify the program to suit. 

I've have provided the script i used to put the images into their respective folders programmatically but this assumes a specific naming convention to start with.


## Character specific script

You'll need to write a script specific to your character with all the characteristic names and directives on how to manage each image layer. This is simpler than it may sound - I have provided the script for the Elysium Titans character Zeus so you can model your script off of it. You'll need to create some directories to get this working e.g. the `charFinal` directory

There is a test script for zeus that you can use as well to safeguard the quality of your code.

## Running the script - IMPORTANT PREREQUISITES

1. There is a `tar` file in the `merger` directory, the program will use this to process the images. Run `tar -xzvf pillow.tar.gz` to unpack
2. `cd` into the folder that was just unpacked - likely called `Pillow-8.3.2` and run `python3 -m pip install Pillow`. You should now be able to use the `pil` package for image processing.
3. Run `go get` in `merger` directory
4. Run the script to generate your images `go run main.go zeus.go`





Created and maintained by Naeri Adam Fernandez of Elysium Titans
Contact us on "https://discord.gg/gYgNcXz3Zc" for assistance.
