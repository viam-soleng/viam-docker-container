# Viam Docker Container Best Practices

Running applications and services inside Docker containers is popular. Within this repository we provide a few tips and tricks around how to use Viam with Docker containers.
The repository contains a section about creating a Docker image with the Viam RDK and how to use it in different ways.
If you are more interested in connecting to a Viam service, have a look at the Viam client section further down in this README.md.

## Viam RDK Docker Container

### Build the Docker Image

The easiest way to get an up to date docker image is to use the Dockerfile in the ```server/light/``` folder. 
This image, when instantiated, will connect to the internet and download/install/start the Viam RDK.

Under some circumstances however you might want to build a container image, which has all the binaries already installed to avoid donloading large amounts of data.
The Dockerfile in the ```server/heavy/``` folder will create an image with a preinstalled Viam server, downloading the files during build time.

```
docker build -t viam-rdk .
```


### Configure the Container [.env]

The ```.env``` file is an easy way to parametrize your Viam Docker container.
All you have to do first is to create a robot at [app.viam.com](https://app.viam.com]) and from the setup instructions for Linux, copy the Config command.

Then in the ```.env``` file replace the existing placeholder config command [key: Config]. 

### Instantiate Viam Server Container

The following command will use the previously built container image ```viam-rdk``` and the updated ```.env``` file and create a running instance of the container.

```
docker run -d --env-file .env viam-rdk
```

Instantiate the "heavy" docker image by providing the ```viam.json``` file at startup.

```
docker run -d -v ./config:/config/ viam-rdk
```

### Container with Persistent Volume

```viam```&emsp;&emsp;&emsp;&emsp;&emsp;The directory on the Docker host machine

```/root/.viam/```&emsp;Viam default data capture folder. If you change this as part of the smart machine data manager configuration, you will have to update it here too.

```
# Run detached container:

docker run --env-file .env --mount source=viam,target=/root/.viam/ -h viam-rdk --name viam-rdk viam-rdk

# Run container with interactive shell. Setup script setup.sh won't be executed automatically!
# Use this to tweak the image before installing/starting Viam server.

docker run --env-file .env --mount source=viam,target=/root/.viam/ -it -h viam-rdk --name viam-rdk viam-rdk bash
```




### Container Port Settings

The Viam RDK by default binds to port ```8080```. This setting can be overriden in the ```viam.json``` config file through [app.viam.com](https://app.viam.com).
Should you have other services conflicting with the default port, you can simply change the port any time. Viam uses WebRTC and signaling servers and thus allows you to connect from anywhere without any changes to the container.

However if the container is not connected to the internet, it may be required to use the local direct URL to establish connectivity. In that case you will have to expose the port of the container.




## Viam Client Applications as Docker Container

Viam provides you with a wide variety of SDKs to write client applications. While you can install the prerequisites on your local system 
to execute these client applications, Docker containers are a very handy tool to test code without changes to your local machine.

The easiest way to get started with a client application is when you have a connected smart machine in [app.viam.com](https://app.viam.com). Simply navigate to the ```Code Sample``` tab and check out the code samples in different languages.

The following instruction will show how easy you can execute these code samples with Docker containers.


### Build a "Go" Client Image

The folder ```client/go/src``` contains the ```client.go``` file. Simply replace the existing code with your own or to get started easily with the sample code from [app.viam.com](https://app.viam.com).
Then build the container by running the following command from within ```client/go/```.

```
docker build -t viam-client .
```

### Instantiate Client Container

Once you have built the container image in the previous step, you can simply instantiate the container and execute your client code.

```
# Run your code automatically
docker run viam-client

# Start your container with a shell and execute the code manually
docker run -it viam-client bash
```
