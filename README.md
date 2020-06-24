## Requirements
- awscli 
  - https://pypi.org/project/awscli/
- jq 
  - https://stedolan.github.io/jq/

## How Tos

### aws-auth-become:
- prerequisites:
  - an active profile in `$HOME/.aws.credentials`
  - the profile credentials are trusted in the role to assume

- usage:
  - ```aws-auth-become profile_from profile_to role_arn [session_name]```

### aws-auth-mfa:
- prerequisites:
  - a default profile in `$HOME/.aws.credentials`

- usage:
  - ```aws-auth-mfa new_profile mfa_code```

### aws-profile:

- usage:
  - ```aws-profile profile_name access_key_id secret_access_key session_token```


#### Note: This was designed to be copied into $HOME/bin. If placed anywhere else, please update the`$HOME/aws-profile` reference to the new path