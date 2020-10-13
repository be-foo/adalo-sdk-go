# Contributing

Contributions in the form of bug reports, code contributions or even feature requests are very welcome.

To make a contribution, fork the project and checkout a branch in the pattern `feature/<FEATURE-DESCRIPTION>`
or `bugfix/<BUG-DESCRIPTION>` from the `develop` branch.

Commit your changes to this branch, but be sure to rebase the commits at the end. Before merge, the
branch should contain only one commit with the same name as your branch (without the prefix).

It is desirable to post an according issue in our issue board first and then reference that issue in your PR.

## Testing

Please add tests when extending the project with new features!
Functionality that connects with the Adalo API will require real credentials to an Adalo app.
That's where we use environment variables. When testing locally:

````sh
# Create the .env file and edit the .env file and insert your keys
cp .env.example .env

# Apply environment variables to your terminal
export $(cat .env | xargs)
````

And then you can run the test suite like
```sh
go test -v ./...
```

Our CI will run the tests using real Adalo keys too.
