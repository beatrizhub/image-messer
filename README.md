# Image Messer

## Description

This is a simple website to mess with images. It can be used to create a new image from a given image by applying a selected transformation to it.

## Usage

### Run locally
```
docker build -t image-messer-image .
docker run --network host -d --name image-messer-container -p 8080:8080 image-messer-image
```