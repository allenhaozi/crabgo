#!/bin/sh

if [[ "${1#-}" != "$1" ]]; then
    set -- openaios-iam "$@"
fi

exec "$@"