This go file imports pachyderm v1.13 and v2.4, build doesn't work.

```shell
go build .
```

This is due to github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b being inompatible with v1.2.0. 
There is also a problem with k8s dependency, but this one might be fixable
