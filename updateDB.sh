#!/usr/bin/env bash

# Usage:
#   update.sh [username] [command]
# Everything but the first two arguements are ignored

set -e

cd $KANTINE_DB

commit_message="[${1}] - ${2}"

# This is for debugging. Should be left out when pushed to production
if [[ "$GO_ENV" != "production" ]]; then
	echo $commit_message > out.txt
fi

if [ whoami == "kantinebot" ]; then
        git commit -am "${commit_message}"
        git push
fi;
