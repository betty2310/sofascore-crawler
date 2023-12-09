#!/bin/bash

# Specify the directory path
directory="./output"

# Loop through each folder in the directory
for folder in "$directory"/*; do
    # Check if the folder contains a subfolder named "matches"
    if [ ! -d "$folder/matches" ]; then
        # If not, delete the folder
        echo "Deleting $folder"
        rm -r "$folder"
    fi
done
