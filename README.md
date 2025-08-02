# JobFinder Application

JobFinder is a application to search and match jobs, that create to practice, is creating with Vue 3, Vite, Vuetify, Pinia and Go.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (for the API server)
- [Node.js](https://nodejs.org/) (for the frontend application)
- [npm](https://www.npmjs.com/) (Node package manager)

## Initialize

### 1. Set Up the Backend (API)

1. Navigate to the API directory:
   ```bash
   cd api
   ```

2. Install the required Go dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables (if needed):
   - Create a `.env` file in the `api` directory
   - Add any required environment variables

   ```
   API_PORT=
   API_HOST=
   API_VERSION=
   CLIENT_HOST=
   ENVIRONMENT="development"
   ```

4. Start the API server:
   ```bash
   go run main.go
   ```
  or `air` if you have it installed

   The API should now be running on `http://localhost:8080` (or the configured port).

### 2. Set Up the Frontend

1. Open a new terminal window and navigate to the app directory:
   ```bash
   cd ../app
   ```

2. Install the required Node.js dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```
   The frontend should now be running on `http://localhost:5173` (or the configured port).


## Project Structure

- `/api` - Backend API server (Go)
- `/app` - Frontend application (Vue.js/TypeScript)

## Notes

this project is only for practice, and is not a production ready application, not use to comerce or production.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
