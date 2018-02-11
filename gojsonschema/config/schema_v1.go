package config

var SchemaV1 = `{
  "x-kubernetes-group-version-kind": [
    {
      "kind": "Pod",
      "version": "v1",
      "group": ""
    }
  ],
  "description": "Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.",
  "properties": {
    "status": {
      "description": "PodStatus represents information about the status of a pod. Status may trail the actual state of a system.",
      "properties": {
        "initContainerStatuses": {
          "items": {
            "required": [
              "name",
              "ready",
              "restartCount",
              "image",
              "imageID"
            ],
            "description": "ContainerStatus contains details for the current status of this container.",
            "properties": {
              "restartCount": {
                "type": "integer",
                "description": "The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC.",
                "format": "int32"
              },
              "name": {
                "type": "string",
                "description": "This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated."
              },
              "image": {
                "type": "string",
                "description": "The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images"
              },
              "imageID": {
                "type": "string",
                "description": "ImageID of the container's image."
              },
              "state": {
                "description": "ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.",
                "properties": {
                  "terminated": {
                    "required": [
                      "exitCode"
                    ],
                    "description": "ContainerStateTerminated is a terminated state of a container.",
                    "properties": {
                      "containerID": {
                        "type": "string",
                        "description": "Container's ID in the format 'docker://<container_id>'"
                      },
                      "signal": {
                        "type": "integer",
                        "description": "Signal from the last termination of the container",
                        "format": "int32"
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason from the last termination of the container"
                      },
                      "finishedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "message": {
                        "type": "string",
                        "description": "Message regarding the last termination of the container"
                      },
                      "exitCode": {
                        "type": "integer",
                        "description": "Exit status from the last termination of the container",
                        "format": "int32"
                      }
                    }
                  },
                  "running": {
                    "description": "ContainerStateRunning is a running state of a container.",
                    "properties": {
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      }
                    }
                  },
                  "waiting": {
                    "description": "ContainerStateWaiting is a waiting state of a container.",
                    "properties": {
                      "message": {
                        "type": "string",
                        "description": "Message regarding why the container is not yet running."
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason the container is not yet running."
                      }
                    }
                  }
                }
              },
              "ready": {
                "type": "boolean",
                "description": "Specifies whether the container has passed its readiness probe."
              },
              "lastState": {
                "description": "ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.",
                "properties": {
                  "terminated": {
                    "required": [
                      "exitCode"
                    ],
                    "description": "ContainerStateTerminated is a terminated state of a container.",
                    "properties": {
                      "containerID": {
                        "type": "string",
                        "description": "Container's ID in the format 'docker://<container_id>'"
                      },
                      "signal": {
                        "type": "integer",
                        "description": "Signal from the last termination of the container",
                        "format": "int32"
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason from the last termination of the container"
                      },
                      "finishedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "message": {
                        "type": "string",
                        "description": "Message regarding the last termination of the container"
                      },
                      "exitCode": {
                        "type": "integer",
                        "description": "Exit status from the last termination of the container",
                        "format": "int32"
                      }
                    }
                  },
                  "running": {
                    "description": "ContainerStateRunning is a running state of a container.",
                    "properties": {
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      }
                    }
                  },
                  "waiting": {
                    "description": "ContainerStateWaiting is a waiting state of a container.",
                    "properties": {
                      "message": {
                        "type": "string",
                        "description": "Message regarding why the container is not yet running."
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason the container is not yet running."
                      }
                    }
                  }
                }
              },
              "containerID": {
                "type": "string",
                "description": "Container's ID in the format 'docker://<container_id>'."
              }
            }
          },
          "type": "array",
          "description": "The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status"
        },
        "qosClass": {
          "type": "string",
          "description": "The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://github.com/kubernetes/kubernetes/blob/master/docs/design/resource-qos.md"
        },
        "containerStatuses": {
          "items": {
            "required": [
              "name",
              "ready",
              "restartCount",
              "image",
              "imageID"
            ],
            "description": "ContainerStatus contains details for the current status of this container.",
            "properties": {
              "restartCount": {
                "type": "integer",
                "description": "The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC.",
                "format": "int32"
              },
              "name": {
                "type": "string",
                "description": "This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated."
              },
              "image": {
                "type": "string",
                "description": "The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images"
              },
              "imageID": {
                "type": "string",
                "description": "ImageID of the container's image."
              },
              "state": {
                "description": "ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.",
                "properties": {
                  "terminated": {
                    "required": [
                      "exitCode"
                    ],
                    "description": "ContainerStateTerminated is a terminated state of a container.",
                    "properties": {
                      "containerID": {
                        "type": "string",
                        "description": "Container's ID in the format 'docker://<container_id>'"
                      },
                      "signal": {
                        "type": "integer",
                        "description": "Signal from the last termination of the container",
                        "format": "int32"
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason from the last termination of the container"
                      },
                      "finishedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "message": {
                        "type": "string",
                        "description": "Message regarding the last termination of the container"
                      },
                      "exitCode": {
                        "type": "integer",
                        "description": "Exit status from the last termination of the container",
                        "format": "int32"
                      }
                    }
                  },
                  "running": {
                    "description": "ContainerStateRunning is a running state of a container.",
                    "properties": {
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      }
                    }
                  },
                  "waiting": {
                    "description": "ContainerStateWaiting is a waiting state of a container.",
                    "properties": {
                      "message": {
                        "type": "string",
                        "description": "Message regarding why the container is not yet running."
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason the container is not yet running."
                      }
                    }
                  }
                }
              },
              "ready": {
                "type": "boolean",
                "description": "Specifies whether the container has passed its readiness probe."
              },
              "lastState": {
                "description": "ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.",
                "properties": {
                  "terminated": {
                    "required": [
                      "exitCode"
                    ],
                    "description": "ContainerStateTerminated is a terminated state of a container.",
                    "properties": {
                      "containerID": {
                        "type": "string",
                        "description": "Container's ID in the format 'docker://<container_id>'"
                      },
                      "signal": {
                        "type": "integer",
                        "description": "Signal from the last termination of the container",
                        "format": "int32"
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason from the last termination of the container"
                      },
                      "finishedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "message": {
                        "type": "string",
                        "description": "Message regarding the last termination of the container"
                      },
                      "exitCode": {
                        "type": "integer",
                        "description": "Exit status from the last termination of the container",
                        "format": "int32"
                      }
                    }
                  },
                  "running": {
                    "description": "ContainerStateRunning is a running state of a container.",
                    "properties": {
                      "startedAt": {
                        "type": "string",
                        "format": "date-time"
                      }
                    }
                  },
                  "waiting": {
                    "description": "ContainerStateWaiting is a waiting state of a container.",
                    "properties": {
                      "message": {
                        "type": "string",
                        "description": "Message regarding why the container is not yet running."
                      },
                      "reason": {
                        "type": "string",
                        "description": "(brief) reason the container is not yet running."
                      }
                    }
                  }
                }
              },
              "containerID": {
                "type": "string",
                "description": "Container's ID in the format 'docker://<container_id>'."
              }
            }
          },
          "type": "array",
          "description": "The list has one entry per container in the manifest. Each entry is currently the output of 'docker inspect'. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status"
        },
        "reason": {
          "type": "string",
          "description": "A brief CamelCase message indicating details about why the pod is in this state. e.g. 'Evicted'"
        },
        "podIP": {
          "type": "string",
          "description": "IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated."
        },
        "startTime": {
          "type": "string",
          "format": "date-time"
        },
        "hostIP": {
          "type": "string",
          "description": "IP address of the host to which the pod is assigned. Empty if not yet scheduled."
        },
        "phase": {
          "type": "string",
          "description": "Current condition of the pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase"
        },
        "message": {
          "type": "string",
          "description": "A human readable message indicating details about why the pod is in this condition."
        },
        "conditions": {
          "items": {
            "required": [
              "type",
              "status"
            ],
            "description": "PodCondition contains details for the current condition of this pod.",
            "properties": {
              "status": {
                "type": "string",
                "description": "Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions"
              },
              "lastTransitionTime": {
                "type": "string",
                "format": "date-time"
              },
              "reason": {
                "type": "string",
                "description": "Unique, one-word, CamelCase reason for the condition's last transition."
              },
              "message": {
                "type": "string",
                "description": "Human-readable message indicating details about last transition."
              },
              "type": {
                "type": "string",
                "description": "Type is the type of the condition. Currently only Ready. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions"
              }
            }
          },
          "x-kubernetes-patch-merge-key": "type",
          "type": "array",
          "description": "Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions",
          "x-kubernetes-patch-strategy": "merge"
        }
      }
    },
    "kind": {
      "type": "string",
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
    },
    "spec": {
      "required": [
        "containers"
      ],
      "description": "PodSpec is a description of a pod.",
      "properties": {
        "dnsPolicy": {
          "type": "string",
          "description": "Set DNS policy for the pod. Defaults to \"ClusterFirst\". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Note that 'None' policy is an alpha feature introduced in v1.9 and CustomPodDNS feature gate must be enabled to use it."
        },
        "hostNetwork": {
          "type": "boolean",
          "description": "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false."
        },
        "restartPolicy": {
          "type": "string",
          "description": "Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy"
        },
        "automountServiceAccountToken": {
          "type": "boolean",
          "description": "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted."
        },
        "priorityClassName": {
          "type": "string",
          "description": "If specified, indicates the pod's priority. \"SYSTEM\" is a special keyword which indicates the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default."
        },
        "securityContext": {
          "description": "PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.",
          "properties": {
            "runAsNonRoot": {
              "type": "boolean",
              "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence."
            },
            "fsGroup": {
              "type": "integer",
              "description": "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:\n\n1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----\n\nIf unset, the Kubelet will not modify the ownership and permissions of any volume.",
              "format": "int64"
            },
            "seLinuxOptions": {
              "description": "SELinuxOptions are the labels to be applied to the container",
              "properties": {
                "role": {
                  "type": "string",
                  "description": "Role is a SELinux role label that applies to the container."
                },
                "type": {
                  "type": "string",
                  "description": "Type is a SELinux type label that applies to the container."
                },
                "user": {
                  "type": "string",
                  "description": "User is a SELinux user label that applies to the container."
                },
                "level": {
                  "type": "string",
                  "description": "Level is SELinux level label that applies to the container."
                }
              }
            },
            "runAsUser": {
              "type": "integer",
              "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
              "format": "int64"
            },
            "supplementalGroups": {
              "items": {
                "type": "integer",
                "format": "int64"
              },
              "type": "array",
              "description": "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container."
            }
          }
        },
        "nodeName": {
          "type": "string",
          "description": "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements."
        },
        "hostAliases": {
          "items": {
            "description": "HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.",
            "properties": {
              "ip": {
                "type": "string",
                "description": "IP address of the host file entry."
              },
              "hostnames": {
                "items": {
                  "type": "string"
                },
                "type": "array",
                "description": "Hostnames for the above IP address."
              }
            }
          },
          "x-kubernetes-patch-merge-key": "ip",
          "type": "array",
          "description": "HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.",
          "x-kubernetes-patch-strategy": "merge"
        },
        "hostname": {
          "type": "string",
          "description": "Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value."
        },
        "serviceAccount": {
          "type": "string",
          "description": "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead."
        },
        "nodeSelector": {
          "additionalProperties": {
            "type": "string"
          },
          "type": "object",
          "description": "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/"
        },
        "priority": {
          "type": "integer",
          "description": "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",
          "format": "int32"
        },
        "affinity": {
          "description": "Affinity is a group of affinity scheduling rules.",
          "properties": {
            "podAffinity": {
              "description": "Pod affinity is a group of inter pod affinity scheduling rules.",
              "properties": {
                "requiredDuringSchedulingIgnoredDuringExecution": {
                  "items": {
                    "required": [
                      "topologyKey"
                    ],
                    "description": "Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running",
                    "properties": {
                      "labelSelector": {
                        "description": "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
                        "properties": {
                          "matchLabels": {
                            "additionalProperties": {
                              "type": "string"
                            },
                            "type": "object",
                            "description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed."
                          },
                          "matchExpressions": {
                            "items": {
                              "required": [
                                "key",
                                "operator"
                              ],
                              "description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
                              "properties": {
                                "operator": {
                                  "type": "string",
                                  "description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist."
                                },
                                "values": {
                                  "items": {
                                    "type": "string"
                                  },
                                  "type": "array",
                                  "description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch."
                                },
                                "key": {
                                  "x-kubernetes-patch-merge-key": "key",
                                  "type": "string",
                                  "description": "key is the label key that the selector applies to.",
                                  "x-kubernetes-patch-strategy": "merge"
                                }
                              }
                            },
                            "type": "array",
                            "description": "matchExpressions is a list of label selector requirements. The requirements are ANDed."
                          }
                        }
                      },
                      "namespaces": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means \"this pod's namespace\""
                      },
                      "topologyKey": {
                        "type": "string",
                        "description": "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed."
                      }
                    }
                  },
                  "type": "array",
                  "description": "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied."
                },
                "preferredDuringSchedulingIgnoredDuringExecution": {
                  "items": {
                    "required": [
                      "weight",
                      "podAffinityTerm"
                    ],
                    "description": "The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)",
                    "properties": {
                      "podAffinityTerm": {
                        "required": [
                          "topologyKey"
                        ],
                        "description": "Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running",
                        "properties": {
                          "labelSelector": {
                            "description": "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
                            "properties": {
                              "matchLabels": {
                                "additionalProperties": {
                                  "type": "string"
                                },
                                "type": "object",
                                "description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed."
                              },
                              "matchExpressions": {
                                "items": {
                                  "required": [
                                    "key",
                                    "operator"
                                  ],
                                  "description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
                                  "properties": {
                                    "operator": {
                                      "type": "string",
                                      "description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist."
                                    },
                                    "values": {
                                      "items": {
                                        "type": "string"
                                      },
                                      "type": "array",
                                      "description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch."
                                    },
                                    "key": {
                                      "x-kubernetes-patch-merge-key": "key",
                                      "type": "string",
                                      "description": "key is the label key that the selector applies to.",
                                      "x-kubernetes-patch-strategy": "merge"
                                    }
                                  }
                                },
                                "type": "array",
                                "description": "matchExpressions is a list of label selector requirements. The requirements are ANDed."
                              }
                            }
                          },
                          "namespaces": {
                            "items": {
                              "type": "string"
                            },
                            "type": "array",
                            "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means \"this pod's namespace\""
                          },
                          "topologyKey": {
                            "type": "string",
                            "description": "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed."
                          }
                        }
                      },
                      "weight": {
                        "type": "integer",
                        "description": "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
                        "format": "int32"
                      }
                    }
                  },
                  "type": "array",
                  "description": "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred."
                }
              }
            },
            "nodeAffinity": {
              "description": "Node affinity is a group of node affinity scheduling rules.",
              "properties": {
                "requiredDuringSchedulingIgnoredDuringExecution": {
                  "required": [
                    "nodeSelectorTerms"
                  ],
                  "description": "A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.",
                  "properties": {
                    "nodeSelectorTerms": {
                      "items": {
                        "required": [
                          "matchExpressions"
                        ],
                        "description": "A null or empty node selector term matches no objects.",
                        "properties": {
                          "matchExpressions": {
                            "items": {
                              "required": [
                                "key",
                                "operator"
                              ],
                              "description": "A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
                              "properties": {
                                "operator": {
                                  "type": "string",
                                  "description": "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt."
                                },
                                "values": {
                                  "items": {
                                    "type": "string"
                                  },
                                  "type": "array",
                                  "description": "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch."
                                },
                                "key": {
                                  "type": "string",
                                  "description": "The label key that the selector applies to."
                                }
                              }
                            },
                            "type": "array",
                            "description": "Required. A list of node selector requirements. The requirements are ANDed."
                          }
                        }
                      },
                      "type": "array",
                      "description": "Required. A list of node selector terms. The terms are ORed."
                    }
                  }
                },
                "preferredDuringSchedulingIgnoredDuringExecution": {
                  "items": {
                    "required": [
                      "weight",
                      "preference"
                    ],
                    "description": "An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).",
                    "properties": {
                      "preference": {
                        "required": [
                          "matchExpressions"
                        ],
                        "description": "A null or empty node selector term matches no objects.",
                        "properties": {
                          "matchExpressions": {
                            "items": {
                              "required": [
                                "key",
                                "operator"
                              ],
                              "description": "A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
                              "properties": {
                                "operator": {
                                  "type": "string",
                                  "description": "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt."
                                },
                                "values": {
                                  "items": {
                                    "type": "string"
                                  },
                                  "type": "array",
                                  "description": "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch."
                                },
                                "key": {
                                  "type": "string",
                                  "description": "The label key that the selector applies to."
                                }
                              }
                            },
                            "type": "array",
                            "description": "Required. A list of node selector requirements. The requirements are ANDed."
                          }
                        }
                      },
                      "weight": {
                        "type": "integer",
                        "description": "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
                        "format": "int32"
                      }
                    }
                  },
                  "type": "array",
                  "description": "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred."
                }
              }
            },
            "podAntiAffinity": {
              "description": "Pod anti affinity is a group of inter pod anti affinity scheduling rules.",
              "properties": {
                "requiredDuringSchedulingIgnoredDuringExecution": {
                  "items": {
                    "required": [
                      "topologyKey"
                    ],
                    "description": "Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running",
                    "properties": {
                      "labelSelector": {
                        "description": "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
                        "properties": {
                          "matchLabels": {
                            "additionalProperties": {
                              "type": "string"
                            },
                            "type": "object",
                            "description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed."
                          },
                          "matchExpressions": {
                            "items": {
                              "required": [
                                "key",
                                "operator"
                              ],
                              "description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
                              "properties": {
                                "operator": {
                                  "type": "string",
                                  "description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist."
                                },
                                "values": {
                                  "items": {
                                    "type": "string"
                                  },
                                  "type": "array",
                                  "description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch."
                                },
                                "key": {
                                  "x-kubernetes-patch-merge-key": "key",
                                  "type": "string",
                                  "description": "key is the label key that the selector applies to.",
                                  "x-kubernetes-patch-strategy": "merge"
                                }
                              }
                            },
                            "type": "array",
                            "description": "matchExpressions is a list of label selector requirements. The requirements are ANDed."
                          }
                        }
                      },
                      "namespaces": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means \"this pod's namespace\""
                      },
                      "topologyKey": {
                        "type": "string",
                        "description": "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed."
                      }
                    }
                  },
                  "type": "array",
                  "description": "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied."
                },
                "preferredDuringSchedulingIgnoredDuringExecution": {
                  "items": {
                    "required": [
                      "weight",
                      "podAffinityTerm"
                    ],
                    "description": "The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)",
                    "properties": {
                      "podAffinityTerm": {
                        "required": [
                          "topologyKey"
                        ],
                        "description": "Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running",
                        "properties": {
                          "labelSelector": {
                            "description": "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
                            "properties": {
                              "matchLabels": {
                                "additionalProperties": {
                                  "type": "string"
                                },
                                "type": "object",
                                "description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed."
                              },
                              "matchExpressions": {
                                "items": {
                                  "required": [
                                    "key",
                                    "operator"
                                  ],
                                  "description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
                                  "properties": {
                                    "operator": {
                                      "type": "string",
                                      "description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist."
                                    },
                                    "values": {
                                      "items": {
                                        "type": "string"
                                      },
                                      "type": "array",
                                      "description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch."
                                    },
                                    "key": {
                                      "x-kubernetes-patch-merge-key": "key",
                                      "type": "string",
                                      "description": "key is the label key that the selector applies to.",
                                      "x-kubernetes-patch-strategy": "merge"
                                    }
                                  }
                                },
                                "type": "array",
                                "description": "matchExpressions is a list of label selector requirements. The requirements are ANDed."
                              }
                            }
                          },
                          "namespaces": {
                            "items": {
                              "type": "string"
                            },
                            "type": "array",
                            "description": "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means \"this pod's namespace\""
                          },
                          "topologyKey": {
                            "type": "string",
                            "description": "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed."
                          }
                        }
                      },
                      "weight": {
                        "type": "integer",
                        "description": "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
                        "format": "int32"
                      }
                    }
                  },
                  "type": "array",
                  "description": "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred."
                }
              }
            }
          }
        },
        "tolerations": {
          "items": {
            "description": "The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.",
            "properties": {
              "operator": {
                "type": "string",
                "description": "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category."
              },
              "key": {
                "type": "string",
                "description": "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys."
              },
              "tolerationSeconds": {
                "type": "integer",
                "description": "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
                "format": "int64"
              },
              "effect": {
                "type": "string",
                "description": "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute."
              },
              "value": {
                "type": "string",
                "description": "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string."
              }
            }
          },
          "type": "array",
          "description": "If specified, the pod's tolerations."
        },
        "subdomain": {
          "type": "string",
          "description": "If specified, the fully qualified Pod hostname will be \"<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>\". If not specified, the pod will not have a domainname at all."
        },
        "containers": {
          "items": {
            "required": [
              "name"
            ],
            "description": "A single application container that you want to run within a pod.",
            "properties": {
              "livenessProbe": {
                "description": "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
                "properties": {
                  "httpGet": {
                    "required": [
                      "port"
                    ],
                    "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                    "properties": {
                      "path": {
                        "type": "string",
                        "description": "Path to access on the HTTP server."
                      },
                      "host": {
                        "type": "string",
                        "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                      },
                      "scheme": {
                        "type": "string",
                        "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                      },
                      "httpHeaders": {
                        "items": {
                          "required": [
                            "name",
                            "value"
                          ],
                          "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                          "properties": {
                            "name": {
                              "type": "string",
                              "description": "The header field name"
                            },
                            "value": {
                              "type": "string",
                              "description": "The header field value"
                            }
                          }
                        },
                        "type": "array",
                        "description": "Custom headers to set in the request. HTTP allows repeated headers."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "timeoutSeconds": {
                    "type": "integer",
                    "description": "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "exec": {
                    "description": "ExecAction describes a \"run in container\" action.",
                    "properties": {
                      "command": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                      }
                    }
                  },
                  "initialDelaySeconds": {
                    "type": "integer",
                    "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "tcpSocket": {
                    "required": [
                      "port"
                    ],
                    "description": "TCPSocketAction describes an action based on opening a socket",
                    "properties": {
                      "host": {
                        "type": "string",
                        "description": "Optional: Host name to connect to, defaults to the pod IP."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "periodSeconds": {
                    "type": "integer",
                    "description": "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
                    "format": "int32"
                  },
                  "successThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
                    "format": "int32"
                  },
                  "failureThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
                    "format": "int32"
                  }
                }
              },
              "args": {
                "items": {
                  "type": "string"
                },
                "type": "array",
                "description": "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell"
              },
              "terminationMessagePath": {
                "type": "string",
                "description": "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated."
              },
              "name": {
                "type": "string",
                "description": "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated."
              },
              "envFrom": {
                "items": {
                  "description": "EnvFromSource represents the source of a set of ConfigMaps",
                  "properties": {
                    "prefix": {
                      "type": "string",
                      "description": "An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER."
                    },
                    "configMapRef": {
                      "description": "ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.\n\nThe contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.",
                      "properties": {
                        "optional": {
                          "type": "boolean",
                          "description": "Specify whether the ConfigMap must be defined"
                        },
                        "name": {
                          "type": "string",
                          "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                        }
                      }
                    },
                    "secretRef": {
                      "description": "SecretEnvSource selects a Secret to populate the environment variables with.\n\nThe contents of the target Secret's Data field will represent the key-value pairs as environment variables.",
                      "properties": {
                        "optional": {
                          "type": "boolean",
                          "description": "Specify whether the Secret must be defined"
                        },
                        "name": {
                          "type": "string",
                          "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                        }
                      }
                    }
                  }
                },
                "type": "array",
                "description": "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated."
              },
              "workingDir": {
                "type": "string",
                "description": "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated."
              },
              "image": {
                "type": "string",
                "description": "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets."
              },
              "stdin": {
                "type": "boolean",
                "description": "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false."
              },
              "volumeMounts": {
                "items": {
                  "required": [
                    "name",
                    "mountPath"
                  ],
                  "description": "VolumeMount describes a mounting of a Volume within a container.",
                  "properties": {
                    "readOnly": {
                      "type": "boolean",
                      "description": "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false."
                    },
                    "mountPropagation": {
                      "type": "string",
                      "description": "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationHostToContainer is used. This field is alpha in 1.8 and can be reworked or removed in a future release."
                    },
                    "subPath": {
                      "type": "string",
                      "description": "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root)."
                    },
                    "name": {
                      "type": "string",
                      "description": "This must match the Name of a Volume."
                    },
                    "mountPath": {
                      "type": "string",
                      "description": "Path within the container at which the volume should be mounted.  Must not contain ':'."
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "mountPath",
                "type": "array",
                "description": "Pod volumes to mount into the container's filesystem. Cannot be updated.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "tty": {
                "type": "boolean",
                "description": "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false."
              },
              "terminationMessagePolicy": {
                "type": "string",
                "description": "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated."
              },
              "volumeDevices": {
                "items": {
                  "required": [
                    "name",
                    "devicePath"
                  ],
                  "description": "volumeDevice describes a mapping of a raw block device within a container.",
                  "properties": {
                    "devicePath": {
                      "type": "string",
                      "description": "devicePath is the path inside of the container that the device will be mapped to."
                    },
                    "name": {
                      "type": "string",
                      "description": "name must match the name of a persistentVolumeClaim in the pod"
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "devicePath",
                "type": "array",
                "description": "volumeDevices is the list of block devices to be used by the container. This is an alpha feature and may change in the future.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "command": {
                "items": {
                  "type": "string"
                },
                "type": "array",
                "description": "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell"
              },
              "env": {
                "items": {
                  "required": [
                    "name"
                  ],
                  "description": "EnvVar represents an environment variable present in a Container.",
                  "properties": {
                    "valueFrom": {
                      "description": "EnvVarSource represents a source for the value of an EnvVar.",
                      "properties": {
                        "secretKeyRef": {
                          "required": [
                            "key"
                          ],
                          "description": "SecretKeySelector selects a key of a Secret.",
                          "properties": {
                            "optional": {
                              "type": "boolean",
                              "description": "Specify whether the Secret or it's key must be defined"
                            },
                            "name": {
                              "type": "string",
                              "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                            },
                            "key": {
                              "type": "string",
                              "description": "The key of the secret to select from.  Must be a valid secret key."
                            }
                          }
                        },
                        "fieldRef": {
                          "required": [
                            "fieldPath"
                          ],
                          "description": "ObjectFieldSelector selects an APIVersioned field of an object.",
                          "properties": {
                            "fieldPath": {
                              "type": "string",
                              "description": "Path of the field to select in the specified API version."
                            },
                            "apiVersion": {
                              "type": "string",
                              "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\"."
                            }
                          }
                        },
                        "configMapKeyRef": {
                          "required": [
                            "key"
                          ],
                          "description": "Selects a key from a ConfigMap.",
                          "properties": {
                            "optional": {
                              "type": "boolean",
                              "description": "Specify whether the ConfigMap or it's key must be defined"
                            },
                            "name": {
                              "type": "string",
                              "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                            },
                            "key": {
                              "type": "string",
                              "description": "The key to select."
                            }
                          }
                        },
                        "resourceFieldRef": {
                          "required": [
                            "resource"
                          ],
                          "description": "ResourceFieldSelector represents container resources (cpu, memory) and their output format",
                          "properties": {
                            "containerName": {
                              "type": "string",
                              "description": "Container name: required for volumes, optional for env vars"
                            },
                            "resource": {
                              "type": "string",
                              "description": "Required: resource to select"
                            },
                            "divisor": {
                              "oneOf": [
                                {
                                  "type": "string"
                                },
                                {
                                  "type": "integer"
                                }
                              ]
                            }
                          }
                        }
                      }
                    },
                    "name": {
                      "type": "string",
                      "description": "Name of the environment variable. Must be a C_IDENTIFIER."
                    },
                    "value": {
                      "type": "string",
                      "description": "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\"."
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "name",
                "type": "array",
                "description": "List of environment variables to set in the container. Cannot be updated.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "imagePullPolicy": {
                "type": "string",
                "description": "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images"
              },
              "readinessProbe": {
                "description": "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
                "properties": {
                  "httpGet": {
                    "required": [
                      "port"
                    ],
                    "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                    "properties": {
                      "path": {
                        "type": "string",
                        "description": "Path to access on the HTTP server."
                      },
                      "host": {
                        "type": "string",
                        "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                      },
                      "scheme": {
                        "type": "string",
                        "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                      },
                      "httpHeaders": {
                        "items": {
                          "required": [
                            "name",
                            "value"
                          ],
                          "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                          "properties": {
                            "name": {
                              "type": "string",
                              "description": "The header field name"
                            },
                            "value": {
                              "type": "string",
                              "description": "The header field value"
                            }
                          }
                        },
                        "type": "array",
                        "description": "Custom headers to set in the request. HTTP allows repeated headers."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "timeoutSeconds": {
                    "type": "integer",
                    "description": "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "exec": {
                    "description": "ExecAction describes a \"run in container\" action.",
                    "properties": {
                      "command": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                      }
                    }
                  },
                  "initialDelaySeconds": {
                    "type": "integer",
                    "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "tcpSocket": {
                    "required": [
                      "port"
                    ],
                    "description": "TCPSocketAction describes an action based on opening a socket",
                    "properties": {
                      "host": {
                        "type": "string",
                        "description": "Optional: Host name to connect to, defaults to the pod IP."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "periodSeconds": {
                    "type": "integer",
                    "description": "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
                    "format": "int32"
                  },
                  "successThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
                    "format": "int32"
                  },
                  "failureThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
                    "format": "int32"
                  }
                }
              },
              "securityContext": {
                "description": "SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.",
                "properties": {
                  "readOnlyRootFilesystem": {
                    "type": "boolean",
                    "description": "Whether this container has a read-only root filesystem. Default is false."
                  },
                  "runAsUser": {
                    "type": "integer",
                    "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                    "format": "int64"
                  },
                  "allowPrivilegeEscalation": {
                    "type": "boolean",
                    "description": "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN"
                  },
                  "capabilities": {
                    "description": "Adds and removes POSIX capabilities from running containers.",
                    "properties": {
                      "add": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Added capabilities"
                      },
                      "drop": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Removed capabilities"
                      }
                    }
                  },
                  "runAsNonRoot": {
                    "type": "boolean",
                    "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence."
                  },
                  "seLinuxOptions": {
                    "description": "SELinuxOptions are the labels to be applied to the container",
                    "properties": {
                      "role": {
                        "type": "string",
                        "description": "Role is a SELinux role label that applies to the container."
                      },
                      "type": {
                        "type": "string",
                        "description": "Type is a SELinux type label that applies to the container."
                      },
                      "user": {
                        "type": "string",
                        "description": "User is a SELinux user label that applies to the container."
                      },
                      "level": {
                        "type": "string",
                        "description": "Level is SELinux level label that applies to the container."
                      }
                    }
                  },
                  "privileged": {
                    "type": "boolean",
                    "description": "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false."
                  }
                }
              },
              "lifecycle": {
                "description": "Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.",
                "properties": {
                  "preStop": {
                    "description": "Handler defines a specific action that should be taken",
                    "properties": {
                      "httpGet": {
                        "required": [
                          "port"
                        ],
                        "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                        "properties": {
                          "path": {
                            "type": "string",
                            "description": "Path to access on the HTTP server."
                          },
                          "host": {
                            "type": "string",
                            "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                          },
                          "scheme": {
                            "type": "string",
                            "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                          },
                          "httpHeaders": {
                            "items": {
                              "required": [
                                "name",
                                "value"
                              ],
                              "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                              "properties": {
                                "name": {
                                  "type": "string",
                                  "description": "The header field name"
                                },
                                "value": {
                                  "type": "string",
                                  "description": "The header field value"
                                }
                              }
                            },
                            "type": "array",
                            "description": "Custom headers to set in the request. HTTP allows repeated headers."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "tcpSocket": {
                        "required": [
                          "port"
                        ],
                        "description": "TCPSocketAction describes an action based on opening a socket",
                        "properties": {
                          "host": {
                            "type": "string",
                            "description": "Optional: Host name to connect to, defaults to the pod IP."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "exec": {
                        "description": "ExecAction describes a \"run in container\" action.",
                        "properties": {
                          "command": {
                            "items": {
                              "type": "string"
                            },
                            "type": "array",
                            "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                          }
                        }
                      }
                    }
                  },
                  "postStart": {
                    "description": "Handler defines a specific action that should be taken",
                    "properties": {
                      "httpGet": {
                        "required": [
                          "port"
                        ],
                        "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                        "properties": {
                          "path": {
                            "type": "string",
                            "description": "Path to access on the HTTP server."
                          },
                          "host": {
                            "type": "string",
                            "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                          },
                          "scheme": {
                            "type": "string",
                            "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                          },
                          "httpHeaders": {
                            "items": {
                              "required": [
                                "name",
                                "value"
                              ],
                              "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                              "properties": {
                                "name": {
                                  "type": "string",
                                  "description": "The header field name"
                                },
                                "value": {
                                  "type": "string",
                                  "description": "The header field value"
                                }
                              }
                            },
                            "type": "array",
                            "description": "Custom headers to set in the request. HTTP allows repeated headers."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "tcpSocket": {
                        "required": [
                          "port"
                        ],
                        "description": "TCPSocketAction describes an action based on opening a socket",
                        "properties": {
                          "host": {
                            "type": "string",
                            "description": "Optional: Host name to connect to, defaults to the pod IP."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "exec": {
                        "description": "ExecAction describes a \"run in container\" action.",
                        "properties": {
                          "command": {
                            "items": {
                              "type": "string"
                            },
                            "type": "array",
                            "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                          }
                        }
                      }
                    }
                  }
                }
              },
              "ports": {
                "items": {
                  "required": [
                    "containerPort"
                  ],
                  "description": "ContainerPort represents a network port in a single container.",
                  "properties": {
                    "hostIP": {
                      "type": "string",
                      "description": "What host IP to bind the external port to."
                    },
                    "protocol": {
                      "type": "string",
                      "description": "Protocol for port. Must be UDP or TCP. Defaults to \"TCP\"."
                    },
                    "containerPort": {
                      "type": "integer",
                      "description": "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
                      "format": "int32"
                    },
                    "name": {
                      "type": "string",
                      "description": "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services."
                    },
                    "hostPort": {
                      "type": "integer",
                      "description": "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
                      "format": "int32"
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "containerPort",
                "type": "array",
                "description": "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \"0.0.0.0\" address inside a container will be accessible from the network. Cannot be updated.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "resources": {
                "description": "ResourceRequirements describes the compute resource requirements.",
                "properties": {
                  "requests": {
                    "additionalProperties": {
                      "oneOf": [
                        {
                          "type": "string"
                        },
                        {
                          "type": "integer"
                        }
                      ]
                    },
                    "type": "object",
                    "description": "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/"
                  },
                  "limits": {
                    "additionalProperties": {
                      "oneOf": [
                        {
                          "type": "string"
                        },
                        {
                          "type": "integer"
                        }
                      ]
                    },
                    "type": "object",
                    "description": "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/"
                  }
                }
              },
              "stdinOnce": {
                "type": "boolean",
                "description": "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false"
              }
            }
          },
          "x-kubernetes-patch-merge-key": "name",
          "type": "array",
          "description": "List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.",
          "x-kubernetes-patch-strategy": "merge"
        },
        "serviceAccountName": {
          "type": "string",
          "description": "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/"
        },
        "schedulerName": {
          "type": "string",
          "description": "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler."
        },
        "hostIPC": {
          "type": "boolean",
          "description": "Use the host's ipc namespace. Optional: Default to false."
        },
        "dnsConfig": {
          "description": "PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.",
          "properties": {
            "nameservers": {
              "items": {
                "type": "string"
              },
              "type": "array",
              "description": "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed."
            },
            "searches": {
              "items": {
                "type": "string"
              },
              "type": "array",
              "description": "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed."
            },
            "options": {
              "items": {
                "description": "PodDNSConfigOption defines DNS resolver options of a pod.",
                "properties": {
                  "name": {
                    "type": "string",
                    "description": "Required."
                  },
                  "value": {
                    "type": "string"
                  }
                }
              },
              "type": "array",
              "description": "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy."
            }
          }
        },
        "activeDeadlineSeconds": {
          "type": "integer",
          "description": "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
          "format": "int64"
        },
        "terminationGracePeriodSeconds": {
          "type": "integer",
          "description": "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",
          "format": "int64"
        },
        "hostPID": {
          "type": "boolean",
          "description": "Use the host's pid namespace. Optional: Default to false."
        },
        "volumes": {
          "items": {
            "required": [
              "name"
            ],
            "description": "Volume represents a named volume in a pod that may be accessed by any container in the pod.",
            "properties": {
              "portworxVolume": {
                "required": [
                  "volumeID"
                ],
                "description": "PortworxVolumeSource represents a Portworx volume resource.",
                "properties": {
                  "readOnly": {
                    "type": "boolean",
                    "description": "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
                  },
                  "volumeID": {
                    "type": "string",
                    "description": "VolumeID uniquely identifies a Portworx volume"
                  },
                  "fsType": {
                    "type": "string",
                    "description": "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\". Implicitly inferred to be \"ext4\" if unspecified."
                  }
                }
              },
              "glusterfs": {
                "required": [
                  "endpoints",
                  "path"
                ],
                "description": "Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.",
                "properties": {
                  "path": {
                    "type": "string",
                    "description": "Path is the Glusterfs volume path. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod"
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod"
                  },
                  "endpoints": {
                    "type": "string",
                    "description": "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod"
                  }
                }
              },
              "gitRepo": {
                "required": [
                  "repository"
                ],
                "description": "Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.",
                "properties": {
                  "directory": {
                    "type": "string",
                    "description": "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name."
                  },
                  "repository": {
                    "type": "string",
                    "description": "Repository URL"
                  },
                  "revision": {
                    "type": "string",
                    "description": "Commit hash for the specified revision."
                  }
                }
              },
              "flocker": {
                "description": "Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.",
                "properties": {
                  "datasetName": {
                    "type": "string",
                    "description": "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated"
                  },
                  "datasetUUID": {
                    "type": "string",
                    "description": "UUID of the dataset. This is unique identifier of a Flocker dataset"
                  }
                }
              },
              "storageos": {
                "description": "Represents a StorageOS persistent volume resource.",
                "properties": {
                  "volumeName": {
                    "type": "string",
                    "description": "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace."
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
                  },
                  "volumeNamespace": {
                    "type": "string",
                    "description": "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to \"default\" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created."
                  },
                  "secretRef": {
                    "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
                    "properties": {
                      "name": {
                        "type": "string",
                        "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                      }
                    }
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified."
                  }
                }
              },
              "iscsi": {
                "required": [
                  "targetPortal",
                  "iqn",
                  "lun"
                ],
                "description": "Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "targetPortal": {
                    "type": "string",
                    "description": "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260)."
                  },
                  "portals": {
                    "items": {
                      "type": "string"
                    },
                    "type": "array",
                    "description": "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260)."
                  },
                  "secretRef": {
                    "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
                    "properties": {
                      "name": {
                        "type": "string",
                        "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                      }
                    }
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi"
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false."
                  },
                  "chapAuthSession": {
                    "type": "boolean",
                    "description": "whether support iSCSI Session CHAP authentication"
                  },
                  "initiatorName": {
                    "type": "string",
                    "description": "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection."
                  },
                  "iscsiInterface": {
                    "type": "string",
                    "description": "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp)."
                  },
                  "iqn": {
                    "type": "string",
                    "description": "Target iSCSI Qualified Name."
                  },
                  "chapAuthDiscovery": {
                    "type": "boolean",
                    "description": "whether support iSCSI Discovery CHAP authentication"
                  },
                  "lun": {
                    "type": "integer",
                    "description": "iSCSI Target Lun number.",
                    "format": "int32"
                  }
                }
              },
              "projected": {
                "required": [
                  "sources"
                ],
                "description": "Represents a projected volume source",
                "properties": {
                  "sources": {
                    "items": {
                      "description": "Projection that may be projected along with other supported volume types",
                      "properties": {
                        "configMap": {
                          "description": "Adapts a ConfigMap into a projected volume.\n\nThe contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.",
                          "properties": {
                            "items": {
                              "items": {
                                "required": [
                                  "key",
                                  "path"
                                ],
                                "description": "Maps a string key to a path within a volume.",
                                "properties": {
                                  "path": {
                                    "type": "string",
                                    "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'."
                                  },
                                  "mode": {
                                    "type": "integer",
                                    "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                    "format": "int32"
                                  },
                                  "key": {
                                    "type": "string",
                                    "description": "The key to project."
                                  }
                                }
                              },
                              "type": "array",
                              "description": "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'."
                            },
                            "optional": {
                              "type": "boolean",
                              "description": "Specify whether the ConfigMap or it's keys must be defined"
                            },
                            "name": {
                              "type": "string",
                              "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                            }
                          }
                        },
                        "secret": {
                          "description": "Adapts a secret into a projected volume.\n\nThe contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.",
                          "properties": {
                            "items": {
                              "items": {
                                "required": [
                                  "key",
                                  "path"
                                ],
                                "description": "Maps a string key to a path within a volume.",
                                "properties": {
                                  "path": {
                                    "type": "string",
                                    "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'."
                                  },
                                  "mode": {
                                    "type": "integer",
                                    "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                    "format": "int32"
                                  },
                                  "key": {
                                    "type": "string",
                                    "description": "The key to project."
                                  }
                                }
                              },
                              "type": "array",
                              "description": "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'."
                            },
                            "optional": {
                              "type": "boolean",
                              "description": "Specify whether the Secret or its key must be defined"
                            },
                            "name": {
                              "type": "string",
                              "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                            }
                          }
                        },
                        "downwardAPI": {
                          "description": "Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.",
                          "properties": {
                            "items": {
                              "items": {
                                "required": [
                                  "path"
                                ],
                                "description": "DownwardAPIVolumeFile represents information to create the file containing the pod field",
                                "properties": {
                                  "path": {
                                    "type": "string",
                                    "description": "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'"
                                  },
                                  "fieldRef": {
                                    "required": [
                                      "fieldPath"
                                    ],
                                    "description": "ObjectFieldSelector selects an APIVersioned field of an object.",
                                    "properties": {
                                      "fieldPath": {
                                        "type": "string",
                                        "description": "Path of the field to select in the specified API version."
                                      },
                                      "apiVersion": {
                                        "type": "string",
                                        "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\"."
                                      }
                                    }
                                  },
                                  "mode": {
                                    "type": "integer",
                                    "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                                    "format": "int32"
                                  },
                                  "resourceFieldRef": {
                                    "required": [
                                      "resource"
                                    ],
                                    "description": "ResourceFieldSelector represents container resources (cpu, memory) and their output format",
                                    "properties": {
                                      "containerName": {
                                        "type": "string",
                                        "description": "Container name: required for volumes, optional for env vars"
                                      },
                                      "resource": {
                                        "type": "string",
                                        "description": "Required: resource to select"
                                      },
                                      "divisor": {
                                        "oneOf": [
                                          {
                                            "type": "string"
                                          },
                                          {
                                            "type": "integer"
                                          }
                                        ]
                                      }
                                    }
                                  }
                                }
                              },
                              "type": "array",
                              "description": "Items is a list of DownwardAPIVolume file"
                            }
                          }
                        }
                      }
                    },
                    "type": "array",
                    "description": "list of volume projections"
                  },
                  "defaultMode": {
                    "type": "integer",
                    "description": "Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                    "format": "int32"
                  }
                }
              },
              "secret": {
                "description": "Adapts a Secret into a volume.\n\nThe contents of the target Secret's Data field will be presented in a volume as files using the keys in the Data field as the file names. Secret volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "items": {
                    "items": {
                      "required": [
                        "key",
                        "path"
                      ],
                      "description": "Maps a string key to a path within a volume.",
                      "properties": {
                        "path": {
                          "type": "string",
                          "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'."
                        },
                        "mode": {
                          "type": "integer",
                          "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                          "format": "int32"
                        },
                        "key": {
                          "type": "string",
                          "description": "The key to project."
                        }
                      }
                    },
                    "type": "array",
                    "description": "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'."
                  },
                  "optional": {
                    "type": "boolean",
                    "description": "Specify whether the Secret or it's keys must be defined"
                  },
                  "defaultMode": {
                    "type": "integer",
                    "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                    "format": "int32"
                  },
                  "secretName": {
                    "type": "string",
                    "description": "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret"
                  }
                }
              },
              "flexVolume": {
                "required": [
                  "driver"
                ],
                "description": "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
                "properties": {
                  "readOnly": {
                    "type": "boolean",
                    "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
                  },
                  "secretRef": {
                    "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
                    "properties": {
                      "name": {
                        "type": "string",
                        "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                      }
                    }
                  },
                  "driver": {
                    "type": "string",
                    "description": "Driver is the name of the driver to use for this volume."
                  },
                  "options": {
                    "additionalProperties": {
                      "type": "string"
                    },
                    "type": "object",
                    "description": "Optional: Extra command options if any."
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". The default filesystem depends on FlexVolume script."
                  }
                }
              },
              "photonPersistentDisk": {
                "required": [
                  "pdID"
                ],
                "description": "Represents a Photon Controller persistent disk resource.",
                "properties": {
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified."
                  },
                  "pdID": {
                    "type": "string",
                    "description": "ID that identifies Photon Controller persistent disk"
                  }
                }
              },
              "azureDisk": {
                "required": [
                  "diskName",
                  "diskURI"
                ],
                "description": "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
                "properties": {
                  "diskName": {
                    "type": "string",
                    "description": "The Name of the data disk in the blob storage"
                  },
                  "cachingMode": {
                    "type": "string",
                    "description": "Host Caching mode: None, Read Only, Read Write."
                  },
                  "kind": {
                    "type": "string",
                    "description": "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared"
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified."
                  },
                  "diskURI": {
                    "type": "string",
                    "description": "The URI the data disk in the blob storage"
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
                  }
                }
              },
              "fc": {
                "description": "Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "targetWWNs": {
                    "items": {
                      "type": "string"
                    },
                    "type": "array",
                    "description": "Optional: FC target worldwide names (WWNs)"
                  },
                  "wwids": {
                    "items": {
                      "type": "string"
                    },
                    "type": "array",
                    "description": "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously."
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
                  },
                  "lun": {
                    "type": "integer",
                    "description": "Optional: FC target lun number",
                    "format": "int32"
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified."
                  }
                }
              },
              "scaleIO": {
                "required": [
                  "gateway",
                  "system",
                  "secretRef"
                ],
                "description": "ScaleIOVolumeSource represents a persistent ScaleIO volume",
                "properties": {
                  "storageMode": {
                    "type": "string",
                    "description": "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned."
                  },
                  "secretRef": {
                    "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
                    "properties": {
                      "name": {
                        "type": "string",
                        "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                      }
                    }
                  },
                  "protectionDomain": {
                    "type": "string",
                    "description": "The name of the ScaleIO Protection Domain for the configured storage."
                  },
                  "volumeName": {
                    "type": "string",
                    "description": "The name of a volume already created in the ScaleIO system that is associated with this volume source."
                  },
                  "sslEnabled": {
                    "type": "boolean",
                    "description": "Flag to enable/disable SSL communication with Gateway, default false"
                  },
                  "system": {
                    "type": "string",
                    "description": "The name of the storage system as configured in ScaleIO."
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified."
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
                  },
                  "storagePool": {
                    "type": "string",
                    "description": "The ScaleIO Storage Pool associated with the protection domain."
                  },
                  "gateway": {
                    "type": "string",
                    "description": "The host address of the ScaleIO API Gateway."
                  }
                }
              },
              "emptyDir": {
                "description": "Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "sizeLimit": {
                    "oneOf": [
                      {
                        "type": "string"
                      },
                      {
                        "type": "integer"
                      }
                    ]
                  },
                  "medium": {
                    "type": "string",
                    "description": "What type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir"
                  }
                }
              },
              "persistentVolumeClaim": {
                "required": [
                  "claimName"
                ],
                "description": "PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).",
                "properties": {
                  "readOnly": {
                    "type": "boolean",
                    "description": "Will force the ReadOnly setting in VolumeMounts. Default false."
                  },
                  "claimName": {
                    "type": "string",
                    "description": "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims"
                  }
                }
              },
              "configMap": {
                "description": "Adapts a ConfigMap into a volume.\n\nThe contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "items": {
                    "items": {
                      "required": [
                        "key",
                        "path"
                      ],
                      "description": "Maps a string key to a path within a volume.",
                      "properties": {
                        "path": {
                          "type": "string",
                          "description": "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'."
                        },
                        "mode": {
                          "type": "integer",
                          "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                          "format": "int32"
                        },
                        "key": {
                          "type": "string",
                          "description": "The key to project."
                        }
                      }
                    },
                    "type": "array",
                    "description": "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'."
                  },
                  "optional": {
                    "type": "boolean",
                    "description": "Specify whether the ConfigMap or it's keys must be defined"
                  },
                  "defaultMode": {
                    "type": "integer",
                    "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                    "format": "int32"
                  },
                  "name": {
                    "type": "string",
                    "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                  }
                }
              },
              "cephfs": {
                "required": [
                  "monitors"
                ],
                "description": "Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.",
                "properties": {
                  "secretRef": {
                    "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
                    "properties": {
                      "name": {
                        "type": "string",
                        "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                      }
                    }
                  },
                  "secretFile": {
                    "type": "string",
                    "description": "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it"
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it"
                  },
                  "user": {
                    "type": "string",
                    "description": "Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it"
                  },
                  "path": {
                    "type": "string",
                    "description": "Optional: Used as the mounted root, rather than the full Ceph tree, default is /"
                  },
                  "monitors": {
                    "items": {
                      "type": "string"
                    },
                    "type": "array",
                    "description": "Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it"
                  }
                }
              },
              "name": {
                "type": "string",
                "description": "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
              },
              "azureFile": {
                "required": [
                  "secretName",
                  "shareName"
                ],
                "description": "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
                "properties": {
                  "shareName": {
                    "type": "string",
                    "description": "Share Name"
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts."
                  },
                  "secretName": {
                    "type": "string",
                    "description": "the name of secret that contains Azure Storage Account Name and Key"
                  }
                }
              },
              "quobyte": {
                "required": [
                  "registry",
                  "volume"
                ],
                "description": "Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.",
                "properties": {
                  "volume": {
                    "type": "string",
                    "description": "Volume is a string that references an already created Quobyte volume by name."
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false."
                  },
                  "group": {
                    "type": "string",
                    "description": "Group to map volume access to Default is no group"
                  },
                  "registry": {
                    "type": "string",
                    "description": "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes"
                  },
                  "user": {
                    "type": "string",
                    "description": "User to map volume access to Defaults to serivceaccount user"
                  }
                }
              },
              "hostPath": {
                "required": [
                  "path"
                ],
                "description": "Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.",
                "properties": {
                  "path": {
                    "type": "string",
                    "description": "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath"
                  },
                  "type": {
                    "type": "string",
                    "description": "Type for HostPath Volume Defaults to \"\" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath"
                  }
                }
              },
              "nfs": {
                "required": [
                  "server",
                  "path"
                ],
                "description": "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
                "properties": {
                  "path": {
                    "type": "string",
                    "description": "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs"
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs"
                  },
                  "server": {
                    "type": "string",
                    "description": "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs"
                  }
                }
              },
              "gcePersistentDisk": {
                "required": [
                  "pdName"
                ],
                "description": "Represents a Persistent Disk resource in Google Compute Engine.\n\nA GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.",
                "properties": {
                  "readOnly": {
                    "type": "boolean",
                    "description": "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk"
                  },
                  "partition": {
                    "type": "integer",
                    "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
                    "format": "int32"
                  },
                  "pdName": {
                    "type": "string",
                    "description": "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk"
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk"
                  }
                }
              },
              "cinder": {
                "required": [
                  "volumeID"
                ],
                "description": "Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "readOnly": {
                    "type": "boolean",
                    "description": "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md"
                  },
                  "volumeID": {
                    "type": "string",
                    "description": "volume id used to identify the volume in cinder More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md"
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md"
                  }
                }
              },
              "awsElasticBlockStore": {
                "required": [
                  "volumeID"
                ],
                "description": "Represents a Persistent Disk resource in AWS.\n\nAn AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "readOnly": {
                    "type": "boolean",
                    "description": "Specify \"true\" to force and set the ReadOnly property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore"
                  },
                  "partition": {
                    "type": "integer",
                    "description": "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty).",
                    "format": "int32"
                  },
                  "volumeID": {
                    "type": "string",
                    "description": "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore"
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore"
                  }
                }
              },
              "rbd": {
                "required": [
                  "monitors",
                  "image"
                ],
                "description": "Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "secretRef": {
                    "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
                    "properties": {
                      "name": {
                        "type": "string",
                        "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                      }
                    }
                  },
                  "image": {
                    "type": "string",
                    "description": "The rados image name. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it"
                  },
                  "keyring": {
                    "type": "string",
                    "description": "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it"
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd"
                  },
                  "readOnly": {
                    "type": "boolean",
                    "description": "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it"
                  },
                  "user": {
                    "type": "string",
                    "description": "The rados user name. Default is admin. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it"
                  },
                  "monitors": {
                    "items": {
                      "type": "string"
                    },
                    "type": "array",
                    "description": "A collection of Ceph monitors. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it"
                  },
                  "pool": {
                    "type": "string",
                    "description": "The rados pool name. Default is rbd. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it"
                  }
                }
              },
              "downwardAPI": {
                "description": "DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.",
                "properties": {
                  "items": {
                    "items": {
                      "required": [
                        "path"
                      ],
                      "description": "DownwardAPIVolumeFile represents information to create the file containing the pod field",
                      "properties": {
                        "path": {
                          "type": "string",
                          "description": "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'"
                        },
                        "fieldRef": {
                          "required": [
                            "fieldPath"
                          ],
                          "description": "ObjectFieldSelector selects an APIVersioned field of an object.",
                          "properties": {
                            "fieldPath": {
                              "type": "string",
                              "description": "Path of the field to select in the specified API version."
                            },
                            "apiVersion": {
                              "type": "string",
                              "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\"."
                            }
                          }
                        },
                        "mode": {
                          "type": "integer",
                          "description": "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                          "format": "int32"
                        },
                        "resourceFieldRef": {
                          "required": [
                            "resource"
                          ],
                          "description": "ResourceFieldSelector represents container resources (cpu, memory) and their output format",
                          "properties": {
                            "containerName": {
                              "type": "string",
                              "description": "Container name: required for volumes, optional for env vars"
                            },
                            "resource": {
                              "type": "string",
                              "description": "Required: resource to select"
                            },
                            "divisor": {
                              "oneOf": [
                                {
                                  "type": "string"
                                },
                                {
                                  "type": "integer"
                                }
                              ]
                            }
                          }
                        }
                      }
                    },
                    "type": "array",
                    "description": "Items is a list of downward API volume file"
                  },
                  "defaultMode": {
                    "type": "integer",
                    "description": "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
                    "format": "int32"
                  }
                }
              },
              "vsphereVolume": {
                "required": [
                  "volumePath"
                ],
                "description": "Represents a vSphere volume resource.",
                "properties": {
                  "storagePolicyName": {
                    "type": "string",
                    "description": "Storage Policy Based Management (SPBM) profile name."
                  },
                  "fsType": {
                    "type": "string",
                    "description": "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified."
                  },
                  "storagePolicyID": {
                    "type": "string",
                    "description": "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName."
                  },
                  "volumePath": {
                    "type": "string",
                    "description": "Path that identifies vSphere volume vmdk"
                  }
                }
              }
            }
          },
          "x-kubernetes-patch-merge-key": "name",
          "type": "array",
          "description": "List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes",
          "x-kubernetes-patch-strategy": "merge,retainKeys"
        },
        "initContainers": {
          "items": {
            "required": [
              "name"
            ],
            "description": "A single application container that you want to run within a pod.",
            "properties": {
              "livenessProbe": {
                "description": "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
                "properties": {
                  "httpGet": {
                    "required": [
                      "port"
                    ],
                    "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                    "properties": {
                      "path": {
                        "type": "string",
                        "description": "Path to access on the HTTP server."
                      },
                      "host": {
                        "type": "string",
                        "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                      },
                      "scheme": {
                        "type": "string",
                        "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                      },
                      "httpHeaders": {
                        "items": {
                          "required": [
                            "name",
                            "value"
                          ],
                          "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                          "properties": {
                            "name": {
                              "type": "string",
                              "description": "The header field name"
                            },
                            "value": {
                              "type": "string",
                              "description": "The header field value"
                            }
                          }
                        },
                        "type": "array",
                        "description": "Custom headers to set in the request. HTTP allows repeated headers."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "timeoutSeconds": {
                    "type": "integer",
                    "description": "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "exec": {
                    "description": "ExecAction describes a \"run in container\" action.",
                    "properties": {
                      "command": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                      }
                    }
                  },
                  "initialDelaySeconds": {
                    "type": "integer",
                    "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "tcpSocket": {
                    "required": [
                      "port"
                    ],
                    "description": "TCPSocketAction describes an action based on opening a socket",
                    "properties": {
                      "host": {
                        "type": "string",
                        "description": "Optional: Host name to connect to, defaults to the pod IP."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "periodSeconds": {
                    "type": "integer",
                    "description": "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
                    "format": "int32"
                  },
                  "successThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
                    "format": "int32"
                  },
                  "failureThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
                    "format": "int32"
                  }
                }
              },
              "args": {
                "items": {
                  "type": "string"
                },
                "type": "array",
                "description": "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell"
              },
              "terminationMessagePath": {
                "type": "string",
                "description": "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated."
              },
              "name": {
                "type": "string",
                "description": "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated."
              },
              "envFrom": {
                "items": {
                  "description": "EnvFromSource represents the source of a set of ConfigMaps",
                  "properties": {
                    "prefix": {
                      "type": "string",
                      "description": "An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER."
                    },
                    "configMapRef": {
                      "description": "ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.\n\nThe contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.",
                      "properties": {
                        "optional": {
                          "type": "boolean",
                          "description": "Specify whether the ConfigMap must be defined"
                        },
                        "name": {
                          "type": "string",
                          "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                        }
                      }
                    },
                    "secretRef": {
                      "description": "SecretEnvSource selects a Secret to populate the environment variables with.\n\nThe contents of the target Secret's Data field will represent the key-value pairs as environment variables.",
                      "properties": {
                        "optional": {
                          "type": "boolean",
                          "description": "Specify whether the Secret must be defined"
                        },
                        "name": {
                          "type": "string",
                          "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                        }
                      }
                    }
                  }
                },
                "type": "array",
                "description": "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated."
              },
              "workingDir": {
                "type": "string",
                "description": "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated."
              },
              "image": {
                "type": "string",
                "description": "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets."
              },
              "stdin": {
                "type": "boolean",
                "description": "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false."
              },
              "volumeMounts": {
                "items": {
                  "required": [
                    "name",
                    "mountPath"
                  ],
                  "description": "VolumeMount describes a mounting of a Volume within a container.",
                  "properties": {
                    "readOnly": {
                      "type": "boolean",
                      "description": "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false."
                    },
                    "mountPropagation": {
                      "type": "string",
                      "description": "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationHostToContainer is used. This field is alpha in 1.8 and can be reworked or removed in a future release."
                    },
                    "subPath": {
                      "type": "string",
                      "description": "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root)."
                    },
                    "name": {
                      "type": "string",
                      "description": "This must match the Name of a Volume."
                    },
                    "mountPath": {
                      "type": "string",
                      "description": "Path within the container at which the volume should be mounted.  Must not contain ':'."
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "mountPath",
                "type": "array",
                "description": "Pod volumes to mount into the container's filesystem. Cannot be updated.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "tty": {
                "type": "boolean",
                "description": "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false."
              },
              "terminationMessagePolicy": {
                "type": "string",
                "description": "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated."
              },
              "volumeDevices": {
                "items": {
                  "required": [
                    "name",
                    "devicePath"
                  ],
                  "description": "volumeDevice describes a mapping of a raw block device within a container.",
                  "properties": {
                    "devicePath": {
                      "type": "string",
                      "description": "devicePath is the path inside of the container that the device will be mapped to."
                    },
                    "name": {
                      "type": "string",
                      "description": "name must match the name of a persistentVolumeClaim in the pod"
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "devicePath",
                "type": "array",
                "description": "volumeDevices is the list of block devices to be used by the container. This is an alpha feature and may change in the future.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "command": {
                "items": {
                  "type": "string"
                },
                "type": "array",
                "description": "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell"
              },
              "env": {
                "items": {
                  "required": [
                    "name"
                  ],
                  "description": "EnvVar represents an environment variable present in a Container.",
                  "properties": {
                    "valueFrom": {
                      "description": "EnvVarSource represents a source for the value of an EnvVar.",
                      "properties": {
                        "secretKeyRef": {
                          "required": [
                            "key"
                          ],
                          "description": "SecretKeySelector selects a key of a Secret.",
                          "properties": {
                            "optional": {
                              "type": "boolean",
                              "description": "Specify whether the Secret or it's key must be defined"
                            },
                            "name": {
                              "type": "string",
                              "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                            },
                            "key": {
                              "type": "string",
                              "description": "The key of the secret to select from.  Must be a valid secret key."
                            }
                          }
                        },
                        "fieldRef": {
                          "required": [
                            "fieldPath"
                          ],
                          "description": "ObjectFieldSelector selects an APIVersioned field of an object.",
                          "properties": {
                            "fieldPath": {
                              "type": "string",
                              "description": "Path of the field to select in the specified API version."
                            },
                            "apiVersion": {
                              "type": "string",
                              "description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\"."
                            }
                          }
                        },
                        "configMapKeyRef": {
                          "required": [
                            "key"
                          ],
                          "description": "Selects a key from a ConfigMap.",
                          "properties": {
                            "optional": {
                              "type": "boolean",
                              "description": "Specify whether the ConfigMap or it's key must be defined"
                            },
                            "name": {
                              "type": "string",
                              "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
                            },
                            "key": {
                              "type": "string",
                              "description": "The key to select."
                            }
                          }
                        },
                        "resourceFieldRef": {
                          "required": [
                            "resource"
                          ],
                          "description": "ResourceFieldSelector represents container resources (cpu, memory) and their output format",
                          "properties": {
                            "containerName": {
                              "type": "string",
                              "description": "Container name: required for volumes, optional for env vars"
                            },
                            "resource": {
                              "type": "string",
                              "description": "Required: resource to select"
                            },
                            "divisor": {
                              "oneOf": [
                                {
                                  "type": "string"
                                },
                                {
                                  "type": "integer"
                                }
                              ]
                            }
                          }
                        }
                      }
                    },
                    "name": {
                      "type": "string",
                      "description": "Name of the environment variable. Must be a C_IDENTIFIER."
                    },
                    "value": {
                      "type": "string",
                      "description": "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\"."
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "name",
                "type": "array",
                "description": "List of environment variables to set in the container. Cannot be updated.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "imagePullPolicy": {
                "type": "string",
                "description": "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images"
              },
              "readinessProbe": {
                "description": "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
                "properties": {
                  "httpGet": {
                    "required": [
                      "port"
                    ],
                    "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                    "properties": {
                      "path": {
                        "type": "string",
                        "description": "Path to access on the HTTP server."
                      },
                      "host": {
                        "type": "string",
                        "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                      },
                      "scheme": {
                        "type": "string",
                        "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                      },
                      "httpHeaders": {
                        "items": {
                          "required": [
                            "name",
                            "value"
                          ],
                          "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                          "properties": {
                            "name": {
                              "type": "string",
                              "description": "The header field name"
                            },
                            "value": {
                              "type": "string",
                              "description": "The header field value"
                            }
                          }
                        },
                        "type": "array",
                        "description": "Custom headers to set in the request. HTTP allows repeated headers."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "timeoutSeconds": {
                    "type": "integer",
                    "description": "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "exec": {
                    "description": "ExecAction describes a \"run in container\" action.",
                    "properties": {
                      "command": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                      }
                    }
                  },
                  "initialDelaySeconds": {
                    "type": "integer",
                    "description": "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
                    "format": "int32"
                  },
                  "tcpSocket": {
                    "required": [
                      "port"
                    ],
                    "description": "TCPSocketAction describes an action based on opening a socket",
                    "properties": {
                      "host": {
                        "type": "string",
                        "description": "Optional: Host name to connect to, defaults to the pod IP."
                      },
                      "port": {
                        "oneOf": [
                          {
                            "type": "string"
                          },
                          {
                            "type": "integer"
                          }
                        ]
                      }
                    }
                  },
                  "periodSeconds": {
                    "type": "integer",
                    "description": "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
                    "format": "int32"
                  },
                  "successThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
                    "format": "int32"
                  },
                  "failureThreshold": {
                    "type": "integer",
                    "description": "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
                    "format": "int32"
                  }
                }
              },
              "securityContext": {
                "description": "SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.",
                "properties": {
                  "readOnlyRootFilesystem": {
                    "type": "boolean",
                    "description": "Whether this container has a read-only root filesystem. Default is false."
                  },
                  "runAsUser": {
                    "type": "integer",
                    "description": "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
                    "format": "int64"
                  },
                  "allowPrivilegeEscalation": {
                    "type": "boolean",
                    "description": "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN"
                  },
                  "capabilities": {
                    "description": "Adds and removes POSIX capabilities from running containers.",
                    "properties": {
                      "add": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Added capabilities"
                      },
                      "drop": {
                        "items": {
                          "type": "string"
                        },
                        "type": "array",
                        "description": "Removed capabilities"
                      }
                    }
                  },
                  "runAsNonRoot": {
                    "type": "boolean",
                    "description": "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence."
                  },
                  "seLinuxOptions": {
                    "description": "SELinuxOptions are the labels to be applied to the container",
                    "properties": {
                      "role": {
                        "type": "string",
                        "description": "Role is a SELinux role label that applies to the container."
                      },
                      "type": {
                        "type": "string",
                        "description": "Type is a SELinux type label that applies to the container."
                      },
                      "user": {
                        "type": "string",
                        "description": "User is a SELinux user label that applies to the container."
                      },
                      "level": {
                        "type": "string",
                        "description": "Level is SELinux level label that applies to the container."
                      }
                    }
                  },
                  "privileged": {
                    "type": "boolean",
                    "description": "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false."
                  }
                }
              },
              "lifecycle": {
                "description": "Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.",
                "properties": {
                  "preStop": {
                    "description": "Handler defines a specific action that should be taken",
                    "properties": {
                      "httpGet": {
                        "required": [
                          "port"
                        ],
                        "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                        "properties": {
                          "path": {
                            "type": "string",
                            "description": "Path to access on the HTTP server."
                          },
                          "host": {
                            "type": "string",
                            "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                          },
                          "scheme": {
                            "type": "string",
                            "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                          },
                          "httpHeaders": {
                            "items": {
                              "required": [
                                "name",
                                "value"
                              ],
                              "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                              "properties": {
                                "name": {
                                  "type": "string",
                                  "description": "The header field name"
                                },
                                "value": {
                                  "type": "string",
                                  "description": "The header field value"
                                }
                              }
                            },
                            "type": "array",
                            "description": "Custom headers to set in the request. HTTP allows repeated headers."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "tcpSocket": {
                        "required": [
                          "port"
                        ],
                        "description": "TCPSocketAction describes an action based on opening a socket",
                        "properties": {
                          "host": {
                            "type": "string",
                            "description": "Optional: Host name to connect to, defaults to the pod IP."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "exec": {
                        "description": "ExecAction describes a \"run in container\" action.",
                        "properties": {
                          "command": {
                            "items": {
                              "type": "string"
                            },
                            "type": "array",
                            "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                          }
                        }
                      }
                    }
                  },
                  "postStart": {
                    "description": "Handler defines a specific action that should be taken",
                    "properties": {
                      "httpGet": {
                        "required": [
                          "port"
                        ],
                        "description": "HTTPGetAction describes an action based on HTTP Get requests.",
                        "properties": {
                          "path": {
                            "type": "string",
                            "description": "Path to access on the HTTP server."
                          },
                          "host": {
                            "type": "string",
                            "description": "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead."
                          },
                          "scheme": {
                            "type": "string",
                            "description": "Scheme to use for connecting to the host. Defaults to HTTP."
                          },
                          "httpHeaders": {
                            "items": {
                              "required": [
                                "name",
                                "value"
                              ],
                              "description": "HTTPHeader describes a custom header to be used in HTTP probes",
                              "properties": {
                                "name": {
                                  "type": "string",
                                  "description": "The header field name"
                                },
                                "value": {
                                  "type": "string",
                                  "description": "The header field value"
                                }
                              }
                            },
                            "type": "array",
                            "description": "Custom headers to set in the request. HTTP allows repeated headers."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "tcpSocket": {
                        "required": [
                          "port"
                        ],
                        "description": "TCPSocketAction describes an action based on opening a socket",
                        "properties": {
                          "host": {
                            "type": "string",
                            "description": "Optional: Host name to connect to, defaults to the pod IP."
                          },
                          "port": {
                            "oneOf": [
                              {
                                "type": "string"
                              },
                              {
                                "type": "integer"
                              }
                            ]
                          }
                        }
                      },
                      "exec": {
                        "description": "ExecAction describes a \"run in container\" action.",
                        "properties": {
                          "command": {
                            "items": {
                              "type": "string"
                            },
                            "type": "array",
                            "description": "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy."
                          }
                        }
                      }
                    }
                  }
                }
              },
              "ports": {
                "items": {
                  "required": [
                    "containerPort"
                  ],
                  "description": "ContainerPort represents a network port in a single container.",
                  "properties": {
                    "hostIP": {
                      "type": "string",
                      "description": "What host IP to bind the external port to."
                    },
                    "protocol": {
                      "type": "string",
                      "description": "Protocol for port. Must be UDP or TCP. Defaults to \"TCP\"."
                    },
                    "containerPort": {
                      "type": "integer",
                      "description": "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
                      "format": "int32"
                    },
                    "name": {
                      "type": "string",
                      "description": "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services."
                    },
                    "hostPort": {
                      "type": "integer",
                      "description": "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
                      "format": "int32"
                    }
                  }
                },
                "x-kubernetes-patch-merge-key": "containerPort",
                "type": "array",
                "description": "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \"0.0.0.0\" address inside a container will be accessible from the network. Cannot be updated.",
                "x-kubernetes-patch-strategy": "merge"
              },
              "resources": {
                "description": "ResourceRequirements describes the compute resource requirements.",
                "properties": {
                  "requests": {
                    "additionalProperties": {
                      "oneOf": [
                        {
                          "type": "string"
                        },
                        {
                          "type": "integer"
                        }
                      ]
                    },
                    "type": "object",
                    "description": "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/"
                  },
                  "limits": {
                    "additionalProperties": {
                      "oneOf": [
                        {
                          "type": "string"
                        },
                        {
                          "type": "integer"
                        }
                      ]
                    },
                    "type": "object",
                    "description": "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/"
                  }
                }
              },
              "stdinOnce": {
                "type": "boolean",
                "description": "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false"
              }
            }
          },
          "x-kubernetes-patch-merge-key": "name",
          "type": "array",
          "description": "List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
          "x-kubernetes-patch-strategy": "merge"
        },
        "imagePullSecrets": {
          "items": {
            "description": "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
            "properties": {
              "name": {
                "type": "string",
                "description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names"
              }
            }
          },
          "x-kubernetes-patch-merge-key": "name",
          "type": "array",
          "description": "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
          "x-kubernetes-patch-strategy": "merge"
        }
      }
    },
    "apiVersion": {
      "type": "string",
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources"
    },
    "metadata": {
      "description": "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
      "properties": {
        "ownerReferences": {
          "items": {
            "required": [
              "apiVersion",
              "kind",
              "name",
              "uid"
            ],
            "description": "OwnerReference contains enough information to let you identify an owning object. Currently, an owning object must be in the same namespace, so there is no namespace field.",
            "properties": {
              "kind": {
                "type": "string",
                "description": "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
              },
              "uid": {
                "type": "string",
                "description": "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
              },
              "apiVersion": {
                "type": "string",
                "description": "API version of the referent."
              },
              "controller": {
                "type": "boolean",
                "description": "If true, this reference points to the managing controller."
              },
              "blockOwnerDeletion": {
                "type": "boolean",
                "description": "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned."
              },
              "name": {
                "type": "string",
                "description": "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names"
              }
            }
          },
          "x-kubernetes-patch-merge-key": "uid",
          "type": "array",
          "description": "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
          "x-kubernetes-patch-strategy": "merge"
        },
        "uid": {
          "type": "string",
          "description": "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
        },
        "deletionTimestamp": {
          "type": "string",
          "format": "date-time"
        },
        "clusterName": {
          "type": "string",
          "description": "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request."
        },
        "deletionGracePeriodSeconds": {
          "type": "integer",
          "description": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
          "format": "int64"
        },
        "labels": {
          "additionalProperties": {
            "type": "string"
          },
          "type": "object",
          "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels"
        },
        "namespace": {
          "type": "string",
          "enum": ["default"],
          "description": "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces"
        },
        "finalizers": {
          "items": {
            "type": "string"
          },
          "type": "array",
          "description": "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.",
          "x-kubernetes-patch-strategy": "merge"
        },
        "generation": {
          "type": "integer",
          "description": "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
          "format": "int64"
        },
        "initializers": {
          "required": [
            "pending"
          ],
          "description": "Initializers tracks the progress of initialization.",
          "properties": {
            "result": {
              "x-kubernetes-group-version-kind": [
                {
                  "kind": "Status",
                  "version": "v1",
                  "group": ""
                }
              ],
              "description": "Status is a return value for calls that don't return other objects.",
              "properties": {
                "status": {
                  "type": "string",
                  "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status"
                },
                "kind": {
                  "type": "string",
                  "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
                },
                "code": {
                  "type": "integer",
                  "description": "Suggested HTTP return code for this status, 0 if not set.",
                  "format": "int32"
                },
                "apiVersion": {
                  "type": "string",
                  "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources"
                },
                "reason": {
                  "type": "string",
                  "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it."
                },
                "details": {
                  "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
                  "properties": {
                    "kind": {
                      "type": "string",
                      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
                    },
                    "group": {
                      "type": "string",
                      "description": "The group attribute of the resource associated with the status StatusReason."
                    },
                    "name": {
                      "type": "string",
                      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described)."
                    },
                    "retryAfterSeconds": {
                      "type": "integer",
                      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
                      "format": "int32"
                    },
                    "causes": {
                      "items": {
                        "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
                        "properties": {
                          "field": {
                            "type": "string",
                            "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\""
                          },
                          "message": {
                            "type": "string",
                            "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader."
                          },
                          "reason": {
                            "type": "string",
                            "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available."
                          }
                        }
                      },
                      "type": "array",
                      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes."
                    },
                    "uid": {
                      "type": "string",
                      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids"
                    }
                  }
                },
                "message": {
                  "type": "string",
                  "description": "A human-readable description of the status of this operation."
                },
                "metadata": {
                  "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
                  "properties": {
                    "continue": {
                      "type": "string",
                      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response."
                    },
                    "selfLink": {
                      "type": "string",
                      "description": "selfLink is a URL representing this object. Populated by the system. Read-only."
                    },
                    "resourceVersion": {
                      "type": "string",
                      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
                    }
                  }
                }
              }
            },
            "pending": {
              "items": {
                "required": [
                  "name"
                ],
                "description": "Initializer is information about an initializer that has not yet completed.",
                "properties": {
                  "name": {
                    "type": "string",
                    "description": "name of the process that is responsible for initializing this object."
                  }
                }
              },
              "x-kubernetes-patch-merge-key": "name",
              "type": "array",
              "description": "Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.",
              "x-kubernetes-patch-strategy": "merge"
            }
          }
        },
        "resourceVersion": {
          "type": "string",
          "description": "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency"
        },
        "generateName": {
          "type": "string",
          "description": "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency"
        },
        "annotations": {
          "additionalProperties": {
            "type": "string"
          },
          "type": "object",
          "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations"
        },
        "selfLink": {
          "type": "string",
          "description": "SelfLink is a URL representing this object. Populated by the system. Read-only."
        },
        "name": {
          "type": "string",
          "description": "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names"
        }
      }
    }
  }
}
`