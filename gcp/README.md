Usage
==========================

# prepare

- create default credentials via https://developers.google.com/identity/protocols/application-default-credentials
- download json key, for example `hyper-test-4be009a5b924.json`

```
$ export http_proxy=http://127.0.0.1:8118
$ export https_proxy=http://127.0.0.1:8118

$ export GOOGLE_APPLICATION_CREDENTIALS=~/gcp/hyper-test-4be009a5b924.json
```

# usage

## list disk

```
//build 
$ go build main.go disk.go address.go

//run
$ ./main --resource disk --action list
2018/02/08 16:26:48 there are 10 disk in zone:us-central1-a
2018/02/08 16:26:48 list disk:
page.Items: []*compute.Disk
us-central1-a	pd-standard	60	READY	gcetest-us-central1-ceph-a-1(ubuntu1604-ceph)
us-central1-a	pd-standard	60	READY	gcetest-us-central1-control-a-1(centos7-hyper)
us-central1-a	pd-standard	100	READY	gke-cluster-1-default-pool-4ffdb34e-6t8j(cos-stable-63-10032-71-0-p)
us-central1-a	pd-standard	100	READY	gke-cluster-1-default-pool-4ffdb34e-dkfc(cos-stable-63-10032-71-0-p)
us-central1-a	pd-standard	100	READY	gke-cluster-1-default-pool-4ffdb34e-j74j(cos-stable-63-10032-71-0-p)
us-central1-a	pd-standard	1	READY	v-31bfb7dbd0494646b4a096bdfe0de098-8b88aa8521a4837ded8367c41ba4()
us-central1-a	pd-standard	1	READY	v-31bfb7dbd0494646b4a096bdfe0de098-accf608129cc716b728e697e2f91()
us-central1-a	pd-standard	1	READY	v-31bfb7dbd0494646b4a096bdfe0de098-b2f84ef05f80b66d9197c15c5541()
us-central1-a	pd-standard	1	READY	v-31bfb7dbd0494646b4a096bdfe0de098-testvol-conflict-name()
us-central1-a	pd-standard	1	READY	v-31bfb7dbd0494646b4a096bdfe0de098-testvol-default-zone()
```

## list address

```
//filter by name
$ ./main --resource address --action list --filter="(name=ip-*)"
or
$ ./main --resource address --action list --filter='(name="ip-*")'
2018/02/08 17:02:40 there are 4 address in region:us-central1
2018/02/08 17:02:40 list address:
page.Items: []*compute.Address
us-central1	IN_USE	35.226.238.133()	ip-00a54ebcc0444bb384e48f6fd7b5597b-gateway
us-central1	IN_USE	35.188.36.134()	ip-ab43de68c97a412e999249c0aff1eef3-gateway
us-central1	IN_USE	104.197.95.6()	ip-c9f0474bb4bd4612b2b1dc2b87ecbd62-gateway
us-central1	IN_USE	35.226.242.108()	ip-default-gateway


//filter by ip
./main --resource address --action list --filter='(address="35.202.162.162")'
2018/02/08 17:43:23 filter:(addressType!=INTERNAL) (address="35.202.162.162")
2018/02/08 17:43:25 there are 1 address in region:us-central1
2018/02/08 17:43:25 list address:
page.Items: []*compute.Address
us-central1	IN_USE	35.202.162.162()	gcetest-us-central1-control-b-1
```