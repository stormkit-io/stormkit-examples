console.log(
  import("./dist/angular/server/server.mjs").then(({ app }) => {
    console.log(app());
  })
);
