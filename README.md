# Flyer Splicer

`flyer-splicer` is a small command line tool for processing groups of digital signage video files.
It reads an input directory, copies and renames files into an output directory, and optionally pushes the final output into a mounted volume such as a USB drive.

It is built for use cases where a single promotional “flyer” video needs to be duplicated and renamed to match other files in a set. The program automates that workflow and provides clear terminal output with colored logging.

## Features

- Copies all non-flyer files from an input directory to an output directory.
- Detects a flyer file (any filename containing the substring flyer) and duplicates it using prefix patterns from the other files.
- Automatically generates new filenames like A.mp4, B2.mp4, etc.
- Displays clear colored logs, including errors, warnings, and success messages.
- Optional volume handling
  - Confirms with the user before erasing files in a target volume.
  - Copies the final output files into the volume.
- Includes utilities for cleaning directories, sorting files, and gathering filenames.

## How It Works

1. The program scans your input directory.
2. Every file that does not contain the word flyer is copied directly into the output directory.
  - While doing this, the program builds a list of prefixes from those filenames.
3. The single flyer file is copied repeatedly into the output directory using each prefix, generating a set of renamed flyer videos.
4. If you provide a volume path, the program optionally cleans the volume (after confirmation) and copies the final output files into it.

The result is a full set of properly structured promotional files ready for deployment or upload.

## Installation
Be sure to have Make installed.

1. Clone the repository:
```bash
git clone https://github.com/jaredthecomputerguy/flyer-splicer.git
cd flyer-splicer
```

2. Build the binary:
```bash
make
```

## Usage
```bash
flyer-splicer <input_dir> <output_dir> [volume_dir]
```

### Arguments

- `input_dir` - Directory containing the raw source files.
- `output_dir` - Directory where processed files will be written.
- `volume_dir` (optional) - A mounted USB or external drive where the final files should be copied.


### Example
```bash
flyer-splicer ./raw ./out /Volumes/USBDrive
```

This will:
1. Process and restructure the files from ./raw into ./out.
2. Ask for confirmation before erasing files in /Volumes/USBDrive.
3. Copy all processed files into the volume.

Input Directory should contain something like:
```bash
../
Apple.mp4
Banana.mp4
Banana2.mp4
Cherry.mp4
flyer.mp4
```

Output Directory (after processing):
```bash
../
Apple.mp4 # (copied from input)
A.mp4 # (flyer)
Banana.mp4
B.mp4 # (flyer)
Blueberry.mp4
Bl.mp4 # (flyer)
Cherry.mp4
C.mp4 # (flyer)
```

## Error Handling
All filesystem operations are validated. Any unrecoverable error triggers a logged message followed by program exit.

Directory cleanup operations collect errors, report them at the end, and continue processing other files.

## Project Structure
```bash
internal/
  log.go             Logging utilities
  files.go           Core processing, copying, volume handling
  file-manager.go    Directory scanning, cleaning, sorting
cmd/
  main.go            Program entry point
  args.go            CLI parsing and usage text
```

## License

MIT License. See LICENSE for details.
