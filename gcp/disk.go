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
)

var (
	project = flag.String("project", "hyper-test-142007", "project name")
	zone    = flag.String("zone", "asia-east1-a", "zone name")
	disk    = flag.String("disk", "test-disk", "disk name")
	filter  = flag.String("filter", "", "filter expression for filtering listed resources")
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

	err = createDisk(ctx, computeService)
	if err != nil {
		log.Fatal(err)
	}

	err = listDisk(ctx, computeService)
	if err != nil {
		log.Fatal(err)
	}

	err = getDisk(ctx, computeService)
	if err != nil {
		log.Fatal(err)
	}

	err = deleteDisk(ctx, computeService)
	if err != nil {
		log.Fatal(err)
	}

}

func createDisk(ctx context.Context, computeService *compute.Service) error {
	rb := &compute.Disk{
		SizeGb:      1,
		Name:        *disk,
		Description: "test",
	}
	resp, err := computeService.Disks.Insert(*project, *zone, rb).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("create disk", resp)
	return nil
}

func listDisk(ctx context.Context, computeService *compute.Service) error {
	req := computeService.Disks.List(*project, *zone)
	if err := req.Filter(*filter).Pages(ctx, func(page *compute.DiskList) error {
		log.Printf("there are %v disk in zone:%v\n", len(page.Items), *zone)
		log.Printf("list disk:\npage.Items: %T", page.Items)
		for _, disk := range page.Items {
			diskJson, err := json.MarshalIndent(disk, "", "\t")
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("disk: %T - %v\n", disk, string(diskJson))
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func getDisk(ctx context.Context, computeService *compute.Service) error {
	resp, err := computeService.Disks.Get(*project, *zone, *disk).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("get disk", resp)
	return nil
}

func deleteDisk(ctx context.Context, computeService *compute.Service) error {
	resp, err := computeService.Disks.Delete(*project, *zone, *disk).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("delete disk", resp)
	return nil
}

func printJson(action string, data interface{}) {
	buf, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s: %T - %v\n", action, data, string(buf))
}
