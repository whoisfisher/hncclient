package hncclient

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/whoisfisher/hncclient/versioned"
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
	aliveHost = "192.168.111.100"
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
		Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImtrcUdScWx1QS1LbFpicTl3cjhDN2xVd080Q185TmNSaVVrVEtOcUlTbUkifQ",
	})
	if err != nil {
		return client, err
	}
	return client, nil
}

func main() {
	client, err := getMonitorClient("mm")
	if err != nil {
		_ = fmt.Sprintf("====%s", err.Error())
	}
	snps, err := client.HNCV1Alpha2().SubnamespaceAnchors("parent-a").Get(context.TODO(), "child-b", metav1.GetOptions{})
	if err != nil {
		_ = fmt.Sprintf("====%s", err.Error())
	}
	_ = fmt.Sprintf("***%v", snps)
}
