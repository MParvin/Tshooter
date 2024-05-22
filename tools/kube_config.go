package tools

import (
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Config struct {
	CurrentContext string
	Contexts       clientcmdapi.Context
	Clusters       clientcmdapi.Cluster
	Users          clientcmdapi.AuthInfo
}

func ReadConfig(filePath string) (*Config, error) {
	return nil, nil
}
