Output diff from `kubectl get * -w -o yaml`


example:
1. run `k get po db-tidb-0 -o yaml -w | ydiff`
2. delete the po in another terminal.

see:
```
...
...
tidbfull/db-tidb-0 diff:
  &unstructured.Unstructured{
  	Object: map[string]interface{}{
  		"apiVersion": string("v1"),
  		"kind":       string("Pod"),
  		"metadata": map[string]interface{}{
  			"annotations":                map[string]interface{}{"kubernetes.io/psp": string("eks.privileged"), "prometheus.io/path": string("/metrics"), "prometheus.io/port": string("10080"), "prometheus.io/scrape": string("true"), ...},
  			"creationTimestamp":          string("2022-06-17T09:17:40Z"),
+ 			"deletionGracePeriodSeconds": int64(30),
+ 			"deletionTimestamp":          string("2022-06-17T09:31:43Z"),
  			"generateName":               string("db-tidb-"),
  			"labels":                     map[string]interface{}{"app.kubernetes.io/component": string("tidb"), "app.kubernetes.io/instance": string("db"), "app.kubernetes.io/managed-by": string("tidb-operator"), "app.kubernetes.io/name": string("tidb-cluster"), ...},
  			"name":                       string("db-tidb-0"),
  			"namespace":                  string("tidbfull"),
  			"ownerReferences":            []interface{}{map[string]interface{}{"apiVersion": string("apps.pingcap.com/v1"), "blockOwnerDeletion": bool(true), "controller": bool(true), "kind": string("StatefulSet"), ...}},
- 			"resourceVersion":            string("1967374"),
+ 			"resourceVersion":            string("1971677"),
  			"uid":                        string("93a96d2a-a62e-4734-a96f-dd364a91e439"),
  		},
...
...
```
