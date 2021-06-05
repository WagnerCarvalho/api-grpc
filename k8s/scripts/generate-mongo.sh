
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
kubectl delete -f ../mongodb/resources/mongo-secret.yaml
kubectl apply -f ../mongodb/resources/mongo-secret.yaml
kubectl get secret -n mongo-grpc
sleep 3

kubectl delete -f ../mongodb/resources/mongo-deployment.yaml
kubectl apply -f ../mongodb/resources/mongo-deployment.yaml
kubectl get all -n mongo-grpc | grep mongo
sleep 3

kubectl delete -f ../mongodb/resources/mongo-configmap.yaml
kubectl apply -f ../mongodb/resources/mongo-configmap.yaml
sleep 3

kubectl delete -f ../mongodb/resources/mongo-express-deployment.yaml
kubectl apply -f ../mongodb/resources/mongo-express-deployment.yaml
sleep 3

kubectl get all - n mongo-grpc | grep mongo
sleep 3