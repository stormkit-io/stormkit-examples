import React from "react";
import { renderToString } from "react-dom/server";
import App from "./App.js";

// Simply export an async function which accepts
// request and response as arguments.
export default async (req, res) => {
  return res.send({
    // Specify the response headers.
    // If you want to redirect the user, simply return a Location header.
    headers: {
      "X-Powered-By": "stormkit.io",
      "X-My-Custom-Header": Date.now().toString()
    },
    // You can either return a string, or an object as a response body.
    // Specify where you would like to mount your app in your index.html
    // file with a namespaced element like <sk:OBJECT_KEY />
    // where OBJECT_KEY is one of the keys in the next statement.
    // For instance, <sk:content/> in this case.
    body: { content: renderToString(<App />) },
    // Return the response code. If not specified, will default to 200.
    status: 200
  });
};
