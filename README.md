## Welcome to Stormkit Examples

This repository contains example applications that can be hosted on [Stormkit](https://www.stormkit.io).

| app name         | type       | endpoint                                                                                             |
| ---------------- | ---------- | ---------------------------------------------------------------------------------------------------- |
| create-react-app | static      | [https://examples--create-react-app.stormkit.dev/](https://examples--create-react-app.stormkit.dev/) |
| next-js          | prerendered | [https://examples--next.stormkit.dev/](https://examples--next.stormkit.dev/) |
| nuxt             | serverless  | [https://examples--nuxt.stormkit.dev/](https://examples--nuxt.stormkit.dev/) |
| static           | static      | [https://examples--static.stormkit.dev/](https://examples--static.stormkit.dev/) |
| angular         | universal  | [https://examples--ng-universal.stormkit.dev/](https://examples--angular.stormkit.dev/) |

## Application Structure

In order to host this example on Stormkit:

1. Create an application by connecting this repository
1. Create an environment for each folder
1. Go to **Environment** > **Config** > **General** 
1. Configure the `branch`. It should point to `main`.
1. Go to **Environment** > **Config** > **Build**
2. Configure the `root directory`. It should point to the subfolder of the example.
