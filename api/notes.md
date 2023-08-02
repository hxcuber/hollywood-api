api/internal/api/rest contains handlers. These are used to validate query parameters, then call the controllers of the specific request. Handlers usually return a HandleFunc, which can then be used in routing.

Controllers handle the business logic, which is to say that they handle the actual queries, running them and returning the data. Handlers call controllers once all parameters have been validated.
