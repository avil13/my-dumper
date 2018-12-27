# My DB Dumper

## A simple utility for creating mysql dumps.

> wip

> Just 4 fun

### Is able to:

- Dump the database with the time of creation of the file in the title
- Separates table creation files and data insertion files
- Execute SQL data from files
- Ignores the dump tables specified in the config
- Uses `.env` file to set up environment


## Create dump
```bash
$ ./dmpr
```

## Execute SQL data from files
```bash
$ ./dmpr -import path/to/file.sql
```

## `.env` config

For create exemple .env file
```bash
$ ./dmpr -make-env
```

### env parameters

.env param | required | example | description
|---|---|---|---|
TITLE | true | example.com |
DB_HOST | true | mysql |
DB_PORT | true | 3306 |
DB_DATABASE | true | homestead |
DB_USERNAME | true | homestead |
DB_PASSWORD | true | secret |
DUMP_CREATE | true | true | create tables file
DUMP_INSERT| true | true | insert data file
DUMP_DIR | false | dumps | Folder for dump
IGNORE_TABLES |false| peoples\|likes | if you want to ignore some tables, specify them with the symbol \| 
DEBUG | false | false | for debugging the file names do not contain the creation date


### flags

```bash
$ ./dmpr -make-env # create .env file
```

```bash
$ ./dmpr -all # Ignore .env parameters DUMP_CREATE and DUMP_INSERT
```

```bash
$ ./dmpr -import path/to/file.sql # Execute sql file
```

```bash
$ ./dmpr -help
Usage of ./dmpr:
  -all
    	create all dump files
  -import string
    	import dump file
  -make-env
    	create .env boilerplate file
```


### example of a generated file

```bash 
$ tree .
.
├── dmpr
└── dumps
    └── 2018-12-27
        └── 2018-12-27_03:57_insert.sql

```