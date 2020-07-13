#!/usr/bin/env bash

usage() {
	cat <<-EOF
		$(basename "$0") cert-path   # print the text info of [cert-path]
	EOF
}

x509_print() {
	echo "---- $1"
	openssl x509 -in "$1" -text
}

entry() {
	if [[ $# -gt 0 ]]; then
		local f=$1 && shift
		if [[ -f $f ]]; then
			x509_print "$f"
		elif [[ -f "./ci/certs/$f" ]]; then
			x509_print "./ci/certs/$f"
		fi
	else
		usage
	fi
}

entry "$@"
