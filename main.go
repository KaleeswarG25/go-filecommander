package main
import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main(){
	scan:= bufio.NewScanner(os.Stdin)
	fmt.Println("==================================================")
	fmt.Println("🚀 Welcome to GoFileCommander Pro v1.1")
	fmt.Println("   Advanced Low-Level Filesystem Forensics Tool   ")
	fmt.Println("==================================================")

	for{
		fmt.Println("\n📂 Main Menu:")
		fmt.Println("  1. List directory contents")
		fmt.Println("  2. Forensic Deep Inspect (OS & Storage Layers)")
		fmt.Println("  3. Exit")
		fmt.Print("👉 Select an option (1-3): ")

		if !scan.Scan(){
			break
		}

		choice:= strings.TrimSpace(scan.Text())
		switch choice{
		case 1:
			listfile()
		case 2:
			fmt.Print("🔍 Enter target path/filename: ")
			if scan.Scan(){
				inspectfile(strings.TrimSpace(scan.Text()))
			}
		case 3:
			fmt.Println("\n👋 Exiting GoFileCommander Pro. Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice. Select 1, 2, or 3.")
		}
	}
}
		
func listfile(){
	stFiles() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Println("\n📁 Directory Items:")
	for _, f := range files {
		tag := "📄 [FILE]"
		if f.IsDir() {
			tag = "📁 [DIR] "
		}
		fmt.Printf("  %s %s\n", tag, f.Name())
	}
}	

func inspectfile(path string){
	info,err:=os.Lstat(path)
	if os.IsNotExist(err){
		fmt.Println("❌ Error: Path does not exist.")
		return
	} else if err != nil {
		fmt.Printf("❌ Scan Error: %v\n", err)
		return
	}
	fmt.Println("\n==============================================")
	fmt.Printf("📊 CORE PROPERTIES FOR: %s\n", info.Name())
	fmt.Println("==============================================")
	fmt.Printf("  • Absolute Path:    %s\n", filepath.Clean(path))
	fmt.Printf("  • File Permissions: %s (Octal: %O)\n", info.Mode(), info.Mode().Perm())

	// 1. Run cross-platform deep OS metadata scan
	getPlatformMetadata(info)

	// 2. Run structural forensic magic-bytes scan
	declaredExt := filepath.Ext(path)
	VerifyFileIntegrity(path, declaredExt)
}
