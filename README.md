# insect-detection-coap

## Summary
A workflow for using a rasberry pi zero 2 w to capture images of leaves, trasfer them to a server using CoAP and process them using PyTorch for insect damage detection

## Setup 
### Hardware requirements
- An Nvidia graphics card with at least 6 gigs of video memmory

### Software requirements
- Go version 1.22 
- Docker
- Cuda 11.8 drivers
- Python 3.11
- Conda

### Running the project
- Go to https://github.com/jasonbuchanan145/Insect-Damage-Finder and download the project
- Create the environment ```conda env create -f env.yaml```
- Train the model ```conda run -n env python training.py```
- When completed it will output model.pth
- After training copy the resulting pth file to this project's ./pytorch_model/model.pth
- Compile the go server
    - go to ./server
    - run go build
- Run ```Docker-compose up```
    - Known issues:
        - The pytorch model container is several gigabytes due to cuda, pytorch, etc, download times could take a bit
        - Because of the size of the dependencies during conda enviornment creation you may see this error ```'Connection broken: IncompleteRead``` a dependency broke mid download, you just need to run the compose up again
- In another tab go to ./client
- compile the client `go build`

