
  When ENVs are not enough - Configmap

* run `kubectl logs -f --tail=10 <RedPodName>`
* `Blue` uses file with encryption configuration
  * bad thing: it is embedded in the image
* ConfigMap to the rescue:
  * use it as ENV
  * mount is as file - let's talk volumes briefly
* provide our own configuration to `Blue`
* configure `Red` to use the same configuration
  * red image: kskitek/k8sfirststeps:red-v2

- Docs: https://kubernetes.io/docs/concepts/configuration/configmap/
- Reference: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#configmap-v1-core
