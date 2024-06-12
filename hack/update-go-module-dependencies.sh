#!/bin/bash -ex

# Background: https://groups.google.com/forum/#!topic/golang-nuts/51-D_YFC78k
#
# TLDR: OCP consists of many repos where for each release a new release branch gets created (release-X.Y).
# When we update vendors we want to get latest changes from the release branch for all of the dependencies.
# With Go modules we can't easily do it, but there is a workaround which consists of multiple steps:
# 	1. Get the latest commit from the branch using `go list -mod=mod -m MODULE@release-x.y`.
# 	2. Using `sed`, transform output of the above command into format accepted by `go mod edit -replace`.
#	3. Modify `go.mod` by calling `go mod edit -replace`.
#
# This needs to happen for each module that uses this branching strategy: all these repos
# live under github.com/openshift organisation.
#
# There are however, some exceptions:
# 	* Some repos under github.com/openshift do not use this strategy.
#     We should skip them in this script and manage directly with `go mod`.
# 	* Some dependencies pin their own dependencies to older commits.
#     For example, dependency Foo from release-4.7 branch requires
#	  dependency Bar at older commit which is
#     not compatible with Bar@release-4.7.
#
# Note that github.com/openshift org also contains forks of K8s upstream repos and we
# use these forks (indirectly in most cases). This means that
# we also must take care of replacing modules such as sigs.k8s.io/cluster-api-provider-azure
# with github.com/openshift/cluster-api-provider-azure (just an example, there are more).

RELEASE=release-4.15
GO_VERSION=1.20

for x in vendor/github.com/openshift/*; do
	case $x in
		# Review the list of special cases on each release.

		# Do not update Hive: it is not part of OCP
		vendor/github.com/openshift/hive)
			;;

		# Replace the installer with our own fork below in this script.
		vendor/github.com/openshift/installer)
			;;

		# It is only used indirectly and intermediate dependencies pin to different incompatible commits.
		# We force a specific commit here to make all dependencies happy.
		vendor/github.com/openshift/cloud-credential-operator)
			go mod edit -replace github.com/openshift/cloud-credential-operator=github.com/openshift/cloud-credential-operator@v0.0.0-20200316201045-d10080b52c9e
			;;

		# MCO is pinned to an old version of MCO, and newer versions don' contain MCO's API anymore, it moved to openshift/api.
		# In order to prevent dependency issues we need to pin to the same version.
		vendor/github.com/openshift/machine-config-operator)
			go mod edit -replace github.com/openshift/machine-config-operator=github.com/openshift/machine-config-operator@v0.0.1-0.20201009041932-4fe8559913b8
			;;

		# This repo doesn't follow release-x.y branching strategy
		# We skip it and let go mod resolve it
		vendor/github.com/openshift/custom-resource-status)
			;;

		*)
			go mod edit -replace "${x##vendor/}"="$(go list -mod=mod -m ${x##vendor/}@$RELEASE | sed -e 's/ /@/')"
			;;
	esac
done

go mod edit -replace sigs.k8s.io/cluster-api="$(go list -mod=mod -m github.com/openshift/cluster-api@$RELEASE | sed -e 's/ /@/')"
for x in aws/v2 azure openstack; do
	go mod edit -replace sigs.k8s.io/cluster-api-provider-$x="$(go list -mod=mod -m github.com/openshift/cluster-api-provider-$x@$RELEASE | sed -e 's/ /@/')"
done

for x in baremetal-operator baremetal-operator/apis baremetal-operator/pkg/hardwareutils cluster-api-provider-baremetal cluster-api-provider-metal3 cluster-api-provider-metal3/api; do
  go mod edit -replace github.com/metal3-io/$x="$(go list -mod=mod -m github.com/openshift/$x@$RELEASE | sed -e 's/ /@/')"
done

go mod edit -replace github.com/openshift/installer="$(go list -mod=mod -m github.com/slintes/installer-aro@aro-hack-4.15 | sed -e 's/ /@/')"

go get ./...

go mod tidy -compat=$GO_VERSION
go mod vendor
