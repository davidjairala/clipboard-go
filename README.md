# clipboard-go
clipboard access for golang

# Requirements

[pngpaste](https://github.com/jcsalterego/pngpaste) is required to paste screenshots.

```bash
brew install pngpaste
```

# Usage

```go
import (
  "github.com/davidjairala/clipboard-go"
)

func main() {
  resultFile, err := clipboard.GetClipboard()
  fmt.Println(resultFile.Filename)
}
```

The file type will be either `txt` or `png` depending on what you had in your clipboard.
