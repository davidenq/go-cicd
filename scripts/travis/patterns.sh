#!/bin/bash

ENVVAR=${env}

if [[ $ENVVAR != *":"* ]]; then
  echo -e $COLOR_RED"Error: You must define env as follows: env=KEY:VALUE${COLOR_NC}"
fi