# 02_02 Use Starter Workflows in GitHub Actions
If you’re creating a new workflow, GitHub will analyze any files stored in your repo and suggest workflows based on the names of the files, the programming language the files contain, and any frameworks that might be configured for your project.

## Identifying Projects by Files
GitHub provides starter workflows for a variety of languages and tooling.

The following file extensions and dependency file names are used to identify projects for recommended workflows.

| Programming Language | File Extensions | Dependency File Names          |
|----------------------|-----------------|--------------------------------|
| JavaScript           | .js, .mjs       | package.json, package.lock.json|
| Python               | .py             | requirements.txt, Pipfile.lock |
| Java                 | .java           | pom.xml, build.gradle          |
| Ruby                 | .rb             | Gemfile, Gemfile.lock          |
| PHP                  | .php            | composer.json, composer.lock   |
| C#                   | .cs             | packages.config, project.json  |
| Go                   | .go             | go.mod, go.sum                 |
| Swift                | .swift          | Package.swift                  |
| Kotlin               | .kt, .kts       | build.gradle, build.gradle.kts |
| TypeScript           | .ts             | package.json, yarn.lock        |
| C/C++                | .c, .cpp        | CMakeLists.txt, Makefile       |
| Rust                 | .rs             | Cargo.toml, Cargo.lock         |
| Scala                | .scala          | build.sbt                      |
| Lua                  | .lua            | rockspec                       |
| Elixir               | .ex             | mix.exs                        |
| Haskell              | .hs             | package.yaml, stack.yaml       |
| Julia                | .jl             | Project.toml, Manifest.toml    |
| R                    | .r, .R          | DESCRIPTION, NAMESPACE         |
| Dart                 | .dart           | pubspec.yaml                   |
| Clojure              | .clj            | project.clj, deps.edn          |
| Groovy               | .groovy         | build.gradle, build.gradle.kts |
| Perl                 | .pl             | META.json, META.yml            |
| Objective-C          | .m              | Podfile, Cartfile              |

## Recommended Reading
- [Using starter workflows](https://docs.github.com/en/actions/using-workflows/using-starter-workflows)

- [Starter Workflows - Repo](https://github.com/actions/starter-workflows)

- [Using workflows](https://docs.github.com/en/actions/using-workflows)

- []()

