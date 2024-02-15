# Viam RDK Lightweight Docker Image

This Dockerfile allows you to create a ligthweight Docker image which will download all required binaries when instantiated.
You can configure the the container using the ```env``` file.

Build the container image:
```
docker build -t viam-rdk-light .
```

Update the ```env``` file with the config command from [app.viam.com](https://app.viam.com).

Run the container image:
```
docker run -d --env-file .env --name viam-rdk-light viam-rdk-light
```
