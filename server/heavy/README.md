# Viam RDK Heavyweight Docker Image

This Dockerfile allows you to create a Docker image which includes the Viam server binary.
You can configure the container / Viam RDK using the ```./config/viam.json``` file.

The easiest way to get started is to create a robot at [app.viam.com](https://app.viam.com]) and then copy the configuration from the Setup tab into ```./config/viam.json```.

Build the container image:
```
docker build -t viam-rdk-heavy .
```

If you haven't done it yet, update the ```./config/viam.json``` file.

Instantiate the "heavy" docker image by providing the ```viam.json``` file at startup.

```
docker run -d -v ./config:/config/ viam-rdk-heavy
```

Add a persistent volume

```viam```&emsp;&emsp;&emsp;&emsp;&emsp;The directory on the Docker host machine

```/root/.viam/```&emsp;Viam default data capture folder. If you change this as part of the smart machine data manager configuration, you will have to update it here too.

```
# Run detached container:

docker run -d --mount source=viam,target=/root/.viam/ --name viam-rdk-heavy viam-rdk-heavy

# Run container with interactive shell. Setup script setup.sh won't be executed automatically!
# Use this to tweak the image before installing/starting Viam server.

docker run --env-file .env --mount source=viam,target=/root/.viam/ -it -h viam-rdk --name viam-rdk viam-rdk bash
```
