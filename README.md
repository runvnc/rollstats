(Image of output)

## What this does

This will perform a standard Dungeons & Dragons ability score roll simulating 4d6 with the lowest dropped and calculating your bonus.
It runs as a stateful program on the Algorand blockchain. Has a few ANSI colors to enhance the output.
You can try it out with the included scripts that call `goal` if you have an Algorand node or sandbox setup.

I use **`fish` shell, so that is required** with these exact scripts (not hard to modify them for bash). You will need the `gpp` program installed.
Also, before running, you will need to **first set the environment variable `$APP_CREATOR`**.

```shell
./init.sh
```

```shell
./call.sh
```

## How

It is written in [Tealang](https://github.com/pzbitskiy/tealang), also using the GPP macro preprocessor.  The macros are helpful to conserve instruction executions, since stateful programs
on the Algorand blockchain only get a maximum of 700 instructions executed.

## Why?

Not to offend anyone, but the most accurate technical term is "pure fuckery".

Note: tealang is a work in progress, and has a lot of gotchas at the moment I am writing this.  But they seem to be working
diligently to update it to the new AVM version.
