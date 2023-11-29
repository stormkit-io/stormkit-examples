## Welcome to Stormkit Examples

This repository contains example applications that can be hosted on [Stormkit](https://www.stormkit.io).

| app name         | type       | endpoint                                                                                             |
| ---------------- | ---------- | ---------------------------------------------------------------------------------------------------- |
| create-react-app | static      | [https://examples--create-react-app.stormkit.dev/](https://examples--create-react-app.stormkit.dev/) |
| next-js          | prerendered | [https://examples--next.stormkit.dev/](https://examples--next.stormkit.dev/)         |
| nuxt             | serverless  | [https://examples--nuxt.stormkit.dev/](https://examples--nuxt.stormkit.dev/)         |
| ng-universal     | serverless  | [https://examples--ng-universal.stormkit.dev/](https://examples--ng-universal.stormkit.dev/)         |

## Application Structure

In order to host this example on Stormkit:

1. Create an application by connecting this repository
1. Create an environment for each folder
1. Specify the build configuration
1. Create an extra environment variable on the config page called `SK_CWD` with the folder name being the value

The `SK_CWD` environment variable tells Stormkit to change the working directory. The build command will be executed directly within that directory.
