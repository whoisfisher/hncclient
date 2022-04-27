package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/whoisfisher/hncclient"
	"github.com/whoisfisher/hncclient/versioned"
	"github.com/whoisfisher/hncclient/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type Host string

type Config struct {
	Hosts []Host
	Token string
}

func NewMonitoringClient(c *Config) (*versioned.Clientset, error) {
	var aliveHost Host
	aliveHost = "192.168.111.100:6443"
	kubeConf := &rest.Config{
		Host:        string(aliveHost),
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	client, err := versioned.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new monitoring client with config failed: %v", err))
	}
	return client, nil
}

func getMonitorClient(clusterName string) (*versioned.Clientset, error) {
	var client *versioned.Clientset
	client, err := NewMonitoringClient(&Config{
		Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImtrcUdScWx1QS1LbFpicTl3cjhDN2xVd080Q185TmNSaVVrVEtOcUlTbUkifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrby1hZG1pbi10b2tlbi0yZDJ4ZyIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJrby1hZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6Ijc2YTIzNWRlLWEzZWEtNDQwMy05MjNhLTdiYTFkMThhMDkxNyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTprby1hZG1pbiJ9.bqoeOq_WttP8bTzajW4lZkJ33EC0SIW6Rp5lSTEoW3nirQKtdZGKo7ldcLZWrr9gQSAYspoUzFWUKppfiyPmPR0HQAB_MSU7ODd_uTJtEmSRrxiN69Rl7HAafRA61GTOftH6lwhoRo-CmofBH6VpaYvL5Y-wj_3fuQTVwSiOwqk3Y6oy1SWnYYFQ2I8ng0f03Hicu5T7OJNbbvrza6DIem3ctS1iPf_ElWCaL_qwV0VrugyaaiJIlJNz1CWbVzu5cpoPgOkin6OSBCIJc60aPuaENGO0A5s1lWAH5Ggr12HsXacg6JDDT_xesMHDy4P3R1Pg7ZgmKcLC44ivRjlLNg",
	})
	if err != nil {
		return client, err
	}
	return client, nil
}

func main() {
	hncclient.AddToScheme(scheme.Scheme)
	client, err := getMonitorClient("mm")
	if err != nil {
		errMsg := fmt.Sprintf("====%s", err.Error())
		fmt.Println(errMsg)
	}
	snps, err := client.HNCV1Alpha2().SubnamespaceAnchors("parent-a").Get(context.TODO(), "child-b", metav1.GetOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("====%s", err.Error())
		fmt.Println(errMsg)
	}
	res := fmt.Sprintf("***%v", snps)
	fmt.Println(res)
}
