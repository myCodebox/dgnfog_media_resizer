# DGNFOG Media Resizer

A small image processor script written in go.

## Usage

Start the script
```bash
go run .
```

---
## Welcome screen
```
 +-+ +-+ +-+ +-+ +-+ +-+   +-+ +-+ +-+ +-+ +-+
 |D| |N| |G| |F| |O| |G|   |M| |e| |d| |i| |a|
 +-+ +-+ +-+ +-+ +-+ +-+   +-+ +-+ +-+ +-+ +-+
Use the arrow keys to navigate: ↓ ↑ → ←
? What do you want to do:
  > Fetch from JSON to input folder
    Start resizing
    Cleanup input folder
    Cleanup output folder
    Exit
```

---
## Fetch from JSON to input folder
Update the JSON file `./src/json/media.json` or get it from the Redaxo backend page (addon Digitalocean) and overwrite it.

The app fetch the images and store it in the input folder `./src/in/`

---
## Start resizing
Resize all images from the input folder `./src/in/`and save 4 versions of each image.

Name | width | height
--- | --- | ---:
normal | auto | 512
small | auto | 53
normal_mark | auto | 512
small_mark | auto | 53

---
## Cleanup input folder
This cleanup the output folder `./src/in/` and remove all file.

---
## Cleanup output folder
This cleanup the output folder `./src/out/` and remove all file.

---
## Exit
Stop the running script. Also `Ctrl + C` works =)

---
## Custom folder
Use the `.env` file to set your own in and out folder.

Default settings:
```
# Folder
FOLDER_IN="./src/in/" 
FOLDER_OUT="./src/out/"

# Json
JSON_FILE="./src/json/media.json"
```

---
### Image from unsplash.com (v0.1-alpha)
* blossom.jpg - alex-blajan-FlUbZ-2S014-unsplash.jpg
* koala.jpg - christine-ellsay-YerVHy1nXq8-unsplash.jpg
* cactus.jpg - david-sola-3guU1kCxxy0-unsplash.jpg
* penguins.jpg - derek-oyen-3Xd5j9-drDA-unsplash.jpg
* desert.jpg - ganapathy-kumar-L75D18aVal8-unsplash.jpg
* tulips.jpg - ilona-frey-lVgR-jwkK7E-unsplash.jpg
* castel.jpg - jacek-dylag-DcQ8dSqEosA-unsplash.jpg
* jellyfish.jpg - jeffrey-hamilton-JtVyK2Sej2I-unsplash.jpg
* waterflower.jpg - jon-geng-MM1FpBrhBPE-unsplash.jpg

#### Todo
* ~~check if a file is an image~~