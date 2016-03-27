package clipboard

import (
  "io/ioutil"
  "os"
  "os/exec"
  "path"
  "time"
)

type ResultFile struct {
  Ext, Data string
  Err error
}

const TimeFormat = "20060102150405"

// Return 'extension', 'data', error
//    extension: png, txt
func GetClipboard() (file ResultFile) {
  // Try to fetch image
  data, err := GetClipboardImage()

  // nothing went wrong with png, lets return that
  if(err == nil) {
    return ResultFile{"png", data, err}
  }

  // Try to fetch text if anything went wrong with png
  data, err = GetClipboardText()
  return ResultFile{"txt", data, err}
}

func GetClipboardText() (string, error) {
  pasteCommand := exec.Command("pbpaste")
  pasteOutput, err := pasteCommand.Output()

  if(err != nil) {
    return "", err
  }

  return string(pasteOutput), nil
}

func GetClipboardImage() (string, error) {
  imageFilename := filename("png")
  imageFilename = path.Join("/tmp", imageFilename)
  pasteCommand := exec.Command("pngpaste", imageFilename)
  _, err := pasteCommand.Output()

  if(err != nil) {
    return "", err
  }

  data, dataErr := ioutil.ReadFile(imageFilename)

  if(err != nil) {
    return "", err
  }

  dataErr = os.Remove(imageFilename)

  if(dataErr != nil) {
    return "", dataErr
  }

  return string(data), nil
}

func filename(ext string) (string) {
  epoch := time.Now().Format("20060102150405")
  return epoch + "." + ext
}
