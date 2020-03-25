package kubernetes

type IKubernetesService interface {
	CreateCluster(clusterName string) error
	DestroyCluster(clusterName string) error
	GetClusterIP(clusterName string) (string, error)
	GetCertificateAuthority(clusterName string) (string, error)
	GetClientCertificate(clusterName string) (string, error)
	GetClientKey(clusterName string) (string, error)
}
