https://golang.org/doc/tutorial/compile-install
https://golang.org/doc/modules/developing#discovery
https://golang.org/doc/modules/developing#workflow
https://golang.org/doc/modules/release-workflow#unpublished
https://golang.org/doc/modules/version-numbers
https://golang.org/doc/modules/publishing
https://golang.org/doc/modules/developing
https://golang.org/doc/modules/managing-dependencies#removing_dependency


# Removing a dependency

When your code no longer uses any packages in a module, you can stop tracking the module as a dependency.

To stop tracking all unused modules, run the go mod tidy command. This command may also add missing dependencies needed to build packages in your module.

```
$ go mod tidy
```

To remove a specific dependency, use the go get command, specifying the module's module path and appending @none, as in the following example:

```
$ go get github.com/andrewlunde/thetaoffchaingo_common@none
```

The go get command will also downgrade or remove other dependencies that depend on the removed module.
