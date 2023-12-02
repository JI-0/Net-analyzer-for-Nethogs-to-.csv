# Net analyzer for Nethogs to .csv
A script tool for changing text output from Nethogs into .csv files for further processing and visualization in tools like LibreOffice Spreadsheet and Excel.

This script is used to process files created with [Nethogs](https://github.com/raboof/nethogs#readme) using [sed, a stream writer](https://www.gnu.org/software/sed/manual/sed.html). 

Essentially sed is a writing tool that can periodically take a snapshot of the Nethogs dashboard and stream it to a text file, which can then, using this script be converted into a .csv file.

## Instructions
To generate the original text file to be processed:
1. Install Nethogs and sed, and clone this repository.
2. Run the command `sudo nethogs -t -d 1 | sed 's/[^[:print:][:cntrl:]]//g' > timestamps.txt`, where `-d n` n is the duration between snapshots in seconds (time of change).

To then process it to a .csv file:
1. Place the output file (`timestamps.txt`) in the directory of this cloned repository.
2. Modify the variables in the script (`original_file_name`, `sum`, `includes`, `cleanup`) as desired.
3. Run the script with `go run .`

## Variables
| Variables             | Explanation                                                           |
| -------------         |:---------------------------------------------------------------------:|
| original_file_name    | Name of the original file without .txt                                |
| sum                   | Sum of all network traffic (for each time duration)                   |
| includes              | Source of network traffic (only get speeds from lines that include)   |
| cleanup               | Clean data for postprocessing (modify to only have up/down speeds)    |