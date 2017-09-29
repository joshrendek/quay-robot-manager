# Installation

```
go get github.com/joshrendek/quay-robot-manager
```

You need to setup an OAUTH access token with the appropriate permissions to create/delete/list robots for your org.

Quay's OAUTH system is backwards and you can only do this through their UI.

Once you have that token set it as `BEARER_TOKEN` in your environment.


There is also an optional `--json` flag for all commands that will output in JSON if you want.

## Listing

```
quay-robot-manager list --org YOURORG --json
```

## Create

```
quay-robot-manager create --name testing12311 --org YOURORG --json
```

## DELETE

```
quay-robot-manager delete --name testing12311 --org YOURORG
```

# Building

See [Jenkinsfile](Jenkinsfile)

# License

[MIT LICENSE](LICENSE)

