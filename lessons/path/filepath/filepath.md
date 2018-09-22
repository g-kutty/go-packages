---
layout: page
title: path/filepath
---
***

Package __filepath__ implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.

&nbsp;

## List Files
***

* The `filepath` package provides the handy `Walk` function. It automatically scans subdirectories too.

    ```go
    func Walk(root string, walkFn WalkFunc) error
    ```

* __filepath.Walk__ accepts a `string` pointing to the root folder, and a `func` with signature:

    ```go
    type WalkFunc func(path string, info os.FileInfo, err error) error
    ```

* WalkFunc function is called for each iteration of the folder scan.

* The `info` variable of type [os.FileInfo](https://golang.org/pkg/os/#FileInfo). This variable will provide many `information` about `files` like name, size, mode, last-modified time and underlying data source.

* We can avoid processing folder by adding:

    ```go
    if info.IsDir() {
        return nil
    }
    ```

* You can exclude files too based on their extension, by using [filepath.Ext](https://golang.org/pkg/path/filepath/#Ext).

    ```go
    if filepath.Ext(path) == ".dat" {
        return nil
    }
    ```

* We could also define the __WalkFunc__ in a separate closure. We just need to pass a pointer to files in visit:

    ```go
    func main()  {
        err := filepath.Walk(root, visit(&files))
        if err != nil {
            panic(err)
        }
    }

    func visit(files *[]string) filepath.WalkFunc {
        return func(path string, info os.FileInfo, err error) error {
            ...

            return nil
        }
    }
    ```