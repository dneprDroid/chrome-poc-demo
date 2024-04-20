module chrome-poc

go 1.14

require (
    boringssl.googlesource.com/boringssl v0.0.0-20221208190510-1ccef4908ce0
)

replace boringssl.googlesource.com/boringssl => github.com/google/boringssl v0.0.0-20221208190510-1ccef4908ce0
