# project02

Names: Jeongseok Suh (jsuh3), Christina Youn (cyoun)

## Goal of the project
We wanted to created a virtual pet experience for the user. The application allows the user to interact with the panda in three different ways: (1) Feed, (2) Walk, (3) Play. These actions will change the pet's attributes (hunger, health, mood) and will also display different animations. If the pet is not attended to for a period of 30 seconds, the pet's attributes will decrement by 5.

## How to run the code
After getting the code, compile the server-side code by running:
> go build panda.go

Then, start the server side code:
> ./panda

Now, connect to the server on port `3000`

## Unique features of Golang
We were initially going to use Swift but decided to use Golang once we realized that the standard library included a web server. This made building the web app very simple. We created handlers for the different pages and were able to pass a struct called PageVariables to the HTML code to change the page view. We also really liked that Golang has short variable declarations. Altough this was confusing at first, it made our lives a lot easier when deploying code quickly.
