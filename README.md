# Viam Docker Container Best Practices

Running applications and services inside Docker containers is popular. Within this repository we provide a few tips and tricks around how to use Viam with Docker containers.
The repository contains a section about creating a Docker image with the Viam RDK and how to use it in different ways.
If you are more interested in connecting to a Viam service, have a look at the Viam client section further down in this README.md.

## Viam RDK Docker Container






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
