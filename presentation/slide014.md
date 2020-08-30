
  Where is my data? - Volumes

* `kubectl exec -it <redPodName> ls /var/log/red.log`
* `kubectl delete pod <redPodName>`
* why old logs are no longer there?
* deploy `redPVC.yaml`
* `kubectl get pv,pvc`
* configure `Red` to use new volume
* again: `kubectl get pv,pvc`

- Docs:
  - https://kubernetes.io/docs/concepts/storage/persistent-volumes/
  - https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims
- Reference:
  - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#persistentvolumeclaim-v1-core
