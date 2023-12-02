# Tentang Anak - Pokedex API

## Overview

This is the repository for the Pokedex API, a RESTful API service for managing monsters as part of the Tentang Anak
test.

The Pokedex API currently supports operations like listing all monsters, adding a new monster, updating an existing
monster, deleting a monster, and more.

## Features

- Add a new monster to the database
- List all available monsters or filter them by various parameters
- Get detailed information about specific monster by id
- Update a monster's details and status
- Delete a monster by id

## API Documentation

Detailed API documentation is provided via an OpenAPI Specification 3.1.0. The API documentation is available in a JSON file
named `api-v1.json`.

## Endpoints
The following endpoints are available:

- `POST /monsters`: Add a new monster
- `GET /monsters`: Get a list of all available monsters with optional filters
- `GET /monsters/{id}`: Get details of a specific monster by id
- `PATCH /monsters/{id}`: Update a specific monster's details
- `PATCH /monsters/{id}/options`: Update a specific monster's options
- `DELETE /monsters/{id}`: Delete a specific monster by id

## API Responses
The API uses standard HTTP response codes as well as return well-structured response bodies. See the API documentation for the structure of responses from each endpoint.
