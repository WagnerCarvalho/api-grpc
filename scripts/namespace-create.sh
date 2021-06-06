
#!/bin/sh
##
# Script to create namespace in Kubernetes.
##

TMPFILE=$(mktemp)
/usr/bin/openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE
rm $TMPFILE

echo "========================================================================================================="
echo "Create namespace"
echo "========================================================================================================="

kubectl apply -f ../k8s/namespace/resources/mongo-namespace.yaml  --validate=false
kubectl get namespace
sleep 3