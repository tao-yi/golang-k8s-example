## Label

```sh
# 为pod资源打标签
$ kubectl label pod nginx-pod version=1.0 -n dev
pod/nginx-pod labeled

# 为pod资源更新标签
$ kubectl label pod nginx-pod version=2.0 -n dev --overwrite
pod/nginx-pod labeled

# 查看标签
$ kubectl get pod nginx-pod  -n dev --show-labels
NAME        READY   STATUS    RESTARTS   AGE   LABELS
nginx-pod   1/1     Running   0          10m   version=2.0

# 筛选标签
$ kubectl get pod -n dev -l version=2.0  --show-labels
NAME        READY   STATUS    RESTARTS   AGE   LABELS
nginx-pod   1/1     Running   0          17m   version=2.0

$ kubectl get pod -n dev -l version!=2.0 --show-labels
No resources found in dev namespace.

#删除标签
$ kubectl label pod nginx-pod version- -n dev
pod/nginx-pod labeled
```
