# PowerSSO
PowerSSO is a fundamental piece that authenticates and manages users with the possibility of integration between systems using a Rest API and gRPC

[If you want to use a web interface, run the project.](https://github.com/isaqueveras/powersso-ui)

<p>
  <img alt="Go report" src="https://goreportcard.com/badge/isaqueveras/power-sso">
  <img alt="" src="https://github.com/isaqueveras/power-sso/actions/workflows/go.yml/badge.svg">
  <img alt="Repository size" src="https://img.shields.io/github/languages/top/isaqueveras/power-sso">
  <img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/isaqueveras/power-sso">
  <img alt="GitHub Sponsors" src="https://img.shields.io/github/sponsors/isaqueveras">  
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/isaqueveras/power-sso?color=%2304D361">
  <img alt="GitHub top language" src="https://img.shields.io/github/repo-size/isaqueveras/power-sso">
  <a href="https://github.com/isaqueveras/power-sso/commits/main">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/isaqueveras/power-sso">
  </a>
  <a href="https://github.com/isaqueveras/power-sso/stargazers">
    <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/isaqueveras/power-sso">
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
$ git clone https://github.com/isaqueveras/power-sso

# Access the project folder in your terminal/cmd
$ cd power-sso

# Install the dependencies
$ go mod tidy

# Run postgres database
$ make local

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
<a href="https://github.com/isaqueveras/power-sso/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=isaqueveras/power-sso&max=100" alt="List of contributors to the powerSSO project"/>
</a>

[reactjs]: https://reactjs.org
[typescript]: https://www.typescriptlang.org/
[nodejs]: https://nodejs.org/
[vscode]: https://code.visualstudio.com/
[golang]: https://go.dev/
