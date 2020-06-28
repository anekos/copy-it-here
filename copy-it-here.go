package main

import (
  "bufio"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "time"
)


func Die(err error) {
  fmt.Fprintln(os.Stderr, "Error: ", err)
  fmt.Fprintln(os.Stderr, "Press enter to exit")
  bufio.NewScanner(os.Stdin).Scan()
  os.Exit(1)
}

func copyFile(from string, to string) {
  content, err := ioutil.ReadFile(from)
  if err != nil {
    Die(err)
  }

  err = ioutil.WriteFile(to, content, 0644)
  if err != nil {
    Die(err)
  }
}

func AppDir() string {
  return filepath.Dir(os.Args[0])
}

func ConfigPath() string {
  return filepath.Join(AppDir(), "config.txt")
}

func WriteTarget(path string) {
  err := ioutil.WriteFile(ConfigPath(), []byte(path), 0644)
  if err != nil {
    Die(err)
  }
}

func ReadTarget() string {
  configPath := ConfigPath()

  _, err := os.Stat(configPath)
  if err == nil {
    content, err := ioutil.ReadFile(configPath)
    if err != nil {
      Die(err)
    }
    return string(content)
  }

  return filepath.Dir(os.Args[0])
}

func main() {
  flag.Parse()
  if flag.NArg() == 0 {
    fmt.Printf("Not enough argument")
    os.Exit(1)
  }

  it := flag.Args()[0]

  if info, err := os.Stat(it); err == nil && info.IsDir() {
    fmt.Printf("Set target directory: %s\n", it)
    WriteTarget(it)
  } else {
    dest := ReadTarget()

    fmt.Printf("Copy %s to %s\n", it, dest)
    copyFile(it, filepath.Join(dest, time.Now().Format("20060102-150405") + filepath.Ext(it)))
    copyFile(it, filepath.Join(dest, "latest" + filepath.Ext(it)))
  }
}
