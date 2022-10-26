package main

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/bgaechter/ddns/internal/ddnscli"
	"github.com/bgaechter/ddns/pkg/ddnsroute53"
)


func main() {

	ddnscli.LoadConfig()

	// Load the Shared AWS Configuration (~/.aws/config)
	svc := route53.New(session.New())

	myPublicIP := ddnscli.GetPublicIPAddress()
	zones := ddnsroute53.GetHostedZones(svc)

	for _, zone := range zones {
		if strings.Contains(*zone.Name, "taldril.ch") {
			ddnsroute53.CreateOrUpdateDNSEntry(svc, *zone, myPublicIP)
		}
	}
}
