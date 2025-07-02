# Go Hello World Example

This repository contains a simple "Hello, World!" application written in Go.

## ⚡️ Deployment with Stormkit

1. Login to your self-hosted Stormkit Instance. Here's [a guide](https://www.stormkit.io/tutorials/how-to-self-host-stormkit-on-hetzner-cloud) to set it up.
2. Import this project from URL
3. Configure the Build Settings:
   - Build command: `go build -o dist/server .`
   - Build root: `./go-api`
   - Output folder: `./dist`
4. Configure the Server Settings:
   - Start command: `./dist/server.go`
5. Click Save and Deploy

- 🛠️ [Setup Stormkit on Hetzner](https://www.stormkit.io/tutorials/how-to-self-host-stormkit-on-hetzner-cloud)
- 🍰 [How to deploy your Strapi CMS](https://www.stormkit.io/tutorials/how-to-deploy-your-self-hosted-strapi-instance)
- 📑 [Stormkit documentation](https://www.stormkit.io/docs/welcome/getting-started)
- 🌞 [Stormkit changelog](https://www.stormkit.io/blog/whats-new)
