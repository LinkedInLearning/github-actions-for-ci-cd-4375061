const express = require('express');
const app = express();
const fs = require('fs');
const winston = require('winston');

// Load data from file
const data = JSON.parse(fs.readFileSync('data.json'));

// Configure the logger to write to STDOUT
const logger = winston.createLogger({
  level: 'info',
  format: winston.format.combine(
      winston.format.timestamp(),
      winston.format.colorize(),
      winston.format.simple(),
      winston.format((info, opts) => {
        if (info.res) {
          info.message = `${info.message} ${info.res.statusCode}`;
        }
        return info;
      })(),
  ),
  transports: [
    new winston.transports.Console(),
  ],
});

// Logging middleware using the logger
app.use((req, res, next) => {
  res.on('finish', () => {
    logger.info(`${res.statusCode} ${req.method} ${req.url}`);
  });
  next();
});

// Define a route that returns all data
app.get('/', (req, res) => {
  res.send(data);
});

// Define a route that returns a single item of data based on GUID
app.get('/:guid', (req, res) => {
  const guid = req.params.guid;
  const item = data.find((item) => item.guid === guid);
  if (item) {
    res.send(item);
  } else {
    logger.error(`Item not found: ${guid}`);
    res.status(404).send('Item not found');
  }
});

module.exports = app.listen(3000);

