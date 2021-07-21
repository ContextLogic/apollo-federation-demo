const { ApolloServer, gql } = require("apollo-server");
const { buildFederatedSchema } = require("@apollo/federation");
const { join } = require("path")
const { readFileSync } = require("fs")

const typeDefs = gql`${readFileSync(join(__dirname, "graph/reviews.graphql"), "utf8")}`;

const resolvers = {
  Review: {
    author(review) {
      return { __typename: "User", id: review.authorID };
    },
    product(review) {
      return { __typename: "Product", upc: review.upc }
    }
  },
  User: {
    reviews(user) {
      return reviews.filter(review => review.authorID === user.id);
    },
    numberOfReviews(user) {
      return reviews.filter(review => review.authorID === user.id).length;
    },
    username(user) {
      const found = usernames.find(username => username.id === user.id);
      return found ? found.username : null;
    }
  },
  Product: {
    reviews(product) {
      return reviews.filter(review => review.product.upc === product.upc);
    }
  }
};

const server = new ApolloServer({
  schema: buildFederatedSchema([
    {
      typeDefs,
      resolvers
    }
  ])
});

server.listen({ port: 4004 }).then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});

const usernames = [
  { id: "1", username: "@ada" },
  { id: "2", username: "@complete" }
];
const reviews = [
  {
    id: "1",
    authorID: "1",
    upc: "1",
    body: "Love it!"
  },
  {
    id: "2",
    authorID: "1",
    upc: "2",
    body: "Too expensive."
  },
  {
    id: "3",
    authorID: "2",
    upc: "3",
    body: "Could be better."
  },
  {
    id: "4",
    authorID: "2",
    upc: "1",
    body: "Prefer something else."
  }
];
