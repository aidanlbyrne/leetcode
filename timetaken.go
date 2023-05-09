func timeTaken() func() {
    start := time.Now()

    return func() {
        elapsed := time.Since(start)
        fmt.Println(elapsed)
    }
}

defer timeTaken()()