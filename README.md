# This Is My Go Lang Repository

## It Contains All The Basic Program's

---

- `I've Created This For Self Understanding`
- `This Repo Contains Simple Basic Concepts Of Golang`

---

# Data Conversion in Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=atharvaskadam&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

Go's basic types are

- boolean
- string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32 and also represents a Unicode code point
- float32 float64
- complex64 complex128

The example below shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

```go
var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

> Reference: [A Tour of Go](https://tour.golang.org/basics/11)

| Conversion   | int                                                               | float                                | bool                 | string                                                                |
| ------------ | ----------------------------------------------------------------- | ------------------------------------ | -------------------- | --------------------------------------------------------------------- |
| var i int    | i                                                                 | float64(i) <br> float32(in)          |                      | strconv.Itoa(i)<br>strconv.FormatInt(i, [10](base))                   |
| var f float  | int(f)                                                            | f                                    |                      | strconv.FormatFloat(f, ['E'](format), [-1](precision), [64](bitsize)) |
| var b bool   |                                                                   |                                      | b                    | strconv.FormatBool(b)                                                 |
| var s string | strconv.Atoi(s)<br>strconv.ParseInt(s, [10](base), [64](bitsize)) | strconv.ParseFloat(s, [32](bitsize)) | strconv.ParseBool(s) | s                                                                     |

> Reference: [strconv - Go Doc](https://golang.org/pkg/strconv/)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=atharvaskadam&label=atharvaskadam&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=atharvaskadam)
# VSCode Extensions List

[![Visitors](https://api.visitorbadge.io/api/visitors?path=atharvaskadam&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=atharvaskadam)

## List of Useful VSCode Extension

| Extension Name | Id  | Type | Description |
| -------------- | --- | ---- | ----------- |
| Rainbow Brackets | `2gua.rainbow-brackets` | General | A rainbow brackets extension for VS Code |
| Terraform | `4ops.terraform` | Terraform | Terraform configuration language support (includes Terragrunt) |
| Better Comments | `aaron-bond.better-comments` | General | Improve your code commenting by annotating with alert, informational, TODOs, and more! |
| HTML Snippets | `abusaidm.html-snippets` | HTML | Full HTML tags including HTML5 Snippets |
| CSS Formatter | `aeschli.vscode-css-formatter` | CSS | Formatter for CSS |
| YAML to JSON | `ahebrank.yaml2json` | YAML/JSON | Convert YAML from clipboard or current document/selection to JSON and vice versa. |
| JavaScript Snippet Pack | `akamud.vscode-javascript-snippet-pack` | JS | A snippet pack to make you more productive working with JavaScript |
| Bookmarks | `alefragnani.Bookmarks` | General | Mark lines and jump to them |
| Project Manager | `alefragnani.project-manager` | General | Easily switch between projects |
| AWS Toolkit | `amazonwebservices.aws-toolkit-vscode` | AWS | Amazon Web Services toolkit for browsing and updating cloud resources |
| AWS Serverless Code Snippets | `AntStack.aws-js-code-snippet` | AWS | Javascript Code snippets for AWS services |
| Alexa Skills Kit (ASK) Toolkit | `ask-toolkit.alexa-skills-kit-toolkit` | Alexa | Build and manage Alexa skills using Visual Studio Code |
| CloudFormation | `aws-scripting-guy.cform` | AWS | VS Code Plugin for CloudFormation |
| Custom CSS and JS Loader | `be5invis.vscode-custom-css` | CSS/JS | Custom CSS and JS for Visual Studio Code |
| HTTP/s and relative link checker | `blackmist.LinkCheckMD` | General | Checks Markdown links for the presence of a country-code as you type and flags as a warning. Checks whether HTTP/s links or relative links are reachable when you press Alt+L. |
| Twitch Chat | `BlitZz.twitch-chat` | General | Display Twitch chat as a tab |
| Path Intellisense | `christian-kohler.path-intellisense` | General | Visual Studio Code plugin that autocompletes filenames |
| Bracket Pair Colorizer | `CoenraadS.bracket-pair-colorizer` | General | A customizable extension for colorizing matching brackets |
| vscode-jq | `dandric.vscode-jq` | JSON | jq plugin for Visual Studio Code |
| aws-cloudformation-yaml | `DanielThielking.aws-cloudformation-yaml` | AWS | Adds YAML and JSON snippets for AWS CloudFormation to VS Code |
| CloudFormation Snippets for VS Code | `dannysteenman.cloudformation-yaml-snippets` | AWS | CloudFormation Snippets for VS Code |
| Dart | `Dart-Code.dart-code` | Dart | Dart language support and debugger for Visual Studio Code. |
| Flutter | `Dart-Code.flutter` | Flutter | Flutter support and debugger for Visual Studio Code. |
| markdownlint | `DavidAnson.vscode-markdownlint` | Markdown | Markdown linting and style checking for Visual Studio Code |
| ESLint | `dbaeumer.vscode-eslint` | JS | Integrates ESLint JavaScript into VS Code |
| docs-article-templates | `docsmsft.docs-article-templates` | General | Docs article templates |
| Docs Authoring Pack | `docsmsft.docs-authoring-pack` | General | Collection of extensions to assist with content development for docs.microsoft.com |
| docs-validation | `docsmsft.docs-build` | General | Docs Validation Extension which enables you to run build validation on a Docs conceptual or Learn repo at authoring time in VS Code |
| docs-images | `docsmsft.docs-images` | General | Docs Images Extension |
| docs-linting | `docsmsft.docs-linting` | General | Docs Linting Extension |
| docs-markdown | `docsmsft.docs-markdown` | Markdown | Docs Markdown Extension |
| docs-metadata | `docsmsft.docs-metadata` | General | Simplifies editing metadata within docs. |
| docs-preview | `docsmsft.docs-preview` | Markdown | Docs Markdown Preview Extension |
| docs-scaffolding | `docsmsft.docs-scaffolding` | General | A Visual Studio Code extension developed to aid with scaffolding and updating Learn modules. |
| docs-visual-areas | `docsmsft.docs-visual-areas` | General | Docs Visual Areas: Visualize Microsoft Docs content inside tabs and zones in Visual Studio Code |
| docs-yaml | `docsmsft.docs-yaml` | YAML | YAML schema validation and auto-completion for docs.microsoft.com authoring |
| Git History | `donjayamanne.githistory` | Git | View git log, file history, compare branches or commits |
| jQuery Code Snippets | `donjayamanne.jquerysnippets` | jQuery | Over 130 jQuery Code Snippets |
| XML Tools | `DotJoshJohnson.xml` | XML | XML Formatting, XQuery, and XPath Tools for Visual Studio Code |
| ES7+ React/Redux/React-Native snippets | `dsznajder.es7-react-js-snippets` | React | Extensions for React, React-Native and Redux in JS/TS with ES7+ syntax. Customizable. Built-in integration with prettier. |
| GitLens — Git supercharged | `eamodio.gitlens` | Git | Supercharge Git within VS Code — Visualize code authorship at a glance via Git blame annotations and CodeLens, seamlessly navigate and explore Git repositories, gain valuable insights via rich visualizations and powerful comparison commands, and so much more |
| Prettier - Code formatter | `esbenp.prettier-vscode` | General | Code formatter using prettier |
| Comment Anchors | `ExodiusStudios.comment-anchors` | General | Place anchor tags within comments for easy file & workspace navigation. |
| Diff | `fabiospampinato.vscode-diff` | Git | Diff 2 opened files with ease. Because running `code --diff path1 path2` is too slow. |
| Todo+ | `fabiospampinato.vscode-todo-plus` | General | Manage todo lists with ease. Powerful, easy to use and customizable. |
| Auto Rename Tag | `formulahendry.auto-rename-tag` | HTML | Auto rename paired HTML/XML tag |
| Code Runner | `formulahendry.code-runner` | General | Run code snippet or code file for multiple languages |
| Terminal | `formulahendry.terminal` | General | Terminal for Visual Studio Code |
| markmap-vscode | `gera2ld.markmap-vscode` | Markdown | This extension integrates [markmap](https://markmap.js.org/) into VSCode |
| GitHub Copilot | `GitHub.copilot` | General | Your AI pair programmer |
| GitHub Theme | `GitHub.github-vscode-theme` | Theme | GitHub theme for VS Code |
| GitHub Pull Requests and Issues | `GitHub.vscode-pull-request-github` | General | Pull Request and Issue Provider for GitHub |
| Go | `golang.go` | Go | Rich Go language support for Visual Studio Code |
| HashiCorp Terraform | `hashicorp.terraform` | Terraform | Syntax highlighting and autocompletion for Terraform |
| Beautify | `HookyQR.beautify` | General | Beautify code in place for VS Code |
| Minify | `HookyQR.minify` | General | Minify for VS Code. Minify with command, and (optionally) re-minify on save. |
| Git History Diff | `huizhou.githd` | Git | View git history. View diff of committed files. View git blame info. View stash details. |
| SVG | `jock.svg` | SVG | SVG Coding, Minify, Pretty, Preview All-In-One |
| AWS CloudFormation Auto-template Generator | `john-goldsmith.vscode-aws-cloudformation-auto-template-generator` | AWS | Automatically generate fully-expanded AWS CloudFormation resource templates |
| Sublime Material Theme | `jprestidge.theme-material-theme` | Theme | Port of the Material Theme for Sublime Text. |
| jq Syntax Highlighting | `jq-syntax-highlighting.jq-syntax-highlighting` | JSON | Syntax support for the jq JSON command-line processor. |
| CloudFormation Linter | `kddejong.vscode-cfn-lint` | AWS | AWS CloudFormation template Linter |
| GitHub | `KnisterPeter.vscode-github` | Git | Integrates github and its workflows into vscode |
| CloudFormation Resource Actions | `ljacobsson.cfn-resource-actions` | AWS | Integrate with deployed resources directly from your CloudFormation template |
| Rainbow CSV | `mechatroner.rainbow-csv` | CSV | Highlight CSV and TSV files, Run SQL-like queries |
| Git Graph | `mhutchie.git-graph` | Git | View a Git Graph of your repository, and perform Git actions from the graph. |
| Mintlify Doc Writer | `mintlify.document` | General | AI based comment / doc writer |
| Azure App Service | `ms-azuretools.vscode-azureappservice` | Azure | An Azure App Service management extension for Visual Studio Code. |
| Azure Resources | `ms-azuretools.vscode-azureresourcegroups` | Azure | An extension for viewing and managing Azure resources. |
| Docker | `ms-azuretools.vscode-docker` | Docker | Makes it easy to create, manage, and debug containerized applications. |
| Data Workspace | `ms-mssql.data-workspace-vscode` | General | Additional common functionality for database projects |
| SQL Server (mssql) | `ms-mssql.mssql` | SQL | Develop Microsoft SQL Server, Azure SQL Database and SQL Data Warehouse everywhere |
| SQL Bindings | `ms-mssql.sql-bindings-vscode` | SQL | Enables users to develop and publish Azure Functions with Azure SQL bindings |
| SQL Database Projects | `ms-mssql.sql-database-projects-vscode` | SQL | Enables users to develop and publish database schemas for MSSQL Databases |
| Anaconda Extension Pack | `ms-python.anaconda-extension-pack` | Python | The Anaconda Extension Pack is a set of extensions that enhance the experience of Anaconda customers using Visual Studio Code |
| isort | `ms-python.isort` | General | Import Organization support for python files using `isort`. |
| Python | `ms-python.python` | Python | IntelliSense (Pylance), Linting, Debugging (multi-threaded, remote), Jupyter Notebooks, code formatting, refactoring, unit tests, and more. |
| Pylance | `ms-python.vscode-pylance` | Python | A performant, feature-rich language server for Python in VS Code |
| Jupyter | `ms-toolsai.jupyter` | Python | Jupyter notebook support, interactive programming and computing that supports Intellisense, debugging and more. |
| Jupyter Keymap | `ms-toolsai.jupyter-keymap` | Python | Jupyter keymaps for notebooks |
| Jupyter Notebook Renderers | `ms-toolsai.jupyter-renderers` | Python | Renderers for Jupyter Notebooks (with plotly, vega, gif, png, svg, jpeg and other such outputs) |
| Jupyter Cell Tags | `ms-toolsai.vscode-jupyter-cell-tags` | General  | Jupyter Cell Tags support for VS Code |
| Jupyter Slide Show | `ms-toolsai.vscode-jupyter-slideshow` | General  | Jupyter Slide Show support for VS Code |
| Remote - Containers | `ms-vscode-remote.remote-containers` | Docker | Open any folder or repository inside a Docker container and take advantage of Visual Studio Code's full feature set. |
| Remote - WSL | `ms-vscode-remote.remote-wsl` | General | Open any folder in the Windows Subsystem for Linux (WSL) and take advantage of Visual Studio Code's full feature set. |
| Azure Account | `ms-vscode.azure-account` | Azure | A common Sign In and Subscription management extension for VS Code. |
| C/C++ | `ms-vscode.cpptools` | General | C/C++ IntelliSense, debugging, and code browsing. |
| Live Preview | `ms-vscode.live-server` | Web | Hosts a local server in your workspace for you to preview your webpages on. |
| PowerShell | `ms-vscode.powershell` | General | Develop PowerShell modules, commands and scripts in Visual Studio Code! |
| Live Share | `ms-vsliveshare.vsliveshare` | General | Real-time collaborative development from the comfort of your favorite tools. |
| React Native Tools | `msjsdiag.vscode-react-native` | React Native | Debugging and integrated commands for React Native |
| SQLTools | `mtxr.sqltools` | SQL | Database management done right. Connection explorer, query runner, intellisense, bookmarks, query history. Feel like a database hero! |
| SQLTools MySQL/MariaDB | `mtxr.sqltools-driver-mysql` | SQL | SQLTools MySQL/MariaDB |
| SQLTools PostgreSQL/Redshift Driver | `mtxr.sqltools-driver-pg` | SQL | SQLTools PostgreSQL/Redshift Driver |
| autoDocstring - Python Docstring Generator | `njpwerner.autodocstring` | Python | Generates python docstrings automatically |
| indent-rainbow | `oderwat.indent-rainbow` | General | Makes indentation easier to read |
| Material Icon Theme | `PKief.material-icon-theme` | Theme | Material Design Icons for Visual Studio Code |
| PlatformIO IDE | `platformio.platformio-ide` | General | Professional development environment for Embedded, IoT, Arduino, CMSIS, ESP-IDF, FreeRTOS, libOpenCM3, mbed OS, Pulp OS, SPL, STM32Cube, Zephyr RTOS, ARM, AVR, Espressif (ESP8266/ESP32), FPGA, MCS-51 (8051), MSP430, Nordic (nRF51/nRF52), PIC32, RISC-V, STMicroelectronics (STM8/STM32), Teensy |
| TypeScript Importer | `pmneo.tsimporter` | TS | Automatically searches for TypeScript definitions in workspace files and provides all known symbols as completion item to allow code completion. |
| Polacode | `pnp.polacode` | General | Polaroid for your code |
| View In Browser | `qinjia.view-in-browser` | Web | view a html file in system's default browser |
| Serverless Framework Snippets | `rafwilinski.serverless-vscode-snippets` | Serverless | Code snippets for Serverless Framework in VS Code editor |
| Thunder Client | `rangav.vscode-thunder-client` | API | Lightweight Rest API Client for VS Code |
| Language Support for Java(TM) by Red Hat | `redhat.java` | Java | Java Linting, Intellisense, formatting, refactoring, Maven/Gradle support and more... |
| Red Hat Commons | `redhat.vscode-commons` | Red Hat | Base utilities for Red Hat extensions. Notably controls telemetry settings. |
| YAML | `redhat.vscode-yaml` | YAML | YAML Language Support by Red Hat, with built-in Kubernetes syntax support |
| Live Server | `ritwickdey.LiveServer` | Web | Launch a development local Server with live reload feature for static & dynamic pages |
| SynthWave '84 | `RobbOwen.synthwave-vscode` | Theme | A Synthwave-inspired colour theme to satisfy your neon dreams |
| Terraform doc snippets | `run-at-scale.terraform-doc-snippets` | Terraform | Terraform code snippets (>8000) pulled from examples in the terraform registry provider docs. resources and data sources. All official and partner providers are scraped for examples. |
| serverless-completer | `securisec.serverless-completer` | Serverless | serverless.yml file completer for VSCode |
| Trailing Spaces | `shardulm94.trailing-spaces` | General | Highlight trailing spaces and delete them in a flash! |
| vscode-note | `shinhwagk.vscode-note` | Notes | A simple note-taking extension. |
| Indenticator | `SirTori.indenticator` | General | Highlights your current indent depth |
| SonarLint | `SonarSource.sonarlint-vscode` | General | SonarLint is an IDE extension that helps you detect and fix quality issues as you write code in JavaScript, TypeScript, Python, Java, HTML and PHP. |
| Code Spell Checker | `streetsidesoftware.code-spell-checker` | General | Spelling checker for source code |
| Tabnine AI Autocomplete for Javascript, Python, Typescript, PHP, Go, Java, Ruby & more | `TabNine.tabnine-vscode` | General | JavaScript, Python, Java, Typescript & all other languages - AI Code completion plugin. Tabnine makes developers more productive by auto-completing their code. |
| GitLive | `TeamHub.teamhub` | General | Extend VS Code with real-time collaborative superpowers |
| ChatGPT: write and improve code using AI | `timkmecl.chatgpt` | General | Use ChatGPT right inside the IDE to enhance and automate your coding with AI-powered assistance (unofficial) |
| Serverless Framework | `TimVaneker.serverless-snippets` | Serverless | Serverless Framework snippets and autocomplete |
| Line Note | `tkrkt.linenote` | General | Add notes to the line of source code |
| VSCode Pets | `tonybaloney.vscode-pets` | General | Add pets to VS Code for fun |
| Trunk | `trunk.io` | General | blazingly fast meta code checker and formatter |
| Sort Lines | `Tyriar.sort-lines` | General | Sorts lines of text |
| Error Lens | `usernamehw.errorlens` | General | Improve highlighting of errors, warnings and other language diagnostics. |
| Code Navigation | `vikas.code-navigation` | General | Visual Studio Code Navigation |
| vscode-icons | `vscode-icons-team.vscode-icons` | Theme | Icons for Visual Studio Code |
| Quokka.js | `WallabyJs.quokka-vscode` | JS/TS | JavaScript and TypeScript playground in your editor. |
| Web Template Studio (Preview) | `WASTeamAccount.WebTemplateStudio-dev-nightly` | Web | Web Template Studio enables developers to quickly scaffold full-stack web applications with cloud services. |
| Cobalt2 Theme Official | `wesbos.theme-cobalt2` | Theme | Cobalt2 Theme for VS Code |
| JavaScript (ES6) code snippets | `xabikos.JavaScriptSnippets` | JS | Code snippets for JavaScript in ES6 syntax |
| Txt Syntax | `xshrim.txt-syntax` | General | highlight text files(.txt, .out .tmp, .log, .ini, .cnf ...) and provide general utility tools for text documents |
| Markdown All in One | `yzhang.markdown-all-in-one` | Markdown | All you need to write Markdown (keyboard shortcuts, table of contents, auto preview and more) |
| [Deprecated] Debugger for Chrome | `msjsdiag.debugger-for-chrome` | Web | Debug your JavaScript code in the Chrome browser, or any other target that supports the Chrome Debugger protocol. |

## Command to List Extensions

Run this command in VSCode Terminal to list all your extensions `code --list-extensions | % { "code --install-extension $_" }`

```powershell
code --list-extensions | % { "code --install-extension $_" }
---------------------------------------------------------------
List Update Date : 2020-06-17
---------------------------------------------------------------
code --install-extension 2gua.rainbow-brackets
code --install-extension 4ops.terraform
code --install-extension aaron-bond.better-comments
code --install-extension abusaidm.html-snippets
code --install-extension aeschli.vscode-css-formatter
code --install-extension ahebrank.yaml2json
code --install-extension akamud.vscode-javascript-snippet-pack
code --install-extension alefragnani.Bookmarks
code --install-extension alefragnani.project-manager
code --install-extension amazonwebservices.aws-toolkit-vscode
code --install-extension AntStack.aws-js-code-snippet
code --install-extension ask-toolkit.alexa-skills-kit-toolkit
code --install-extension aws-scripting-guy.cform
code --install-extension be5invis.vscode-custom-css
code --install-extension blackmist.LinkCheckMD
code --install-extension BlitZz.twitch-chat
code --install-extension christian-kohler.path-intellisense
code --install-extension CoenraadS.bracket-pair-colorizer
code --install-extension dandric.vscode-jq
code --install-extension DanielThielking.aws-cloudformation-yaml
code --install-extension dannysteenman.cloudformation-yaml-snippets
code --install-extension Dart-Code.dart-code
code --install-extension Dart-Code.flutter
code --install-extension DavidAnson.vscode-markdownlint
code --install-extension dbaeumer.vscode-eslint
code --install-extension docsmsft.docs-article-templates
code --install-extension docsmsft.docs-authoring-pack
code --install-extension docsmsft.docs-build
code --install-extension docsmsft.docs-images
code --install-extension docsmsft.docs-linting
code --install-extension docsmsft.docs-markdown
code --install-extension docsmsft.docs-metadata
code --install-extension docsmsft.docs-preview
code --install-extension docsmsft.docs-scaffolding
code --install-extension docsmsft.docs-visual-areas
code --install-extension docsmsft.docs-yaml
code --install-extension donjayamanne.githistory
code --install-extension donjayamanne.jquerysnippets
code --install-extension DotJoshJohnson.xml
code --install-extension dsznajder.es7-react-js-snippets
code --install-extension eamodio.gitlens
code --install-extension esbenp.prettier-vscode
code --install-extension ExodiusStudios.comment-anchors
code --install-extension fabiospampinato.vscode-diff
code --install-extension fabiospampinato.vscode-todo-plus
code --install-extension formulahendry.auto-rename-tag
code --install-extension formulahendry.code-runner
code --install-extension formulahendry.terminal
code --install-extension gera2ld.markmap-vscode
code --install-extension GitHub.copilot
code --install-extension GitHub.github-vscode-theme
code --install-extension GitHub.vscode-pull-request-github
code --install-extension golang.go
code --install-extension hashicorp.terraform
code --install-extension HookyQR.beautify
code --install-extension HookyQR.minify
code --install-extension huizhou.githd
code --install-extension jock.svg
code --install-extension john-goldsmith.vscode-aws-cloudformation-auto-template-generator
code --install-extension jprestidge.theme-material-theme
code --install-extension jq-syntax-highlighting.jq-syntax-highlighting
code --install-extension kddejong.vscode-cfn-lint
code --install-extension KnisterPeter.vscode-github
code --install-extension ljacobsson.cfn-resource-actions
code --install-extension mechatroner.rainbow-csv
code --install-extension mhutchie.git-graph
code --install-extension mintlify.document
code --install-extension ms-azuretools.vscode-azureappservice
code --install-extension ms-azuretools.vscode-azureresourcegroups
code --install-extension ms-azuretools.vscode-docker
code --install-extension ms-mssql.data-workspace-vscode
code --install-extension ms-mssql.mssql
code --install-extension ms-mssql.sql-bindings-vscode
code --install-extension ms-mssql.sql-database-projects-vscode
code --install-extension ms-python.anaconda-extension-pack
code --install-extension ms-python.isort
code --install-extension ms-python.python
code --install-extension ms-python.vscode-pylance
code --install-extension ms-toolsai.jupyter
code --install-extension ms-toolsai.jupyter-keymap
code --install-extension ms-toolsai.jupyter-renderers
code --install-extension ms-toolsai.vscode-jupyter-cell-tags
code --install-extension ms-toolsai.vscode-jupyter-slideshow
code --install-extension ms-vscode-remote.remote-containers
code --install-extension ms-vscode-remote.remote-wsl
code --install-extension ms-vscode.azure-account
code --install-extension ms-vscode.cpptools
code --install-extension ms-vscode.live-server
code --install-extension ms-vscode.powershell
code --install-extension ms-vsliveshare.vsliveshare
code --install-extension msjsdiag.vscode-react-native
code --install-extension mtxr.sqltools
code --install-extension mtxr.sqltools-driver-mysql
code --install-extension mtxr.sqltools-driver-pg
code --install-extension njpwerner.autodocstring
code --install-extension oderwat.indent-rainbow
code --install-extension PKief.material-icon-theme
code --install-extension platformio.platformio-ide
code --install-extension pmneo.tsimporter
code --install-extension pnp.polacode
code --install-extension qinjia.view-in-browser
code --install-extension rafwilinski.serverless-vscode-snippets
code --install-extension rangav.vscode-thunder-client
code --install-extension redhat.java
code --install-extension redhat.vscode-commons
code --install-extension redhat.vscode-yaml
code --install-extension ritwickdey.LiveServer
code --install-extension RobbOwen.synthwave-vscode
code --install-extension run-at-scale.terraform-doc-snippets
code --install-extension securisec.serverless-completer
code --install-extension shardulm94.trailing-spaces
code --install-extension shinhwagk.vscode-note
code --install-extension SirTori.indenticator
code --install-extension SonarSource.sonarlint-vscode
code --install-extension streetsidesoftware.code-spell-checker
code --install-extension TabNine.tabnine-vscode
code --install-extension TeamHub.teamhub
code --install-extension timkmecl.chatgpt
code --install-extension TimVaneker.serverless-snippets
code --install-extension tkrkt.linenote
code --install-extension tonybaloney.vscode-pets
code --install-extension trunk.io
code --install-extension Tyriar.sort-lines
code --install-extension usernamehw.errorlens
code --install-extension vikas.code-navigation
code --install-extension vscode-icons-team.vscode-icons
code --install-extension WallabyJs.quokka-vscode
code --install-extension WASTeamAccount.WebTemplateStudio-dev-nightly
code --install-extension wesbos.theme-cobalt2
code --install-extension xabikos.JavaScriptSnippets
code --install-extension xshrim.txt-syntax
code --install-extension yzhang.markdown-all-in-one
```

[![Visitors](https://api.visitorbadge.io/api/visitors?path=atharvaskadam&label=atharvaskadam&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=atharvaskadam)



# SSO Login With VSCode

[![Visitors](https://api.visitorbadge.io/api/visitors?path=atharvaskadam.vscode.atharvaskadam&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=atharvaskadam.vscode.awsssologin)

- [SSO Login With VSCode](#sso-login-with-vscode)
  - [Prerequisites](#prerequisites)
  - [Configure VSCode](#configure-vscode)
  - [Configure AWS Credentials](#configure-aws-credentials)
  - [SSO Login](#sso-login)
  - [Reference](#reference)

## Prerequisites

- Python 3 [Download Link](https://www.python.org/downloads/windows/)
- AWS CLI v2 [Download Link](https://awscli.amazonaws.com/AWSCLIV2.msi)
- Git [Download Link](https://git-scm.com/downloads)

Make sure to Uninstall any previous versions of above mentioned tools before installing new ones. Also make sure to remove the old path variable.

After installing everything, restart your system.

## Configure VSCode

- Open VS Code
- Then Open Terminal in VSCode
- Run `aws --version` command and make sure that AWS CLI V2 version is shown, for example `aws-cli/2.1.0 Python/3.7.7 Windows/11 exe/AMD64`
- Run `pip install pipx` command
- Run `python -m pipx ensurepath` Command
- Run `pipx install aws-sso-credential-process` Command
- Run `pip install git-remote-codecommit` Command

## Configure AWS Credentials

Open file `c:\users\[user]\.aws\credentials` and paste below code in the file

```config
[default]
credential_process = aws-sso-credential-process --profile default
sso_start_url = https://my-sso-portal.awsapps.com/start
sso_region = us-east-1
sso_account_id = 123456789011
sso_role_name = readOnly
region = us-west-2
output = json
```

or you can also run following command: `aws configure sso --profile default` this will insert all values to credential file

## SSO Login

Now to login use command `aws sso login`

You can use following command to Describes a user's SSH information `aws opsworks --region us-west-2 describe-my-user-profile` which will generate an output as shown below

```json
{
  "UserProfile": {
    "IamUserArn": "arn:aws:iam::123456789012:user/myusername",
    "SshPublicKey": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQ...3LQ4aX9jpxQw== rsa-key-20141104",
    "Name": "myusername",
    "SshUsername": "myusername"
  }
}
```

## Reference

- [Configuring the AWS CLI to use AWS IAM Identity Center (successor to AWS Single Sign-On)](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html)
- [describe-my-user-profile](https://docs.aws.amazon.com/cli/latest/reference/opsworks/describe-my-user-profile.html)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=atharvaskadam&label=atharvaskadam&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=atharvaskadam)
