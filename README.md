# Insect Damage Detection

## Summary
A toolchain for light weight IOT devices to transmit images taken of plant leaves to a server using CoAP and process them using PyTorch to detect insect damage

## Setup 
### Hardware requirements
#### Server
- An Nvidia graphics card with at least 4gigs of video memory
- At least 8 gigs of available ram
#### Client
- Anything that has network access and can run Go binaries.

### Software requirements
- Go version 1.22 
- Docker
- Cuda 11.8 drivers
- Python 3.11
- Conda

### Running the project
#### Server
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
#### Client
- In another tab go to ./client
- compile the client `go build`
- As of right now plug and play network discovery is not enabled and only local communication is supported. To run use ./client _yourimage.jpg_

### Example
- After running the server configurations above go to the client directory
- Place your photo in ./payloads (this directory is hardcoded as it's the directory being used the zero configurations)
- example photo
  
  ![image](https://github.com/jasonbuchanan145/leaf-analyzer-server/assets/83380304/1a4237b2-b8f2-4ba4-8258-796350bc8566)
- Run ./client test.jpg
- Wait for a couple seconds
- Go to ./processed (this directory is created by docker)
- View the output, each box is labeled with a level of percentage of confidence. 
  
 ![image](https://github.com/jasonbuchanan145/leaf-analyzer-server/assets/83380304/9d71f6ae-a36c-4d1b-8d5f-939d90ecad24)

### Archetecture diagram
- TODO:
