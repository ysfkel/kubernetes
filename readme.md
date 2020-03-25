
kubectl config current-context does the trick (it outputs little bit more, like project name, region, etc., but it should give you the answer you need).


In order to interact with a specific cluster, you only need to specify the cluster name as a context in kubectl:

kubectl cluster-info --context kind-kind
kubectl cluster-info --context kind-2