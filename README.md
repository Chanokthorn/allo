# allo

allo is a versatile photography-based tool designed to organize your files based on their creation dates. It supports various allocation steps, including sorting by year, month, day, or a combination of these. Additionally, it offers specialized handling for raw and JPEG files, making it an ideal solution for photographers and anyone looking to bring order to their file collections.

## Features

- **Flexible Date Allocation**: Allocate files into directories based on year, month, and day of creation.
- **Support for Images**: Special handling for raw and JPEG files to ensure your photos are neatly organized.
- **Customizable Steps**: Choose specific allocation steps to suit your organizational needs.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. Allocate-Date is built with Go, so you'll need it to compile and run the tool.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/https://github.com/Chanokthorn/allo/allocate-date.git
   ```
2. Navigate to the project directory:
   ```bash
   cd allocate-date
   ```
3. Build the project:
   ```bash
   go build -o allocate-date
   ```

### Usage

To use Allocate-Date, you'll need to specify the steps for allocation and the path to the directory containing the files you want to organize.

```bash
./allocate-date -steps "y,m,d,raw-jpeg" -path "/path/to/directory"
```

The command will create directories in this order:

```bash
.
├── 2024
│   ├── 2024-05
│   │   ├── 2024-05-27
│   │   │   ├── jpeg
│   │   │   │   ├── DSCF9184.JPG
│   │   │   │   └── DSCF9193.JPG
│   │   │   └── raw
│   │   │       ├── DSCF9184.RAF
│   │   │       └── DSCF9193.RAF
│   │   └── 2024-05-28
│   │       └── jpeg
│   │           ├── DSCF9194.JPG
│   │           └── DSCF9206.RAF
│   └── 2024-06
│       └── 2024-06-23
│           ├── jpeg
│           │   ├── DSCF9279.JPG
│           │   └── DSCF9280.JPG
│           └── raw
│               ├── DSCF9279.RAF
│               ├── DSCF9280.RAF
│               └── DSCF9281.RAF
├── DSCF9207.MOV
└── DSCF9241.MOV
```

#### Available Allocation Steps

- `create-date`: Allocate files based on year, month, and day.
- `y`: Allocate files based on the year of creation.
- `m`: Allocate files based on the month of creation.
- `d`: Allocate files based on the day of creation.
- `raw-jpeg`: Special handling for raw and JPEG files.

#### Supported File Types
We use file signatures for file type detection. The service will currently detect these file types

- JPEG
- PNG
- RAF

## Contributing

Contributions are welcome! If you have a feature request, bug report, or a suggestion, please open an issue in the repository.

## License

Allocate-Date is open-source software licensed under the MIT license. See the LICENSE file for more details.
