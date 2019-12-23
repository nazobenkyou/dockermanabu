# Base image, the container needs an image which it runs from
# you can build any image, upload it to docker registry and use it
# in this example we will use alpine image
FROM alpine:3.10 AS req-image

# The AS syntax is for building multi stage docker images
# Our production image should need the essential

# You can install a lot of dependencies here, the RUN command will execute any command inside
# the container
RUN apk --update add ca-certificates

# we specify the workspace, this workspace is like our home directory, if it 
# does not exist, it will create it
WORKDIR /src

# because we are using a dockerignore file, we will copy ONLY the binary file
# created from the go build command
# the file will be copied into the WORKDIR directory
COPY . .

# This will do the magic and will copy everything from the image we defined
# before, so before we built the image it was around 15MB, now it could be around 7MB
FROM scratch
COPY --from=req-image /src /app

# Exposes the port outside the container
EXPOSE 8080

# Although we set up the workdir, to execute the app we need to specify the absolute path
CMD ["/app/dockermanabu"]
