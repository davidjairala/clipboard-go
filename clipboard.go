package clipboard

import (
  "os/exec"
)

func getClipboard() (string, error) {
  pasteCommand := exec.Command("pbpaste")
  pasteOutput, err := pasteCommand.Output()

  if(err != nil) {
    return "", err
  }

  return string(pasteOutput), nil
}
