const serverless = require("@stormkit/serverless");
const app = require("./app.js");

module.exports = {
  handler: serverless(app),
};
