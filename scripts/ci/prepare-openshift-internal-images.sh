#!/usr/bin/env bash

# Prepares OpenShift API tests with
# Pushes select images to the given OpenShift Internal Registry in the qa project.
# Usage:
#   ./push-openshift-internal-images.sh <registry> <username> <password>
# Examples:
#   ./push-openshift-internal-images.sh image-registry.openshift-image-registry.svc:5000 kubeadmin mypassword
#   ./push-openshift-internal-images.sh docker-registry.default.svc:5000 admin mypassword

set -euo pipefail

die() {
  echo >&2 "$@"
  exit 1
}

[[ "$#" == 3 ]] || die "Usage: $0 <registry> <username> <password>"

registry="$1"
username="$2"
password="$3"

images=(
  "nginx:1.18.0@sha256:e90ac5331fe095cea01b121a3627174b2e33e06e83720e9a934c7b8ccc9c55a0"
  "quay.io/rhacs-eng/qa:sandbox-jenkins-agent-maven-35-rhel7"
  "quay.io/rhacs-eng/qa:sandbox-nodejs-10"
  "quay.io/rhacs-eng/qa:sandbox-log4j-2-12-2"
)

quay_repo="quay.io/rhacs-eng/"

docker login -u "$username" -p "$password" "$registry"

for full_image in "${images[@]}"
do
  docker pull "$full_image"
  image=${full_image/$quay_repo/""}
  docker tag "$full_image" "$registry/qa/$image"
  docker push "$registry/qa/$image"
done
