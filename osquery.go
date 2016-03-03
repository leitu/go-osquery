package main

import (
  "os/exec"
  "log"
  "bufio"
)


func query(sql string)string{

  bs, err := exec.Command("osqueryi", sql, "--json").Output()

  if err != nil {
    panic(err)
  }

  return string(bs)
}

func getOnetable(table string)string{
  sql := "select * from " + table
  query(sql)
}

func getAlltables()[]string{

  cmd := exec.Command("osqueryi", ".table")
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }

  if err := cmd.Start(); err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(stdout)
  scanner.Split(bufio.ScanWords)
  var lines []string

  for scanner.Scan() {
    line := scanner.Text()
    if len(line) == 2{
      continue
    }
    lines = append(lines,line)
  }

  return lines

}
