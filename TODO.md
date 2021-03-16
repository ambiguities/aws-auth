Target State:

```
aws-auth

  assume source_profile target_profile role_arn

  mfa source_profile target_profile mfa_serial token_code 
    
  profile

    add target_profile profile_key profile_secret [profile_session]
  
    delete target_profile
  
    update target_profile profile_key profile_secret [profile_session]
  
    view target_profile
```