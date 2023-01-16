## Namespace

Namespace 是 kubernetes 系统中的一种非常重要资源，它的主要作用是用来实现多套环境的资源隔离或者多租户的资源隔离。

默认情况下，kubernetes 集群中的所有的 Pod 都是可以相互访问的。但是在实际中，可能不想让两个 Pod 之间进行互相的访问，那此时就可以将两个 Pod 划分到不同的 namespace 下。kubernetes 通过将集群内部的资源分配到不同的 Namespace 中，可以形成逻辑上的"组"，以方便不同的组的资源进行隔离使用和管理。

kubernetes 在集群启动之后，会默认创建几个 namespace

```
[root@master ~]# kubectl  get namespace
NAME              STATUS   AGE
default           Active   45h     #  所有未指定Namespace的对象都会被分配在default命名空间
kube-node-lease   Active   45h     #  集群节点之间的心跳维护，v1.13开始引入
kube-public       Active   45h     #  此命名空间下的资源可以被所有人访问（包括未认证用户）
kube-system       Active   45h     #  所有由Kubernetes系统创建的资源都处于这个命名空间
```

```sh
# 1 查看所有的ns
$ kubectl get ns

# 2 查看指定的ns kubectl get ns ns名称
$ kubectl get ns default

# 3 指定输出格式  命令：kubectl get ns ns名称  -o 格式参数
# kubernetes支持的格式有很多，比较常见的是wide、json、yaml
$ kubectl get ns default -o yaml

# 4 查看ns详情  命令：kubectl describe ns ns名称
$ kubectl describe ns default
```

```sh
# 创建namespace
$ kubectl create ns dev

# 删除namespace
$ kubectl delete ns dev
```
