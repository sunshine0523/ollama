export ANDROID_HOME=/home/kindbrave/Android/Sdk
export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/28.0.12916984
export SYSROOT=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/linux-x86_64/sysroot
export GOARCH=amd64
export GOOS=android
export CGO_ENABLED=1
export CGO_CFLAGS="--sysroot=$SYSROOT"
export CGO_CXXFLAGS="--sysroot=$SYSROOT"
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android23-clang
export CXX=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android23-clang++
# go build -buildmode=c-shared -o output/android/arm64-v8a/libollama.so main.go
go build -o output/android/arm64-v8a/ollama main.go