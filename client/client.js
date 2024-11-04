import net from "net";
import fs from "fs";

const HOST = "localhost";
const PORT = 6969;

/**
 * @param {string} filePath
 **/
function client(filePath) {
    const client = new net.Socket();
    const fileType = filePath.split('.').pop()

    client.on('error', function(error) {
        console.error(`Could not connect to server on ${HOST}:${PORT}. Error: ${error.message}`);
    });

    client.connect(PORT, HOST, function() {
        console.log(`Connected to server on : ${HOST}:${PORT}`);

        fs.stat(filePath, (err, stats) => {
            if (err) {
                console.error(`Error getting file stats : ${err}`);
                client.destroy();
                return;
            }

            const fileSize = stats.size;
            const header = createHeader(fileType, fileSize);

            client.write(header, (err) => {
                if (err) {
                    console.error(`Error sending header : ${err}`);
                    client.destroy();
                    return;
                }
            });

            fs.readFile(filePath, (err, data) => {
                if (err) {
                    console.error(`Error reading file : ${err}`);
                    client.destroy();
                    return;
                }

                client.write(data, (err) => {
                    if (err) {
                        console.error(`Error sending file data : ${err}`);
                        client.destroy();
                    } else {
                        console.log("File data sent");
                    }
                });
            });
        });
    });

    client.on("close", function() {
        console.log("Connection closed.");
    });
}

/**
 * @param {string} fileType
 * @param {number} fileSize
 * @returns {Buffer}
 **/
function createHeader(fileType, fileSize) {
    const fileTypeBuffer = Buffer.from(fileType + '\0');  
    const fileSizeBuffer = Buffer.alloc(8);  
    fileSizeBuffer.writeBigUInt64LE(BigInt(fileSize), 0);  

    const header = Buffer.concat([Buffer.alloc(4), fileTypeBuffer, fileSizeBuffer]); 
    header.writeUInt32LE(header.length, 0); 
    return header;
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
