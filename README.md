# EmployeeService Backend

This repository contains the code for a simple backend for the API used in the Quick Start Guide for [WSO2 APK.](https://apk.docs.wso2.com/en/latest/get-started/quick-start-guide/)

## Code Explanation

This repository is a Go-based HTTP server designed for managing employee records. It facilitates CRUD (Create, Read, Update, Delete) operations directly in memory.

There are four endpoints:
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

### Multi Architecture Image

If you wish to build multi-arch images for arm64 and amd64, you can use Docker Buildx.

Note: Docker Buildx is included in Docker 19.03 and later versions. You can check if Buildx is available by running 

```bash
docker buildx version
```

Then run the following commands. 
NOTE: Ensure that you replace {dockerhub-name}, {image-name} and {tag} with your DockerHub name, the image name and the version you wish to tag the image with. This is because multi-arch builds require pushing the image to a registry to store the different architecture versions.

```bash
docker buildx create --name empsvcbuilder --use
docker buildx build --platform linux/amd64,linux/arm64 -t {dockerhub-name}/{image-name}:{tag} . --push
```

## Kubernetes Deployments

### Creating the deployment and service
A sample Kubernetes yaml file for the Service and the Deployment have been provided in the k8s directory of this repository. To apply it, use the following command.

```bash
cd k8s
```

```bash
kubectl apply -f .
```
You can use the following command to see if the pod has spun up.
```bash
kubectl get pods
```

### Test the service locally

To test the backend locally, you have to port-forward the service's port as follows. Do not close the terminal you ran this command in until you are done testing it.

```bash
kubectl port-forward svc/employee-service 9000:80
```

You can now access the service at port 9000 on localhost. Sample cURL commands for each of the endpoints have been provided below.

#### GET request
```bash
curl -X GET http://localhost:9000/employee -k
```

#### POST request
```bash
curl -X POST http://localhost:9000/employee -H "Content-Type: application/json" -d '{"id":"56789123", "name":"John Doe", "department":"Marketing"}' -k
```

#### PUT request
```bash
curl -X PUT http://localhost:9000/employee/56789123 -H "Content-Type: application/json" -d '{"id":"56789123", "name":"John Updated", "department":"Marketing"}' -k
```

#### DELETE request
```bash
curl -X DELETE http://localhost:9000/employee/56789123 -k
```

#### Facing any problems or bugs? Let me know!
Create an issue and let me know if you face any problems.
