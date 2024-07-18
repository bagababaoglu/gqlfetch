# gqlfetch
Introspect GraphQL schema from a user-defined server. It uses `github.com/suessflorian/gqlfetch` on the background.

## Flags

### token
Bearer token to use it in the header of the http request in case the server is protected.
### host
URL of the host
### api-path 
Path of the graphql api, defaults to api/graphql/
### output
Destination to save fetched schema file, defaults to schema.graphql in the current working directory
