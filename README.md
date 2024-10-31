# WC-tool

A command-line utility tool written in Go that provides word count functionalities similar to the Unix `wc` command. This tool can count words, lines, characters, and bytes in a given file or text input, and supports multiple command options for customized output.

## Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/wahlly/WC-tool.git
   cd WC-tool

2. **Ensure you have Golang installed, then build the project by running**
   ```bash
   go build -o wc

3. **Run the project(from the root directory)**
   ```bash
   go run .

## Usage

The wc tool can be executed with various options, each with its use. here is a list of all available options:
### Command Options
The following commands are available:

- -w : Outputs the number of words in the file
- -l : Outputs the number of lines in the file
- -c : Outputs the number of bytes in the file
- -m : Outputs the number of characters in the file

### Examples

1. **Character count**
   ```bash
   ccwc -c test.txt

2. **Line count**
   ```bash
   ccwc -l test.txt

3. **Word count**
   ```bash
   ccwc -w test.txt

4. **Character count**
   ```bash
   ccwc -m test.txt

   wc -m test.txt

The default option i.e. when no options are provided, is equivalent to -c, -l and -w options.
 ```bash
 ccwc test.txt
```
It is also possible to read from standard input, if no file is specified. An example is shown below. This reads the `test.txt` file in the stdin and then outputs the number of lines in it
```bash
 cat test.txt | ccwc -l
```

