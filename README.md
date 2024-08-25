# kia-logix
A smart logistics solution for order tracking, status updates, and automated customer notifications.
## 🛠️ How to Run the Project

### Using Docker
**Prerequisites**: Docker

1. Clone the repository from GitHub.
2. Rename the `config.yaml.example` file to `config.yaml`. (change the values if you want)(daily job time is set here)
3. Rename the `.env.example` file to `.env`. (change the values if you want)
4. Run the `make dockerize` command.
5. Now import postman collection from `docs` directory into your postman. (Change the BASE_URL in the collection variables if you changed .env or config.yaml data)
6. Call the APIs from your postman and enjoy 😍 (Don't worry about jwt token at all, it handles using post request option in postman😉. you need just register and login)
6. APIs flow (using default postman bodies is recommended):
   - Register
   - Login
   - Create Provider
   - Create at least two addresses
   - Create order
   - Get Provider Reports



# AS a developer:
## Installation Steps for `pre-commit`
1. **Install `pre-commit`**

   If you don’t have `pre-commit` installed, you can install it using pip:

   ```bash
   sudo apt install python3-pip
   ```
   ```bash
   pip3 install pre-commit
   ```

   Alternatively, you can use brew on macOS:
      ```bash
       brew install pre-commit
      ```

2. **Install pre-commit hooks**

   ```bash
   pre-commit install
   ```
