package postgres

//go:generate pgsearchbindings-wrapper --table processindicators --write-options=false --options-path "pkg/search/options/processindicators" --type ProcessIndicator --singular ProcessIndicator --search-category PROCESS_INDICATORS