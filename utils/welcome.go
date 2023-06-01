package utils

import "fmt"

func Welcome() {
	fmt.Print("Welcome back!👋\n")
	PrintUsage()
}

func PrintUsage() {
	fmt.Print("You can use the following methods: [Create 💡| List 📓 | Un/Check ✅ | Delete 🗑️ ]\n")

}
