package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/motty-mio2/comn/src/functions"
)

func ComposeWrapper(backend string, pwd string, name string, args string) {
	cmd := exec.Command(backend, "compose", args)
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
	files, _ := os.ReadDir(functions.ExpandHome(compose_dir))

	var names []string
	for _, entry := range files {
		names = append(names, entry.Name())
	}

	return functions.FzfSelector(names)
}
