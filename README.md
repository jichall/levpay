# Levpay

The company is a fintech in the field of digital payments. For more information
head onto their [website](https://levpay.com/https://levpay.com/).

# Usage

To use the project you will have to follow the instructions on each respective
topic below.

## Build

To build the software you will need to have [go 1.14](https://golang.org) and
`make`. Make is a building tool that comes within any Linux distribution
nowadays.

When go 1.14 is installed and make is present on the system you can run `make`
at the root directory of this project and the compiler will generate a binary
called `hero.out` inside the **bin** folder.

## Set up the environment

First you have to set up the database. The project uses
[postgresql](https://postgresql.org/) and with that installed you will have to
create the database of the application, this step can be automated by using the
script inside the database folder.

To run the software you have to export some variables to the environment, only
when the environment variables `HOST`, `PORT`, `DATABASE_URL` the program can be
ran and work properly.

With the env variables exported you can just execute the binary inside the bin
folder as simple as `./hero.out` inside the **bin** folder.

## API

The API is RESTful and uses HTTP verbs and status code heavily (as one would
expect from such API). The following routes describes the resources that can be
used.

|   URL   |   Methods    |                       Description                           |
|---------|--------------|-------------------------------------------------------------|
| /super/ | POST, DELETE | Responsible for adding or removing a superhero from the API |
| /super/ | GET          | Responsible for returning a super given a param             |
| /supers | GET          | Responsible for returning supers given a param              |

