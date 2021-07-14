## Apollo Federation Demo

This repository is a demo of using Apollo Federation to build a single schema on top of multiple services. The microservices are located under the [`./services`](./services/) folder and the gateway that composes the overall schema is in the [`gateway.js`](./gateway.js) file.

### Installation

Install `rover` cli first in order to build supergraph

```sh
npm install -g @apollo/rover
```

To run this demo locally, pull down the repository then run the following commands:

```sh
npm install
```

This will install all of the dependencies for the gateway and each underlying service.

```sh
npm run start-services
```

This command will run all of the microservices at once. They can be found at http://localhost:4001, http://localhost:4002, http://localhost:4003, and http://localhost:4004.

In another terminal window, run the gateway by running this command:

```sh
npm run start-gateway
```

This will start up the gateway and serve it at http://localhost:4000

### Managed mode
This example runs Apollo Federation with the managed mode by periodically polling supergraph from `./supergraph.graphql` at runtime. If you update subgraph services, you can simply run

```sh
rover supergraph compose --config ./supergraph-config.yaml > supergraph.graphql
```
to update supergraph, **without a need to restart the gateway service**. Subgraph services might need to be reloaded to understand the new subgraph by simply do `npm run start-services` again.

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
