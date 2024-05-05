## Job Scheduling Application

### Folder Structure

- client - Fronte with React + Typescipt
- server - Backend with Golang and websockets

#### Environment Setup

1. **Create Environment Files**: 
 - Navigate to the root directory of your project.
 - Create `.env` files for both the frontend and backend environments.

2. **Frontend Environment Setup**:
 - In the `.env` file within the `client/job-sheduler-fe` directory, add:
   ```
   VITE_BACKEND_URL=localhost:8080/api/v1
   ```

3. **Backend Environment Setup**:
 - In the `.env` file within the `server` directory, add:
   ```
   APP_PORT=8080
   SERVER_API_PREFIX_V1=api/v1
   SERVER_BASE_PATH=localhost
   ```

4. **Running the Frontend Server**:
 - Navigate to the `client/job-sheduler-fe` directory.
 - Run:
   ```
   npm run dev
   ```

5. **Running the Backend Server**:
 - Navigate to the `server` directory.
 - Run:
   ```
   make server
   ```

6. **Testing**:
 - Access the frontend server URL (`http://localhost:3000`) in your browser and interact with the application's frontend interface.

