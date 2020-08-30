
  Staging vs Production - why the hell Helm?

* what if:
  * production secrets need to be different
  * ExternalName service should point to something else
  * a Service should not be deployed on customer's cluster
  * the same port is used in: deployment, service, configMap and ingress
  * we need the same database  deployed multiple times with some changes
* Helm can help with that..
* Helm is basically:
  * templating mechanism for .yaml files
  * repository/package manager of production ready "charts"
  * deployed releases manager

- Templating:
  - [golang templates](https://golang.org/pkg/text/template/)
  - [sprig](http://masterminds.github.io/sprig/)
  - [helm docs](https://helm.sh/docs/intro/quickstart/)
