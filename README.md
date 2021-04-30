# Example of a Dockerized Go Application

A simple Dockerized Go App

## Docker Images

Build development image:

```
docker build --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) -t mathapp-development .
```

Run container:

```
docker run -it --rm -p 8010:8010 -v $PWD/src:/go/src/mathapp mathapp-development
```
