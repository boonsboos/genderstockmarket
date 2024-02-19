# the gender (and sexuality) stock market

you trade in the commodities of genders and sexualities to become the richest gaytrader of the world

missing a gender or sexuality? please open an issue!

---
a build-your-own-client incremental game where you you need to trade to make money

compete for the leaderboards and grow your firm

the Spectrum 300 is waiting for you!

## building

- go 1.21 (or greater)
- postgres 16 (or greater)

run `Database/Migrations/Up.pgsql` on your database

create a file in the project root called `options.json` in the following format:

```json
{
    "databaseURL": "theURLToYourPostgresDatabaseHere",
    "databaseName": "yourDatabaseNameHere",
    "githubID": "yourGithubOAuthAppIDHere",
    "githubToken": "yourGithubOAuthSecretHere"
}
```

build the project with `go build .` and run the executable