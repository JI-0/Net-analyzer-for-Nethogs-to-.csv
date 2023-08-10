package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const original_file_name = "webrtc" //Name of origin without .txt
const sum = true                    //Sum of all network traffic

const includes = "64.20.34.178:8088" //Source of network traffic
const cleanup = true                 //Clean data for postprocessing

func main() {
	original_file, original_file_err := os.Open(original_file_name + ".txt")
	if original_file_err != nil {
		log.Fatal(original_file_err)
	}
	defer original_file.Close()
	new_name := original_file_name + ".csv"
	if sum {
		new_name = "SUM_" + new_name
	} else if cleanup {
		new_name = "CLEAN_" + new_name
	}
	new_file, new_file_err := os.OpenFile(new_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if new_file_err != nil {
		log.Fatal(new_file_err)
	}

	scanner := bufio.NewScanner(original_file)
	writer := bufio.NewWriter(new_file)

	up := 0.0
	down := 0.0
	for scanner.Scan() {
		if sum {
			if strings.Contains(scanner.Text(), "Refreshing:") {
				l := strconv.FormatFloat(up, 'g', 5, 64)

				r := strconv.FormatFloat(down, 'g', 5, 64)

				writer.WriteString(l + "," + r + "\n")
				up = 0.0
				down = 0.0
			} else if scanner.Text() == "" {
				continue
			} else {
				new_line_no_tab := strings.ReplaceAll(scanner.Text(), "\t", ",")
				new_line_no_space := strings.ReplaceAll(new_line_no_tab, " ", ",")

				parts := strings.Split(new_line_no_space, ",")

				if len(parts) < 3 {
					continue
				}

				r, err := strconv.ParseFloat(parts[len(parts)-1], 64)
				if err != nil {
					continue
				}

				down += r

				l, err := strconv.ParseFloat(parts[len(parts)-2], 64)
				if err != nil {
					continue
				}

				up += l
			}
		} else {
			if strings.Contains(scanner.Text(), includes) {
				new_line_no_tab := strings.ReplaceAll(scanner.Text(), "\t", ",")
				new_line_no_space := strings.ReplaceAll(new_line_no_tab, " ", ",")
				if !cleanup {
					writer.WriteString(new_line_no_space + "\n")
					continue
				}
				parts := strings.Split(new_line_no_space, ",")
				writer.WriteString(parts[len(parts)-2] + "," + parts[len(parts)-1] + "\n")
			}
		}
	}

	writer.Flush()
	new_file.Close()
}
