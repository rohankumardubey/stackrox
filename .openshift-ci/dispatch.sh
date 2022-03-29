#!/usr/bin/env bash

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd)"
source "$ROOT/scripts/ci/lib.sh"

set -euo pipefail

# shellcheck disable=SC1090
source /tmp/secret/stackrox-stackrox-*/credentials

if pr_has_label "delay-tests"; then
    function hold() {
        info "Holding on for debug"
        sleep 3600
    }
    trap hold EXIT
fi

if pr_has_label "debug-tests"; then
    set -x
fi

if [[ "$#" -lt 1 ]]; then
    die "usage: dispatch <ci-job> [<...other parameters...>]"
fi

ci_job="$1"
shift

case "$ci_job" in
    gke-upgrade-tests)
        "$ROOT/.openshift-ci/gke_upgrade_test.py"
        ;;
    *)
        # For ease of initial integration this function does not fail.
        echo "nothing to see here"
        exit 0
esac
