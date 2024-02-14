# Viam in Docker Container Best Practices

Running applications and service inside Docker containers is popular. Within this repository we provide a few tips and tricks around how to use Viam with Docker containers.
The repository contains a section about creating a Docker image with the Viam RDK and how to use it in different ways.
If you are more interested in connecting to a Viam service, have a look at the Viam client section further down in this README.md.

## Build and Run the Viam RDK Docker Container


### Build the Docker Image

```
docker build -t viam-rdk .
```

### Configure the Container [.env]

The ```.env``` file is an easy way to parametrize your Viam Docker container.
All you have to do first is to create a robot at [app.viam.com](https://app.viam.com]) and from the setup instructions for Linux, copy the Config command.

Then in the ```.env``` file replace the existing placeholder config command [key: Config]. 

### Instantiate Viam Server Container

The followin command will use the previously built container image ```viam-rdk``` and the updated ```.env``` file and create a running instance of the container.

```
docker run -d --env-file .env viam-rdk
```

## Instantiate Viam Server Container with Persistent Volume

```viam```:          The directory on the Docker host machine

```/root/.viam/```:  Viam default data capture folder. If you change this as part of the smart machine data manager configuration, you will have to update it here too.

```
# Run detached container:
docker run --env-file .env --mount source=viam,target=/root/.viam/ viam-rdk

# Run container with interactive shell. Setup script setup.sh won't be executed automatically!
# Use this to tweak the image before installing/starting Viam server
docker run --env-file .env --mount source=viam,target=/root/.viam/ -it viam-rdk bash
```


## Build and Run Viam Client Applicatinos as Docker Container



### Build Client Image

```
docker build -t viam-client .
```

### Instantiate Client Container

```
docker run viam-client
```