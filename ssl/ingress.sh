#without ingress controller

helm install  nginx-ing stable/nginx-ingress


kubectl run nginx --image=nginx

kubectl expose deployment nginx --port 80

kubectl apply -f ingress.yaml

kubectl get ing nginx -o yaml


echo 35.231.119.153 chrisge.net | sudo tee -a /etc/hosts

cat /etc/hosts | tail -n 1

curl -k https://example.com

#only work with minikube
#curl --cacert tls.crt  https://example.com