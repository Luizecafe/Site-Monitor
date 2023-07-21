# Site Monitor - GO Lang Project

Welcome to the Site Monitor project! This is my first software project developed in GO Lang, and I'm excited to introduce it to you.

## Description

The Site Monitor project is a powerful tool designed to monitor multiple websites by making HTTP requests and capturing important information. Its main function is to store logs of the queries made, allowing users to easily determine when a site was online or offline.

## Features

- Website monitoring: The program monitors multiple websites by sending HTTP requests at regular intervals.
- Logging: It captures and stores logs of the queries made, including the timestamp, website status, and response information.
- Offline/online status: Users can check the online and offline status of a specific website based on the stored logs.
- Easy configuration: The program reads the list of websites to monitor from the `sites.txt` file, making it easy to update and customize the monitored sites.
- Flexible monitoring settings: The time interval and frequency of monitoring are defined as constants in the code, allowing for customization based on specific needs.

## Installation

1. Clone the repository: `git clone [https://github.com/Luizecafe/Site-Monitor.git]`
2. Navigate to the project directory: `cd SiteMonitor`
3. Build the project: `go build`
4. Run the program: `./SiteMonitor`

## Usage

1. Open the `sites.txt` file and add the websites you want to monitor, each on a separate line.
2. Configure the monitoring time and frequency by adjusting the constants defined in the code.
3. Run the program.
4. The program will start monitoring the websites and storing logs.
5. Use the provided functions and commands to retrieve the website status and logs.

**Important Notes:**

1. Time and Frequency Configuration: Take caution when modifying the monitoring time and frequency constants defined in the code. Altering these values without careful consideration may affect the accuracy and performance of the monitoring process.
2. `sites.txt` File: Use the `sites.txt` file to add the websites you wish to monitor. Each website should be added on a separate line.

## Dependencies

The following packages were used in the development of this project:

- bufio
- fmt
- io
- io/ioutil
- net/http
- os
- strconv
- strings
- time

## Contribution

Contributions are welcome! If you would like to contribute to the Site Monitor project, please follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b my-feature-branch`
3. Make your changes and commit them: `git commit -am 'Add new feature'`
4. Push the branch to your forked repository: `git push origin my-feature-branch`
5. Submit a pull request, describing your changes in detail.


## Contact

If you have any questions or feedback, feel free to reach out to me at luizgmoreirahg@gmail.com.

Thank you for checking out the Site Monitor project! Happy monitoring!
