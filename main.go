package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "net/http"

    "github.com/esimov/caire"
    "github.com/pkg/errors"
)

const(
    IMAGE_URL string = "https://bit.ly/2QGPDkr"
)

func ContentAwareResize(url string) ([]byte, error) {
    fmt.Printf("Download starting for url %s\n", url)
    response, err := http.Get(url)
    if err != nil {
        return nil, errors.Wrap(err, "Failed to read the image")
    }
    defer response.Body.Close()
    converted := &bytes.Buffer{}
    fmt.Printf("Download complete %s", url)

    shrinkFactor := 30
    fmt.Printf("Resize in progress %s, shrinking width by %d percent...\n", url, shrinkFactor)
    p := &caire.Processor{
        NewWidth: shrinkFactor,
        Percentage: true,
    }

    err = p.Process(response.Body, converted)
    if err != nil {
        return nil, errors.Wrapp(err, "Failed to apply seam carving to the image")
    }
    fmt.Printf("Seam carving completed for %s\n", url)
    return converted.Bytes(), nil
}


