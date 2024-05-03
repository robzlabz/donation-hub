# DonationHub

DonationHub is a web application designed to facilitate connections between donors and recipients, allowing donors to contribute to projects and recipients to create projects to receive donations. It's primarily built as a learning project for mastering backend development with the Go programming language. While the emphasis is on backend development, the frontend interface is provided as a reference, requiring no modifications.

## Your Missions

1. Implement all features defined in the provided API documentation and the web prototype with [Hexagonal Architecture](https://github.com/Haraj-backend/hex-monscape/blob/master/docs/reference/hex-architecture.md), including:
    - User registration for donors and recipients.
    - User login functionality.
    - Home page interface for guests, donors, and recipients, displaying relevant information such as total donations distributed and a list of donation projects.
    - Donation project detail page, featuring donation details, a donate button, and a list of donors along with their messages.
    - Admin interface with access to lists of donation projects, donors, and recipients for administrative purposes.
2. Implement qualified unit tests to ensure the reliability and correctness of your code.
3. Implement local deployment for easy testing and development.
4. Ensure maintainability of the codebase:
    - Write clear and concise code comments to explain the functionality of complex sections of code.
    - Adhere to consistent coding conventions and naming conventions to make the codebase more understandable for other developers.
    - Provide additional documentation that you think is necessary to help other developers understand the codebase.

## Important Files

When starting DonationHub, please refer to the following files for essential information:

- `cmd/main.go`: Entry point of the application.
- `cmd/web/index.html`: Prototype frontend interface for this project.
- `docs/http_api.md`: Documentation outlining the API endpoints and usage.
- `docs/schema.sql`: Definition of the database schema for DonationHub.

## Evaluation

We will evaluate your submission based on the following criteria:

1. The completeness and correctness of your implementation of all defined features.
2. The quality and coverage of your unit tests.
3. The effectiveness of your local deployment setup.
4. The maintainability of the codebase, including clarity of code comments, adherence to coding conventions, modularity of code organization, and helpfulness of additional documentation.

## Submission

1. Fork this repository & do your work in your own forked repository.
2. We will review your submission progress weekly, so please push your code to your forked repository regularly.
3. The due date for the final submission is on 24th May 2024.
