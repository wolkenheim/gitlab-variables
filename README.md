# gitlab-variable update tool

Ever spent hours copy & pasting files to your Gitlab CI variables section? And then even more hours debugging your 
containers as environment variables may or may not been set? Lo and behold... 

Frankly speaking, this is just a small project and will not fix all possible scenarios and edge cases.

## What it does
* Enable versioning for Gitlab CI variables
* Automated updating of Gitlab CI variables via REST API

## Caveats
There are a lot of caveats here. The ones I can think of are:
Renaming the names of a variable will not be possible. There is no good way to detect that a key has changed. The name
IS the unique key. If you rename a variable, a new variable in Gitlab will be created and the old one deleted.

A word on security: this app will download sensitive data like passwords and store it as backup files on your computer.
In some cases this might be okay, in others this absolutely not what you want.

## What I learned
The app is missing an important part. Somehow the variables from your project are stored somewhere
and need to be formatted into the Gitlab json format for the API. I figured out quickly that here there are not a few 
edge cases but only edge cases. Every one of our own project was different. E.g. in one project certificates needed
to be base64 decoded and stored as variables. In other projects various .env Files like .staging.env and .production.env
with the same variable names needed to be mapped. This is far beyond the scope of this little project and the 
easiest way is probably to solve this with shell scripting. As variables hardly ever change so a fully automated 
solution is a bit of an overhead.

## Test
```
go test -v ./...
```

## Run locally
```
export GITLAB_ENV=demo_project
go run . init
go run . update
```