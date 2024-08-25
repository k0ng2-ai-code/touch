# touch

This Go project is a simple implementation of the Unix `touch` command. It allows you to create files and update their access and modification times. The program supports various options to set specific times, reference an existing file's timestamps, or prevent file creation if it doesn't exist.

## Features
- **Create Files:** Automatically create a file if it does not exist, unless suppressed with the `-c` flag.
- **Modify Timestamps:** Update the access and modification times of existing files.
- **Custom Datetime Parsing:** Supports various datetime formats, as well as Unix timestamps.
- **Reference File:** Use another file's timestamps as a reference with the `-r` flag.
- **Selective Time Update:** Update only the access time (`-a`) or modification time (`-m`).
- **Suppress Creation:** Use the `-c` flag to suppress the creation of a file if it doesn't exist.

## Installation
1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/touch-utility.git
   cd touch-utility
   ```

2. **Build the project:**
   ```sh
   go build -o touch
   ```

3. **(Optional) Install the utility globally:**
   ```sh
   sudo mv touch /usr/local/bin/
   ```

## Usage
To use the `touch` utility, you can run the following command:

```sh
./touch [OPTION]... FILE...
```

### Options:
- `-a` : Change only the access time.
- `-c` : Do not create any files if they do not exist.
- `-d` : Set the access and modification times using the specified string (supports multiple datetime formats).
- `-m` : Change only the modification time.
- `-r` : Use this file's times instead of the current time.

### Examples:
- **Create a new file or update its timestamp:**
  ```sh
  ./touch filename.txt
  ```

- **Set a specific datetime as the access and modification time:**
  ```sh
  ./touch -d "2024-08-25T12:34:56" filename.txt
  ```

- **Update only the access time:**
  ```sh
  ./touch -a filename.txt
  ```

- **Use the timestamps from another file:**
  ```sh
  ./touch -r reference.txt filename.txt
  ```

- **Suppress file creation if it doesn’t exist:**
  ```sh
  ./touch -c filename.txt
  ```
