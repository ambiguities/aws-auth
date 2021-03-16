#/usr/bin/env bash

if [ "$AWS_SHARED_CREDENTIALS_FILE" != "" ]; then
    CREDENTIAL_FILE="$AWS_SHARED_CREDENTIALS_FILE"
else
    CREDENTIAL_FILE="$HOME/.aws/credentials"
fi

_commands () {
  echo "assume help mfa profile"
}

_profiles () {
  echo $(cat $CREDENTIAL_FILE | grep -e '\[.*\]' | sed 's/[][]//g')
}

_aws_auth_completions()
{
  if [ "${COMP_CWORD}" -eq "1" ]; then
    COMPREPLY+=($( compgen -W "$( _commands )" -- ${COMP_WORDS[COMP_CWORD]} ))
  elif [ "${COMP_CWORD}" -gt "1" ]; then
    local COMMAND=$(echo ${COMP_LINE} | awk '{ print $2 }')

    if [ "$COMMAND" == "profile" ]; then
      if [ "${COMP_CWORD}" -eq "2" ]; then COMPREPLY+=($( compgen -W "$( _profiles )" -- ${COMP_WORDS[COMP_CWORD]} )); fi
    fi
  fi
}
complete -F _aws_auth_completions aws-auth