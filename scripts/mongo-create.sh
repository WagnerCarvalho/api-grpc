
#!/bin/sh
##
# Script to deploy a Kubernetes project with a Deployment and Service running a MongoDB Replica Set, to a local Minikube environment.
##

TMPFILE=$(mktemp)
/usr/bin/openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE
rm $TMPFILE

echo "========================================================================================================="
echo "Create Secret for MongoDb"
echo "========================================================================================================="
kubectl apply -f ../k8s/mongodb/resources/mongo-secret.yaml
kubectl get secret -n mongo-grpc
sleep 3

echo "========================================================================================================="
echo "Create ConfigMap for MongoDb"
echo "========================================================================================================="
kubectl apply -f ../k8s/mongodb/resources/mongo-configmap.yaml
sleep 3

echo "========================================================================================================="
echo "Create Deployment and Service for MongoDb"
echo "========================================================================================================="
kubectl apply -f ../k8s/mongodb/resources/mongo-deployment.yaml
kubectl get all -n mongo-grpc | grep mongo
sleep 3

echo "========================================================================================================="
echo "Create Deployment and Service for Mongo Express"
echo "========================================================================================================="
kubectl apply -f ../k8s/mongodb/resources/mongo-express-deployment.yaml
sleep 3

echo "========================================================================================================="
echo "Verify MongoDb and Mongo Express"
echo "========================================================================================================="
kubectl get all -n mongo-grpc | grep mongo
sleep 3