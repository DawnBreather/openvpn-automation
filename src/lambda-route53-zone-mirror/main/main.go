package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)


var vpnCidr =  os.Getenv("VPN_CIDR")
var baseZoneId = os.Getenv("BASE_ZONE_ID")
var mirrorZoneId = os.Getenv("MIRROR_ZONE_ID")
var awsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
var awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
var runInterval, err = strconv.Atoi(os.Getenv("RUN_INTERVAL"))

func main(){
	if err == nil {
		for {
			run()
			log.Printf("INFO: Interval %d seconds", runInterval)
			time.Sleep(time.Duration(runInterval) * time.Second)
		}
	} else {
		log.Printf("ERROR: Error parsing RUN_INTERVAL environment variable:\n%v", err)
	}
}

func run() {
	//flag.Parse()

	if vpnCidr == "" || baseZoneId == "" || mirrorZoneId == "" || awsAccessKeyId == "" || awsSecretAccessKey == "" {
		fmt.Println("Not all required configuration parameters passed over environment variables")
		fmt.Println("VPN_CIDR")
		fmt.Println("BASE_ZONE_ID")
		fmt.Println("MIRROR_ZONE_ID")
		fmt.Println("AWS_ACCESS_KEY_ID")
		fmt.Println("AWS_SECRET_ACCESS_KEY")
	}

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	_, vpnNet, _ := net.ParseCIDR(vpnCidr)

	svc := route53.New(sess)

	records := listAll(svc, baseZoneId) //"Z00362693229M1IVYVTVZ")
	for _, r := range records.ResourceRecordSets {

		var status = ""

		if *r.Type == "A" || *r.Type == "CNAME" || *r.Type == "MX" || *r.Type == "TXT" || *r.Type == "AAAA" {
			// if it's a simple A-record (not alias)
			if *r.Type == "A" && r.AliasTarget == nil {

				existentRecord := listOne(svc, mirrorZoneId, "A", *r.Name)

				// if entry already exists in target zone
				if existentRecord != nil {

					// if entry doesn't belong to VPN subnet
					var vpnEntry = false
					for _, v := range existentRecord.ResourceRecords {
						if vpnNet.Contains(net.ParseIP(*v.Value)) {
							vpnEntry = true
						}
					}

					if !vpnEntry {
						//fmt.Println(existentRecord.ResourceRecords)
						status = create(svc, r, mirrorZoneId)
					} else {
						status = "skipped due to existing vpn entry"
					}

					fmt.Printf("%s: %s\n", *r.Name, status)
				}
			}
		}
	}
}

func create(svc *route53.Route53, r *route53.ResourceRecordSet, zoneId string) string {

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("UPSERT"), // Required
					ResourceRecordSet: r,
				},
			},
			Comment: aws.String("Sample update."),
		},
		HostedZoneId: aws.String(zoneId), // Required
	}

	if r.AliasTarget != nil {
		if *r.AliasTarget.HostedZoneId == baseZoneId {
			r.AliasTarget.HostedZoneId = aws.String(zoneId)
		}
	}

	/*resp*/_, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(r.AliasTarget)

		return "failure"
	}

	// Pretty-print the response data.
	//fmt.Println("Change Response:")
	//fmt.Println(resp)

	return "success"
}


func listAll(svc *route53.Route53, zoneId string) *route53.ListResourceRecordSetsOutput {
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneId), // Required
		MaxItems:              aws.String("1024"),
		//StartRecordIdentifier: aws.String("Sample update."),
		//StartRecordName:       aws.String("com."),
		//StartRecordType:       aws.String("CNAME"),
	}
	respList, err := svc.ListResourceRecordSets(listParams)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return respList
}

func listOne(svc *route53.Route53, zoneId, recordType, name string) *route53.ResourceRecordSet {
	listParams := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneId), // Required
		MaxItems:              aws.String("1"),
		//StartRecordIdentifier: aws.String("Sample update."),
		StartRecordName:       aws.String(name),
		StartRecordType:       aws.String(recordType),
	}
	respList, err := svc.ListResourceRecordSets(listParams)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if len(respList.ResourceRecordSets) > 0 {
		return respList.ResourceRecordSets[0]
	} else {
		return nil
	}
}
