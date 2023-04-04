const express = require("express");
const app = express();
const bodyParser = require("body-parser");

// Middleware to parse JSON requests
app.use(bodyParser.json());

// POST route to handle JSON request
app.all("*", (req, res) => {
  const data = req.body;
  console.log(data);
  // console.log(typeof data);
  // console.log(data.requestHeaders);
  // console.log(data["requestHeaders"]);
  //res.send(data);
  if (data["requestHeaders"][":path"] === "/http-bin-api-basic/1.0.8/get") {
    // if (data["requestHeaders"]["x-user-id"] === "bob") {
    //   res.send({
    //     throttleKeys: { rlkey_usergroup: "admin", rlkey_user: "bob" },
    //   });
    // } else {
    //   res.send({
    //     throttleKeys: { rlkey_usergroup: "admin" },
    //   });
    // }
    res.send({
      // throttleKeys: { rlkey_usergroup: "admin", rlkey_user: "bob" },
      rateLimitKeys: { rlkey_usergroup: "admin", rlkey_user: "bob" },
    });
  } else {
    res.send({
      // throttleKeys: { rlkey_usergroup: "admin" },
      rateLimitKeys: { rlkey_usergroup: "admin" },
    });
  }
});

// Start the server on port 8080
app.listen(8080, () => {
  console.log("Server is listening on port 8080");
});
