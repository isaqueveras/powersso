# PowerSSO
PowerSSO is a authenticator and user manager for systems (under construction) 

<p>
  <img alt="Go report" src="https://goreportcard.com/badge/isaqueveras/power-sso">
  <img alt="" src="https://github.com/isaqueveras/powersso/actions/workflows/go.yml/badge.svg">
  <img alt="Repository size" src="https://img.shields.io/github/languages/top/isaqueveras/powersso">
  <img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/isaqueveras/powersso">
  <a href="https://github.com/isaqueveras/powersso/commits/main">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/isaqueveras/powersso">
  </a>
</p>

## 🛠 Technologies

Some tools being used in this project: [Golang][golang], [React][reactjs], [TypeScript][typescript]

## 🚀 How to run the project

### 📌 Prerequisites

Before you begin, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Node.js][nodejs] and [Golang][golang].

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

# Generate documentation
$ make swag

# Run the application in development mode
$ go run main.go

# Access the project folder in your terminal/cmd
$ cd ui

# Install the dependencies
$ npm install

# Run the application in development mode
$ npm run start
```

```bash
- The frontend will open on the port:3000       # access http://localhost:3000
- The backend will open on the port:5000        # access http://localhost:5000
- The mailcatcher will open on the port:1080    # access http://localhost:1080
- The documentation will open on the port:5000: # access http://localhost:5000/swagger/index.html
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

[reactjs]: https://reactjs.org
[typescript]: https://www.typescriptlang.org/
[nodejs]: https://nodejs.org/
[vscode]: https://code.visualstudio.com/
[golang]: https://go.dev/
