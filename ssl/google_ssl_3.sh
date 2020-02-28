 kubectl apply -f google_ssl.yaml

 kubectl describe managedcertificate chrisge-net-certificate

kubectl apply -f ingress_google_ssl.yaml


kubectl delete managedcertificate chrisge-net-certificate