const express = require('express');
const setupSwagger = require('./swagger');
const dotenv = require('dotenv');
const connectDB = require('./config/db');
const authRoutes = require('./routes/authRoutes');
const cors = require('cors');

dotenv.config();
connectDB();

const app = express();
app.use(cors());
app.use(express.json());

app.use('/api/auth', authRoutes);
setupSwagger(app);
console.log("📘 Swagger disponible sur http://localhost:3000/api/docs");

app.listen(process.env.PORT, () => {
    console.log(`🚀 Auth Service en écoute sur le port ${process.env.PORT}`);
});
