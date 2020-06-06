# Golang Application Template with JSON Configurations

Simple Application with Configurations given in a JSON named <i>config.json</i>

## Change these values to add more Configurations to your app
Fields that can be found inside the <i>config.json</i> file:
``` Go
type appConfig struct {
	Name             string
	Description      string
	ServerAddress    string
	OptionalArgument string
}
```

***

Define optional fields (i.e: specify that these fields are not required inside the <i>config.json</i> file)
``` Go
optionalFields = append(optionalFields, "OptionalArgument",
		"OptionalArgument2",
		"OptionalArgument3")
```

***
## How to Run:

To run the template locally:
``` Go
go run .
```

***

To run the tests locally:
``` Go
go test ./...
```

Or run the script:
``` Go
sh ./test_config.sh
```

***

To run the template in Docker:
```
docker build -t <image_name> .
docker run <image_name>
```

***