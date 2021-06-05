
#!/bin/sh
##
# Script to deploy a Kubernetes project with a StatefulSet running a MongoDB Replica Set, to a local Minikube environment.
##

# Create keyfile for the MongoD cluster as a Kubernetes shared secret
TMPFILE=$(mktemp)
/usr/bin/openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE
rm $TMPFILE

# Create mongodb service with mongod stateful-set
# TODO: Temporarily added no-valudate due to k8s 1.8 bug: https://github.com/kubernetes/kubernetes/issues/53309

kubectl delete -f ../services/resources/app-config.yaml 
kubectl apply -f ../services/resources/app-config.yaml
kubectl get configmap -n mongo-grpc
sleep 3

kubectl delete -f ../services/resources/app-secrets.yaml
kubectl apply -f ../services/resources/app-secrets.yaml
kubectl get secrets -n mongo-grpc
sleep 3

kubectl delete -f ../services/resources/authentication.yaml
kubectl apply -f ../services/resources/authentication.yaml
kubectl get pods -n mongo-grpc
sleep 3