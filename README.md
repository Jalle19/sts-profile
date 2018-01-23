# sts-profile

[![Build Status](https://travis-ci.org/Jalle19/sts-profile.svg?branch=master)](https://travis-ci.org/Jalle19/sts-profile)

So you're using Amazon Web Services and you have MFA enabled for programmatical access too, not just for console 
access. You have a profile specifically for called `aws sts get-session-token`, which then spits out temporary 
credentials.

The problem is that `get-session-token` outputs JSON while your `~/.aws/credentials` file has an INI structure. This 
means copy-pasting, and this tool solves that

## Usage

Pipe the output from `aws sts get-session-token` to this tool:

```bash
cat example.json | sts-profile
``` 

This will output:

```
aws_access_key_id = ASIAULTRA_SECRETA
aws_secret_access_key = b14DA7BULTRA_SECRETrk3zxYeQtbrh
aws_session_token = FQoDYULTRA_SECRET
```

## Installation

Download the latest release for your platform from the releases page here on GitHub. Unzip the package, place the 
file somewhere and make it executable.

## License

GPL 2.0