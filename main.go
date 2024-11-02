package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	numBuckets = 5
)

func addUserDetails(userId int, userIdMap map[int][]int) {
	shardIndex := getShardIndex(userId)
	if _, exists := userIdMap[shardIndex]; !exists {
		userIdMap[shardIndex] = []int{}
	}
	userIdMap[shardIndex] = append(userIdMap[shardIndex], userId)
}

func contains(slice []int, target int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func getUserDetails(wg *sync.WaitGroup, userId int, userIdMap map[int][]int) {
	defer wg.Done()
	shardIndex := getShardIndex(userId)
	validStr := "VALID"
	if !contains(userIdMap[shardIndex], userId) {
		validStr = "INVALID"
	}
	fmt.Printf("UserId: %v => Invoking shard %d :: %v\n", userId, shardIndex, validStr)
}

func getShardIndex(userId int) int {
	userIdStr := strconv.Itoa(userId)
	hasher := sha256.New()
	hasher.Write([]byte(userIdStr))
	hashBytes := hasher.Sum(nil)
	hashInt := new(big.Int).SetBytes(hashBytes)
	bucket := new(big.Int).Mod(hashInt, big.NewInt(int64(numBuckets)))
	return int(bucket.Int64())
}

func main() {

	userIdMap := make(map[int][]int)
	fmt.Println("Adding Users to DB...")
	for i := 1; i <= 100; i++ {
		addUserDetails(i, userIdMap)
	}
	fmt.Println("Adding Users completed: ", userIdMap)

	reqSize := 20
	var wg sync.WaitGroup
	rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Querying users from the DB...")
	for i := 1; i <= reqSize; i++ {
		userId := rand.Intn(100)

		wg.Add(1)
		go getUserDetails(&wg, userId, userIdMap)
	}

	// wait for all the go-routines to complete
	wg.Wait()
}
