package main

import (
	"encoding/json"
	"log"
	"fmt"

	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v0.beta"
)

type DiskTable struct {
	Name string
	Size int64
	Zone string
	Type string
	Image string
	Status string
}

func createDisk(ctx context.Context, computeService *compute.Service, zone, name string) error {
	rb := &compute.Disk{
		SizeGb:      1,
		Name:        name,
		Description: "for test",
	}
	resp, err := computeService.Disks.Insert(*project, zone, rb).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("create disk", resp)
	return nil
}

func listDisk(ctx context.Context, computeService *compute.Service, zone, filter string) error {
	req := computeService.Disks.List(*project, zone)
	if err := req.Filter(filter).MaxResults(maxResults).Pages(ctx, func(page *compute.DiskList) error {
		log.Printf("there are %v disk in zone:%v\n", len(page.Items), zone)
		log.Printf("list disk:\npage.Items: %T", page.Items)
		diskTable := make([]DiskTable, 0, len(page.Items))
		for _, disk := range page.Items {
			if *format == "json" {
				diskJson, err := json.MarshalIndent(disk, "", "\t")
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("disk: %T - %v\n", disk, string(diskJson))
			} else {
				_disk := DiskTable {
					Name: disk.Name,
					Size: disk.SizeGb,
					Zone: getShortValue(disk.Zone),
					Type: getShortValue(disk.Type),
					Image: getShortValue(disk.SourceImage),
					Status: disk.Status,
				}
				diskTable = append(diskTable, _disk)
			}
		}
		if *format != "json" {
			for _, dt := range diskTable {
				fmt.Printf("%v\t%v\t%v\t%v\t%v(%v)\n", dt.Zone, dt.Type, dt.Size, dt.Status, dt.Name, dt.Image)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func getDisk(ctx context.Context, computeService *compute.Service, zone, name string) error {
	resp, err := computeService.Disks.Get(*project, zone, name).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("get disk", resp)
	return nil
}

func deleteDisk(ctx context.Context, computeService *compute.Service, zone, name string) error {
	resp, err := computeService.Disks.Delete(*project, zone, name).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("delete disk", resp)
	return nil
}

