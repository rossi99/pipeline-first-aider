load('ext://uibutton', 'cmd_button', 'text_input', 'choice_input')

# Vars
wd = os.getcwd()

# Resources
local_resource(
    labels=["docker"],
    name="Build",
    cmd="docker build -t first-aider:latest -f Dockerfile .",
    auto_init=False
)

local_resource(
    labels=["docker"],
    name="Run",
    cmd="docker run --rm -v '{0}/logs/pass.log:/app/build.log:ro' first-aider:latest --log-path /app/logs/pass.log".format(wd),
    auto_init=False
)
cmd_button(
    name="choose dummy log",
    text="Select dummy logs",
    resource="Run",
    argv=[
        "bash", "-c",
        "docker run --rm -v '{0}/logs/$log_type.log:/app/build.log:ro' first-aider:latest --log-path /app/logs/$log_type.log".format(wd)
    ],
    inputs=[
        choice_input(
            "log_type",
            label="log type",
            choices=["fail", "mixed", "pass"],
        ),
    ]
)