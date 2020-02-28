#without ingress controller

helm install  nginx-ing stable/nginx-ingress


kubectl run nginx --image=nginx

kubectl expose deployment nginx --port 80 --target-port=80  --type=NodePort

kubectl apply -f ingress.yaml

kubectl get ing nginx -o yaml


echo 35.231.119.153 chrisge.net | sudo tee -a /etc/hosts

cat /etc/hosts | tail -n 1

curl -k https://example.com

#only work with minikube
#curl --cacert tls.crt  https://example.com

#The options on this curl command will provide verbose output, following any redirects, show the TLS headers in the output, and not error on insecure certificates. With nginx-ingress-controller, the service will be available with a TLS certificate, but it will be using a self-signed certificate provided as a default from the nginx-ingress-controller. Browsers will show a warning that this is an invalid certificate. This is expected and normal, as we have not yet used cert-manager to get a fully trusted certificate for our site.
curl -kivL -H 'Host: example.your-domain.com' 'http://35.199.164.14'