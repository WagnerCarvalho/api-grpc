
#!/bin/sh
##
# Script to deploy a Kubernetes project with a StatefulSet running a MongoDB Replica Set, to a local Minikube environment.
##

TMPFILE=$(mktemp)
/usr/bin/openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE
rm $TMPFILE

echo "========================================================================================================="
echo "Create configMap"
echo "========================================================================================================="
kubectl apply -f ../k8s/services/resources/app-config.yaml
kubectl get configmap -n mongo-grpc
sleep 3

echo "========================================================================================================="
echo "Create Secrets"
echo "========================================================================================================="
kubectl apply -f ../k8s/services/resources/app-secrets.yaml
kubectl get secrets -n mongo-grpc
sleep 3

echo "========================================================================================================="
echo "Create Deployment and Service to MS Authentication"
echo "========================================================================================================="
kubectl apply -f ../k8s/services/resources/authentication.yaml
kubectl get pods -n mongo-grpc
sleep 3

echo "========================================================================================================="
echo "Create Deployment and Service to MS Api"
echo "========================================================================================================="
kubectl apply -f ../k8s/services/resources/api.yaml
kubectl get pods -n mongo-grpc
sleep 3