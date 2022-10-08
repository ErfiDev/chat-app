# Command line chat with websocket

---

simple project with websockets, users can join to the rooms and chat with others.

## Built using rich tech-stack:

---

- Api-server built using [Go-fiber](https://gofiber.io/).
- [FastHttp](https://github.com/fasthttp/websocket) as websocket pkg.
- [Docker](https://docker.com/) for running application.
- [Gocui](https://github.com/jroimartin/gocui) command line ui.
- [uuid](https://github.com/google/uuid) google uuid.

## Installation & setup :-

---

- Go,Docker,Docker compose & Make should be pre-installed.
- Clone the repository: `git clone https://github.com/erfidev/chat-app`.
- Run the command `make build_server` (this will build the server Dockerfile).
- Run the command `make run_server` (this command run the server image).
- Run the command `docker ps` to ensure server container started.
- Open a new terminal & run the command `make build_client` to build the client.
- run the client binary and enjoy!

# Contributing:

---

- Performance improvements, bug fixes, better design approaches are welcome. Please discuss any changes by raising an issue, beforehand.

# Maintainer

---

By Erfan Hanifezade 2022 Aug

[Linkedin](https://www.linkedin.com/in/erfan-hanifezade-07239b201/) <br>
[Email](erfanhanifezade@gmail.com)
