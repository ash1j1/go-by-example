# TAKEAWAYS

**1. go run vs go build**
- `go run` is mainly used for local development to quickly debug/test as it executes the program without saving the binary in the working directory by quickly compiling and running the `main` package
- `go build` produces an executable from a `main` package into a single and statically linked executable program. this is mainly used in production.

**2. main packages and library packages**
- a package that is named as `main` has an entry point at the `main()` function and a `main` package is compiled as an executable program
- any other named packages is a `library package` which exports functionalities that can be used by other packages and don't have any entry points

**3. static linked binaries** [Yet to understand clearly]
