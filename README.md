# gql
`gql` is a collection of tools to manage GraphQL schema. The tool can be used for linting schema and finding breaking changes in the schema.

## Usage
`gql` can be installed with `go install github.com/CrowdStrike/gql`  
```shell
~ $ gql -h
gql is a CLI built for federated GraphQL services' schemas

Usage:
  gql [command]

Available Commands:
  compare     compare two graphql schemas
  help        Help about any command
  lint        lints given GraphQL schema

Flags:
  -h, --help   help for gql

Use "gql [command] --help" for more information about a command.
```

## linter
The linter is inspired from [graphql-schema-linter](https://github.com/cjoudrey/graphql-schema-linter/) with changes for supporting Apollo federation. 
### How to use?
Help command prints help message. This has all the supported rules. 
```shell
~ $ gql lint -h
lints given GraphQL schema

Usage:
  gql lint [flags]

Flags:
  -f, --filepath string   Path to your GraphQL schema
  -h, --help              help for lint
  -r, --rules strings     Rules you want linter to use e.g.(-r type-desc,field-desc); available rules:
                           	type-desc => type-desc checks whether all the types defined have description
                          	args-desc => args-desc checks whether arguments have description
                          	field-desc => field-desc checks whether fields have description
                          	enum-caps => enum-caps checks whether Enum values are all UPPER_CASE
                          	enum-desc => enum-desc checks whether Enum values have description
                          	field-camel => field-camel checks whether fields defined are all camelCase
                          	type-caps => type-caps checks whether types defined are Capitalized
                          	relay-conn-type => relay-conn-type checks if Connection Types follow the Relay Cursor Connections Specification
                          	relay-conn-args => relay-conn-args checks if Connection Args follow of the Relay Cursor Connections Specification
```
Specifying the schema file:
```shell
~ $ gql lint -f schema.graphqls
2022/05/28 12:53:01 16 errors occurred:
	* 6:3 field Todo.id does not have description
	* 7:3 field Todo.text does not have description
	* 8:3 field Todo.done does not have description
	* 9:3 field Todo.user does not have description
	* 13:3 field User.id does not have description
	* 14:3 field User.todos does not have description
	* 14:9 argument User.todos.offset does not have description
	* 14:21 argument User.todos.limir does not have description
	* 15:3 field User.color does not have description
	* 19:3 field Query.todos does not have description
	* 19:9 argument Query.todos.offset does not have description
	* 19:22 argument Query.todos.limit does not have description
	* 23:3 field NewTodo.text does not have description
	* 24:3 field NewTodo.userId does not have description
	* 28:3 field Mutation.createTodo does not have description
	* 28:14 argument Mutation.createTodo.input does not have description
```

Reading schema from stdin:
```shell
~ $ cat schema.graphqls | gql lint
2022/05/28 12:54:48 16 errors occurred:
	* 6:3 field Todo.id does not have description
	* 7:3 field Todo.text does not have description
	* 8:3 field Todo.done does not have description
	* 9:3 field Todo.user does not have description
	* 13:3 field User.id does not have description
	* 14:3 field User.todos does not have description
	* 14:9 argument User.todos.offset does not have description
	* 14:21 argument User.todos.limir does not have description
	* 15:3 field User.color does not have description
	* 19:3 field Query.todos does not have description
	* 19:9 argument Query.todos.offset does not have description
	* 19:22 argument Query.todos.limit does not have description
	* 23:3 field NewTodo.text does not have description
	* 24:3 field NewTodo.userId does not have description
	* 28:3 field Mutation.createTodo does not have description
	* 28:14 argument Mutation.createTodo.input does not have description
```
Passing subset of rules to be applied
```shell
~ $ gql lint -f schema.graphqls -r types-have-description
2022/05/28 13:23:20 1 error occurred:
	* 3 errors occurred:
	* 5:6 type Todo does not have description
	* 22:7 type NewTodo does not have description
	* 27:6 type Mutation does not have description
```
Specifying wildcards for schema file paths
```shell
graphql-linter -f '*.graphqls' -r types-have-description
2022/05/28 13:23:20 1 error occurred:
	* 3 errors occurred:
	* 5:6 type Todo does not have description
	* 22:7 type NewTodo does not have description
	* 27:6 type Mutation does not have description
```

> Note: If your argument has wildcards your shell can execute the glob and provide individual values to graphql-linter. 
> So don't forget the quotes around path with wildcards.

## Available rules 
Following table describes all the lint rules supported by the linter

| Lint Rule       | Description   |
| :-------------: |:--------------|
| type-desc       | type-desc checks whether all the types defined have description |
| args-desc       | args-desc checks whether arguments have description |
| field-desc      | field-desc checks whether fields have description |
| enum-caps       | enum-caps checks whether enum values are all UPPER_CASE |
| enum-desc       | enum-desc checks whether enum values have description |
| field-camel     | field-camel checks whether fields defined are all camelCase |
| type-caps       | type-caps checks whether types defined are Capitalized |
| relay-conn-type | relay-conn-type checks whether types defined are following relay cursor connection spec |
| relay-conn-args | relay-conn-args checks whether args defined are following relay cursor connection spec |


### Enabling and disabling certain rules
It is possible that you want certain part of your schema to be ignored for linting. You can do it with comments in your schema: 

#### ignoring multiple lines:
If you use `#lint-disable rule1` on line x and `#lint-enable rule1` for line y where x < y then errors because of rule1 
between line x and y will be ignored from the output.  

#### ignoring single line:
You can use `#lint-disable-line rule1` to disable rule1 for a specific line. This only is applicable if rule1 is applied 
for rest of the schema. If rule1 is not one among the rules passed then this has no effect on output.

##compare
compare command compares two schema files and returns all the differences. It is also built to support Apollo federation
specification and can be used to find breaking changes in schema.
  
 ### How to use?
 ```shell 
$gql compare -h
compare two graphql schemas

Usage:
  gql compare [flags]

Flags:
  -b, --breaking-change-only     Get breaking change only
  -e, --exclude-print-filepath   Exclude printing schema filepath positions
  -h, --help                     help for compare
  -n, --newversion string        Path to your new version of GraphQL schema
  -o, --oldversion string        Path to your older version of GraphQL schema
```

Compare the schema
```shell 
~ $ gql compare -o oldSchema.graphql -n newSchema.graphql
❌   /Users/spahariya/code/temp/newSchema/library.graphql:2  Argument 'from' type changed from 'String' to 'String!' in directive '@transform'
❌  /Users/spahariya/code/temp/newSchema/library.graphql:20  Field 'Book.title' type changed from 'String!' to 'String' in OBJECT
❌  /Users/spahariya/code/temp/newSchema/library.graphql:18  Field 'Book.year' was removed from OBJECT
❌  /Users/spahariya/code/temp/newSchema/user.graphql:12  Input field 'UserInput.adBooks' type changed from '[Book]' to '[Book]!' in input object type
❌  /Users/spahariya/code/temp/newSchema/user.graphql:11  Input field 'UserInput.newBooks' type changed from '[Book]!' to '[Book!]!' in input object type
✋   /Users/spahariya/code/temp/newSchema/library.graphql:4  Member 'text' was added to Union type 'body'
✅   /Users/spahariya/code/temp/newSchema/library.graphql:19  Field 'Book.isbn' type changed from 'String' to 'String!' in OBJECT
✅  /Users/spahariya/code/temp/newSchema/library.graphql:15  Field 'Library.books' was added to OBJECT
✅  /Users/spahariya/code/temp/newSchema/library.graphql:8  Field 'Query.books' was added to OBJECT
✅  /Users/spahariya/code/temp/newSchema/user.graphql:6  Field 'User.address' was added to OBJECT
✅  /Users/spahariya/code/temp/newSchema/user.graphql:5  Field 'User.name' type changed from 'String' to 'String!' in OBJECT
Breaking errors in schema: 5
```
If user want to exclude schema filepath positions from the stdout, pass -e ot --exclude-print-filepath option with the command
``` 
~ $ gql compare -o "oldGraphql/*.graphql" -n "newGraphql/*.graphql" -e
❌  Field 'Book.isbn' type changed from 'String!' to 'String' in OBJECT
❌ Field 'Library.books' was removed from OBJECT
❌ Field 'Query.books' was removed from OBJECT
✋ Member 'text' was added to Union type 'body'
✅  Argument 'from' type changed from 'String!' to 'String' in directive '@transform'
✅ Field 'Book.title' type changed from 'String' to 'String!' in OBJECT
✅ Field 'Book.year' was added to OBJECT
Breaking errors in schema: 3
```

## Type of changes in schema
Generally speaking either a change can break API contract with client or it won't. But in case of GraphQL there's another category of changes,
which won't actually break clients but will change their behavior and if not handled properly in code will cause client-side errors. Thus developers need 
to pay special attention to not only what are the breaking changes but also the dangerous ones. Following is the list of breaking and dangerous changes which can be made 
in graphql schema. 

* **Breaking❌**: Changes that will break existing queries to the GraphQL API. Below are the type of changes that comes in breaking change category. 
   - Schema root operation type (`Query`, `Mutation`, `Subscription`) changed or removed
   - `type`/`field`/`directive`/`interface` removed
   - `type` kind changed
   - `directive` location changed 
   - `field` type changed or made optional
   - `input` type changed or made required or required fields added
   - argument type changed or removed or made required or required arguments added
   - `enum` value removed
   - Union member removed
  
  
* **Dangerous✋** : Changes that won't break existing queries but could affect the runtime behavior of clients. Below are the type of changes that comes in dangerous change category.
   - Argument default value changed
   - `directive` optional argument added
   - Deprecation added/removed to `field`/`enum` values
   - `enum` value added
   - Interfaces added to object implements
   - Union member added
   - Input fields added