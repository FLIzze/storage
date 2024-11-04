import net from "net";
import fs from "fs";

const host = "localhost";
const port = 6969;

/**
 * @param {string} filePath
 **/
function client(filePath) {
    const client = new net.Socket();

    client.on('error', function(error) {
        console.error(`Could not connect to server on ${host}:${port}. Error: ${error.message}`);
        client.end();
    });

    client.connect(port, host, function() {
        console.log(`Connected to server on : ${host}:${port}`);

        fs.readFile(filePath, (err, data) => {
            if (err) {
                console.error(`Error reading file ${err}`);
                client.destroy();
                return;
            }

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
