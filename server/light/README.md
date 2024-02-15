# Viam RDK Lightweight Docker Image

This Dockerfile allows you to create a ligthweight Docker image which will download all required binaries when instantiated.
You can configure the container using the ```.env``` file.

Build the container image:
```
docker build -t viam-rdk-light .
```

Update the ```.env``` file with the config command from [app.viam.com](https://app.viam.com).

Run the container image:
```
docker run -d --env-file .env --name viam-rdk-light viam-rdk-light
```

Add a persistent volume:

```viam```&emsp;&emsp;&emsp;&emsp;&emsp;The directory on the Docker host machine

```/root/.viam/```&emsp;Viam default data capture folder. If you change this as part of the smart machine data manager configuration, you will have to update it here too.

```
# Run detached container:

docker run --env-file .env --mount source=viam,target=/root/.viam/ -h viam-rdk --name viam-rdk viam-rdk

# Run container with interactive shell. Setup script setup.sh won't be executed automatically!
# Use this to tweak the image before installing/starting Viam server.

docker run --env-file .env --mount source=viam,target=/root/.viam/ -it -h viam-rdk --name viam-rdk viam-rdk bash
```
