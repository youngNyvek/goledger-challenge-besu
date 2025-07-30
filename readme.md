# SmartContract API

A Go-based REST API to interact with a Hyperledger Besu private blockchain network. Features:

- **Set** a value in the `SimpleStorage` smart contract.
- **Get** the current on-chain value.
- **Sync** the on-chain value to a PostgreSQL database.
- **Check** if the database value matches the on-chain value.

The project follows an **MVC-inspired** structure with clear separation of concerns:

- **Controller**: Handles HTTP requests (Gin).
- **Service**: Business logic and blockchain interactions.
- **Repository**: Database access.
- **Infrastructure**: Clients for Ethereum RPC and PostgreSQL connection.

---

## 📁 Project Structure

```plaintext
.
├── app/                   # Go application
│   ├── main.go            # Entry point
│   ├── go.mod             # Module definition
│   ├── abi/               # Contract ABI files (e.g., SimpleStorage.json)
│   ├── internal/          # Application layers
│   │   ├── controller/    # HTTP handlers
│   │   ├── service/       # Business logic
│   │   ├── repository/    # Data access
│   │   └── infra/         # Infrastructure clients
│   │       ├── database/  # PostgreSQL connection logic
│   │       └── ethereum/  # Ethereum client & contract binding
│   └── docker/            # Docker config and init scripts
│       ├── Dockerfile
│       ├── docker-compose.yml
│       └── db-init/
│           └── init.sql
├── besu/                  # Besu network scripts
│   └── startDev.sh        # Launches local QBFT network
├── run-all.sh             # Orchestrator script (Besu, DB, API)
├── example.http           # HTTP request examples
└── README.md              # This file
```

---

## 🏗 Architecture Overview

1. **Controller** (HTTP Layer)
   - Defines REST endpoints using Gin.
2. **Service** (Business Logic)
   - Interacts with the Besu network via generated contract bindings.
3. **Repository** (Data Access)
   - Persists contract values to PostgreSQL.
4. **Infrastructure**
   - Ethereum client setup and PostgreSQL connection with retry logic.

This layered approach ensures single responsibility per component and improves maintainability.

---

## 🚀 Getting Started

### Prerequisites

- **Go** 1.22+
- **Node.js** 14+ (includes NPM/NPX)
- **Docker** & **Docker Compose**
- **Hardhat** (Ethereum development)
- **Hyperledger Besu** binary installed locally

> **Hardhat must be installed *****after***** cloning the repo**, inside the `besu/` directory:
>
> ```bash
> cd besu
> npm install --save-dev hardhat
> ```

### 1. Clone the repository

```bash
git clone https://github.com/your-user/goledger-challenge-besu-api.git
cd goledger-challenge-besu-api
```

### 2. Prepare the Go application

After the Besu network is running and PostgreSQL is up, navigate to the `app/` folder and initialize dependencies before running:

```bash
cd app
# Download Go module dependencies
go mod tidy
```

### 3. Run the all-in-one script

The `run-all.sh` script performs the following:

1. Starts the Besu network (`besu/startDev.sh`)
2. Launches PostgreSQL (`app/docker/docker-compose.yml`)
3. Waits until the database is ready
4. Initializes Go modules and builds the application (`go mod tidy` + `go run`)
5. Runs the Go API

Make sure `run-all.sh` is executable and then run:

```bash
chmod +x run-all.sh
./run-all.sh
```