(Image of output)

## What this does

This will perform a standard Dungeons & Dragons ability score roll simulating 4d6 with the lowest dropped and calculating your bonus.
It runs as a stateful program on the Algorand blockchain. Has a few ANSI colors to enhance the output.
You can try it out with the included bash script (without even having to clone the repo):

```shell
wget https://raw.github.com/
bash roll.sh
```

## How

It is written in tealang (repo), also using the GPP macro preprocessor.  The macros are helpful to conserve instruction executions, since stateful programs
on the Algorand blockchain only get a maximum of 700 instructions executed.

## Why?

Not to offend anyone, but the most accurate technical term is "pure fuckery".

Note: tealang is a work in progress, and has a lot of gotchas at the moment I am writing this.  But they seem to be working
diligently to update it to the new AVM version.
