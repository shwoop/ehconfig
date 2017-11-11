package main

// main is the entrypoint fo the application.
func main() {
	c := buildConfig()
	c.action(c)
}
