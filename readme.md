# DB Sharing Using Go

This Go program simulates a basic sharded user database where user data is distributed across 5 shards (buckets) using a **SHA-256** hashing mechanism. Users are added to the database and queried concurrently using **goroutines**.

## Features
- **User Sharding**: User IDs are distributed across 5 shards using a hash-based partitioning scheme.
- **Concurrent User Queries**: Queries are handled concurrently using Go's `sync.WaitGroup` and goroutines.
- **Random User Querying**: A random set of users are queried to check whether they exist in the database.

## Project Structure

The project consists of the following major functions:

- **`addUserDetails(userId int, userIdMap map[int][]int)`**: Adds a user ID to the appropriate shard in the user map based on a hash of the user ID.

- **`getUserDetails(wg *sync.WaitGroup, userId int, userIdMap map[int][]int)`**: Retrieves a user ID from the user map, checking whether the user exists in the correct shard. This function is designed to be run concurrently in goroutines.

- **`getShardIndex(userId int) int`**: Uses SHA-256 to hash the user ID and assigns it to a shard (bucket) based on the hash value.

- **`contains(slice []int, target int) bool`**: Helper function to check if a user ID exists in a slice.

## How It Works

1. **Adding Users**:
   The program first adds 100 user IDs (from 1 to 100) to a sharded database. User IDs are hashed using SHA-256, and the resulting hash value is used to assign the user ID to one of the 5 shards.

2. **Querying Users**:
   After adding the users, the program randomly queries 20 user IDs to check whether they exist in the database. These queries are executed concurrently using goroutines, and the result (whether the user exists or not) is printed to the console.

## Installation and Running the Program

### Prerequisites
- Install Go on your machine: https://golang.org/doc/install

### Running the Program

1. Clone the repository or copy the code to your local machine.

2. Create a Go file (e.g., `main.go`) and paste the provided code.

3. Open a terminal in the directory where your Go file is located.

4. Run the program using the following command:

   ```bash
   go run main.go
   ```

## Code Explanation

### Constants
```go
const numBuckets = 5
```
Defines the number of shards (buckets) where users will be distributed.

### Main Logic

1. **Sharding Users**:
   - Users are added to the map `userIdMap` based on their shard index, calculated by `getShardIndex(userId)`.

2. **Concurrency with Goroutines**:
   - The program uses `sync.WaitGroup` to manage concurrency. Each user query is executed in a separate goroutine, and `wg.Wait()` ensures that the main program waits until all goroutines have completed.

3. **SHA-256 Hashing**:
   - The program hashes each `userId` to determine the shard where it will be stored. This ensures even distribution across the shards.

### Functions

- **`addUserDetails()`**: Adds users to the correct shard based on the hash.
- **`getShardIndex()`**: Computes the hash and assigns the user to a shard.
- **`getUserDetails()`**: Retrieves a user from the map and prints if the user is valid (exists in the shard).
- **`contains()`**: Helper function to check if a user ID is in a slice.
