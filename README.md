# Static Site Generator

👋 Hi! This is a static site generator made in [Gin](https://github.com/gin-gonic/gin). It checks for markdown files in the markdown directory, renders them to HTML and serves them locally. For example the file `markdown/test/example.md` can be viewed in the browser at `example.com/test/example.md`. This app was created for me to try to tidy my university notes (and start finally learning golang and gin).

---

## ✨ Getting Started

### 1. Get the prerequisites

🛠 This project requires `golang`, if you're unsure how to set it up I recommend  [this page](https://go.dev/doc/install).

### 2. Clone the repository and enter it

```sh
git clone --depth=1 https://github.com/TypicalAM/Static-Site-Generator.git && cd Static-Site-Generator
```

### 3. Run the app!

🌟 To run in a dev environment

```sh
go run .
```

## ✨ Starting the app in a cointainer

If you have a possiblity to run docker containers on your machine, you can also run the dev version of this product directly in the container using the provided `Dockerfile`.

```sh
docker build -t go-docker . 
docker run -d -p 3006:3004 go-docker
```

---
