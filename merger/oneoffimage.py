# import PIL
from PIL import Image

pathToSaveImg = "../godsfinal/medusa/pngs/000916.png"
imageLayers = ['../gods/backgrounds/4.png', '../gods/MEDUSA/1/3.png', '../gods/MEDUSA/2/3/A.png', '../gods/MEDUSA/3/1.png', '../gods/MEDUSA/4/5/D.png', '../gods/MEDUSA/5/4/E.png', '../gods/MEDUSA/6/2.png', '../gods/MEDUSA/7/3/B.png', '../gods/MEDUSA/8/1.png']

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