
  Pod ~= []Container
  Deployment = Pod

* check what is running on the cluster:
  * `kubectl get pod,deploy`
* deploy Red and Blue
  * `kubectl apply -f deploymentRed.yaml`
* check its status:
  * `kubectl get pod`
  * `kubectl logs <podname>`
  * `kubectl delete pod <podname>`
* there is little bit more into it..
  * `kubectl get replicaset`

- Docs:
  - https://kubernetes.io/docs/concepts/workloads/pods/
  - https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
- Reference:
  - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#deployment-v1-apps
