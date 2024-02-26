# Readme

This project shows the behavior of gqlgen's generated resolver code with enums.

## Schema

The schema has an enum `Priority`. The `Todo` type has a property `priority` that is of type `priority`.

### Generated

## As Input

The mutation `createTodo` takes the Priority enum as a value. It is strictly validated at runtime.

## As Output

The resolver constructor is seeded with a Todo struct that contains a value for `Priority` that is not a member of the enum. This is acceptable at runtime.
