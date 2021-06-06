
#!/bin/sh
##
# Script to delete a Kubernetes project with a StatefulSet running a MongoDB Replica Set, to a local Minikube environment.
##

TMPFILE=$(mktemp)
/usr/bin/openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE
rm $TMPFILE


echo "========================================================================================================="
echo "Delete configMap"
echo "========================================================================================================="
kubectl delete -f ../k8s/services/resources/app-config.yaml 
sleep 3

echo "========================================================================================================="
echo "Delete Secrets"
echo "========================================================================================================="
kubectl delete -f ../k8s/services/resources/app-secrets.yaml
sleep 3

echo "========================================================================================================="
echo "Delete MS Authentication"
echo "========================================================================================================="
kubectl delete -f ../k8s/services/resources/authentication.yaml
sleep 3

echo "========================================================================================================="
echo "Delete MS Api"
echo "========================================================================================================="
kubectl delete -f ../k8s/services/resources/api.yaml
sleep 3