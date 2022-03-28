#!/usr/bin/env bash

SCRIPTS_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd)"
source "$SCRIPTS_ROOT/scripts/ci/lib.sh"

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

if [[ "$ci_job" == "push-check" ]]; then
    push_and_check_image "$@"
fi

# For ease of initial integration this function does not fail.
echo "nothing to see here"
exit 0
