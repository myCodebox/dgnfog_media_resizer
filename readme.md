# DGNFOG Media Resizer

A small image processor script written in go.

## Usage

Start the script
```bash
# run the script
go run .

# build the script
go build .
```

---
## Welcome screen
```bash
# +-+ +-+ +-+ +-+ +-+ +-+   +-+ +-+ +-+ +-+ +-+
# |D| |N| |G| |F| |O| |G|   |M| |e| |d| |i| |a|
# +-+ +-+ +-+ +-+ +-+ +-+   +-+ +-+ +-+ +-+ +-+
Use the arrow keys to navigate: ↓ ↑ → ←
? What do you want to do:
  > Fetch images from Server (use JSON)
    Resize images
    Upload to s3 bucket
    Cleanup input folder
    Cleanup output folder
    Exit
```

---
## Fetch images from Server (use JSON)
Update the JSON file `./src/json/media.json` or get it from the Redaxo backend page (addon Digitalocean) and overwrite it.

The app fetch the images and store it in the input folder `./tmp/in/`

---
## Resize images
Resize all images from the input folder `./tmp/in/`and save 4 versions of each image.

Name | width | height
--- | --- | ---:
normal | auto | 512
small | auto | 53
normal-mark | auto | 512
small-mark | auto | 53

---
### Upload to s3 bucket
Upload all images from the output folder `./tmp/out/` to a s3 bucket set in the .env file.

---
## Cleanup input folder
This cleanup the output folder `./tmp/in/` and remove all file.

---
## Cleanup output folder
This cleanup the output folder `./tmp/out/` and remove all file.

---
## Exit
Stop the running script. Also `Ctrl + C` works =)

---
## Custom folder
Use the `.env` file to set your own in and out folder and some other stuff.

Default settings:
```bash
# Folder
FOLDER_IN="./tmp/in/" 
FOLDER_OUT="./tmp/out/"

# Json
JSON_FILE="./src/json/media.json"

# HTTP
INSECURE_SKIP_VERIFY=true
MAX_IDLE_CONNS=20
MAX_IDLE_CONNS_PER_HOST=20

# s3
S3_REGION="s3"
S3_BUCKET="go_upload"
S3_ENDPOINT="http://localhost:9444/s3"
S3_AKID="AKIAIOSFODNN7EXAMPLE"
S3_SECRET_KEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
S3_TOKEN=""
```

---
## Local testing
For local testing i use [s3ninja](https://s3ninja.net/) running in Docker. Start the Docker container and open http://localhost:9444/s3/ui in your browser and create a new private  bucket called `go_upload`.

Fetch the test images, resize it and upload it to the s3 bucket.

Done 🐞



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

---
### Todo
* ~~check if a file is an image~~
* ~~upload files in output folder to s3~~