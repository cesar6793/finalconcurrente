// index.js
const express = require('express');
const app = express();
const PORT = 3000;

// Ruta de ejemplo
app.get('/', (req, res) => {
  res.send('¡Hola desde mi API con Docker!');
});

// Escucha del servidor
app.listen(PORT, () => {
  console.log(`Servidor API escuchando en http://localhost:${PORT}`);
});
