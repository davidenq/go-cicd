#!/bin/bash
ACTION=${action}
ENVVAR=${env}


if [ "$ACTION" == "login" ]; then
  travis login
elif [ "$ACTION" == "encrypt" ]; then

  read KEY VALUE <<<$(IFS=":"; echo $ENVVAR)
  
  if [ "$KEY" == "" ]; then
    echo -e "${COLOR_RED}Error: Key is empty${COLOR_NC}"
  elif [ "$VALUE" == "" ]; then
    echo -e "${COLOR_RED}Error: Value is empty${COLOR_NC}"
  else
    travis encrypt $KEY=\"$VALUE\" --add
  fi
fi