import React from "react";
import { renderToString } from "react-dom/server";
import App from "./App.js";

export default async (req, res) => {
  // We don't need bots indexing this page
  if (req.path === "/robots.txt") {
    return res.send({
      body: "User-agent: *\nDisallow: /",
      status: 200,
      headers: {
        "Content-Type": "text/plain; charset=utf-8"
      }
    });
  }

  return res.send({
    headers: {
      "X-Powered-By": "stormkit.io",
      "X-My-Custom-Header": Date.now().toString()
    },
    body: { content: renderToString(<App />) },
    status: 200
  });
};
