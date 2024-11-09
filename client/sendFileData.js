import net from "net";

const host = "192.168.1.25";
const port = 6969;

const extension_len = 2;
const file_name_len = 2;
const data_len = 4;

/**
 * @param {byte[]} fileData
 * @param {string} fileExtName
 **/
export function uploadFile(fileData, fileExtName) {
    const client = new net.Socket();

    client.on('error', function(error) {
        console.error(`Could not connect to server on ${host}:${port}. Error: ${error.message}`);
        client.end();
    });

    client.connect(port, host, function() {
        if (!fileData || !fileExtName) {
            console.error("Error reading file.");
            client.destroy();
        }

        const [fileName, fileExtension] = fileExtName.split(".");

        const dataSize = fileData.length; 
        const fName = Buffer.byteLength(fileName);
        const fExt = Buffer.byteLength(fileExtension);

        const totalHeaderLen = extension_len + fExt + file_name_len + fName + data_len;
        const header = Buffer.alloc(totalHeaderLen);

        header.writeUInt16BE(fExt, 0);
        header.write(fileExtension, extension_len);

        header.writeUInt16BE(fName, extension_len + fExt);
        header.write(fileName, extension_len + fExt + file_name_len);

        header.writeUInt32BE(dataSize, extension_len + fExt + file_name_len + fName);

        client.write(header); 
        client.write(fileData); 

        console.log(`Sent ${fileName}.`);
    });
}
