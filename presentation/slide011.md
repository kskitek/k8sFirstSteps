
  Help me to find my container - Services

* run `kubectl delete pod -l app=rvb`
* why containers cannot connect to each other now?
* let's deploy a service:
  * `kubectl apply -f serviceRed.yaml`
  * `kubectl get svc,endpoints`
* types of services:
  * ClusterIP
  * NodePort
  * LoadBalancer
  * ExternalName
  * Headless
  * and not really a Service type: Ingress
* Since we are talking about services - `kubectl port-forward ..`

- Docs: https://kubernetes.io/docs/concepts/services-networking/service/
- Reference: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#service-v1-core
