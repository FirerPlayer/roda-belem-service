<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Roda Belem Service" />

&#xa0;

  <!-- <a href="https://rodabelemservice.netlify.app">Demo</a> -->
</div>

<h1 align="center">Roda Belem Service</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/FirerPlayer/roda-belem-service?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/FirerPlayer/roda-belem-service?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/FirerPlayer/roda-belem-service?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/FirerPlayer/roda-belem-service?color=56BEB8">

  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/FirerPlayer/roda-belem-service?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/FirerPlayer/roda-belem-service?color=56BEB8" /> -->

  <!-- <img alt="Github stars" src="https://img.shields.io/github/stars/FirerPlayer/roda-belem-service?color=56BEB8" /> -->
</p>

<!-- Status -->

<!-- <h4 align="center">
	🚧  Roda Belem Service 🚀 Under construction...  🚧
</h4>

<hr> -->

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/FirerPlayer" target="_blank">Author</a>
</p>

<br>

## :dart: About

The backend service for the Roda Belém app has a hexagonal architecture with a REST API and consumes the Google Places API. It is designed for high-performance communication with the Roda Belém frontend.

## :sparkles: Features

:heavy_check_mark: Hexagonal Architecture\
:heavy_check_mark: Scalable and easy maintenance\
:heavy_check_mark: Cuckoo filter implemented to reduce latency and improve server response\
:heavy_check_mark: Jwt authentication\
:heavy_check_mark: Data persistence in MySQL database\
:heavy_check_mark: No ORM approach, with SQLC, for faster database interaction

## :rocket: Technologies

The following tools were used in this project:

- [Go](https://go.dev) \
- [Chi](https://github.com/go-chi/chi)
- [sqlc](https://sqlc.dev)
- [viper](https://github.com/spf13/viper)
- [jwt](https://github.com/golang-jwt/jwt)
- [Google Maps Services](https://github.com/googlemaps/google-maps-services-go)
- [Docker](https://www.docker.com) \
- [Docker Compose](https://docs.docker.com/compose/)
- [MySQL](https://www.mysql.com)

## :white_check_mark: Requirements

Before starting :checkered_flag:, you need to have [Docker](https://www.docker.com) (with composer) and some database manager or MySQL managers (like MySQL Workbench) installed.

Create a .env file based on .env.example, you can configure the database environment variables. Remember to add your own APIs keys.
Open the database and create tables according to the sql in schemas folder.

## :checkered_flag: Runing up the service

```bash
# Clone this project
$ git clone https://github.com/FirerPlayer/roda-belem-service

# Access
$ cd roda-belem-service

# Run containers
$ docker compose up -d
```

## :memo: License

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.

Made with :heart: by <a href="https://github.com/FirerPlayer" target="_blank">Micael</a>

&#xa0;

<a href="#top">Back to top</a>
