const express = require("express");
const app = express();
const Redis = require("ioredis");
const redis = new Redis(6379, "redis");
const fs = require("fs");

async function setFlag() {
  try {
    const data = fs.readFileSync("flag", "utf8");
    const flag = {
      id: "1",
      message: data,
      email: "FLAG@HERE.rc",
    };

    await redis.hset("Issue:1", flag);
    console.log("flag updated: " + data);
  } catch (err) {
    console.error(err);
  }
}

app.get("/issues", async (req, res) => {
  const keys = await redis.keys("Issue:*");
  const data = [];
  for (let key of keys) {
    data.push({
      id: await redis.hget(key, "id"),
      email: await redis.hget(key, "email"),
    });
  }
  return res.send(data);
});
app.get("/issues/:id", async (req, res) => {
  const data = await redis.hgetall("Issue:" + req.params.id);
  await setFlag();
  return res.send(data);
});
app.get("/reviews", async (req, res) => {
  const keys = await redis.keys("Review:*");
  const data = [];
  for (let key of keys) {
    data.push(await redis.hgetall(key));
    console.log(await redis.hgetall(key));
  }
  return res.send(data);
});
app.get("/keys", async (req, res) => {
  const data = await redis.keys();
  return res.send(data);
});

app.listen(3017, setFlag());
