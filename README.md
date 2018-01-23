# sts-profile

[![Build Status](https://travis-ci.org/Jalle19/sts-profile.svg?branch=master)](https://travis-ci.org/Jalle19/sts-profile)

So you're using Amazon Web Services and you have MFA enabled for programmatical access too, not just for console 
access. You have a profile specifically for calling `aws sts get-session-token`, which then spits out temporary 
credentials that you use in your "real" profile.

The problem is that `aws sts get-session-token` outputs JSON while your `~/.aws/credentials` file has an INI structure. 
Copy-pasting the three relevant strings manually quickly becomes tedious and error-prone. This tool aims to solve that.

## Usage

Pipe the output from `aws sts get-session-token` to this tool:

```bash
aws sts get-session-token --serial-number arn:aws:foo --token-code 123456 | sts-profile
``` 

This will output something like this:

```
aws_access_key_id = ASIAULTRA_SECRETA
aws_secret_access_key = b14DA7BULTRA_SECRETrk3zxYeQtbrh
aws_session_token = FQoDYULTRA_SECRET
```

Now you just need to copy-paste this into your AWS profile of choice

## Installation

Download the latest release for your platform from the releases page here on GitHub. Unzip the package, place the 
file somewhere and make it executable.

## License

GNU GENERAL PUBLIC LICENSE Version 3