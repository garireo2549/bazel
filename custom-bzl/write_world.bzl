def _world_impl(ctx):
    print(ctx.files.srcs)

world = rule(
    implementation = _world_impl,
    attrs = {
        "srcs": attr.label_list(allow_files = [".txt"]),
        "deps": attr.label_list(),
    },
#    outputs = {
#      "main": "%{name}.txt",
#    },
)