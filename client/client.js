import net from "net";
import fs from "fs";
import path from "path";

const host = "localhost";
const port = 6969;

const extension_len = 2;
const file_name_len = 2;
const data_len = 4;

function client(filePath) {
    const client = new net.Socket();

    client.on('error', function(error) {
        console.error(`Could not connect to server on ${host}:${port}. Error: ${error.message}`);
        client.end();
    });

    client.connect(port, host, function() {
        console.log(`Connected to server on: ${host}:${port}`);

        fs.readFile(filePath, (err, data) => {
            if (err) {
                console.error(`Error reading file: ${err}`);
                client.destroy();
                return;
            }

            const fileName = path.basename(filePath, path.extname(filePath)); 
            const fileExtension = path.extname(filePath).slice(1); 

            const dataSize = data.length; 
            const fName = Buffer.byteLength(fileName);
            const fExt = Buffer.byteLength(fileExtension);

            const totalHeaderLen = extension_len + fExt + file_name_len + fName + data_len;
            const header = Buffer.alloc(totalHeaderLen);

            header.writeUInt16BE(fExt, 0);
            header.write(fileExtension, extension_len);

            header.writeUInt16BE(fName, extension_len + fExt);
            header.write(fileName, extension_len + fExt + file_name_len);

            header.writeUInt32BE(dataSize, extension_len + fExt + file_name_len + fName);

            console.log(`Header : { fExt: ${fExt}, fileExtension: ${fileExtension}, fNameL ${fName}, fileName: ${fileName}, dataSize: ${dataSize} }`);

            client.write(header); 
            client.write(data); 
        });
    });

    client.on("close", function() {
        console.log("Connection closed.");
    });
}

function main() {
    const filePath = process.argv.slice(2)[0];

    if (!filePath) {
        console.error("You must provide a file path as an argument.");
        process.exit(1);
    }

    client(filePath);
}

main();
