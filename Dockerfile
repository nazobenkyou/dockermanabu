# Base image, the container needs an image which it runs from
# you can build any image, upload it to docker registry and use it
# in this example we will use alpine image
FROM alpine:3.10

# You can install a lot of dependencies here, the RUN command will execute any command inside
# the container
RUN apk --update add ca-certificates

# we specify the workspace, this workspace is like our home directory, if it 
# does not exist, it will create it
WORKDIR /app

# because we are using a dockerignore file, we will copy ONLY the binary file
# created from the go build command
# the file will be copied into the WORKDIR directory
COPY . .

# Exposes the port outside the container
EXPOSE 8080

# We specify volume directory to share it with our local machine, this volume will be created in build
# phase, but we still need to specify the path when running it
VOLUME [ "/data" ]

# Although we set up the workdir, to execute the app we need to specify the absolute path
CMD ["/app/dockermanabu"]
