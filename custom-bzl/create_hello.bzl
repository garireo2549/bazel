def _hello_impl(ctx):
    ctx.actions.write(ctx.outputs.main, "hello")

hello = rule(
    implementation = _hello_impl,
    outputs = {
      "main": "%{name}.txt",
    },
)
