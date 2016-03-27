package clipboard

import (
  "io/ioutil"
  "os"
  "os/exec"
  "path"
  "time"
)

type ResultFile struct {
  Filename, Data string
}

const TimeFormat = "20060102150405"

// Return 'extension', 'data', error
//    extension: png, txt
func GetClipboard() (ResultFile, error) {
  // Try to fetch image
  resultFile, err := GetClipboardImage()

  // nothing went wrong with png, lets return that
  if(err == nil) {
    return resultFile, nil
  }

  // Try to fetch text if anything went wrong with png
  resultFile, err = GetClipboardText()
  return resultFile, err
}

func GetClipboardText() (ResultFile, error) {
  imageFilename := filename("txt")
  pasteCommand := exec.Command("pbpaste")
  pasteOutput, err := pasteCommand.Output()

  if(err != nil) {
    return ResultFile{}, err
  }

  return ResultFile{imageFilename, string(pasteOutput)}, nil
}

func GetClipboardImage() (ResultFile, error) {
  imageFilename := filename("png")
  pasteCommand := exec.Command("pngpaste", imageFilename)
  _, err := pasteCommand.Output()

  if(err != nil) {
    return ResultFile{}, err
  }

  data, dataErr := ioutil.ReadFile(imageFilename)

  if(dataErr != nil) {
    return ResultFile{}, dataErr
  }

  dataErr = os.Remove(imageFilename)

  if(dataErr != nil) {
    return ResultFile{}, dataErr
  }

  return ResultFile{imageFilename, string(data)}, nil
}

func filename(ext string) (string) {
  epoch := time.Now().Format("20060102150405")
  fileName := epoch + "." + ext
  return path.Join("/tmp", fileName)
}
