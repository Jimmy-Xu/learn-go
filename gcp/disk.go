package main

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
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v0.beta"
)

var (
	project = flag.String("project", "hyper-test-142007", "project name")
	zone    = flag.String("zone", "asia-east1-a", "zone name")
	disk    = flag.String("disk", "test_disk_api", "disk name")
	create  = flag.Bool("create", false, "create new disk or not")
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

	err = listDisk(ctx, computeService)
	if err != nil {
		log.Fatal(err)
	}
}

func listDisk(ctx context.Context, computeService *compute.Service) error {
	req := computeService.Disks.List(*project, *zone)
	if err := req.Pages(ctx, func(page *compute.DiskList) error {
		fmt.Printf("there are %v disk in zone:%v\n", len(page.Items), *zone)
		for _, disk := range page.Items {
			diskJson, err := json.MarshalIndent(disk, "", "\t")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", string(diskJson))
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
