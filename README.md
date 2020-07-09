### Environment Requirements
- `awscli`
  - https://pypi.org/project/awscli/
- `jq`
  - https://stedolan.github.io/jq/

---

## Scripts

#### `aws-auth-become`

- prerequisites:
  - an active profile in `$HOME/.aws/credentials`
  - the profile credentials are trusted in the role to assume

```sh
# Usage:
aws-auth-become $profile_from $profile_to $role_arn [$session_name]
```

#### `aws-auth-mfa`

- prerequisites:
  - a `[default]` profile in `$HOME/.aws/credentials`
  - the profile has an active access key pair from the mfa user

```sh
# Usage:
aws-auth-mfa $new_profile $mfa_serial $mfa_code
```

#### `aws-profile`
``` sh
# Usage:
aws-profile $profile_name $access_key_id $secret_access_key $session_token [--print]
```

---

## Helper Ideas

#### MFA

```sh
# usage: aws-auth-mfa-username target_profile mfa_code
aws-auth-mfa-username () {
  aws-auth-mfa $1 arn:aws:iam::account_number:mfa/username ${@:2}
}
```

#### Becoming a role

```sh
# usage: aws-auth-become-role-alias from_profile to_profile [session_name]
aws-auth-become-role-alias () { 
  aws-auth-become $1 $2 arn:aws:iam::account_number:role/role_name ${@:3}
}
```

---

#### Note: These scripts were designed to be copied into $HOME/bin. If placed anywhere else, please update the`$HOME/aws-profile` reference to the new path