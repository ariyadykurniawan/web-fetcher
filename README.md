# Web Fetcher

Web Fetcher is a command-line program written in Go that allows you to fetch web pages, save them to disk

## Usage

1. Build the Docker image:
   ```bash
   docker build -t web-fetcher .
   
2. Run the docker container
   ```bash
   docker run -it web-fetcher /bin/bash

3. Inside the container, run the ./fetch command to fetch web pages:
    ```bash
    ./fetch https://www.example.com