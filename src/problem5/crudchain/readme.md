# crudchain
**crudchain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Details
I have created the following:
- resource is created with the following fields: name, description + I created this while following the tutorial and also introduce consensus-breaking-change to it in the branch `breaking-consensus-change`.
- film is created with the following fields: name, description, genre + This is the main resource I do CRUD operations with.

### CRUD operations on film:
- **Create** create new film with the fields `name`, `description`, `genre` using the command:
<br>  ```crudchaind tx crudchain create-film [name] [description] [genre] [flags]```
- **Read**  I have created some methods for reading films:
```bash
Usage:
  crudchaind query crudchain [flags]
  crudchaind query crudchain [command]

Available Commands:
  list-film             List all films
  show-film             Shows a film by id
  show-film-genre       Query show-film-genre. This will show all the films whose genre matches the keyword
  show-film-name        Query show-film-name. This will show all the films whose name mathes the keyword
```
- **Update** update existing film using the command:
<br> ```crudchaind tx crudchain update-film [id] [name] [description] [genre] [flags]```

- **Delete** delete existing film using the command:
<br> ```crudchaind tx crudchain delete-film [id] [flags]```
