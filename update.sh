#!/usr/bin/env bash

set -e

if [ whoami == "kantinebot" ]; then
        git commit -am "<generic commit message>"
        git push
fi;
