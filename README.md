# PowerSSO
PowerSSO is a authenticator and user manager for systems (under construction) 

[If you want to use a web interface, run the project.](https://github.com/isaqueveras/powersso-ui)

<p>
  <img alt="Go report" src="https://goreportcard.com/badge/isaqueveras/powersso">
  <img alt="" src="https://github.com/isaqueveras/powersso/actions/workflows/go.yml/badge.svg">
  <img alt="Repository size" src="https://img.shields.io/github/languages/top/isaqueveras/powersso">
  <img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/isaqueveras/powersso">
  <a href="https://github.com/isaqueveras/powersso/commits/main">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/isaqueveras/powersso">
  </a>
</p>

---

![permission architecture](https://github.com/isaqueveras/powersso/assets/46972789/e2c91750-2fcc-4ba9-b4ef-324d7dece6e0)

## 🚀 How to run the project
📌 Before starting, you will need to have the [Golang][golang] language installed on your machine.

### 🧭 Running the application

```bash
# Clone this repository
$ git clone https://github.com/isaqueveras/powersso

# Access the project folder in your terminal/cmd
$ cd powersso

# Install the dependencies
$ go mod tidy

# Run postgres database
$ make dev

# Run the migrations
$ make migrate-up

# Run the application in development mode
$ go run main.go
```

```bash
- The backend will open on the port:5000        # access http://localhost:5000
```
## 😯 How to contribute to the project

1. **Fork** the project.
2. Create a new branch with your changes: `git checkout -b my-feature`
3. Save the changes and create a commit message telling what you did: `git commit -m "feature: My new feature"`
4. Submit your changes: `git push origin my-feature`

If you have any questions, check this [GitHub Contributing Guide](https://github.com/firstcontributions/first-contributions)

## Contributors
<a href="https://github.com/isaqueveras/powersso/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=isaqueveras/powersso&max=100" alt="List of contributors to the powerSSO project"/>
</a>

[golang]: https://go.dev/

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=isaqueveras/powersso&type=Date)](https://star-history.com/#isaqueveras/powersso&Date)
