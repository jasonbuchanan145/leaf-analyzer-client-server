FROM nvidia/cuda:11.8.0-base-ubuntu20.04

# The school super computer requires conda but there isn't a cuda image that includes a conda build
# So need to manually install cuda on the base image
ENV CONDA_DIR /opt/conda
ENV PATH $CONDA_DIR/bin:$PATH
#miniconda is not on the linux repos, have to download manually 
#https://conda.io/projects/conda/en/latest/user-guide/install/linux.html
RUN apt-get update && apt-get install -y wget netcat\
    && wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh -O miniconda.sh \
    && bash miniconda.sh -b -p $CONDA_DIR \
    && rm miniconda.sh
COPY ./env.yaml /app/
RUN conda env create -f /app/env.yaml
WORKDIR /app
#change this value as you edit the python files so it keeps the cache envs but doesn't cache changed python files
ENV cache_bust=5
# Copy the Python script into the container
COPY ./queue_runner.py /app/
COPY ./modelRunner.py /app/
COPY ./model.pth /app/
# Activate the environment and run the Python script with CUDA support
CMD ["conda", "run", "-n", "env", "python", "queue.py"]
