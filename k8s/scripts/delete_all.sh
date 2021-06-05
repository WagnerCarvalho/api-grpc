#!/bin/sh
##
# Script to just undeploy the MongoDB Service & StatefulSet but nothing else.
##

# Just delete mongod stateful set + mongodb service onlys (keep rest of k8s environment in place)
kubectl delete namespace mongo-grpc
sleep 3
kubectl get namespace


# kubectl delete statefulsets mongod
# kubectl delete services mongodb-service

# Show persistent volume claims are still reserved even though mongod stateful-set has been undeployed
#kubectl get persistentvolumes