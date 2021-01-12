# Resource server for image storing
Server for downloading images from a URL in a local folder. With the further ability to receive these images by a special assigned name.
## How to use
Server starting by default on port `:11111` and has few endpoints.
### GET /load<?link>
The `load` endpoint loads one image into the assets folder. `link` is a request parameter. This parameter should contain URL for the image.
Response (converted filename):
```sh
3344ba80a6b5476ce8cb7b8f2e864cb3.png(|jpg)
```
### POST /load/batch
The `/load/batch` endpoint loads multiply images into the assets folder. Data must be in json format.
For example:
```sh
[
    ...
    "https://someresourceurl/image1.png",
    "https://someresourceurl/image2.jpg",
    "https://someresourceurl/image3.png",
    ...
]
```
Response (a hash map of the given url for the new filename):
```sh
{
    ...
    "https://someresourceurl/image1.png": "3344ba80a6b5476ce8cb7b8f2e864cb3.png(|jpg)",
    ...
}
```
### GET /
In development
## Installation
### Docker
Solution still in developing. Dockerfile will be added after all features are complete.
## Build
For building app run next command
```sh
$ make build
```
After that you will see a new folder with Go executable
## Run
To run the application you need to pass the following CMD command
```sh
$ make run
```
or
```sh
$ ./cmd/main -path (optional)
```
`-path` - the path to the folder where you will store the images after upload
By default path is `./resources`


