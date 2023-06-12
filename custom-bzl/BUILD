load(":create_hello.bzl", "hello")

hello(name = "hello")

load(":write_world.bzl", "world")

world(
    name = "world",
    deps = [":hello"],
)
