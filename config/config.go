package config

type ClusterDetail struct {
	CertificateAuthorityData string `mapstructure:"certificate-authority-data"`
	Server                   string `json:"server"`
}
type Cluster struct {
	Name    string        `json:"name"`
	Cluster ClusterDetail `json:"cluster"`
}

type UserDetail struct {
	ClientKeyData         string `mapstructure:"client-key-data"`
	ClientCertificateData string `mapstructure:"client-certificate-data"`
}

type User struct {
	Name string     `json:"name"`
	User UserDetail `json:"user"`
}

type Token struct {
	Jwt string
}

type OrganizationRequest struct {
	CompanyID             string `json:"company_id"`
	CompanyName           string `json:"company_name"`
	OrgName               string `json:"org_name"`
	Description           string `json:"description"`
	ChaincodeCount        uint32 `json:"chaincode_count"`
	PeerCount             uint32 `json:"peer_count"`
	OrdererCount          uint32 `json:"orderer_count"`
	ConsortiumName        string `json:"consortium_name"`
	DbBackend             string `json:"db_backend"`
	ConsortiumDescription string `json:"consortium_description"`
	UserName              string `json:"user_name"`
	UserEmail             string `json:"user_email"`
}
type Configuration struct {
	ApiVersion          string     `json:"apiVersion"`
	Clusters            []*Cluster `json:"clusters"`
	Users               []*User    `json:"users"`
	Token               Token
	OrganizationRequest OrganizationRequest
}
