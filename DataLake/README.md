General Architecture
---
This project is divided into several microservices that are packaged with Docker
and deployed with Kubernetes. Instructions for building and running the
individual microservices can be found below. Instructions for deploying the
project to a GKE cluster are found at the bottom.

Backend Microservices
---
api - Serves JSON data to be ingested by React JS frontend application. To build
    the project run `docker compose up --build api`. The accessible ports can
    be edited in `docker-compose.yml`.
    Endpoints:
    - /transportation_network_trips
    - /building_permits
    - /chicago_ccvi
    - /public_health_stats
    - /covid_19_reports

datapull - Pulls data from Chicago Data Portal using Go and the SODA API. To
    build the project run `docker-compose build --no-cache datapull` or omit the
    `--no-cache` flag to speed things up a bit. To run the service use
    `docker-compose up datapull`.

postgres - Database for the data lake. To run the database locally use
    `docker-compose up postgres`.

Frontend Microservices
---
dashboard - React app to display data stored in the shared database. The 
libraries Material UI and Recharts are used to visualize the data returned from
the api backend service.

Scripts
---
Scripts for development are stored in the "scripts" directory and can be run
with bash. For example, `bash scripts/rebuild-containers.sh`. These may need to
be modified depending on the user's system.

This deployment requires a Docker account and for repositories to have been
created for each microservice image in the project. The repositories should be
named datalake-datapull and datalake-api.

Deployment Dependencies
---
This project has several dependencies for deployment. This section details what
will need to be installed and what accounts created.

Install Docker to manage images. Once this is done, create an account at
https://hub.docker.com/ and create two repositories to hold the images for the
api microservice and datapull microservice. The database service runs on
Postgres and does not need a repository.

Install kubectl and minikube. Once minikube is installed test the installation:
```
minikube start --kubernetes-version=1.31.2 // replace with the kubernetes
minikube ip                                // version you installed
```

And test kubectl with:
```
kubectl get all
```

Below are some helpful commands for debugging and testing your kubernetes
installation
```
kubectl get all / kubectl get services
kubectl describe service/kubernetes
kubectl exec postgres -- ls
kubectl logs <kubernetes resource>
```

The demo project runs on Google Kubernetes Engine (GKE). Create or login to an
existing Google Cloud project and create a GKE cluster. This takes a significant
amount of time, so be patient with this step. While the cluster is being created
, install Google Cloud SDK https://cloud.google.com/sdk/docs/install-sdk

Ensure the gcloud executable is on the PATH before continuing:
```
export PATH=$PATH:/Users/esn2981/google-cloud-sdk/bin
```

Install the Gcloud Auth Plugin:
```
gcloud components install gke-gcloud-auth-plugin
```

Once the plugin is installed, you are ready to deploy the application to the
GKE cluster.

Deploying the App
---
If you have not done so already, create a Google GKE cluster to deploy to with
Google cloud console. This can take some time to complete, so be patient with
this step

Once your cluster has been created, run the following replacing <project-id>,
<region>, and <gke-cluster-name> with the appropriate values from your project
```
./bin/gcloud auth login
./bin/gcloud config set project <project-id>
./bin/gcloud container clusters get-credentials <gke-cluster-name> --zone <zone> --project <project-id>
```

Example with values:
```
./bin/gcloud auth login
./bin/gcloud config set project msds432-441922
./bin/gcloud container clusters get-credentials msds432-datalake --zone us-central1 --project msds432-441922
```

Test the cluster connection to ensure everything worked
```
kubectl cluster-info
kubectl config current-context
```

If something goes wrong, switch back to minikube with
```
kubectl config get-contexts
kubectl config use-context minikube
```
and then try again

Build the images, retag for publishing to Docker Hub, and push to registry with
the following (replace ericesn with your docker hub profile):
```
docker compose build --no-cache
docker tag datalake-datapull ericesn/datalake-datapull
docker tag datalake-api ericesn/datalake-api
docker login
docker push ericesn/datalake-datapull
docker push ericesn/datalake-api
```

To deploy the app with kubernetes, use the following commands (switch to the
kubernetes directory and run the deployment scripts)
```
cd kubernetes
kubectl apply -f postgres-data-persistentvolumeclaim.yaml,postgres-cm1-configmap.yaml,postgres-service.yaml,api-service.yaml,postgres-deployment.yaml,api-deployment.yaml,datapull-deployment.yaml
```

Once this has completed the microservices are deployed to the GKE cluster.

Local Testing
---
To test the api locally, run the deployment scripts with minikube as the target
cluster.

Ensure minikube is up with `minikube ip`, then switch to the kubernetes
directory and run the scripts with the below commands
```
cd kubernetes
kubectl apply -f postgres-data-persistentvolumeclaim.yaml,postgres-cm1-configmap.yaml,postgres-service.yaml,api-service.yaml,postgres-deployment.yaml,api-deployment.yaml,datapull-deployment.yaml
```

Next, create a tunnel to the api service
```
minikube service api
```

And now the API endpoints will be available on the browser. The endpoints can be
found in `api/cmd/main.go`.
