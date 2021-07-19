const { ApolloServer } = require("apollo-server");
const { ApolloGateway, RemoteGraphQLDataSource } = require("@apollo/gateway");
const { join } = require("path")
const { readFileSync } = require("fs")
const md5 = require("md5")
const { ApolloServerPluginLandingPageGraphQLPlayground } = require('apollo-server-core');

//const supergraphSchema = readFileSync(join(__dirname, "supergraph.graphql")).toString();
//console.log(`Using super graph: ${supergraphSchema}`)

class RemoteGraphQLDataSourceWithCustomizedFetcher extends RemoteGraphQLDataSource {
  fetcher = require("./node_modules/make-fetch-happen").defaults({
  })
}

const gateway = new ApolloGateway({
  // // This entire `serviceList` is optional when running in managed federation
  // // mode, using Apollo Graph Manager as the source of truth.  In production,
  // // using a single source of truth to compose a schema is recommended and
  // // prevents composition failures at runtime using schema validation using
  // // real usage-based metrics.
  // serviceList: [
  //   { name: "accounts", url: "http://localhost:4001/graphql" },
  //   { name: "reviews", url: "http://localhost:4002/graphql" },
  //   { name: "products", url: "http://localhost:4003/graphql" },
  //   { name: "inventory", url: "http://localhost:4004/graphql" }
  // ],

  // // composed schema
  // supergraphSdl: supergraphSchema,

  async experimental_updateSupergraphSdl() {
    const supergraphSchema = readFileSync(join(__dirname, "supergraph.graphql")).toString();
    const id = md5(supergraphSchema);

    console.log(`Using super graph: ${supergraphSchema} with id ${id}`);

    return {
      supergraphSdl: supergraphSchema,
      id: id,
    };
  },
  experimental_pollInterval: 10000, // 10 sec

  // Experimental: Enabling this enables the query plan view in Playground.
  __exposeQueryPlanExperimental: true,

  // https://github.com/apollographql/federation/pull/188 for more information about customizing own fetcher
  // this is useful when needing to adjust retry policy and enabling TLS
  buildService: ({ url }) => new RemoteGraphQLDataSourceWithCustomizedFetcher({ url }),
});

(async () => {
  const server = new ApolloServer({
    gateway,

    // Apollo Graph Manager (previously known as Apollo Engine)
    // When enabled and an `ENGINE_API_KEY` is set in the environment,
    // provides metrics, schema management and trace reporting.
    engine: false,

    // Subscriptions are unsupported but planned for a future Gateway version.
    subscriptions: false,

    plugins: [
      // To see the query plan, we have to replicate the legacy playground
      // https://www.apollographql.com/docs/apollo-server/testing/build-run-queries/#graphql-playground
      ApolloServerPluginLandingPageGraphQLPlayground({
      })
    ]
  });

  server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
  });
})();
