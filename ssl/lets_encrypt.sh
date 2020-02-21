#create a DNS entry via google dns
kubectl create -f kube-dns.yaml

#or use scale to 0 then to 1
kubectl delete pod -l k8s-app=kube-dns -n kube-system

kubectl logs kube-dns-5dbbd9cc58-jzmj5 kubedns -n kube-system
kubectl logs kube-dns-5dbbd9cc58-jzmj5 dnsmasq -n kube-system

kubectl apply -f issuer_letsencrypt.yaml
kubectl apply -f ingress_letencrypt.yaml

kubectl describe certificate chrisge-net-letsencrypt-tls
kubectl describe order chrisge-net-letsencrypt-tls-2576135374-292549555
kubectl describe challenge chrisge-net-letsencrypt-tls-2576135374-292549555-4185108436



