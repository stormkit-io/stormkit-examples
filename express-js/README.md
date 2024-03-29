# Serverless Express.js App

This is an example Express.js Starter Application. The application is generated by 
the following command:

```
npx express-generator
```

## Deployment

There are two possible ways to deploy a Serverless Express.js application on [Stormkit](https://www.stormkit.io).

### 1. Quick way

1. Create a [server.js](./server.js) file on the root level of the project that exports the Express app wrapped with `@stormkit/serverless` helper.
2. Hit the deploy button. 

Example:

```js
const serverless = require("@stormkit/serverless");
const app = require("./app.js");

module.exports = {
  handler: serverless(app),
};
```

This method will upload anything, including the `node_modules`, to the Lambda function and will
configure the entry point as `server.js:handler`. 

It works out of the box however note that the more dependencies your project use, the bigger will be your 
deployment package. We have a `250MB` limit on the zipped folder. 

### 2. Advanced 

1. Follow the first step mentioned above. 
2. Use a bundler like Webpack, Vite or Rollup to build and bundle your application. 
3. Set the build folder as follows: `.stormkit/server` 
4. Make sure the `server.js` file is also located in the `.stormkit/server` subfolder. 

This is a more a advanced technique as it requires you to bundle your application. However, the deployment
package will be significantly lower because the bundlers will remove all non-used node modules and only
the content inside the `.stormkit/server` will be uploaded.