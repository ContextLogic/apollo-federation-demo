## Apollo Federation Demo

This repository is a demo of using Apollo Federation to build a single schema on top of multiple services. The microservices are located under the [`./services`](./services/) folder and the gateway that composes the overall schema is in the [`gateway.js`](./gateway.js) file.

### Installation

```sh
make install
```

This will install the following:
1. `rover` cli first in order to build supergraph
2. node modules to run the apollo/gateway 
3. gqlgen binary to generate GraphQL resolvers for the Golang services

### Start Services

```sh
make -j4 accounts inventory products reviews
```

This command will run all of the microservices at once concurrently. They can be found at http://localhost:4001, http://localhost:4002, http://localhost:4003, and http://localhost:4004.

### Start Gateway
In another terminal window, run the gateway by running this command:

```sh
make gateway
```

This will start up the gateway and serve it at http://localhost:4000

### Managed mode
This example runs Apollo Federation with the managed mode by periodically polling supergraph from `./supergraph.graphql` at runtime. If you update subgraph services, you can simply run

```sh
make graph
```

to update supergraph, **without a need to restart the gateway service**. Subgraph service may need a restart to reflect the changes of the subgraph schema, by
```sh
make $(service_name)
```


### Sample query
```
query {
	me {
    username
    reviews {
      body
      product {
        name
        upc
        price
        weight
        shippingEstimate
      }
      author {
        username
      }
    }
  }
}
```

and you should see
```
{
  "data": {
    "me": {
      "username": "@ada",
      "reviews": [
        {
          "body": "Love it!",
          "product": {
            "name": "Table",
            "upc": "1",
            "price": 899,
            "weight": 100,
            "shippingEstimate": 50
          },
          "author": {
            "username": "@ada"
          }
        },
        {
          "body": "Too expensive.",
          "product": {
            "name": "Couch",
            "upc": "2",
            "price": 1299,
            "weight": 1000,
            "shippingEstimate": 0
          },
          "author": {
            "username": "@ada"
          }
        }
      ]
    }
  }
}
```

### What is this?

This demo showcases four partial schemas running as federated microservices. Each of these schemas can be accessed on their own and form a partial shape of an overall schema. The gateway fetches the composed supergraph at runtime without a need to reload.

To see the query plan when running queries against the gateway, click on the `Query Plan` tab in the bottom right hand corner of [GraphQL Playground](http://localhost:4000)

To learn more about Apollo Federation, check out the [docs](https://www.apollographql.com/docs/apollo-server/federation/introduction)
