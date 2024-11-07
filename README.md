## Generating SSL ceritificates

This project requires SSL certificates ('server.crt' and 'server.key') for secure communication. To generate these files, follow the steps below:

1. **Install OpenSSL:**
    If you do not already have OpenSSL installed, [download and install OpenSSL](https://www.openssl.org/) for your operating system

2. **Run the Certificate Generation Script:**
    - On Linux/MacOS, open a terminal and run:
    ```bash
    ./generate_certificates.sh
    ```
    - On Windows, open Command Prompt and run:
    ```batch
    generate_certificates.bat
    ```

3. **What It Does:**
    The script will generate 'server.crt' (SSL certificate) and 'server.key' (private key) in the root directory of the project. These files will be used by the server to enable SSL/TLS encryption.

4. **Next Steps**
    After generating the certificates, you can start the server as usual, and the files generated will be used for secure communication.