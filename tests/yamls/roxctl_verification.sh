#! /bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"

extra_args=()
if [[ -n "$CA" ]]; then
  extra_args+=(--ca "$CA")
fi

FAILED="false"
for yaml in $(ls "$DIR"/*.yaml); do
	NUM_ALERTS="$(roxctl "${extra_args[@]}" -e "$API_ENDPOINT" -p "$ROX_PASSWORD" deployment check --file $yaml --json | \
	    jq '.alerts[].policy.name | select(.=="Latest tag" or .=="No resource requests or limits specified")' | jq -s '. | length')"
	if [[ $NUM_ALERTS != "2" ]]; then
		>&2 echo "Did not find 2 alerts for $yaml"
		FAILED="true"
    else
        echo "Analyzed $yaml successfully"
	fi
done

if [[ "$FAILED" == "true" ]]; then
	echo "Roxctl test failed"
	exit 1
fi
exit 0
