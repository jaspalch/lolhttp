#!/bin/bash

while true; do 
    read input; echo -e "$input" | ncat 127.0.0.1 1456 
done
