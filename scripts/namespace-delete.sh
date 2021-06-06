#!/bin/sh
##
# Script to delete namespace.
##

echo "========================================================================================================="
echo "Delete namespace"
echo "========================================================================================================="
kubectl delete namespace mongo-grpc
sleep 3
kubectl get namespace