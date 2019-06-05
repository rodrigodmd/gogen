#!/bin/bash

SO=$(uname)
BASE_URL="https://raw.githubusercontent.com/rodrigodmd/..."

bash -l <(curl -s $BASE_URL)
exit 0