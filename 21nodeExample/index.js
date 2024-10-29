const express = require("express");
const app = express();
const PORT = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("Welcome");
});

app.get("/get", (req, res) => {
  res.status(200).send("GET Request");
});

app.post("/post", (req, res) => {
  let myJson = req.body;
  res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
  console.log(`Server is running on PORT ${PORT}`);
});
