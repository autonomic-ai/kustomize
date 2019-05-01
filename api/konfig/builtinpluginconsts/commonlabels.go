// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package builtinpluginconsts

const commonLabelFieldSpecs = `
commonLabels:
- path: metadata/labels
  create: true

- path: spec/selector
  create: true
  version: v1
  kind: Service

- path: spec/selector
  create: true
  version: v1
  kind: ReplicationController

- path: spec/template/metadata/labels
  create: true
  version: v1
  kind: ReplicationController

- path: spec/selector/matchLabels
  create: true
  kind: Deployment

- path: spec/template/metadata/labels
  create: true
  kind: Deployment

- path: spec/template/spec/affinity/podAffinity/preferredDuringSchedulingIgnoredDuringExecution/podAffinityTerm/labelSelector/matchLabels
  create: false
  group: apps
  kind: Deployment

- path: spec/template/spec/affinity/podAffinity/requiredDuringSchedulingIgnoredDuringExecution/labelSelector/matchLabels
  create: false
  group: apps
  kind: Deployment

- path: spec/template/spec/affinity/podAntiAffinity/preferredDuringSchedulingIgnoredDuringExecution/podAffinityTerm/labelSelector/matchLabels
  create: false
  group: apps
  kind: Deployment

- path: spec/template/spec/affinity/podAntiAffinity/requiredDuringSchedulingIgnoredDuringExecution/labelSelector/matchLabels
  create: false
  group: apps
  kind: Deployment

- path: spec/selector/matchLabels
  create: true
  kind: ReplicaSet

- path: spec/template/metadata/labels
  create: true
  kind: ReplicaSet

- path: spec/selector/matchLabels
  create: true
  kind: DaemonSet

- path: spec/template/metadata/labels
  create: true
  kind: DaemonSet

- path: spec/selector/matchLabels
  create: true
  group: apps
  kind: StatefulSet

- path: spec/template/metadata/labels
  create: true
  group: apps
  kind: StatefulSet

- path: spec/template/spec/affinity/podAffinity/preferredDuringSchedulingIgnoredDuringExecution/podAffinityTerm/labelSelector/matchLabels
  create: false
  group: apps
  kind: StatefulSet

- path: spec/template/spec/affinity/podAffinity/requiredDuringSchedulingIgnoredDuringExecution/labelSelector/matchLabels
  create: false
  group: apps
  kind: StatefulSet

- path: spec/template/spec/affinity/podAntiAffinity/preferredDuringSchedulingIgnoredDuringExecution/podAffinityTerm/labelSelector/matchLabels
  create: false
  group: apps
  kind: StatefulSet

- path: spec/template/spec/affinity/podAntiAffinity/requiredDuringSchedulingIgnoredDuringExecution/labelSelector/matchLabels
  create: false
  group: apps
  kind: StatefulSet

- path: spec/volumeClaimTemplates[]/metadata/labels
  create: true
  group: apps
  kind: StatefulSet

- path: spec/selector/matchLabels
  create: false
  group: batch
  kind: Job

- path: spec/template/metadata/labels
  create: true
  group: batch
  kind: Job

- path: spec/jobTemplate/spec/selector/matchLabels
  create: false
  group: batch
  kind: CronJob

- path: spec/jobTemplate/metadata/labels
  create: true
  group: batch
  kind: CronJob

- path: spec/jobTemplate/spec/template/metadata/labels
  create: true
  group: batch
  kind: CronJob

- path: spec/selector/matchLabels
  create: false
  group: policy
  kind: PodDisruptionBudget
`
