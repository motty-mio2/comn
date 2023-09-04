package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/koki-develop/go-fzf"
)

func DockerWrapper(pwd string) {
	cmd := exec.Command("ls")

	// cmd.Dir="~"
	cmd.Dir = pwd

	result, err := cmd.Output()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(string(result))
}

func ComposeWrapper(pwd string, name string, args string) {
	cmd := exec.Command("docker", "compose", args)
	cmd.Dir = pwd + "/" + name

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("標準出力をキャプチャできませんでした:", err)
		os.Exit(1)
	}

	// コマンドを開始
	if err := cmd.Start(); err != nil {
		fmt.Println("コマンドを開始できませんでした:", err)
		os.Exit(1)
	}

	// 標準出力を非同期で読み取り、出力を表示
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// コマンドの終了を待つ
	if err := cmd.Wait(); err != nil {
		fmt.Println("コマンドがエラーで終了しました:", err)
		os.Exit(1)
	}
}

func PickupComposeSpace(compose_dir string) string {
	usr, _ := user.Current()

	expandedPath := strings.Replace(compose_dir, "~", usr.HomeDir, 1)

	files, _ := os.ReadDir(expandedPath)

	var names []string
	for _, entry := range files {
		names = append(names, entry.Name())
	}

	return nameSelector(names)

}

func nameSelector(namespaces []string) string {
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(namespaces, func(i int) string { return namespaces[i] })
	if err != nil {
		log.Fatal(err)
	}

	return namespaces[idxs[0]]
}
