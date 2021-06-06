
#!/bin/sh
##
# Script to delete deployment and service in Kubernetes.
##

TMPFILE=$(mktemp)
/usr/bin/openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE
rm $TMPFILE

echo "========================================================================================================="
echo "Delete Secret for MongoDb"
echo "========================================================================================================="
kubectl delete -f ../k8s/mongodb/resources/mongo-secret.yaml
sleep 3

echo "========================================================================================================="
echo "Delete ConfigMap for MongoDb"
echo "========================================================================================================="
kubectl delete -f ../k8s/mongodb/resources/mongo-configmap.yaml
sleep 3

echo "========================================================================================================="
echo "Delete Deployment and Service for MongoDb"
echo "========================================================================================================="
kubectl delete -f ../k8s/mongodb/resources/mongo-deployment.yaml
sleep 3

echo "========================================================================================================="
echo "Delete Deployment and Service for Mongo Express"
echo "========================================================================================================="
kubectl delete -f ../k8s/mongodb/resources/mongo-express-deployment.yaml
sleep 3

echo "========================================================================================================="
echo "Verify MongoDb and Mongo Express"
echo "========================================================================================================="
kubectl get all -n mongo-grpc | grep mongo
sleep 3