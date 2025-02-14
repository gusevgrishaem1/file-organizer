# File Organizer

A simple CLI tool written in Go that organizes files in a directory by moving them into categorized subfolders based on their extensions.

## Features
- Automatically sorts files into predefined categories (Images, Videos, Documents, Archives, etc.).
- Customizable sorting rules via configuration.
- CLI-based with simple usage.
- Efficient and lightweight.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/file-organizer.git
   cd file-organizer
   ```
2. Initialize the Go module:
   ```sh
   go mod init file-organizer
   go mod tidy
   ```
3. Build the executable:
   ```sh
   go build -o file-organizer main.go
   ```

## Usage

Run the program with a directory path to organize:
```sh
./file-organizer --path ~/Downloads
```

### Example Output
```
🔄 Organizing files in: ~/Downloads
📂 Moved image1.jpg → Images
📂 Moved video.mp4 → Videos
📂 Moved document.pdf → Documents
✅ Organization complete!
```

## Configuration
You can define custom sorting rules in a `config.json` file:
```json
{
  "Images": [".jpg", ".jpeg", ".png", ".gif"],
  "Videos": [".mp4", ".mkv", ".mov"],
  "Documents": [".pdf", ".docx", ".txt"],
  "Archives": [".zip", ".rar", ".tar.gz"]
}
```

## Future Improvements
- Implement a **watch mode** to automatically organize new files.
- Add **GUI support** for easier interaction.
- Support **multi-threading** for faster processing.

## License
This project is open-source and available under the MIT License.

