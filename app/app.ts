import express, { Express } from "express";
import { engine } from "express-handlebars";
import templateHelpers from "./utils/template-helpers";
import pageRoutes from "./routes/pages";

export function createApp(): Express {
  const app = express();

  // Configure handlebars
  app.engine(
    "hbs",
    engine({
      extname: "hbs",
      helpers: templateHelpers,
    }),
  );
  app.set("view engine", "hbs");

  app.use(express.static("public"));
  app.use("/", pageRoutes);

  return app;
}
