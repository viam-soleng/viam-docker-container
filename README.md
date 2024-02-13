# Viam Docker Image

A simple dockerfile to run viam server as a container.

## Build Viam Server Docker Image

```
docker build -t viam-rdk .
```

## Configure Viam Server Image [.env]

In the ```.env``` file paste the Viam config command from app.viam.com. 

## Run Viam Server Container

```
docker run --env-file .env viam-rdk
```


## Build Client Image

```
docker build -t viam-client .
```

## Run Client Container

```
docker run viam-client
```
