package main

/*
SDK: https://github.com/google/google-api-go-client
REF:
 - https://cloud.google.com/compute/docs/reference/beta/disks/list#examples
 - https://cloud.google.com/compute/docs/reference/latest/disks/list#example
 - https://godoc.org/google.golang.org/api/compute/v1#DisksListCall.Filter
*/

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Compute Engine API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/compute
// 2. This sample uses Application Default Credentials for authentication.
//    If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run
//    `gcloud beta auth application-default login`.
//    For more information, see
//    https://developers.google.com/identity/protocols/application-default-credentials
// 3. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
	"encoding/json"
	"flag"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v0.beta"
	"fmt"
	"strings"
)

const (
	maxResults = 10
)

var (
	project = flag.String("project", "hyper-test-142007", "project name")
	region  = flag.String("region", "us-central1", "region name, for address")
	zone    = flag.String("zone", "us-central1-a", "zone name, for disk")
	resource= flag.String("resource", "", "resource name, disk/address")
	action  = flag.String("action", "", "resource operation: insert, list, get, delete")
	name    = flag.String("name", "", "resource name")
	format  = flag.String("format", "table", "output format: table, json")
	filter  = flag.String("filter", "", "filter expression for filtering listed resources")
)

func main() {

	flag.Parse()

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	if *resource == "disk" {
		if *zone == "" {
			fmt.Printf("please specify --zone for disk operation\n")
		}
		//zone range
		switch *action {
		case "insert":
			err = createDisk(ctx, computeService, *zone, *name)
		case "list":
			err = listDisk(ctx, computeService, *zone, *filter)
		case "get":
			err = getDisk(ctx, computeService,*zone, *name)
		case "delete":
			err = deleteDisk(ctx, computeService, *zone, *name)
		default:
			fmt.Printf("please specify --action, valid value is 'insert','get','list','delete'\n")
		}
	} else if *resource == "address" {
		if *region == "" {
			fmt.Printf("please specify --region for address operation\n")
		}
		//region range
		switch *action {
		case "insert":
			err = createAddress(ctx, computeService, *region, *name)
		case "list":
			err = listAddress(ctx, computeService, *region, *filter)
		case "get":
			err = getAddress(ctx, computeService, *region, *name)
		case "delete":
			err = deleteAddress(ctx, computeService, *region, *name)
		default:
			fmt.Printf("please specify --action, valid value is 'insert','get','list','delete'\n")
		}
	} else {
		fmt.Printf("please specify --resource, valid value is 'disk' or 'address'\n")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func printJson(action string, data interface{}) {
	buf, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s: %T - %v\n", action, data, string(buf))
}


func getShortValue(val string) string {
	t := strings.Split(val, "/")
	return t[len(t)-1]
}