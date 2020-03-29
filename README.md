# Backend API

Backend server of momocomics.com

**Complete**
  * Cloud Build setup
  * Helm Chart
  * Rest API    
  * gRPC
  * Firestore
  * Skaffold

**Build and Deploy services**
 
  ***Option 1***
  * cd grpc-server
  * gcloud builds submit .
  * helm install grpc-server helm/
  
  ***Option 2***
  * skaffold run

    