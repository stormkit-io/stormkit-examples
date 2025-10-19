# Go Hello World Example

This repository contains a simple "Hello, World!" application written in Go.

## ⚡️ Deployment with Stormkit

1. Log in to your self-hosted Stormkit instance. Here's [a guide](https://www.stormkit.io/tutorials/how-to-self-host-stormkit-on-hetzner-cloud) to set it up.
2. Import this project from the URL
3. Configure the build settings:
   - Build command: `go build -o dist/server .`
   - Build root: `./go-api`
   - Output folder: `./dist`
4. Configure the server settings:
   - Start command: `./dist/server`
5. This example connects to a real-time database, so we need to update our environment variables. These environment variables are the ones Stormkit uses to access its own database, and we can use them directly without needing to create a new database. In production, make sure to create your own database.
   - `POSTGRES_DB=$POSTGRES_DB`
   - `POSTGRES_HOST=$POSTGRES_HOST`
   - `POSTGRES_PASSWORD=$POSTGRES_PASSWORD`
   - `POSTGRES_USER=$POSTGRES_USER`
6. Click Save and Deploy

- 🛠️ [Set up Stormkit on Hetzner](https://www.stormkit.io/tutorials/how-to-self-host-stormkit-on-hetzner-cloud)
- 🍰 [How to deploy your Strapi CMS](https://www.stormkit.io/tutorials/how-to-deploy-your-self-hosted-strapi-instance)
- 📑 [Stormkit documentation](https://www.stormkit.io/docs/welcome/getting-started)
- 🌞 [Stormkit changelog](https://www.stormkit.io/blog/whats-new)
