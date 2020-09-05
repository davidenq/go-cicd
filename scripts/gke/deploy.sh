
#!/bin/bash
gcloud config set project devops-flow-testing
gcloud container clusters get-credentials gocicd --region us-east1

# configmap
kubectl apply -f ./deploy/gke/${SERVICE}/configmap.yaml

# deployment
export TAGS=$(gcloud container images list-tags gcr.io/devops-flow-testing/go-cicd --format="get(tags)")
export REVISION_ID=${TAGS::40}
echo "revision id generated: " $REVISION_ID
sed -i -e "s/REVISION_ID/$REVISION_ID/g" ./deploy/gke/${SERVICE}/deployment.yaml
kubectl apply -f ./deploy/gke/${SERVICE}/deployment.yaml
sed -i -e "s/$REVISION_ID/REVISION_ID/g" ./deploy/gke/${SERVICE}/deployment.yaml
rm -rf ./deploy/gke/${SERVICE}/deployment.yaml-e

# service
kubectl apply -f ./deploy/gke/${SERVICE}/service.yaml