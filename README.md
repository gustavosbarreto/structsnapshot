# structsnapshot

structsnapshot is a Go package that allows you to take a snapshot of a struct,
saving its fields, types, and tags in a JSON format.
It provides functions to create a snapshot, convert it to JSON,
save it to a file, and load a snapshot from a file.

## Installing

```
go install github.com/gustavosbarreto/structsnapshot/cmd/structsnapshot@latest
```

## Using

There are two methods of generating struct snapshots (see bellow).

> Regardless of the method you choose, a special directory called
> `__structsnapshots__` will be created in the current working directory.
> This directory will contain the generated struct snapshots as JSON
> files for each struct you have configured.

### Go generate (RECOMMENDED)

Add a `go:generate` comment before the struct:

```go
//go:generate structsnapshot User
type User struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Timezone string `json:"timezone" validate:"timezone"`
}
```

Run go generate:
```
$ go generate
```

## Standalone

Run the following command on the directory where the User struct is defined:

```
$ structsnapshot User
```

## Unit testing

After generating struct snapshots, you need to ensure that your structs match
the generated snapshots. To accomplish this, you can use the `TakeSnapshot(struct)`
function to generate an in-memory snapshot of the struct.
This snapshot can then be compared with the generated snapshot on the filesystem by
calling `LoadSnapshot(struct)` function.

See [example](./example).
