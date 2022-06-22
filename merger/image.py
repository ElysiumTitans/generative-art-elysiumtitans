# import PIL
from PIL import Image
import sys

pathToSaveImg = sys.argv[1]
imageLayers = sys.argv[2:]

numLayers = len(imageLayers)

image1 = Image.open(imageLayers[0])
image_copy = image1.copy()
position = (0, 0)

for i, x in enumerate(imageLayers):
    if i != 0:
        image = Image.open(imageLayers[i])
        image_copy.paste(image, position, image)

image_copy.save(pathToSaveImg, dpi=(1200, 1200))

exit()