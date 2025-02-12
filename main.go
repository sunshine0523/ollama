package main

import (
	"C"
	"context"

	"github.com/ollama/ollama/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.NewCLI().ExecuteContext(context.Background()))
}

// func RunServe(argc C.int, argv **C.char) C.int {
// 	go func() {
// 		os.Args = make([]string, int(argc))
// 		for i := 0; i < int(argc); i++ {
// 			os.Args[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(i)*unsafe.Sizeof(uintptr(0)))))
// 		}
// 		filePath := filepath.Join(os.Args[0], "ollama.log")
// 		file, e := os.Create(filePath)
// 		if e != nil {
// 			slog.Info("main.go#RunServe()", "Create log file failed", e)
// 		}
// 		defer file.Close()
// 		logger := slog.New(slog.NewTextHandler(file, nil))
// 		slog.SetDefault(logger)

// 		ln, err := net.Listen("tcp", envconfig.Host().Host)
// 		if err != nil {
// 			slog.Info("main.go#RunServe()", "Run net.Listen() failed", err)
// 		}

// 		err = server.Serve(ln)
// 		if errors.Is(err, http.ErrServerClosed) {
// 			slog.Info("main.go#RunServe()", "Run server.Serve() failed", err)
// 		}
// 	}()
// 	return 0
// }
