# EmployeeService Backend

This repository contains the code for the backend for the API used in the Quick Start Guide for WSO2 APK.

## Code Explanation

This Go code exposes four endpoints:
1. GET /employee - Returns the list of employees
2. POST /employee - Adds an employee to the list
3. PUT /employee/{id} - Updates an employee in the list
4. DELETE /employee/{id} - Deletes an employee in the list

## Docker Image

### Single Architecture Image

These commands will build the docker image suited to the architecture of the machine you use. This is suitable if you wish to build, run and test the image locally.

To build the Docker image, you can use the following command.
```bash
docker build .
```

If you wish to tag it with a specific name, you can use the -t flag. Replace {your-image-name} with the name you wish to give your docker image, and {tag} with the version you wish to tag it. However, ensure that you use the correct name for other places like Kubernetes deployment files.
```bash
docker build -t {your-image-name}:{tag} .
```

## Multi Architecture Image

If you wish to build multi-arch images for arm64 and amd64, you can use Docker Buildx.

Note: Docker Buildx is included in Docker 19.03 and later versions. You can check if Buildx is available by running 

```bash
docker buildx version.
```

Then run the following commands. 
NOTE: Ensure that you replace {dockerhub-name}, {image-name} and {tag} with your DockerHub name, the image name and the version you wish to tag the image with. This is because multi-arch builds require pushing the image to a registry to store the different architecture versions.

```bash
docker buildx create --name empsvcbuilder --use
docker buildx build --platform linux/amd64,linux/arm64 -t {dockerhub-name}/{image-name}:{tag} . --push
```
