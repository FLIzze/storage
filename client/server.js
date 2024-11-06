import express from 'express';
import multer from 'multer';
import { uploadFile } from './client.js';

const port = 6969;

const app = express();

const storage = multer.memoryStorage();
const upload = multer({ storage: storage });

app.post('/upload', upload.single('file'), (req, res) => {
    const fileBuffer = req.file.buffer;
    const fileName = req.file.originalname;

    uploadFile(fileBuffer, fileName);
    res.status(200).send("Sent file to 'uploadFile'.");
});

app.listen(port, () => {
    console.log(`Server running on port: ${port}`);
});
