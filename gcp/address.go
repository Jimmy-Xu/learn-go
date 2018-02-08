package main

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v0.beta"
	"fmt"
)

type AddressTable struct {
	Name string
	Region string
	Address string
	Type string
	Status string
}

func createAddress(ctx context.Context, computeService *compute.Service, region, name string) error {
	rb := &compute.Address{
		Name:        name,
		Description: "for test",
	}
	resp, err := computeService.Addresses.Insert(*project, region, rb).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("create address", resp)
	return nil
}

func listAddress(ctx context.Context, computeService *compute.Service, region, filter string) error {
	filter = fmt.Sprintf("(addressType!=INTERNAL) %v", filter)
	log.Printf("filter:%v\n", filter)
	req := computeService.Addresses.List(*project, region)
	if err := req.Filter(filter).MaxResults(maxResults).Pages(ctx, func(page *compute.AddressList) error {
		log.Printf("there are %v address in region:%v\n", len(page.Items), region)
		log.Printf("list address:\npage.Items: %T", page.Items)
		addressTable := make([]AddressTable, 0, len(page.Items))
		for _, address := range page.Items {
			if *format == "json" {
				addressJson, err := json.MarshalIndent(address, "", "\t")
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("address: %T - %v\n", address, string(addressJson))
			} else {
				_address := AddressTable {
					Name: address.Name,
					Region: getShortValue(address.Region),
					Address: getShortValue(address.Address),
					Type: getShortValue(address.AddressType),
					Status: address.Status,
				}
				addressTable = append(addressTable, _address)
			}
		}
		if *format != "json" {
			for _, at := range addressTable {
				fmt.Printf("%v\t%v\t%v(%v)\t%v\n", at.Region, at.Status, at.Address, at.Type, at.Name)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}


func getAddress(ctx context.Context, computeService *compute.Service, region, name string) error {
	resp, err := computeService.Addresses.Get(*project, region, name).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("get address", resp)
	return nil
}

func deleteAddress(ctx context.Context, computeService *compute.Service, region, name string) error {
	resp, err := computeService.Addresses.Delete(*project, region, name).Context(ctx).Do()
	if err != nil {
		return err
	}
	printJson("delete address", resp)
	return nil
}
