# 🗳️ Hyperledger Fabric Voting System Chaincode

A sample Hyperledger Fabric smart contract written in Go for implementing a decentralized voting system. This project demonstrates how to create elections, register voters, cast votes, and retrieve election results — all while using MSP-based role control and Go generics.

---

## 📦 Features

- ✅ Register voters (with eligibility check)
- ✅ Create elections with multiple candidates
- ✅ Cast vote securely (one vote per voter)
- ✅ View live vote tally
- ✅ Retrieve final election result (after end time)
- ✅ Uses CouchDB for rich queries
- ✅ Uses Go generics for reusable state fetching

---

## 🧱 Prerequisites

Make sure you have the following installed:

- Docker & Docker Compose
- Go (>= 1.18)
- Hyperledger Fabric binaries (v2.5+ recommended)
- Fabric test network (`fabric-samples/test-network`)

---

## 🚀 Setup Instructions

1. **Clone Fabric samples repo and navigate:**

   ```bash
   git clone https://github.com/hyperledger/fabric-samples.git
   cd fabric-samples/test-network
