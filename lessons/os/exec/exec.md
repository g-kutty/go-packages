---
layout: page
title: os/exec
---
***

Package __exec__ runs external commands. It wraps `os.StartProcess` to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

&nbsp;

## Execute external command
***

* Easier and efficient way to external command which does not need standard output:

    ```go
    err := exec.Command("touch", file_name).Run()
    if err != nil {
        log.Printf("Something went wrong: %v", err)
    }
    ```

* If you want standard output use method `Output()` instead of `Run()`.