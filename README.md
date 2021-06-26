# GRPC With Golang
MS REST to integration with GRPC Server

>### Requires:
```
1. Docker
2. Docker-Compose
3. Minikube (v1.20.0)
4. Kubectl
5. Golang (go1.15.1)
6. Protoc
```

# Build Proto Files
>### Generate file proto
```
$ cd api-grpc
$ protoc -I=./messages --go_out=plugins=grpc:. ./messages/*.proto
```

# Start Application
>### Start MongoDB
```
$ cd api-grpc
$ docker-compose up -d
```

>### Build Services
```
$ cd api-grpc
$ ./build.sh 
```

>### Start Server GRPC
```
$ cd api-grpc/authentication
$ ./authsvc-local
```

>### Start MS Api
```
$ cd api-grpc/api
$ ./apisvc-local
```

# OR

# Start Kubernets
>### Start Minikube
```
$ minikube start
$ eval $(minikube docker-env)
```
>### Dashboard Minikube
```
$ minikube dashboard
```

# Create Image Docker
```
$ cd api-grpc/k8s/dock
$ ./build.sh
```

# Provision Infra
>### 1. Create NameSpace
```
$ cd api-grpc/scripts
$ ./namespace-create.sh 
```

>### 2. Create Mongo
```
$ cd api-grpc/scripts
$ ./namespace-create.sh 
```

>### 3. Create User Mongo
```
$ kubectl get pods -n mongo-grpc
$ kubectl exec -it mongodb-deployment-8f6675bc5-v4q7r /bin/bash -n mongo-grpc
```
```
$ mongo -u username -p password --authenticationDatabase admin
```
```
$ show users
$ use api_grpc
$ db.createUser({user: 'user', pwd: 'password', roles:[{ 'role': 'readWrite', 'db': 'api_grpc' }] });
$ show users
```

>### 4. Create Services
```
$ cd api-grpc/scripts
$ ./service-create.sh
```

>### 5. Start Tunel
```
$ minikube tunnel
```

# Endpoints
>### Post Signup
```
curl --location --request POST 'http://localhost:9000/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Super Admin",
    "email": "admin@admin.com.br",
    "password": "12345"
}'
```

>### Post Signin
```
curl --location --request POST 'http://localhost:9000/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "admin@admin.com.br",
    "password": "12345"
}'
```

>### Get User by Id
```
curl --location --request GET 'http://localhost:9000/users/${id}' \
--header 'Authorization: Bearer ${token}'
```

>### Get Users All
```
curl --location --request GET 'http://localhost:9000/users/' \
--header 'Authorization: Bearer ${token}'
```

>### Put User
```
curl --location --request PUT 'http://localhost:9000/users/${id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Admin"
}'
```

>### Delete User
```
curl --location --request DELETE 'http://localhost:9000/users/${id}' \
--header 'Authorization: Bearer ${token}'
```


