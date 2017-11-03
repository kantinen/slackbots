#!/bin/bash

# Usage:
#   update.sh [username] [command]
# Everything but the first two arguements are ignored


set -e

commit_message="[${1}] $ ${2}"

# This is for debugging. Should be left out when pushed to production
echo $commit_message > out.txt

if [ whoami == "kantinebot" ]; then
        git commit -am "${commit_message}"
        git push
fi;
