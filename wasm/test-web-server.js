const http = require('http');
const fs = require('fs');
const zlib = require('zlib');
const path = require('path');

const PORT = 8080;

const server = http.createServer((req, res) => {
    const filePath = req.url === '/' ? '/test.html' : req.url;
    const fullPath = path.join(__dirname, filePath);

    fs.readFile(fullPath, (err, data) => {
        if (err) {
            res.writeHead(404);
            res.end('Not found');
            return;
        }

        const ext = path.extname(filePath);
        const contentType = {
            '.html': 'text/html',
            '.js': 'application/javascript',
            '.wasm': 'application/wasm',
        }[ext] || 'text/plain';

        if (ext === '.wasm' && req.headers['accept-encoding']?.includes('gzip')) {
            zlib.gzip(data, (err, compressed) => {
                if (!err) {
                    const formatBytes = (bytes) => {
                        const mb = (bytes / 1024 / 1024).toFixed(2);
                        return `${mb} MB (${bytes.toLocaleString()} bytes)`;
                    };
                    console.log(`Serving ${filePath}: ${formatBytes(data.length)} → ${formatBytes(compressed.length)} (${Math.round(compressed.length / data.length * 100)}% of original)`);
                    res.writeHead(200, {
                        'Content-Type': contentType,
                        'Content-Encoding': 'gzip',
                    });
                    res.end(compressed);
                } else {
                    res.writeHead(200, { 'Content-Type': contentType });
                    res.end(data);
                }
            });
        } else {
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(data);
        }
    });
});

server.listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}/`);
    console.log('Serving WASM with gzip compression when supported');
});
