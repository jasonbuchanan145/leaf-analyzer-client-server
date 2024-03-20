#!/bin/bash

# Example GPU job submission script

# This is a comment.
# Lines beginning with the # symbol are comments and are not interpreted by 
# the Job Scheduler.

# Lines beginning with #SBATCH are special commands to configure the job.
		
### Job Configuration Starts Here #############################################

# Export all current environment variables to the job (Don't change this)
#SBATCH --get-user-env 
#SBATCH --chdir=/home/usd.local/jason.buchanan/Insect-Damage-Finder
# The default is one task per node
#SBATCH --ntasks=1
#SBATCH --nodes=1

# Request 1 GPU
# Each gpu node has two logical GPUs, so up to 2 can be requested per node
# To request 2 GPUs use --gres=gpu:pascal:2
#SBATCH --partition=gpu
#SBATCH --gres=gpu:pascal:1 

#request 10 minutes of runtime - the job will be killed if it exceeds this
#SBATCH --time=3:00:00

# Change email@example.com to your real email address
#SBATCH --mail-user=jason.buchanan@coyotoes.usd.edu
#SBATCH --mail-type=ALL



### Commands to run your program start here ####################################

pwd
srun --gres=gpu -pgpu python3 training.py >test.out
sleep 5

nvidia-smi
