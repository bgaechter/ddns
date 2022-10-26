package ddnsroute53

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/route53"
)


func GetHostedZones(svc *route53.Route53) []*route53.HostedZone {

	hostedZones, err := svc.ListHostedZones(&route53.ListHostedZonesInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, zone := range hostedZones.HostedZones {
		log.Printf("%s: %s", *zone.Name, *zone.Id)
	}
	return hostedZones.HostedZones
}

func CreateOrUpdateDNSEntry(svc *route53.Route53, hostedZone route53.HostedZone, targetIP string) {
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String("dyn." + *hostedZone.Name),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(targetIP),
							},
						},
						TTL:  aws.Int64(60),
						Type: aws.String("A"),
					},
				},
			},
			Comment: aws.String("Dynamic DNS Entry for " + targetIP),
		},
		HostedZoneId: aws.String(*hostedZone.Id),
	}

	result, err := svc.ChangeResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				log.Fatal(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				log.Fatal(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				log.Fatal(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				log.Fatal(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				log.Fatal(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				log.Fatal(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Fatal(err.Error())
		}
		log.Println(result)
	}
}
