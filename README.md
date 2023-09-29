# freemium-engine
Engine for creating Freemium applications with sales automation and a serverless model

## how to import private lib to another project

add this to `.gitconfig` file:

```bash
git config \
 --global \
 url."https://${user}:${personal_access_token}@github.com".insteadOf \
 "https://github.com"
```

and set `user` & `personal_access_token` in environment variables

## TODO

1. Come up with something with the private key so that it is not in the code. Maybe it can be generated and stored somehow based on the data of the current device.

## TODO ideas

1. Add collection of email addresses via a P2P network, to which the developer will have access through asymmetric encryption.
