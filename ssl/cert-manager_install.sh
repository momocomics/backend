#self signed
kubectl apply --validate=false -f https://raw.githubusercontent.com/jetstack/cert-manager/v0.13.1/deploy/manifests/00-crds.yaml

#kubectl create namespace cert-manager

helm repo add jetstack https://charts.jetstack.io

helm repo update

helm install \
  cert-manager jetstack/cert-manager \
  --set rbac.create=false  \
  --version v0.13.1
#  --namespace cert-manager \



openssl genrsa -out ca.key 2048


cp /etc/ssl/openssl.cnf openssl_cp.cnf

openssl req -x509 -new -nodes -key ca.key -sha256 -subj "/CN=sampleissuer.local" -days 1024 -out ca.crt -extensions v3_ca -config openssl_cp.cnf

kubectl create secret tls ca-key-pair --key=ca.key  --cert=ca.crt

kubectl get secret

kubectl apply -f issuer.yaml

kubectl get issuer

kubectl apply -f certificate.yaml

kubectl get certificate

kubectl describe certificate chrisge-net

