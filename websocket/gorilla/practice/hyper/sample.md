# pod events

## run container
$ hyper run -d --name=test busybox top
```
recv[json]: {
  "Action": "ADDED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Pending",
  "Time": 1.476020105e+09,
  "TimeNano": 1.476020105268817e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Pending",
  "Time": 1.476020105e+09,
  "TimeNano": 1.4760201052730767e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Pending",
  "Time": 1.476020105e+09,
  "TimeNano": 1.4760201052776668e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Running",
  "Time": 1.476020107e+09,
  "TimeNano": 1.4760201070706463e+18,
  "Type": "container"
}
```

## stop container

$ hyper stop test
```
recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Succeeded",
  "Time": 1.476020145e+09,
  "TimeNano": 1.4760201455606147e+18,
  "Type": "container"
}
```

## start container

$ hyper start test
```
recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Succeeded",
  "Time": 1.476020183e+09,
  "TimeNano": 1.4760201831480044e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "DELETED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Succeeded",
  "Time": 1.476020183e+09,
  "TimeNano": 1.4760201831485804e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "ADDED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Pending",
  "Time": 1.476020184e+09,
  "TimeNano": 1.4760201841509678e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Pending",
  "Time": 1.476020184e+09,
  "TimeNano": 1.47602018415341e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Pending",
  "Time": 1.476020184e+09,
  "TimeNano": 1.4760201841585713e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Running",
  "Time": 1.476020186e+09,
  "TimeNano": 1.4760201860969262e+18,
  "Type": "container"
}
```

## remove container

$ hyper rm -f test
```
recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Succeeded",
  "Time": 1.476020248e+09,
  "TimeNano": 1.476020248911624e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "MODIFIED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Succeeded",
  "Time": 1.476020249e+09,
  "TimeNano": 1.4760202491858732e+18,
  "Type": "container"
}

recv[json]: {
  "Action": "DELETED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test"
    },
    "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05"
  },
  "From": "busybox",
  "Id": "50af009a9d579517250676afaa3647587fdf822466dce20a5ba9a9fbf32e2c05",
  "Status": "Succeeded",
  "Time": 1.476020249e+09,
  "TimeNano": 1.4760202491866368e+18,
  "Type": "container"
}
```

# pod info
```
-------------------- Pod.ObjectMeta --------------------
0: Name(string)=4ebd063e9866214568c8cdb84cf1d166939e5dd971fee352b4ea33236edafd27
1: GenerateName(string)=
2: Namespace(string)=0b1318054bfe440a97e8258ce8d92d47-bridge
3: SelfLink(string)=/api/v1/namespaces/0b1318054bfe440a97e8258ce8d92d47-bridge/pods/4ebd063e9866
4: UID(string)=b0f50de8-8dfe-11e6-979c-0cc47a15e44a
5: ResourceVersion(string)=137504504
6: Generation(int64)=<int64 Value>
7: CreationTimestamp(struct)=<unversioned.Time Value>
8: DeletionTimestamp(ptr)=<*unversioned.Time Value>
9: DeletionGracePeriodSeconds(ptr)=<*int64 Value>
10: Labels(map)=<map[string]string Value>
11: Annotations(map)=<map[string]string Value>
-------------------- Pod.TypeMeta --------------------
0: Kind(string)=
1: APIVersion(string)=
-------------------- Pod.Spec --------------------
0: Volumes(slice)=<[]api.Volume Value>
1: Containers(slice)=<[]api.Container Value>
2: RestartPolicy(string)=Never
3: TerminationGracePeriodSeconds(ptr)=<*int64 Value>
4: ActiveDeadlineSeconds(ptr)=<*int64 Value>
5: DNSPolicy(string)=ClusterFirst
6: NodeSelector(map)=<map[string]string Value>
7: ServiceAccountName(string)=default
8: NodeName(string)=cns-2
9: SecurityContext(ptr)=<*api.PodSecurityContext Value>
10: ImagePullSecrets(slice)=<[]api.LocalObjectReference Value>
-------------------- Pod.Spec.Containers --------------------
0: Name(string)=test
1: Image(string)=busybox
2: Command(slice)=<[]string Value>
3: Args(slice)=<[]string Value>
4: WorkingDir(string)=
5: Ports(slice)=<[]api.ContainerPort Value>
6: Env(slice)=<[]api.EnvVar Value>
7: Resources(struct)=<api.ResourceRequirements Value>
8: VolumeMounts(slice)=<[]api.VolumeMount Value>
9: LivenessProbe(ptr)=<*api.Probe Value>
10: ReadinessProbe(ptr)=<*api.Probe Value>
11: Lifecycle(ptr)=<*api.Lifecycle Value>
12: TerminationMessagePath(string)=/dev/termination-log
13: ImagePullPolicy(string)=IfNotPresent
14: SecurityContext(ptr)=<*api.SecurityContext Value>
15: Stdin(bool)=<bool Value>
16: StdinOnce(bool)=<bool Value>
17: TTY(bool)=<bool Value>
-------------------- Pod.Status --------------------
0: Phase(string)=Running
1: Conditions(slice)=<[]api.PodCondition Value>
2: Message(string)=
3: Reason(string)=
4: HostIP(string)=147.75.195.41
5: PodIP(string)=172.16.0.32
6: StartTime(ptr)=<*unversioned.Time Value>
7: ContainerStatuses(slice)=<[]api.ContainerStatus Value>
-------------------- db.Container --------------------
0: ID(string)=955fb7fed391d325bed5b7f85c05824e3bd035b0f5d9aa30ca87c6169d075148
1: Name(string)=/test
2: Pod(string)=955fb7fed391d325bed5b7f85c05824e3bd035b0f5d9aa30ca87c6169d075148
3: Namespace(string)=0b1318054bfe440a97e8258ce8d92d47-bridge
4: PublicIP(string)=
5: Tenant(string)=0b1318054bfe440a97e8258ce8d92d47
6: Alias(map)=<map[string]string Value>
7: Links(map)=<map[string]*links.Link Value>
8: Volumes(slice)=<[]*volume.Volume Value>
9: Base(ptr)=<*types.ContainerJSON Value>
10: Restart(ptr)=<*container.RestartCtrl Value>
11: Updating(bool)=<bool Value>
12: Attaching(bool)=<bool Value>
dbContainer.Base.Image: sha256:e02e811dd08fd49e7f6032625495118e63f597eb150403d02e3238af1df240ba
```
