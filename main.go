package main

import (
	"errors"
	"fmt"
	"os"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		//TIP You can try debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>. To start your debugging session,
		// right-click your code in the editor and select the <b>Debug</b> option.
		fmt.Println("i =", 100/i)
	}

	//errorHandle("sample.txt")
	compare()

	printErr("Error log test")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

func errorHandle(fname string) error {
	// エラーは状態であり表現
	// Goにおいてはエラーをerrorインターフェースの実装された値として返却して、それを実装者が愚直にハンドリングすることをエラー処理と見なしているだけで言語としてエラー処理の仕組みがある訳ではない
	f, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("can not create file %s: %w", fname, err)
	}

	defer f.Close()

	_, err = f.WriteAt([]byte("sample"), 0)
	if err != nil {
		return fmt.Errorf("can not write file %s: %w", fname, err)
	}

	return nil
}

// ErrNotFound センチネルエラー
// 慣習的に命名はErrXXX
var ErrNotFound = errors.New("not found")

// エラーのラップ
// errors.Unwrap()でラップされた中身を取り出せる
func wrapError(err error) error {
	return fmt.Errorf("%w", err)
}

// FormatError 独自エラー
// 慣習的に命名はXXXError
type FormatError struct {
	Exp string
	Err error
}

func (e FormatError) Error() string {
	return fmt.Sprintf("%s: %s", e.Exp, e.Err)
}

func (e FormatError) Unwrap() error {
	return e.Err
}

// 複数エラーの集約
func aggregateError(errs []error) error {
	// Since Go 1.20
	return errors.Join(errs...)
}
