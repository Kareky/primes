# Primes

A simple project with a collection of different algorithms to test whether a number is prime, find prime up to a certain number, or other prime-related operations.

## Running the Application

The application provides two distinct entry points:

1. Normal application – starts the main program (primality testing, sieves, etc.).
2. Seeder – populates the database with prime numbers up to a given bound and exits.

---

### Normal Startup

Start the main application:

```bash
go run main.go
```

### Specify a custom configuration file

```bash
go run main.go -config=config.yaml
```

### Flags (main.go)

| Flag          | Default           | Description                          |
| ------------- | ----------------- | ------------------------------------ |
| ```-config``` | ```config.yaml``` | Path to the YAML configuration file. |

---

### Seeding the Database

Run the seeder to generate and store primes:

```bash
go run ./cmd/seed
```

Customize the seeding process:

```bash
go run ./cmd/seed -bound=50000000 -algo=eratosthenes
```

### Override database connection settings

```bash
go run ./cmd/seed -dbPath=./data/test.db -dbType=sqlite -bound=1000000
```

### Flags (cmd/seed)

| Flag          | Default           | Description                                                  |
| ------------- | ----------------- | ------------------------------------------------------------ |
| ```-config``` | ```config.yaml``` | Path to the YAML configuration file.                         |
| ```-path```   | ```""```          | Override the database file path from the config.             |
| ```-type```   | ```""```          | Override the database type (e.g., "sqlite") from the config. |
| ```-bound```  | ```1000000000```  | Upper bound for prime generation (inclusive).                |
| ```-algo```   | ```eratosthenes```| Algorithm to use for generation.                             |

---

#### Notes

- The seeder is idempotent: primes already present in the database are skipped (primary key constraint prevents duplicates).
- Seeding can take a long time for very large bounds. Use reasonable values (≤ 1e9 recommended).
- All flags are optional; if omitted, the values from the configuration file are used.
- If you change the database contents (e.g., by adding new primes via other means), the application's internal size limit is automatically updated.
