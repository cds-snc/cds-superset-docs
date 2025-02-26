import config from "./config";
import { createApp } from "./app";

const app = createApp();

app.listen(config.port, () => {
  console.log(`Server: http://localhost:${config.port}`);
  console.log(`WordPress: ${config.wordpress.url}`);
});
