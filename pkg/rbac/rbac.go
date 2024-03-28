package rbac

// RoleBasedAccessControlService

func IntToBinary(n, length int) []int {
	binary := make([]int, length)
	// log.Printf("n:asdasdasdas %v", n)
	// log.Printf("length: %v", length)
	for i := 0; i < length; i++ {
		binary[i] = n % 2
		n /= 2
	}
	// Log binary
	// log.Printf("binary: %v", binary)
	return binary
}
