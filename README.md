On an mac M1 with docker Desktop installed make sure this setting is enabled (Rosetta):
![Screenshot 2023-09-22 at 9.44.35 AM.png](Screenshot%202023-09-22%20at%209.44.35%20AM.png)

then in a shell run this:
`docker compose up` 

Then in NEW shell run this:
`docker exec -it nakama_debug-nakama-1 /bin/bash`

In that NEW shell run this in the docker:

`/nakama/nakama migrate up --database.address postgres:localdb@postgres:5432/nakama`

and

`./dlv --log --log-output=debugger exec /nakama/nakama -- --config /nakama/data/local.yml --database.address postgres:localdb@postgres:5432/nakama`

and in the debugger run:

`break main.go:181`

You should hit an error:

`Command failed: could not find statement at github.com/heroiclabs/nakama/v3/main.go:181, please use a line with a statement
Set a suspended breakpoint (Delve will try to set this breakpoint when a plugin is loaded) [Y/n]?
`

And then nothing works after this. 